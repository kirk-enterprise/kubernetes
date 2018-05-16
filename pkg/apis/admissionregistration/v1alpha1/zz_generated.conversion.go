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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_AdmissionHookClientConfig_To_admissionregistration_AdmissionHookClientConfig,
		Convert_admissionregistration_AdmissionHookClientConfig_To_v1alpha1_AdmissionHookClientConfig,
		Convert_v1alpha1_ExternalAdmissionHook_To_admissionregistration_ExternalAdmissionHook,
		Convert_admissionregistration_ExternalAdmissionHook_To_v1alpha1_ExternalAdmissionHook,
		Convert_v1alpha1_ExternalAdmissionHookConfiguration_To_admissionregistration_ExternalAdmissionHookConfiguration,
		Convert_admissionregistration_ExternalAdmissionHookConfiguration_To_v1alpha1_ExternalAdmissionHookConfiguration,
		Convert_v1alpha1_ExternalAdmissionHookConfigurationList_To_admissionregistration_ExternalAdmissionHookConfigurationList,
		Convert_admissionregistration_ExternalAdmissionHookConfigurationList_To_v1alpha1_ExternalAdmissionHookConfigurationList,
		Convert_v1alpha1_Initializer_To_admissionregistration_Initializer,
		Convert_admissionregistration_Initializer_To_v1alpha1_Initializer,
		Convert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration,
		Convert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration,
		Convert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList,
		Convert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList,
		Convert_v1alpha1_Rule_To_admissionregistration_Rule,
		Convert_admissionregistration_Rule_To_v1alpha1_Rule,
		Convert_v1alpha1_RuleWithOperations_To_admissionregistration_RuleWithOperations,
		Convert_admissionregistration_RuleWithOperations_To_v1alpha1_RuleWithOperations,
		Convert_v1alpha1_ServiceReference_To_admissionregistration_ServiceReference,
		Convert_admissionregistration_ServiceReference_To_v1alpha1_ServiceReference,
	)
}

func autoConvert_v1alpha1_AdmissionHookClientConfig_To_admissionregistration_AdmissionHookClientConfig(in *AdmissionHookClientConfig, out *admissionregistration.AdmissionHookClientConfig, s conversion.Scope) error {
	if err := Convert_v1alpha1_ServiceReference_To_admissionregistration_ServiceReference(&in.Service, &out.Service, s); err != nil {
		return err
	}
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	return nil
}

// Convert_v1alpha1_AdmissionHookClientConfig_To_admissionregistration_AdmissionHookClientConfig is an autogenerated conversion function.
func Convert_v1alpha1_AdmissionHookClientConfig_To_admissionregistration_AdmissionHookClientConfig(in *AdmissionHookClientConfig, out *admissionregistration.AdmissionHookClientConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmissionHookClientConfig_To_admissionregistration_AdmissionHookClientConfig(in, out, s)
}

func autoConvert_admissionregistration_AdmissionHookClientConfig_To_v1alpha1_AdmissionHookClientConfig(in *admissionregistration.AdmissionHookClientConfig, out *AdmissionHookClientConfig, s conversion.Scope) error {
	if err := Convert_admissionregistration_ServiceReference_To_v1alpha1_ServiceReference(&in.Service, &out.Service, s); err != nil {
		return err
	}
	if in.CABundle == nil {
		out.CABundle = make([]byte, 0)
	} else {
		out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	}
	return nil
}

// Convert_admissionregistration_AdmissionHookClientConfig_To_v1alpha1_AdmissionHookClientConfig is an autogenerated conversion function.
func Convert_admissionregistration_AdmissionHookClientConfig_To_v1alpha1_AdmissionHookClientConfig(in *admissionregistration.AdmissionHookClientConfig, out *AdmissionHookClientConfig, s conversion.Scope) error {
	return autoConvert_admissionregistration_AdmissionHookClientConfig_To_v1alpha1_AdmissionHookClientConfig(in, out, s)
}

func autoConvert_v1alpha1_ExternalAdmissionHook_To_admissionregistration_ExternalAdmissionHook(in *ExternalAdmissionHook, out *admissionregistration.ExternalAdmissionHook, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_v1alpha1_AdmissionHookClientConfig_To_admissionregistration_AdmissionHookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
		return err
	}
	out.Rules = *(*[]admissionregistration.RuleWithOperations)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*admissionregistration.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	return nil
}

// Convert_v1alpha1_ExternalAdmissionHook_To_admissionregistration_ExternalAdmissionHook is an autogenerated conversion function.
func Convert_v1alpha1_ExternalAdmissionHook_To_admissionregistration_ExternalAdmissionHook(in *ExternalAdmissionHook, out *admissionregistration.ExternalAdmissionHook, s conversion.Scope) error {
	return autoConvert_v1alpha1_ExternalAdmissionHook_To_admissionregistration_ExternalAdmissionHook(in, out, s)
}

func autoConvert_admissionregistration_ExternalAdmissionHook_To_v1alpha1_ExternalAdmissionHook(in *admissionregistration.ExternalAdmissionHook, out *ExternalAdmissionHook, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_admissionregistration_AdmissionHookClientConfig_To_v1alpha1_AdmissionHookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
		return err
	}
	out.Rules = *(*[]RuleWithOperations)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	return nil
}

// Convert_admissionregistration_ExternalAdmissionHook_To_v1alpha1_ExternalAdmissionHook is an autogenerated conversion function.
func Convert_admissionregistration_ExternalAdmissionHook_To_v1alpha1_ExternalAdmissionHook(in *admissionregistration.ExternalAdmissionHook, out *ExternalAdmissionHook, s conversion.Scope) error {
	return autoConvert_admissionregistration_ExternalAdmissionHook_To_v1alpha1_ExternalAdmissionHook(in, out, s)
}

func autoConvert_v1alpha1_ExternalAdmissionHookConfiguration_To_admissionregistration_ExternalAdmissionHookConfiguration(in *ExternalAdmissionHookConfiguration, out *admissionregistration.ExternalAdmissionHookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.ExternalAdmissionHooks = *(*[]admissionregistration.ExternalAdmissionHook)(unsafe.Pointer(&in.ExternalAdmissionHooks))
	return nil
}

// Convert_v1alpha1_ExternalAdmissionHookConfiguration_To_admissionregistration_ExternalAdmissionHookConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ExternalAdmissionHookConfiguration_To_admissionregistration_ExternalAdmissionHookConfiguration(in *ExternalAdmissionHookConfiguration, out *admissionregistration.ExternalAdmissionHookConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ExternalAdmissionHookConfiguration_To_admissionregistration_ExternalAdmissionHookConfiguration(in, out, s)
}

func autoConvert_admissionregistration_ExternalAdmissionHookConfiguration_To_v1alpha1_ExternalAdmissionHookConfiguration(in *admissionregistration.ExternalAdmissionHookConfiguration, out *ExternalAdmissionHookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.ExternalAdmissionHooks = *(*[]ExternalAdmissionHook)(unsafe.Pointer(&in.ExternalAdmissionHooks))
	return nil
}

// Convert_admissionregistration_ExternalAdmissionHookConfiguration_To_v1alpha1_ExternalAdmissionHookConfiguration is an autogenerated conversion function.
func Convert_admissionregistration_ExternalAdmissionHookConfiguration_To_v1alpha1_ExternalAdmissionHookConfiguration(in *admissionregistration.ExternalAdmissionHookConfiguration, out *ExternalAdmissionHookConfiguration, s conversion.Scope) error {
	return autoConvert_admissionregistration_ExternalAdmissionHookConfiguration_To_v1alpha1_ExternalAdmissionHookConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ExternalAdmissionHookConfigurationList_To_admissionregistration_ExternalAdmissionHookConfigurationList(in *ExternalAdmissionHookConfigurationList, out *admissionregistration.ExternalAdmissionHookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]admissionregistration.ExternalAdmissionHookConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ExternalAdmissionHookConfigurationList_To_admissionregistration_ExternalAdmissionHookConfigurationList is an autogenerated conversion function.
func Convert_v1alpha1_ExternalAdmissionHookConfigurationList_To_admissionregistration_ExternalAdmissionHookConfigurationList(in *ExternalAdmissionHookConfigurationList, out *admissionregistration.ExternalAdmissionHookConfigurationList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ExternalAdmissionHookConfigurationList_To_admissionregistration_ExternalAdmissionHookConfigurationList(in, out, s)
}

func autoConvert_admissionregistration_ExternalAdmissionHookConfigurationList_To_v1alpha1_ExternalAdmissionHookConfigurationList(in *admissionregistration.ExternalAdmissionHookConfigurationList, out *ExternalAdmissionHookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]ExternalAdmissionHookConfiguration, 0)
	} else {
		out.Items = *(*[]ExternalAdmissionHookConfiguration)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_admissionregistration_ExternalAdmissionHookConfigurationList_To_v1alpha1_ExternalAdmissionHookConfigurationList is an autogenerated conversion function.
func Convert_admissionregistration_ExternalAdmissionHookConfigurationList_To_v1alpha1_ExternalAdmissionHookConfigurationList(in *admissionregistration.ExternalAdmissionHookConfigurationList, out *ExternalAdmissionHookConfigurationList, s conversion.Scope) error {
	return autoConvert_admissionregistration_ExternalAdmissionHookConfigurationList_To_v1alpha1_ExternalAdmissionHookConfigurationList(in, out, s)
}

func autoConvert_v1alpha1_Initializer_To_admissionregistration_Initializer(in *Initializer, out *admissionregistration.Initializer, s conversion.Scope) error {
	out.Name = in.Name
	out.Rules = *(*[]admissionregistration.Rule)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*admissionregistration.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	return nil
}

// Convert_v1alpha1_Initializer_To_admissionregistration_Initializer is an autogenerated conversion function.
func Convert_v1alpha1_Initializer_To_admissionregistration_Initializer(in *Initializer, out *admissionregistration.Initializer, s conversion.Scope) error {
	return autoConvert_v1alpha1_Initializer_To_admissionregistration_Initializer(in, out, s)
}

func autoConvert_admissionregistration_Initializer_To_v1alpha1_Initializer(in *admissionregistration.Initializer, out *Initializer, s conversion.Scope) error {
	out.Name = in.Name
	out.Rules = *(*[]Rule)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	return nil
}

// Convert_admissionregistration_Initializer_To_v1alpha1_Initializer is an autogenerated conversion function.
func Convert_admissionregistration_Initializer_To_v1alpha1_Initializer(in *admissionregistration.Initializer, out *Initializer, s conversion.Scope) error {
	return autoConvert_admissionregistration_Initializer_To_v1alpha1_Initializer(in, out, s)
}

func autoConvert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(in *InitializerConfiguration, out *admissionregistration.InitializerConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Initializers = *(*[]admissionregistration.Initializer)(unsafe.Pointer(&in.Initializers))
	return nil
}

// Convert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(in *InitializerConfiguration, out *admissionregistration.InitializerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(in, out, s)
}

func autoConvert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(in *admissionregistration.InitializerConfiguration, out *InitializerConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Initializers = *(*[]Initializer)(unsafe.Pointer(&in.Initializers))
	return nil
}

// Convert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration is an autogenerated conversion function.
func Convert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(in *admissionregistration.InitializerConfiguration, out *InitializerConfiguration, s conversion.Scope) error {
	return autoConvert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(in *InitializerConfigurationList, out *admissionregistration.InitializerConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]admissionregistration.InitializerConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList is an autogenerated conversion function.
func Convert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(in *InitializerConfigurationList, out *admissionregistration.InitializerConfigurationList, s conversion.Scope) error {
	return autoConvert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(in, out, s)
}

func autoConvert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(in *admissionregistration.InitializerConfigurationList, out *InitializerConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]InitializerConfiguration, 0)
	} else {
		out.Items = *(*[]InitializerConfiguration)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList is an autogenerated conversion function.
func Convert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(in *admissionregistration.InitializerConfigurationList, out *InitializerConfigurationList, s conversion.Scope) error {
	return autoConvert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(in, out, s)
}

func autoConvert_v1alpha1_Rule_To_admissionregistration_Rule(in *Rule, out *admissionregistration.Rule, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1alpha1_Rule_To_admissionregistration_Rule is an autogenerated conversion function.
func Convert_v1alpha1_Rule_To_admissionregistration_Rule(in *Rule, out *admissionregistration.Rule, s conversion.Scope) error {
	return autoConvert_v1alpha1_Rule_To_admissionregistration_Rule(in, out, s)
}

func autoConvert_admissionregistration_Rule_To_v1alpha1_Rule(in *admissionregistration.Rule, out *Rule, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_admissionregistration_Rule_To_v1alpha1_Rule is an autogenerated conversion function.
func Convert_admissionregistration_Rule_To_v1alpha1_Rule(in *admissionregistration.Rule, out *Rule, s conversion.Scope) error {
	return autoConvert_admissionregistration_Rule_To_v1alpha1_Rule(in, out, s)
}

func autoConvert_v1alpha1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in *RuleWithOperations, out *admissionregistration.RuleWithOperations, s conversion.Scope) error {
	out.Operations = *(*[]admissionregistration.OperationType)(unsafe.Pointer(&in.Operations))
	if err := Convert_v1alpha1_Rule_To_admissionregistration_Rule(&in.Rule, &out.Rule, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_RuleWithOperations_To_admissionregistration_RuleWithOperations is an autogenerated conversion function.
func Convert_v1alpha1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in *RuleWithOperations, out *admissionregistration.RuleWithOperations, s conversion.Scope) error {
	return autoConvert_v1alpha1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in, out, s)
}

func autoConvert_admissionregistration_RuleWithOperations_To_v1alpha1_RuleWithOperations(in *admissionregistration.RuleWithOperations, out *RuleWithOperations, s conversion.Scope) error {
	out.Operations = *(*[]OperationType)(unsafe.Pointer(&in.Operations))
	if err := Convert_admissionregistration_Rule_To_v1alpha1_Rule(&in.Rule, &out.Rule, s); err != nil {
		return err
	}
	return nil
}

// Convert_admissionregistration_RuleWithOperations_To_v1alpha1_RuleWithOperations is an autogenerated conversion function.
func Convert_admissionregistration_RuleWithOperations_To_v1alpha1_RuleWithOperations(in *admissionregistration.RuleWithOperations, out *RuleWithOperations, s conversion.Scope) error {
	return autoConvert_admissionregistration_RuleWithOperations_To_v1alpha1_RuleWithOperations(in, out, s)
}

func autoConvert_v1alpha1_ServiceReference_To_admissionregistration_ServiceReference(in *ServiceReference, out *admissionregistration.ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_ServiceReference_To_admissionregistration_ServiceReference is an autogenerated conversion function.
func Convert_v1alpha1_ServiceReference_To_admissionregistration_ServiceReference(in *ServiceReference, out *admissionregistration.ServiceReference, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceReference_To_admissionregistration_ServiceReference(in, out, s)
}

func autoConvert_admissionregistration_ServiceReference_To_v1alpha1_ServiceReference(in *admissionregistration.ServiceReference, out *ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_admissionregistration_ServiceReference_To_v1alpha1_ServiceReference is an autogenerated conversion function.
func Convert_admissionregistration_ServiceReference_To_v1alpha1_ServiceReference(in *admissionregistration.ServiceReference, out *ServiceReference, s conversion.Scope) error {
	return autoConvert_admissionregistration_ServiceReference_To_v1alpha1_ServiceReference(in, out, s)
}
