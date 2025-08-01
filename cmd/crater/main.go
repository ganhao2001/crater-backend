/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"
	"time"

	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/raids-lab/crater/cmd/crater/helper"
)

// @title						Crater API
// @version						1.0.0
// @description					This is the API server for Crater, a Multi-tenant AI Model Training Platform based on Kubernetes.
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description					访问 /login 并获取 TOKEN 后，填入 'Bearer ${TOKEN}' 以访问受保护的接口
func main() {
	// Set global timezone
	time.Local = time.UTC

	// Initialize configuration
	configInit := helper.NewConfigInitializer()
	backendConfig := configInit.GetBackendConfig()

	// Load debug environment if needed
	if err := configInit.LoadDebugEnvironment(); err != nil {
		panic(err.Error())
	}

	// Initialize register config and dependencies
	registerConfig, err := configInit.InitializeRegisterConfig()
	if err != nil {
		panic(err.Error())
	}

	// Setup server runner and logger
	serverRunner := helper.NewServerRunner(backendConfig)
	serverRunner.SetupLogger()

	// Initialize signal handler
	stopCh := ctrl.SetupSignalHandler()

	// Create and setup manager
	managerSetup := helper.NewManagerSetup(registerConfig.KubeConfig, backendConfig)
	mgr, err := managerSetup.CreateCRDManager()
	if err != nil {
		klog.ErrorS(err, "unable to create manager")
		os.Exit(1)
	}

	// Setup manager dependencies
	configInit.SetupManagerDependencies(registerConfig, mgr)

	// Setup custom CRD addons
	err = managerSetup.SetupCustomCRDAddon(mgr, registerConfig, stopCh)
	if err != nil {
		klog.ErrorS(err, "unable to set up custom CRD addon")
		os.Exit(1)
	}

	// Setup health checks
	serverRunner.SetupHealthChecks(mgr)

	// Start manager
	serverRunner.StartManager(mgr, stopCh)

	// Start HTTP server
	serverRunner.StartServer(registerConfig)
}
