//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dashboard) DeepCopyInto(out *Dashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dashboard.
func (in *Dashboard) DeepCopy() *Dashboard {
	if in == nil {
		return nil
	}
	out := new(Dashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardBillboardWidgetThresholdInput) DeepCopyInto(out *DashboardBillboardWidgetThresholdInput) {
	*out = *in
	if in.AlertSeverity != nil {
		in, out := &in.AlertSeverity, &out.AlertSeverity
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardBillboardWidgetThresholdInput.
func (in *DashboardBillboardWidgetThresholdInput) DeepCopy() *DashboardBillboardWidgetThresholdInput {
	if in == nil {
		return nil
	}
	out := new(DashboardBillboardWidgetThresholdInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardList) DeepCopyInto(out *DashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardList.
func (in *DashboardList) DeepCopy() *DashboardList {
	if in == nil {
		return nil
	}
	out := new(DashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardObservation) DeepCopyInto(out *DashboardObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardObservation.
func (in *DashboardObservation) DeepCopy() *DashboardObservation {
	if in == nil {
		return nil
	}
	out := new(DashboardObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardPage) DeepCopyInto(out *DashboardPage) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Widgets != nil {
		in, out := &in.Widgets, &out.Widgets
		*out = make([]DashboardWidget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardPage.
func (in *DashboardPage) DeepCopy() *DashboardPage {
	if in == nil {
		return nil
	}
	out := new(DashboardPage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardParameters) DeepCopyInto(out *DashboardParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Pages != nil {
		in, out := &in.Pages, &out.Pages
		*out = make([]DashboardPage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make([]DashboardVariable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardParameters.
func (in *DashboardParameters) DeepCopy() *DashboardParameters {
	if in == nil {
		return nil
	}
	out := new(DashboardParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpec) DeepCopyInto(out *DashboardSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpec.
func (in *DashboardSpec) DeepCopy() *DashboardSpec {
	if in == nil {
		return nil
	}
	out := new(DashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardStatus) DeepCopyInto(out *DashboardStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardStatus.
func (in *DashboardStatus) DeepCopy() *DashboardStatus {
	if in == nil {
		return nil
	}
	out := new(DashboardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardVariable) DeepCopyInto(out *DashboardVariable) {
	*out = *in
	if in.DefaultValues != nil {
		in, out := &in.DefaultValues, &out.DefaultValues
		*out = new([]DashboardVariableDefaultItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]DashboardVariableDefaultItem, len(*in))
			copy(*out, *in)
		}
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DashboardVariableEnumItem, len(*in))
		copy(*out, *in)
	}
	if in.NRQLQuery != nil {
		in, out := &in.NRQLQuery, &out.NRQLQuery
		*out = new(DashboardVariableNRQLQuery)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardVariable.
func (in *DashboardVariable) DeepCopy() *DashboardVariable {
	if in == nil {
		return nil
	}
	out := new(DashboardVariable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardVariableDefaultItem) DeepCopyInto(out *DashboardVariableDefaultItem) {
	*out = *in
	out.Value = in.Value
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardVariableDefaultItem.
func (in *DashboardVariableDefaultItem) DeepCopy() *DashboardVariableDefaultItem {
	if in == nil {
		return nil
	}
	out := new(DashboardVariableDefaultItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardVariableDefaultValue) DeepCopyInto(out *DashboardVariableDefaultValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardVariableDefaultValue.
func (in *DashboardVariableDefaultValue) DeepCopy() *DashboardVariableDefaultValue {
	if in == nil {
		return nil
	}
	out := new(DashboardVariableDefaultValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardVariableEnumItem) DeepCopyInto(out *DashboardVariableEnumItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardVariableEnumItem.
func (in *DashboardVariableEnumItem) DeepCopy() *DashboardVariableEnumItem {
	if in == nil {
		return nil
	}
	out := new(DashboardVariableEnumItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardVariableNRQLQuery) DeepCopyInto(out *DashboardVariableNRQLQuery) {
	*out = *in
	if in.AccountIDs != nil {
		in, out := &in.AccountIDs, &out.AccountIDs
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardVariableNRQLQuery.
func (in *DashboardVariableNRQLQuery) DeepCopy() *DashboardVariableNRQLQuery {
	if in == nil {
		return nil
	}
	out := new(DashboardVariableNRQLQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardWidget) DeepCopyInto(out *DashboardWidget) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	out.Layout = in.Layout
	if in.RawConfiguration != nil {
		in, out := &in.RawConfiguration, &out.RawConfiguration
		*out = new(DashboardWidgetRawConfiguration)
		(*in).DeepCopyInto(*out)
	}
	out.Visualization = in.Visualization
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardWidget.
func (in *DashboardWidget) DeepCopy() *DashboardWidget {
	if in == nil {
		return nil
	}
	out := new(DashboardWidget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardWidgetLayout) DeepCopyInto(out *DashboardWidgetLayout) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardWidgetLayout.
func (in *DashboardWidgetLayout) DeepCopy() *DashboardWidgetLayout {
	if in == nil {
		return nil
	}
	out := new(DashboardWidgetLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardWidgetNRQLQueryInput) DeepCopyInto(out *DashboardWidgetNRQLQueryInput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardWidgetNRQLQueryInput.
func (in *DashboardWidgetNRQLQueryInput) DeepCopy() *DashboardWidgetNRQLQueryInput {
	if in == nil {
		return nil
	}
	out := new(DashboardWidgetNRQLQueryInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardWidgetRawConfiguration) DeepCopyInto(out *DashboardWidgetRawConfiguration) {
	*out = *in
	if in.NRQLQueries != nil {
		in, out := &in.NRQLQueries, &out.NRQLQueries
		*out = new([]DashboardWidgetNRQLQueryInput)
		if **in != nil {
			in, out := *in, *out
			*out = make([]DashboardWidgetNRQLQueryInput, len(*in))
			copy(*out, *in)
		}
	}
	if in.PlatformOptions != nil {
		in, out := &in.PlatformOptions, &out.PlatformOptions
		*out = new(RawConfigurationPlatformOptions)
		**out = **in
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = new(float64)
		**out = **in
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = new(string)
		**out = **in
	}
	if in.Thresholds != nil {
		in, out := &in.Thresholds, &out.Thresholds
		*out = make([]DashboardBillboardWidgetThresholdInput, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardWidgetRawConfiguration.
func (in *DashboardWidgetRawConfiguration) DeepCopy() *DashboardWidgetRawConfiguration {
	if in == nil {
		return nil
	}
	out := new(DashboardWidgetRawConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardWidgetVisualization) DeepCopyInto(out *DashboardWidgetVisualization) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardWidgetVisualization.
func (in *DashboardWidgetVisualization) DeepCopy() *DashboardWidgetVisualization {
	if in == nil {
		return nil
	}
	out := new(DashboardWidgetVisualization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawConfigurationPlatformOptions) DeepCopyInto(out *RawConfigurationPlatformOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawConfigurationPlatformOptions.
func (in *RawConfigurationPlatformOptions) DeepCopy() *RawConfigurationPlatformOptions {
	if in == nil {
		return nil
	}
	out := new(RawConfigurationPlatformOptions)
	in.DeepCopyInto(out)
	return out
}
