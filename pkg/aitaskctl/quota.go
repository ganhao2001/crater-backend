package aitaskctl

import (
	"context"
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"

	"github.com/raids-lab/crater/dao/model"
	"github.com/raids-lab/crater/dao/query"
)

type QuotaInfo struct {
	mu        sync.RWMutex
	Name      string
	Hard      v1.ResourceList
	Soft      v1.ResourceList
	HardUsed  v1.ResourceList
	SoftUsed  v1.ResourceList
	UsedTasks map[string]*model.AITask
}

type QuotaInfoSnapshot struct {
	Name      string
	Hard      v1.ResourceList
	Soft      v1.ResourceList
	HardUsed  v1.ResourceList
	SoftUsed  v1.ResourceList
	UsedTasks map[string]*model.AITask
}

// AddTask adds Running Job Quota
func (info *QuotaInfo) AddTask(task *model.AITask) {
	info.mu.Lock()
	defer info.mu.Unlock()
	key := keyFunc(task)
	// 没有找到task的时候才添加quota
	if _, ok := info.UsedTasks[key]; !ok {
		info.UsedTasks[key] = task
		resourceRequest, _ := model.JSONToResourceList(task.ResourceRequest)
		switch task.SLO {
		case model.EmiasHighSLO:
			AddResourceList(info.HardUsed, resourceRequest)
		case model.EmiasLowSLO:
			AddResourceList(info.SoftUsed, resourceRequest)
		}
	}
}

func (info *QuotaInfoSnapshot) AddTask(task *model.AITask) {
	key := keyFunc(task)
	// 没有找到task的时候才添加quota
	if _, ok := info.UsedTasks[key]; !ok {
		info.UsedTasks[key] = task
		resourceRequest, _ := model.JSONToResourceList(task.ResourceRequest)
		switch task.SLO {
		case model.EmiasHighSLO:
			AddResourceList(info.HardUsed, resourceRequest)
		case model.EmiasLowSLO:
			AddResourceList(info.SoftUsed, resourceRequest)
		}
	}
}

// DeleteTask deletes Completed or Deleted Job Quota
func (info *QuotaInfo) DeleteTask(task *model.AITask) {
	info.mu.Lock()
	defer info.mu.Unlock()
	key := keyFunc(task)
	// 找到quotainfo里的task时才删除quota
	if task, ok := info.UsedTasks[key]; ok {
		delete(info.UsedTasks, key)
		resourceRequest, _ := model.JSONToResourceList(task.ResourceRequest)
		switch task.SLO {
		case model.EmiasHighSLO:
			SubResourceList(info.HardUsed, resourceRequest)
		case model.EmiasLowSLO:
			SubResourceList(info.SoftUsed, resourceRequest)
		}
	}
}

// CheckHardQuotaExceed 判断作业的hard quota是否超出限制
func (info *QuotaInfo) CheckHardQuotaExceed(task *model.AITask) bool {
	if task.SLO == model.EmiasLowSLO {
		return false
	}
	info.mu.Lock()
	defer info.mu.Unlock()
	resourcelist, _ := model.JSONToResourceList(task.ResourceRequest)
	return CheckResourceListExceed(info.Hard, info.HardUsed, resourcelist)
}

func (info *QuotaInfoSnapshot) CheckHardQuotaExceed(task *model.AITask) bool {
	if task.SLO == model.EmiasLowSLO {
		return false
	}
	resourcelist, _ := model.JSONToResourceList(task.ResourceRequest)
	return CheckResourceListExceed(info.Hard, info.HardUsed, resourcelist)
}

func (info *QuotaInfo) Snapshot() *QuotaInfoSnapshot {
	info.mu.Lock()
	defer info.mu.Unlock()
	return &QuotaInfoSnapshot{
		Name:      info.Name,
		Hard:      info.Hard.DeepCopy(),
		Soft:      info.Soft.DeepCopy(),
		HardUsed:  info.HardUsed.DeepCopy(),
		SoftUsed:  info.SoftUsed.DeepCopy(),
		UsedTasks: make(map[string]*model.AITask),
	}
}

func (c *TaskController) GetQuotaInfo(username string) *QuotaInfo {
	if value, ok := c.quotaInfos.Load(username); ok {
		return value.(*QuotaInfo)
	} else {
		q := query.Account
		quotadb, err := q.WithContext(context.Background()).Where(q.Name.Eq((username))).First()
		if err != nil {
			klog.Errorf("get quota from db failed, err: %v", err)
			return nil
		}
		_, info := c.AddOrUpdateQuotaInfo(quotadb)
		return info
	}
}

// GetQuotaInfoSnapshotByUsername 获取某个用户的QuotaInfo的clone，对quota的增加减少不改变原数据
func (c *TaskController) GetQuotaInfoSnapshotByUsername(username string) *QuotaInfoSnapshot {
	if value, ok := c.quotaInfos.Load(username); ok {
		return value.(*QuotaInfo).Snapshot()
	} else {
		return nil
	}
}

// GetQuotaInfoSnapshotByUsername 获取某个用户的QuotaInfo的clone，对quota的增加减少不改变原数据
func (c *TaskController) ListQuotaInfoSnapshot() []QuotaInfoSnapshot {
	quotaInfos := make([]QuotaInfoSnapshot, 0)
	c.quotaInfos.Range(func(_, value any) bool {
		info := value.(*QuotaInfo)
		infoSnapShot := info.Snapshot()
		quotaInfos = append(quotaInfos, *infoSnapShot)
		return true
	})
	return quotaInfos
}

func (c *TaskController) AddOrUpdateQuotaInfo(queue *model.Account) (added bool, quotaInfo *QuotaInfo) {
	name := queue.Name
	quota := queue.Quota.Data()
	if _, ok := c.quotaInfos.Load(queue.Name); !ok {
		quotaInfo := &QuotaInfo{
			Name:      name,
			Hard:      quota.Deserved,
			Soft:      quota.Capability,
			HardUsed:  v1.ResourceList{},
			SoftUsed:  v1.ResourceList{},
			UsedTasks: make(map[string]*model.AITask),
		}

		// todo: add db tasks
		tasksRunning, err := c.taskDB.ListByUserAndStatuses(name, model.EmiasTaskOcupiedQuotaStatuses)
		//nolint:staticcheck // TODO: remove this line after fixing the error
		if err != nil {
			// todo: handler err
		}

		for i := range tasksRunning {
			quotaInfo.AddTask(tasksRunning[i])
		}

		c.quotaInfos.Store(name, quotaInfo)
		added = true
	} else {
		c.UpdateQuotaInfoHard(name, quota.Deserved)
		added = false
	}
	return added, c.GetQuotaInfo(name)
}

// UpdateQuotaInfoHard updates QuotaInfo Hard
func (c *TaskController) UpdateQuotaInfoHard(username string, hard v1.ResourceList) {
	if value, ok := c.quotaInfos.Load(username); ok {
		info := value.(*QuotaInfo)
		info.mu.Lock()
		defer info.mu.Unlock()
		if !CmpResourceListSame(info.Hard, hard) {
			info.Hard = hard.DeepCopy()
		}
	}
}

// DeleteQuotaInfo deletes QuotaInfo
func (c *TaskController) DeleteQuotaInfo(namespace string) {
	c.quotaInfos.Delete(namespace)
}
