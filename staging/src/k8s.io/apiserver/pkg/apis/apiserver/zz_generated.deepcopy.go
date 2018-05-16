// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package apiserver

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiserver_AdmissionConfiguration, InType: reflect.TypeOf(&AdmissionConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiserver_AdmissionPluginConfiguration, InType: reflect.TypeOf(&AdmissionPluginConfiguration{})},
	)
}

// DeepCopy_apiserver_AdmissionConfiguration is an autogenerated deepcopy function.
func DeepCopy_apiserver_AdmissionConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AdmissionConfiguration)
		out := out.(*AdmissionConfiguration)
		*out = *in
		if in.Plugins != nil {
			in, out := &in.Plugins, &out.Plugins
			*out = make([]AdmissionPluginConfiguration, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*AdmissionPluginConfiguration)
				}
			}
		}
		return nil
	}
}

// DeepCopy_apiserver_AdmissionPluginConfiguration is an autogenerated deepcopy function.
func DeepCopy_apiserver_AdmissionPluginConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AdmissionPluginConfiguration)
		out := out.(*AdmissionPluginConfiguration)
		*out = *in
		// in.Configuration is kind 'Interface'
		if in.Configuration != nil {
			if newVal, err := c.DeepCopy(&in.Configuration); err != nil {
				return err
			} else {
				out.Configuration = *newVal.(*runtime.Object)
			}
		}
		return nil
	}
}
