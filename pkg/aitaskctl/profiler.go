package aitaskctl

import (
	"context"
	"fmt"
	"sync"
	"time"

	"hash/fnv"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/raids-lab/crater/dao/model"
	"github.com/raids-lab/crater/pkg/crclient"
	"github.com/raids-lab/crater/pkg/monitor"
	"github.com/raids-lab/crater/pkg/util/queue"
)

const (
	tickerDuration = 5 * time.Second
)

type Profiler struct {
	mutex            sync.Mutex
	taskQueue        queue.Queue                   //
	taskDB           DBService                     // update profiling status
	prometheusClient monitor.PrometheusInterface   // get monitor data
	podControl       *crclient.ProfilingPodControl // get pod status
	profilingTimeout time.Duration                 // profiling timeout
	profileCache     map[uint64]monitor.PodUtil
}

func NewProfiler(mgr manager.Manager, prometheusClient monitor.PrometheusInterface, profileTimeout int) *Profiler {
	return &Profiler{
		mutex:            sync.Mutex{}, // todo: add lock to taskQueue
		taskQueue:        queue.New(keyFunc, fifoOrdering),
		taskDB:           NewDBService(),
		profilingTimeout: time.Duration(profileTimeout) * time.Second, //todo: configuraion
		podControl:       &crclient.ProfilingPodControl{Client: mgr.GetClient()},
		prometheusClient: prometheusClient,
		profileCache:     make(map[uint64]monitor.PodUtil),
	}
}

func hashString(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func (p *Profiler) checkAndGetCache(task *model.AITask) (monitor.PodUtil, bool) {
	key := fmt.Sprintf("%s-%s", task.TaskType, task.Command)
	cacheKey := hashString(key)
	p.mutex.Lock()
	util, ok := p.profileCache[cacheKey]
	p.mutex.Unlock()
	klog.Infof("get profile cache, key:%v, ok:%v", key, ok)
	return util, ok
}

//nolint:gocritic // Must copy the util object
func (p *Profiler) storeProfileCache(task *model.AITask, util monitor.PodUtil) {
	key := fmt.Sprintf("%s-%s", task.TaskType, task.Command)
	cacheKey := hashString(key)
	p.mutex.Lock()
	p.profileCache[cacheKey] = util
	p.mutex.Unlock()
	klog.Infof("store profile cache, key:%v", key)
}

func (p *Profiler) SubmitProfileTask(taskID uint) {
	task, err := p.taskDB.GetByID(taskID)
	if err != nil {
		klog.Errorf("profiling task not found, taskID: %v", taskID)
		return
	}
	// check cache
	util, ok := p.checkAndGetCache(task)
	if ok {
		klog.Infof("profile cache hit, taskID:%v, taskName:%v, taskType:%v, command:%v",
			task.ID, task.TaskName, task.TaskType, task.Command)
		err = p.taskDB.UpdateProfilingStat(task.ID, model.EmiasProfileFinish, monitor.PodUtilToJSON(util), "")
		if err != nil {
			klog.Errorf("update profiling stat failed, taskID:%v, err:%v", taskID, err)
		}
		return
	}

	if task.ProfileStatus == model.EmiasUnProfiled {
		klog.Infof("submit profiling task, user:%v, taskName:%v, taskID: %v", task.UserName, task.TaskName, taskID)
		err = p.taskDB.UpdateProfilingStat(task.ID, model.EmiasProfileQueued, "", "")
		if err != nil {
			klog.Errorf("update profiling stat failed, taskID:%v, err:%v", taskID, err)
		}
		p.taskQueue.PushIfNotPresent(task)
	}
}
func (p *Profiler) DeleteProfilePodFromTask(taskID uint) {
	task, err := p.taskDB.GetByID(taskID)
	if err != nil {
		klog.Errorf("profiling task not found, taskID: %v", taskID)
		return
	}
	p.taskQueue.Delete(task)
	err = p.podControl.DeleteProfilePodFromTask(task)
	if err != nil {
		klog.Errorf("delete profiling pod failed, taskID:%v, taskName:%v, err:%v", task.ID, task.TaskName, err)
	}
}

func (p *Profiler) Start(ctx context.Context) {
	go p.run(ctx)
}

//nolint:gocyclo // todo: refactor
func (p *Profiler) run(ctx context.Context) {
	ticker := time.NewTicker(tickerDuration)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// create profiling pod
			// todo: check resource free
			// todo: check task status
			if p.taskQueue.Len() > 0 {
				t := p.taskQueue.Top()
				if t == nil {
					continue
				}
				task := t.(*model.AITask)
				// 1. create pod
				// 2. update task status

				err := p.podControl.CreateProfilePodFromTask(ctx, task)
				if err != nil {
					klog.Errorf("create profiling pod failed, taskID:%v, taskName:%v, err:%v", task.ID, task.TaskName, err)
					err = p.taskDB.UpdateProfilingStat(task.ID, model.EmiasProfileFailed, "", "")
					if err != nil {
						klog.Errorf("update profiling stat failed, taskID:%v, err:%v", task.ID, err)
					}
				} else {
					klog.Infof("create profiling pod success, taskID:%v, taskName:%v", task.ID, task.TaskName)
					err = p.taskDB.UpdateProfilingStat(task.ID, model.EmiasProfiling, "", "")
					if err != nil {
						klog.Errorf("update profiling stat failed, taskID:%v, err:%v", task.ID, err)
					}
				}
				p.taskQueue.Delete(task)
			}

			// check profiling pod status
			podList, err := p.podControl.ListProflingPods()
			if err != nil {
				klog.Errorf("list profiling pods failed: %v", err)
			}
			for i := range podList {
				pod := &podList[i]
				// get task
				taskID, err := p.podControl.GetTaskIDFromPod(pod)
				if err != nil {
					klog.Error(err)
					continue
				}
				task, _ := p.taskDB.GetByID(taskID)

				if pod.Status.Phase == corev1.PodPending {
					util, ok := p.checkAndGetCache(task)
					if ok {
						klog.Infof("profile cache hit, taskID:%v, taskName:%v, taskType:%v, command:%v",
							task.ID, task.TaskName, task.TaskType, task.Command)
						err = p.taskDB.UpdateProfilingStat(task.ID, model.EmiasProfileFinish, monitor.PodUtilToJSON(util), "")
						if err != nil {
							klog.Errorf("update profiling stat failed, taskID:%v, err:%v", taskID, err)
						}
						err = p.podControl.Delete(context.Background(), pod)
						if err != nil {
							klog.Errorf("delete profiling pod failed, taskID:%v, pod:%v/%v, err:%v", taskID, pod.Namespace, pod.Name, err)
						}
						continue
					}

					continue
				}
				// todo:
				// pod.Status.ContainerStatuses[0].State.Running.StartedAt?
				// pod.Status.StartTime
				if pod.Status.Phase == corev1.PodRunning && time.Since(pod.Status.StartTime.Time) < p.profilingTimeout {
					// p.taskDB.UpdateProfilingStat(task.ID, model.EmiasProfileFailed, "", "")
					// todo: pod running-> update profiling stat
					continue
				}
				if pod.Status.Phase == corev1.PodUnknown {
					klog.Errorf("profiling pod status unknow, pod: %v/%v", pod.Namespace, pod.Name)
					err = p.podControl.Delete(context.Background(), pod)
					if err != nil {
						klog.Errorf("delete profiling pod failed, taskID:%v, pod:%v/%v, err:%v", taskID, pod.Namespace, pod.Name, err)
					}
					continue
				}

				jobStatus := ""
				switch pod.Status.Phase {
				case corev1.PodFailed:
					jobStatus = model.EmiasTaskFailedStatus
				case corev1.PodSucceeded:
					jobStatus = model.EmiasTaskSucceededStatus
				}
				podUtil, err := p.prometheusClient.QueryPodProfileMetric(pod.Namespace, pod.Name)
				if err != nil {
					klog.Errorf("profile query pod util failed, taskID:%v, pod:%v/%v, err:%v", taskID, pod.Namespace, pod.Name, err)
					err = p.taskDB.UpdateProfilingStat(taskID, model.EmiasProfileFailed, "", jobStatus)
					if err != nil {
						klog.Errorf("update profiling stat failed, taskID:%v, err:%v", taskID, err)
					}
				} else {
					err = p.taskDB.UpdateProfilingStat(taskID, model.EmiasProfileFinish, monitor.PodUtilToJSON(podUtil), jobStatus)
					if err != nil {
						klog.Errorf("update profiling stat failed, taskID:%v, err:%v", taskID, err)
					}
					p.storeProfileCache(task, podUtil)
					// todo: error handle
					klog.Infof("profile query pod util success, taskID:%v, pod:%v/%v, status:%v", taskID, pod.Namespace, pod.Name, jobStatus)
				}
				err = p.podControl.Delete(context.Background(), pod)
				if err != nil {
					klog.Errorf("delete profiling pod failed, taskID:%v, pod:%v/%v, err:%v", taskID, pod.Namespace, pod.Name, err)
				}
			}
		}
	}
}
