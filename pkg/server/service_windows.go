// +build windows

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

package server

import (
	"github.com/containerd/containerd/platforms"
	runtime "k8s.io/kubernetes/pkg/kubelet/apis/cri/runtime/v1alpha2"
)

// isApparmorEnabled is not supported on Windows.
func isApparmorEnabled() bool {
	return false
}

// isSeccompEnabled is not supported on Windows.
func isSeccompEnabled() bool {
	return false
}

// doSelinux is not supported on Windows.
func doSelinux(enable bool) {
}

func (c *criService) getDefaultSnapshotterForSandbox(cfg *runtime.PodSandboxConfig) string {
	if getDefaultPlatform(cfg) == "linux/amd64" {
		return "windows-lcow"
	}
	return c.config.ContainerdConfig.Snapshotter
}

func (c *criService) getDefaultSandboxImage(cfg *runtime.PodSandboxConfig) string {
	if getDefaultPlatform(cfg) == "linux/amd64" {
		return "k8s.gcr.io/pause:3.1"
	}
	return c.config.SandboxImage
}

func getDefaultPlatform(cfg *runtime.PodSandboxConfig) string {
	if cfg != nil {
		if plat, ok := cfg.Labels["sandbox-platform"]; ok {
			return plat
		}
	}
	return platforms.DefaultString()
}
