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
	"github.com/containerd/containerd/oci"
	"github.com/opencontainers/runtime-tools/generate"
	runtime "k8s.io/kubernetes/pkg/kubelet/apis/cri/runtime/v1alpha2"
)

// generateSeccompSpecOpts unsupported on Windows.
func generateSeccompSpecOpts(seccompProf string, privileged, seccompEnabled bool) (oci.SpecOpts, error) {
	return nil, nil
}

// generateApparmorSpecOpts unsupported on Windows.
func generateApparmorSpecOpts(apparmorProf string, privileged, apparmorEnabled bool) (oci.SpecOpts, error) {
	return nil, nil
}

// addDevices set device mapping without privilege.
func (c *criService) addOCIDevices(g *generate.Generator, devs []*runtime.Device) error {
	// TODO: JTERRY75 - WCOW/LCOW support
	return nil
}

// addDevices set device mapping with privilege.
func setOCIDevicesPrivileged(g *generate.Generator) error {
	// TODO: JTERRY75 - WCOW/LCOW support
	return nil
}

// doSelinuxRelabel not supported on Windows.
func doSelinuxRelable(path string, fileLabel string, shared bool) error {
	return nil
}
