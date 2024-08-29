//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v2beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BondConfig) DeepCopyInto(out *BondConfig) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BondConfig.
func (in *BondConfig) DeepCopy() *BondConfig {
	if in == nil {
		return nil
	}
	out := new(BondConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoordinatorSpec) DeepCopyInto(out *CoordinatorSpec) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.PodCIDRType != nil {
		in, out := &in.PodCIDRType, &out.PodCIDRType
		*out = new(string)
		**out = **in
	}
	if in.HijackCIDR != nil {
		in, out := &in.HijackCIDR, &out.HijackCIDR
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodMACPrefix != nil {
		in, out := &in.PodMACPrefix, &out.PodMACPrefix
		*out = new(string)
		**out = **in
	}
	if in.TunePodRoutes != nil {
		in, out := &in.TunePodRoutes, &out.TunePodRoutes
		*out = new(bool)
		**out = **in
	}
	if in.PodDefaultRouteNIC != nil {
		in, out := &in.PodDefaultRouteNIC, &out.PodDefaultRouteNIC
		*out = new(string)
		**out = **in
	}
	if in.HostRuleTable != nil {
		in, out := &in.HostRuleTable, &out.HostRuleTable
		*out = new(int)
		**out = **in
	}
	if in.HostRPFilter != nil {
		in, out := &in.HostRPFilter, &out.HostRPFilter
		*out = new(int)
		**out = **in
	}
	if in.PodRPFilter != nil {
		in, out := &in.PodRPFilter, &out.PodRPFilter
		*out = new(int)
		**out = **in
	}
	if in.TxQueueLen != nil {
		in, out := &in.TxQueueLen, &out.TxQueueLen
		*out = new(int)
		**out = **in
	}
	if in.DetectIPConflict != nil {
		in, out := &in.DetectIPConflict, &out.DetectIPConflict
		*out = new(bool)
		**out = **in
	}
	if in.DetectGateway != nil {
		in, out := &in.DetectGateway, &out.DetectGateway
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoordinatorSpec.
func (in *CoordinatorSpec) DeepCopy() *CoordinatorSpec {
	if in == nil {
		return nil
	}
	out := new(CoordinatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoordinatorStatus) DeepCopyInto(out *CoordinatorStatus) {
	*out = *in
	if in.OverlayPodCIDR != nil {
		in, out := &in.OverlayPodCIDR, &out.OverlayPodCIDR
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceCIDR != nil {
		in, out := &in.ServiceCIDR, &out.ServiceCIDR
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoordinatorStatus.
func (in *CoordinatorStatus) DeepCopy() *CoordinatorStatus {
	if in == nil {
		return nil
	}
	out := new(CoordinatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAllocationDetail) DeepCopyInto(out *IPAllocationDetail) {
	*out = *in
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(string)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(string)
		**out = **in
	}
	if in.IPv4Pool != nil {
		in, out := &in.IPv4Pool, &out.IPv4Pool
		*out = new(string)
		**out = **in
	}
	if in.IPv6Pool != nil {
		in, out := &in.IPv6Pool, &out.IPv6Pool
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(int64)
		**out = **in
	}
	if in.IPv4Gateway != nil {
		in, out := &in.IPv4Gateway, &out.IPv4Gateway
		*out = new(string)
		**out = **in
	}
	if in.IPv6Gateway != nil {
		in, out := &in.IPv6Gateway, &out.IPv6Gateway
		*out = new(string)
		**out = **in
	}
	if in.CleanGateway != nil {
		in, out := &in.CleanGateway, &out.CleanGateway
		*out = new(bool)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAllocationDetail.
func (in *IPAllocationDetail) DeepCopy() *IPAllocationDetail {
	if in == nil {
		return nil
	}
	out := new(IPAllocationDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolSpec) DeepCopyInto(out *IPPoolSpec) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(int64)
		**out = **in
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeIPs != nil {
		in, out := &in.ExcludeIPs, &out.ExcludeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(int64)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
	if in.PodAffinity != nil {
		in, out := &in.PodAffinity, &out.PodAffinity
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceAffinity != nil {
		in, out := &in.NamespaceAffinity, &out.NamespaceAffinity
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MultusName != nil {
		in, out := &in.MultusName, &out.MultusName
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.Disable != nil {
		in, out := &in.Disable, &out.Disable
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolSpec.
func (in *IPPoolSpec) DeepCopy() *IPPoolSpec {
	if in == nil {
		return nil
	}
	out := new(IPPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolStatus) DeepCopyInto(out *IPPoolStatus) {
	*out = *in
	if in.AllocatedIPs != nil {
		in, out := &in.AllocatedIPs, &out.AllocatedIPs
		*out = new(string)
		**out = **in
	}
	if in.TotalIPCount != nil {
		in, out := &in.TotalIPCount, &out.TotalIPCount
		*out = new(int64)
		**out = **in
	}
	if in.AllocatedIPCount != nil {
		in, out := &in.AllocatedIPCount, &out.AllocatedIPCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolStatus.
func (in *IPPoolStatus) DeepCopy() *IPPoolStatus {
	if in == nil {
		return nil
	}
	out := new(IPPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultusCNIConfigSpec) DeepCopyInto(out *MultusCNIConfigSpec) {
	*out = *in
	if in.CniType != nil {
		in, out := &in.CniType, &out.CniType
		*out = new(string)
		**out = **in
	}
	if in.MacvlanConfig != nil {
		in, out := &in.MacvlanConfig, &out.MacvlanConfig
		*out = new(SpiderMacvlanCniConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IPVlanConfig != nil {
		in, out := &in.IPVlanConfig, &out.IPVlanConfig
		*out = new(SpiderIPvlanCniConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SriovConfig != nil {
		in, out := &in.SriovConfig, &out.SriovConfig
		*out = new(SpiderSRIOVCniConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.OvsConfig != nil {
		in, out := &in.OvsConfig, &out.OvsConfig
		*out = new(SpiderOvsCniConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IbSriovConfig != nil {
		in, out := &in.IbSriovConfig, &out.IbSriovConfig
		*out = new(SpiderIBSriovCniConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IpoibConfig != nil {
		in, out := &in.IpoibConfig, &out.IpoibConfig
		*out = new(SpiderIpoibCniConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EnableCoordinator != nil {
		in, out := &in.EnableCoordinator, &out.EnableCoordinator
		*out = new(bool)
		**out = **in
	}
	if in.DisableIPAM != nil {
		in, out := &in.DisableIPAM, &out.DisableIPAM
		*out = new(bool)
		**out = **in
	}
	if in.CoordinatorConfig != nil {
		in, out := &in.CoordinatorConfig, &out.CoordinatorConfig
		*out = new(CoordinatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ChainCNIJsonData != nil {
		in, out := &in.ChainCNIJsonData, &out.ChainCNIJsonData
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomCNIConfig != nil {
		in, out := &in.CustomCNIConfig, &out.CustomCNIConfig
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultusCNIConfigSpec.
func (in *MultusCNIConfigSpec) DeepCopy() *MultusCNIConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MultusCNIConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodIPAllocation) DeepCopyInto(out *PodIPAllocation) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]IPAllocationDetail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodIPAllocation.
func (in *PodIPAllocation) DeepCopy() *PodIPAllocation {
	if in == nil {
		return nil
	}
	out := new(PodIPAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolIPAllocation) DeepCopyInto(out *PoolIPAllocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPAllocation.
func (in *PoolIPAllocation) DeepCopy() *PoolIPAllocation {
	if in == nil {
		return nil
	}
	out := new(PoolIPAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PoolIPAllocations) DeepCopyInto(out *PoolIPAllocations) {
	{
		in := &in
		*out = make(PoolIPAllocations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPAllocations.
func (in PoolIPAllocations) DeepCopy() PoolIPAllocations {
	if in == nil {
		return nil
	}
	out := new(PoolIPAllocations)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolIPPreAllocation) DeepCopyInto(out *PoolIPPreAllocation) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPPreAllocation.
func (in *PoolIPPreAllocation) DeepCopy() *PoolIPPreAllocation {
	if in == nil {
		return nil
	}
	out := new(PoolIPPreAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PoolIPPreAllocations) DeepCopyInto(out *PoolIPPreAllocations) {
	{
		in := &in
		*out = make(PoolIPPreAllocations, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPPreAllocations.
func (in PoolIPPreAllocations) DeepCopy() PoolIPPreAllocations {
	if in == nil {
		return nil
	}
	out := new(PoolIPPreAllocations)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservedIPSpec) DeepCopyInto(out *ReservedIPSpec) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(int64)
		**out = **in
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservedIPSpec.
func (in *ReservedIPSpec) DeepCopy() *ReservedIPSpec {
	if in == nil {
		return nil
	}
	out := new(ReservedIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderCoordinator) DeepCopyInto(out *SpiderCoordinator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderCoordinator.
func (in *SpiderCoordinator) DeepCopy() *SpiderCoordinator {
	if in == nil {
		return nil
	}
	out := new(SpiderCoordinator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderCoordinator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderCoordinatorList) DeepCopyInto(out *SpiderCoordinatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderCoordinator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderCoordinatorList.
func (in *SpiderCoordinatorList) DeepCopy() *SpiderCoordinatorList {
	if in == nil {
		return nil
	}
	out := new(SpiderCoordinatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderCoordinatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderEndpoint) DeepCopyInto(out *SpiderEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderEndpoint.
func (in *SpiderEndpoint) DeepCopy() *SpiderEndpoint {
	if in == nil {
		return nil
	}
	out := new(SpiderEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderEndpointList) DeepCopyInto(out *SpiderEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderEndpointList.
func (in *SpiderEndpointList) DeepCopy() *SpiderEndpointList {
	if in == nil {
		return nil
	}
	out := new(SpiderEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderIBSriovCniConfig) DeepCopyInto(out *SpiderIBSriovCniConfig) {
	*out = *in
	if in.Pkey != nil {
		in, out := &in.Pkey, &out.Pkey
		*out = new(string)
		**out = **in
	}
	if in.LinkState != nil {
		in, out := &in.LinkState, &out.LinkState
		*out = new(string)
		**out = **in
	}
	if in.RdmaIsolation != nil {
		in, out := &in.RdmaIsolation, &out.RdmaIsolation
		*out = new(bool)
		**out = **in
	}
	if in.IbKubernetesEnabled != nil {
		in, out := &in.IbKubernetesEnabled, &out.IbKubernetesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SpiderpoolConfigPools != nil {
		in, out := &in.SpiderpoolConfigPools, &out.SpiderpoolConfigPools
		*out = new(SpiderpoolPools)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderIBSriovCniConfig.
func (in *SpiderIBSriovCniConfig) DeepCopy() *SpiderIBSriovCniConfig {
	if in == nil {
		return nil
	}
	out := new(SpiderIBSriovCniConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderIPPool) DeepCopyInto(out *SpiderIPPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderIPPool.
func (in *SpiderIPPool) DeepCopy() *SpiderIPPool {
	if in == nil {
		return nil
	}
	out := new(SpiderIPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderIPPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderIPPoolList) DeepCopyInto(out *SpiderIPPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderIPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderIPPoolList.
func (in *SpiderIPPoolList) DeepCopy() *SpiderIPPoolList {
	if in == nil {
		return nil
	}
	out := new(SpiderIPPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderIPPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderIPvlanCniConfig) DeepCopyInto(out *SpiderIPvlanCniConfig) {
	*out = *in
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(int32)
		**out = **in
	}
	if in.Bond != nil {
		in, out := &in.Bond, &out.Bond
		*out = new(BondConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SpiderpoolConfigPools != nil {
		in, out := &in.SpiderpoolConfigPools, &out.SpiderpoolConfigPools
		*out = new(SpiderpoolPools)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderIPvlanCniConfig.
func (in *SpiderIPvlanCniConfig) DeepCopy() *SpiderIPvlanCniConfig {
	if in == nil {
		return nil
	}
	out := new(SpiderIPvlanCniConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderIpoibCniConfig) DeepCopyInto(out *SpiderIpoibCniConfig) {
	*out = *in
	if in.SpiderpoolConfigPools != nil {
		in, out := &in.SpiderpoolConfigPools, &out.SpiderpoolConfigPools
		*out = new(SpiderpoolPools)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderIpoibCniConfig.
func (in *SpiderIpoibCniConfig) DeepCopy() *SpiderIpoibCniConfig {
	if in == nil {
		return nil
	}
	out := new(SpiderIpoibCniConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderMacvlanCniConfig) DeepCopyInto(out *SpiderMacvlanCniConfig) {
	*out = *in
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(int32)
		**out = **in
	}
	if in.Bond != nil {
		in, out := &in.Bond, &out.Bond
		*out = new(BondConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SpiderpoolConfigPools != nil {
		in, out := &in.SpiderpoolConfigPools, &out.SpiderpoolConfigPools
		*out = new(SpiderpoolPools)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderMacvlanCniConfig.
func (in *SpiderMacvlanCniConfig) DeepCopy() *SpiderMacvlanCniConfig {
	if in == nil {
		return nil
	}
	out := new(SpiderMacvlanCniConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderMultusConfig) DeepCopyInto(out *SpiderMultusConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderMultusConfig.
func (in *SpiderMultusConfig) DeepCopy() *SpiderMultusConfig {
	if in == nil {
		return nil
	}
	out := new(SpiderMultusConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderMultusConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderMultusConfigList) DeepCopyInto(out *SpiderMultusConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderMultusConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderMultusConfigList.
func (in *SpiderMultusConfigList) DeepCopy() *SpiderMultusConfigList {
	if in == nil {
		return nil
	}
	out := new(SpiderMultusConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderMultusConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderOvsCniConfig) DeepCopyInto(out *SpiderOvsCniConfig) {
	*out = *in
	if in.VlanTag != nil {
		in, out := &in.VlanTag, &out.VlanTag
		*out = new(int32)
		**out = **in
	}
	if in.Trunk != nil {
		in, out := &in.Trunk, &out.Trunk
		*out = make([]*Trunk, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Trunk)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.SpiderpoolConfigPools != nil {
		in, out := &in.SpiderpoolConfigPools, &out.SpiderpoolConfigPools
		*out = new(SpiderpoolPools)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderOvsCniConfig.
func (in *SpiderOvsCniConfig) DeepCopy() *SpiderOvsCniConfig {
	if in == nil {
		return nil
	}
	out := new(SpiderOvsCniConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderReservedIP) DeepCopyInto(out *SpiderReservedIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderReservedIP.
func (in *SpiderReservedIP) DeepCopy() *SpiderReservedIP {
	if in == nil {
		return nil
	}
	out := new(SpiderReservedIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderReservedIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderReservedIPList) DeepCopyInto(out *SpiderReservedIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderReservedIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderReservedIPList.
func (in *SpiderReservedIPList) DeepCopy() *SpiderReservedIPList {
	if in == nil {
		return nil
	}
	out := new(SpiderReservedIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderReservedIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderSRIOVCniConfig) DeepCopyInto(out *SpiderSRIOVCniConfig) {
	*out = *in
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(int32)
		**out = **in
	}
	if in.MinTxRateMbps != nil {
		in, out := &in.MinTxRateMbps, &out.MinTxRateMbps
		*out = new(int)
		**out = **in
	}
	if in.MaxTxRateMbps != nil {
		in, out := &in.MaxTxRateMbps, &out.MaxTxRateMbps
		*out = new(int)
		**out = **in
	}
	if in.SpiderpoolConfigPools != nil {
		in, out := &in.SpiderpoolConfigPools, &out.SpiderpoolConfigPools
		*out = new(SpiderpoolPools)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderSRIOVCniConfig.
func (in *SpiderSRIOVCniConfig) DeepCopy() *SpiderSRIOVCniConfig {
	if in == nil {
		return nil
	}
	out := new(SpiderSRIOVCniConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderSubnet) DeepCopyInto(out *SpiderSubnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderSubnet.
func (in *SpiderSubnet) DeepCopy() *SpiderSubnet {
	if in == nil {
		return nil
	}
	out := new(SpiderSubnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderSubnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderSubnetList) DeepCopyInto(out *SpiderSubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderSubnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderSubnetList.
func (in *SpiderSubnetList) DeepCopy() *SpiderSubnetList {
	if in == nil {
		return nil
	}
	out := new(SpiderSubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderSubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderpoolPools) DeepCopyInto(out *SpiderpoolPools) {
	*out = *in
	if in.IPv4IPPool != nil {
		in, out := &in.IPv4IPPool, &out.IPv4IPPool
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPv6IPPool != nil {
		in, out := &in.IPv6IPPool, &out.IPv6IPPool
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderpoolPools.
func (in *SpiderpoolPools) DeepCopy() *SpiderpoolPools {
	if in == nil {
		return nil
	}
	out := new(SpiderpoolPools)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(int64)
		**out = **in
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeIPs != nil {
		in, out := &in.ExcludeIPs, &out.ExcludeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(int64)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetStatus) DeepCopyInto(out *SubnetStatus) {
	*out = *in
	if in.ControlledIPPools != nil {
		in, out := &in.ControlledIPPools, &out.ControlledIPPools
		*out = new(string)
		**out = **in
	}
	if in.TotalIPCount != nil {
		in, out := &in.TotalIPCount, &out.TotalIPCount
		*out = new(int64)
		**out = **in
	}
	if in.AllocatedIPCount != nil {
		in, out := &in.AllocatedIPCount, &out.AllocatedIPCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetStatus.
func (in *SubnetStatus) DeepCopy() *SubnetStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trunk) DeepCopyInto(out *Trunk) {
	*out = *in
	if in.MinID != nil {
		in, out := &in.MinID, &out.MinID
		*out = new(uint)
		**out = **in
	}
	if in.MaxID != nil {
		in, out := &in.MaxID, &out.MaxID
		*out = new(uint)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(uint)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trunk.
func (in *Trunk) DeepCopy() *Trunk {
	if in == nil {
		return nil
	}
	out := new(Trunk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadEndpointStatus) DeepCopyInto(out *WorkloadEndpointStatus) {
	*out = *in
	in.Current.DeepCopyInto(&out.Current)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEndpointStatus.
func (in *WorkloadEndpointStatus) DeepCopy() *WorkloadEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadEndpointStatus)
	in.DeepCopyInto(out)
	return out
}
