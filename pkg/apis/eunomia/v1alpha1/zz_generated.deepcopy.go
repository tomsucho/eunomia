// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitConfig) DeepCopyInto(out *GitConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitConfig.
func (in *GitConfig) DeepCopy() *GitConfig {
	if in == nil {
		return nil
	}
	out := new(GitConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsConfig) DeepCopyInto(out *GitOpsConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsConfig.
func (in *GitOpsConfig) DeepCopy() *GitOpsConfig {
	if in == nil {
		return nil
	}
	out := new(GitOpsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsConfigList) DeepCopyInto(out *GitOpsConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitOpsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsConfigList.
func (in *GitOpsConfigList) DeepCopy() *GitOpsConfigList {
	if in == nil {
		return nil
	}
	out := new(GitOpsConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitOpsConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsConfigSpec) DeepCopyInto(out *GitOpsConfigSpec) {
	*out = *in
	out.TemplateSource = in.TemplateSource
	out.ParameterSource = in.ParameterSource
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]GitOpsTrigger, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsConfigSpec.
func (in *GitOpsConfigSpec) DeepCopy() *GitOpsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GitOpsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsConfigStatus) DeepCopyInto(out *GitOpsConfigStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsConfigStatus.
func (in *GitOpsConfigStatus) DeepCopy() *GitOpsConfigStatus {
	if in == nil {
		return nil
	}
	out := new(GitOpsConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitOpsTrigger) DeepCopyInto(out *GitOpsTrigger) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitOpsTrigger.
func (in *GitOpsTrigger) DeepCopy() *GitOpsTrigger {
	if in == nil {
		return nil
	}
	out := new(GitOpsTrigger)
	in.DeepCopyInto(out)
	return out
}
