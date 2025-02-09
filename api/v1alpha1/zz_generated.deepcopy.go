//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
	apiv1alpha1 "github.com/openshift/hypershift/api/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSCredentials) DeepCopyInto(out *AWSCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSCredentials.
func (in *AWSCredentials) DeepCopy() *AWSCredentials {
	if in == nil {
		return nil
	}
	out := new(AWSCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSPlatform) DeepCopyInto(out *AWSPlatform) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSPlatform.
func (in *AWSPlatform) DeepCopy() *AWSPlatform {
	if in == nil {
		return nil
	}
	out := new(AWSPlatform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzurePlatform) DeepCopyInto(out *AzurePlatform) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzurePlatform.
func (in *AzurePlatform) DeepCopy() *AzurePlatform {
	if in == nil {
		return nil
	}
	out := new(AzurePlatform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialARNs) DeepCopyInto(out *CredentialARNs) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSCredentials)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialARNs.
func (in *CredentialARNs) DeepCopy() *CredentialARNs {
	if in == nil {
		return nil
	}
	out := new(CredentialARNs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypershiftDeployment) DeepCopyInto(out *HypershiftDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypershiftDeployment.
func (in *HypershiftDeployment) DeepCopy() *HypershiftDeployment {
	if in == nil {
		return nil
	}
	out := new(HypershiftDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HypershiftDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypershiftDeploymentList) DeepCopyInto(out *HypershiftDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HypershiftDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypershiftDeploymentList.
func (in *HypershiftDeploymentList) DeepCopy() *HypershiftDeploymentList {
	if in == nil {
		return nil
	}
	out := new(HypershiftDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HypershiftDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypershiftDeploymentSpec) DeepCopyInto(out *HypershiftDeploymentSpec) {
	*out = *in
	in.Infrastructure.DeepCopyInto(&out.Infrastructure)
	if in.HostedClusterSpec != nil {
		in, out := &in.HostedClusterSpec, &out.HostedClusterSpec
		*out = new(apiv1alpha1.HostedClusterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]*HypershiftNodePools, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HypershiftNodePools)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(CredentialARNs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypershiftDeploymentSpec.
func (in *HypershiftDeploymentSpec) DeepCopy() *HypershiftDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(HypershiftDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypershiftDeploymentStatus) DeepCopyInto(out *HypershiftDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypershiftDeploymentStatus.
func (in *HypershiftDeploymentStatus) DeepCopy() *HypershiftDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(HypershiftDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypershiftNodePools) DeepCopyInto(out *HypershiftNodePools) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypershiftNodePools.
func (in *HypershiftNodePools) DeepCopy() *HypershiftNodePools {
	if in == nil {
		return nil
	}
	out := new(HypershiftNodePools)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraSpec) DeepCopyInto(out *InfraSpec) {
	*out = *in
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(Platforms)
		(*in).DeepCopyInto(*out)
	}
	out.CloudProvider = in.CloudProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraSpec.
func (in *InfraSpec) DeepCopy() *InfraSpec {
	if in == nil {
		return nil
	}
	out := new(InfraSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platforms) DeepCopyInto(out *Platforms) {
	*out = *in
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzurePlatform)
		**out = **in
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSPlatform)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platforms.
func (in *Platforms) DeepCopy() *Platforms {
	if in == nil {
		return nil
	}
	out := new(Platforms)
	in.DeepCopyInto(out)
	return out
}
