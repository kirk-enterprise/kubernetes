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
	"strconv"
	"testing"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/kubernetes/pkg/api"
)

func getPod(name, cpu, gpu, mem string) *api.Pod {
	res := api.ResourceRequirements{}
	res.Requests = api.ResourceList{
		api.ResourceCPU: resource.MustParse(cpu),
		api.ResourceMemory: resource.MustParse(mem),
		api.ResourceNvidiaGPU: resource.MustParse(gpu),
	}
	res.Limits = api.ResourceList{
		api.ResourceCPU: resource.MustParse(cpu),
		api.ResourceMemory: resource.MustParse(mem),
		api.ResourceNvidiaGPU: resource.MustParse(gpu),
	}

	pod := &api.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: "test"},
		Spec:       api.PodSpec{},
	}
	pod.Spec.Containers = make([]api.Container, 0, numContainers)
	for i := 0; i < numContainers; i++ {
		pod.Spec.Containers = append(pod.Spec.Containers, api.Container{
			Image:     "foo:V" + strconv.Itoa(i),
			Resources: res,
		})
	}

	return pod
}

func checkFieldEqual(a, b resource.Quantity) {
	if a.cmp(b) == 0 {
		return true
	} else {
		return false
	}
}

func assertAllFieldsEquality(pod *api.Pod, want *api.Pod) {
	for i, c := range pod.Spec.InitContainers {
		l := c.Resources
		w := want.Spec.InitContainers[i].Resources
		if !checkFieldEqual(l.Limits[api.ResourceCPU], w.Limits[api.ResourceCPU])
		|| !checkFieldEqual(l.Limits[api.ResourceMemory], w.Limits[api.ResourceMemory])
		|| !checkFieldEqual(l.Limits[api.ResourceNvidiaGPU], w.Limits[api.ResourceNvidiaGPU]) {
			return false
		}
		if !checkFieldEqual(l.Requests[api.ResourceCPU], w.Requests[api.ResourceCPU])
		|| !checkFieldEqual(l.Requests[api.ResourceMemory], w.Requests[api.ResourceMemory])
		|| !checkFieldEqual(l.Requests[api.ResourceNvidiaGPU], w.Requests[api.ResourceNvidiaGPU]) {
			return false
		}
	}
	for i, c := range pod.Spec.Containers {
		l := c.Resources
		w := want.Spec.Containers[i].Resources
		if !checkFieldEqual(l.Limits[api.ResourceCPU], w.Limits[api.ResourceCPU])
		|| !checkFieldEqual(l.Limits[api.ResourceMemory], w.Limits[api.ResourceMemory])
		|| !checkFieldEqual(l.Limits[api.ResourceNvidiaGPU], w.Limits[api.ResourceNvidiaGPU]) {
			return false
		}
		if !checkFieldEqual(l.Requests[api.ResourceCPU], w.Requests[api.ResourceCPU])
		|| !checkFieldEqual(l.Requests[api.ResourceMemory], w.Requests[api.ResourceMemory])
		|| !checkFieldEqual(l.Requests[api.ResourceNvidiaGPU], w.Requests[api.ResourceNvidiaGPU]) {
			return false
		}
	}
	return true
}

func TestAdmit(t *testing.T) {
	handler := NewLimitPolicyPlugin()
	testCases := []struct {
		Pod *api.Pod,
		Want, *api.Pod
	}{
		{
			getPod("test", "1", "2", "128Mem")
			getPod("test", 1, "1", "128Mem")
		},
		{
			getPod("test", 1, "2", "512Mem")
			getPod("test", 2, "2", "256Mem")
		}
	}
	for _, c := range testCases {
		if err := handler.Admit(admission.NewAttributesRecord(newPod, nil, 
				api.Kind("Pod").WithVersion("version"), 
				newPod.Namespace, newPod.Name, 
				api.Resource("pods").WithVersion("version"), "", 
				admission.Create, nil)); err != nil {
			t.Errorf("Unexpected error returned from admission handler")
		}
		if check := assertAllFieldsEquality(c.Pod, c.Want); !check {
			t.Errorf("Expect pod: %#v\n But got pod: %#v\n", c.Want, c.Pod)
		} 
	}
}

func TestHandles(t *testing.T) {
	for op, shouldHandle := range map[admission.Operation]bool{
		admission.Create:  true,
		admission.Update:  false,
		admission.Connect: false,
		admission.Delete:  false,
	} {
		handler := NewLimitPolicyPlugin()
		if e, a := shouldHandle, handler.Handles(op); e != a {
			t.Errorf("%v: shouldHandle=%t, handles=%t", op, e, a)
		}
	}
}
