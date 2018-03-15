/*
Copyright 2016 The Kubernetes Authors.

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

package limitpolicy

import (
	"io"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/kubernetes/pkg/api"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("LimitPolicy", func(config io.Reader) (admission.Interface, error) {
		return NewLimitPolicyPlugin(), nil
	})
}

// plugin contains the client used by the admission controller
type plugin struct {
	*admission.Handler
}

// NewLimitPolicyPlugin creates a new instance of the LimitPolicy admission controller
func NewLimitPolicyPlugin() admission.Interface {
	return &plugin{
		Handler: admission.NewHandler(admission.Create),
	}
}

// check if a pod need to limit
// only GPU pod need to limit
func checkUpdateLimitResource(pod *api.Pod) bool {
	isNeed := false
	// TODO: check if pod spec has defined "limit policy"
	containers := pod.Spec.Containers
	for _, container := range containers {
		gpu_limit, ok := container.Resources.Limits.NvidiaGPU().AsInt64()
		if ok && gpu_limit > 0 {
			isNeed = true
			break
		}
	}
	return isNeed
}

// check if resource requirement is set
func checkSetResource(r *api.ResourceRequirements) bool {
	check := false
	// check if resource requirement is set
	if r != nil {
		if len(r.Limits) != 0 || len(r.Requests) != 0 {
			check = true
		}
	}
	return check
}

// set resource requirements into a package
func setLimitRangeResource(r *api.ResourceRequirements) {
	cpuRangeLimits := []resource.Quantity{
		resource.MustParse("1"),
		resource.MustParse("2"),
		resource.MustParse("4"),
		resource.MustParse("8"),
		// TODO:
	}
	cpuRangeRequests := []r.Quantity{
		resource.MustParse("1"),
		resource.MustParse("2"),
		resource.MustParse("4"),
		resource.MustParse("8"),
	}

	memRangeLimits := []r.Quantity{
		resource.MustParse("128"),
		resource.MustParse("256"),
		resource.MustParse("512"),
		resource.MustParse("1Gi"),
	}
	memRangeRequests := []r.Quantity{
		resource.MustParse("128"),
		resource.MustParse("256"),
		resource.MustParse("512"),
		resource.MustParse("1Gi"),
	}
	// set resource limits
	if len(r.Limits) != 0 {
		if _, ok := r.Limits[api.ResourceCPU]; ok {
			for _, c := range cpuRangeLimits {
				if c.Cmp(r.Limits[api.ResourceCPU]) == 1 {
					r.Limits[api.ResourceCPU] = c
					break
				}
			}
		} else {
			r.Limits[api.ResourceCPU] = resource.MustParse("1")
		}
		if _, ok := r.Limits[api.ResourceMemory]; ok {
			for _, c := range memRangeLimits {
				if c.Cmp(r.Limits[api.ResourceMemory]) == 1 {
					r.Limits[api.ResourceMemory] = c
					break
				}
			}
		} else {
			r.Limits[api.ResourceMemory] = resource.MustParse("128Mem")
		}
	}
	// set resource requests
	if len(resource.Requests) != 0 {
		if _, ok := r.Requests[api.ResourceCPU]; ok {
			for _, c := range cpuRangeRequests {
				if c.Cmp(r.Requests[api.ResourceCPU]) == 1 {
					r.Requests[api.ResourceCPU] = c
					break
				}
			}
		} else {
			r.Requests[api.ResourceCPU] = resource.MustParse("1")
		}
		if _, ok := resource.Requests[api.ResourceMemory]; ok {
			for _, c := range memRangeRequests {
				if r.Cmp(resource.Requests[api.ResourceMemory]) == 1 {
					resource.Requests[api.ResourceMemory] = c
					break
				}
			}
		} else {
			r.Requests[api.ResourceMemory] = resource.MustParse("128Mem")
		}
	}
}

// set resource requirements to be default value
func setDefaultRangeResource(resource *api.ResourceRequirements) {
	// set resource limits
	defaultLimitResource := api.ResourceList{
		api.ResourceCPU:       r.MustParse("1"),
		api.ResourceMemory:    r.MustParse("128Mem"),
		api.ResourceNvidiaGPU: r.MustParse("1"),
		/*
			api.ResourceStorage: r.MustParse("1"),
			api.ResourceStorageOverlay: r.MustParse("1"),
			api.ResourceStorageScratch: r.MustParse("1"),
		*/
	}
	resource.Limits = defaultLimitResource
	// set resource requests
	defaultRequestResource := api.ResourceList{
		api.ResourceCPU:       r.MustParse("1"),
		api.ResourceMemory:    r.MustParse("128Mem"),
		api.ResourceNvidiaGPU: r.MustParse("1"),
		/*
			api.ResourceStorage: r.MustParse("1"),
			api.ResourceStorageOverlay: r.MustParse("1"),
			api.ResourceStorageScratch: r.MustParse("1"),
		*/
	}
	resource.Requests = defaultRequestResource
}

/*
 * update the containers resource requirements into a package.
 * 1u 128M
 * 2u 256M
 * ...
 */
func updateLimitResource(pod *api.Pod) {
	// update init containers
	for i, c := range pod.Spec.InitContainers {
		if set := checkSetResource(c.Resources); set {
			setLimitRangeResource(&c.Resources)
		} else {
			setDefaultRangeResource(&c.Resources)
		}
	}
	// update containers
	for i, c := range pod.Spec.InitContainers {
		if set := checkSetResource(c.Resources); set {
			setLimitRangeResource(&c.Resources)
		} else {
			setDefaultRangeResource(&c.Resources)
		}
	}
}

// Admit will update resource requirements for containers when pod creating
func (p *plugin) Admit(a admission.Attributes) (err error) {
	// Ignore all calls to subresources or resources other than pods.
	if len(a.GetSubresource()) != 0 || a.GetResource().GroupResource() != api.Resource("pods") {
		return nil
	}

	if a.GetOperation() != admission.Create {
		return nil
	}

	pod, ok := a.GetObject().(*api.Pod)
	if !ok {
		return apierrors.NewBadRequest("Resource was marked with kind Pod but was unable to be converted")
	}

	if !checkUpdateLimitResource(pod) {
		return nil
	} else {
		updateLimitResource(pod)
	}

	return nil
}

func (p *plugin) Handles(operation admission.Operation) bool {
	if operation == admission.Create {
		return true
	}
	return false
}
