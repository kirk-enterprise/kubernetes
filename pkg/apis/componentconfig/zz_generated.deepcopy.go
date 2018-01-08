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

package componentconfig

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_ClientConnectionConfiguration, InType: reflect.TypeOf(&ClientConnectionConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_GroupResource, InType: reflect.TypeOf(&GroupResource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_IPVar, InType: reflect.TypeOf(&IPVar{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeControllerManagerConfiguration, InType: reflect.TypeOf(&KubeControllerManagerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeProxyConfiguration, InType: reflect.TypeOf(&KubeProxyConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeProxyConntrackConfiguration, InType: reflect.TypeOf(&KubeProxyConntrackConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeProxyIPTablesConfiguration, InType: reflect.TypeOf(&KubeProxyIPTablesConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeSchedulerConfiguration, InType: reflect.TypeOf(&KubeSchedulerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletAnonymousAuthentication, InType: reflect.TypeOf(&KubeletAnonymousAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletAuthentication, InType: reflect.TypeOf(&KubeletAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletAuthorization, InType: reflect.TypeOf(&KubeletAuthorization{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletConfiguration, InType: reflect.TypeOf(&KubeletConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletWebhookAuthentication, InType: reflect.TypeOf(&KubeletWebhookAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletWebhookAuthorization, InType: reflect.TypeOf(&KubeletWebhookAuthorization{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletX509Authentication, InType: reflect.TypeOf(&KubeletX509Authentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_LeaderElectionConfiguration, InType: reflect.TypeOf(&LeaderElectionConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration, InType: reflect.TypeOf(&PersistentVolumeRecyclerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_PortRangeVar, InType: reflect.TypeOf(&PortRangeVar{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_VolumeConfiguration, InType: reflect.TypeOf(&VolumeConfiguration{})},
	)
}

// DeepCopy_componentconfig_ClientConnectionConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_ClientConnectionConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClientConnectionConfiguration)
		out := out.(*ClientConnectionConfiguration)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_GroupResource is an autogenerated deepcopy function.
func DeepCopy_componentconfig_GroupResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupResource)
		out := out.(*GroupResource)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_IPVar is an autogenerated deepcopy function.
func DeepCopy_componentconfig_IPVar(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IPVar)
		out := out.(*IPVar)
		*out = *in
		if in.Val != nil {
			in, out := &in.Val, &out.Val
			*out = new(string)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_componentconfig_KubeControllerManagerConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeControllerManagerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeControllerManagerConfiguration)
		out := out.(*KubeControllerManagerConfiguration)
		*out = *in
		if in.Controllers != nil {
			in, out := &in.Controllers, &out.Controllers
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.GCIgnoredResources != nil {
			in, out := &in.GCIgnoredResources, &out.GCIgnoredResources
			*out = make([]GroupResource, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_componentconfig_KubeProxyConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeProxyConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeProxyConfiguration)
		out := out.(*KubeProxyConfiguration)
		*out = *in
		if err := DeepCopy_componentconfig_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, c); err != nil {
			return err
		}
		if in.OOMScoreAdj != nil {
			in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_componentconfig_KubeProxyConntrackConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeProxyConntrackConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeProxyConntrackConfiguration)
		out := out.(*KubeProxyConntrackConfiguration)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_KubeProxyIPTablesConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeProxyIPTablesConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeProxyIPTablesConfiguration)
		out := out.(*KubeProxyIPTablesConfiguration)
		*out = *in
		if in.MasqueradeBit != nil {
			in, out := &in.MasqueradeBit, &out.MasqueradeBit
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_componentconfig_KubeSchedulerConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeSchedulerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeSchedulerConfiguration)
		out := out.(*KubeSchedulerConfiguration)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_KubeletAnonymousAuthentication is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeletAnonymousAuthentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletAnonymousAuthentication)
		out := out.(*KubeletAnonymousAuthentication)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_KubeletAuthentication is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeletAuthentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletAuthentication)
		out := out.(*KubeletAuthentication)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_KubeletAuthorization is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeletAuthorization(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletAuthorization)
		out := out.(*KubeletAuthorization)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_KubeletConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeletConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletConfiguration)
		out := out.(*KubeletConfiguration)
		*out = *in
		if in.HostNetworkSources != nil {
			in, out := &in.HostNetworkSources, &out.HostNetworkSources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.HostPIDSources != nil {
			in, out := &in.HostPIDSources, &out.HostPIDSources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.HostIPCSources != nil {
			in, out := &in.HostIPCSources, &out.HostIPCSources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.ClusterDNS != nil {
			in, out := &in.ClusterDNS, &out.ClusterDNS
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.PodPidsLimit != nil {
			in, out := &in.PodPidsLimit, &out.PodPidsLimit
			*out = new(int64)
			**out = **in
		}
		if in.RegisterWithTaints != nil {
			in, out := &in.RegisterWithTaints, &out.RegisterWithTaints
			*out = make([]api.Taint, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_Taint(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.NodeLabels != nil {
			in, out := &in.NodeLabels, &out.NodeLabels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.ExperimentalQOSReserved != nil {
			in, out := &in.ExperimentalQOSReserved, &out.ExperimentalQOSReserved
			*out = make(ConfigurationMap)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.AllowedUnsafeSysctls != nil {
			in, out := &in.AllowedUnsafeSysctls, &out.AllowedUnsafeSysctls
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.SystemReserved != nil {
			in, out := &in.SystemReserved, &out.SystemReserved
			*out = make(ConfigurationMap)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.KubeReserved != nil {
			in, out := &in.KubeReserved, &out.KubeReserved
			*out = make(ConfigurationMap)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.EnforceNodeAllocatable != nil {
			in, out := &in.EnforceNodeAllocatable, &out.EnforceNodeAllocatable
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_componentconfig_KubeletWebhookAuthentication is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeletWebhookAuthentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletWebhookAuthentication)
		out := out.(*KubeletWebhookAuthentication)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_KubeletWebhookAuthorization is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeletWebhookAuthorization(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletWebhookAuthorization)
		out := out.(*KubeletWebhookAuthorization)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_KubeletX509Authentication is an autogenerated deepcopy function.
func DeepCopy_componentconfig_KubeletX509Authentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletX509Authentication)
		out := out.(*KubeletX509Authentication)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_LeaderElectionConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_LeaderElectionConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LeaderElectionConfiguration)
		out := out.(*LeaderElectionConfiguration)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PersistentVolumeRecyclerConfiguration)
		out := out.(*PersistentVolumeRecyclerConfiguration)
		*out = *in
		return nil
	}
}

// DeepCopy_componentconfig_PortRangeVar is an autogenerated deepcopy function.
func DeepCopy_componentconfig_PortRangeVar(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PortRangeVar)
		out := out.(*PortRangeVar)
		*out = *in
		if in.Val != nil {
			in, out := &in.Val, &out.Val
			*out = new(string)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_componentconfig_VolumeConfiguration is an autogenerated deepcopy function.
func DeepCopy_componentconfig_VolumeConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*VolumeConfiguration)
		out := out.(*VolumeConfiguration)
		*out = *in
		return nil
	}
}
