// +build !windows

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
	"strings"

	"github.com/containerd/containerd/contrib/apparmor"
	"github.com/containerd/containerd/contrib/seccomp"
	"github.com/containerd/containerd/oci"
	"github.com/pkg/errors"
)

// generateSeccompSpecOpts generates containerd SpecOpts for seccomp.
func generateSeccompSpecOpts(seccompProf string, privileged, seccompEnabled bool) (oci.SpecOpts, error) {
	if privileged {
		// Do not set seccomp profile when container is privileged
		return nil, nil
	}
	// Set seccomp profile
	if seccompProf == runtimeDefault || seccompProf == dockerDefault {
		// use correct default profile (Eg. if not configured otherwise, the default is docker/default)
		seccompProf = seccompDefaultProfile
	}
	if !seccompEnabled {
		if seccompProf != "" && seccompProf != unconfinedProfile {
			return nil, errors.New("seccomp is not supported")
		}
		return nil, nil
	}
	switch seccompProf {
	case "", unconfinedProfile:
		// Do not set seccomp profile.
		return nil, nil
	case dockerDefault:
		// Note: WithDefaultProfile specOpts must be added after capabilities
		return seccomp.WithDefaultProfile(), nil
	default:
		// Require and Trim default profile name prefix
		if !strings.HasPrefix(seccompProf, profileNamePrefix) {
			return nil, errors.Errorf("invalid seccomp profile %q", seccompProf)
		}
		return seccomp.WithProfile(strings.TrimPrefix(seccompProf, profileNamePrefix)), nil
	}
}

// generateApparmorSpecOpts generates containerd SpecOpts for apparmor.
func generateApparmorSpecOpts(apparmorProf string, privileged, apparmorEnabled bool) (oci.SpecOpts, error) {
	if !apparmorEnabled {
		// Should fail loudly if user try to specify apparmor profile
		// but we don't support it.
		if apparmorProf != "" && apparmorProf != unconfinedProfile {
			return nil, errors.New("apparmor is not supported")
		}
		return nil, nil
	}
	switch apparmorProf {
	case runtimeDefault:
		// TODO (mikebrow): delete created apparmor default profile
		return apparmor.WithDefaultProfile(appArmorDefaultProfileName), nil
	case unconfinedProfile:
		return nil, nil
	case "":
		// Based on kubernetes#51746, default apparmor profile should be applied
		// for non-privileged container when apparmor is not specified.
		if privileged {
			return nil, nil
		}
		return apparmor.WithDefaultProfile(appArmorDefaultProfileName), nil
	default:
		// Require and Trim default profile name prefix
		if !strings.HasPrefix(apparmorProf, profileNamePrefix) {
			return nil, errors.Errorf("invalid apparmor profile %q", apparmorProf)
		}
		return apparmor.WithProfile(strings.TrimPrefix(apparmorProf, profileNamePrefix)), nil
	}
}
