//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
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

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteDeleteHTTPHeader) DeepCopyInto(out *RouteDeleteHTTPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteDeleteHTTPHeader.
func (in *RouteDeleteHTTPHeader) DeepCopy() *RouteDeleteHTTPHeader {
	if in == nil {
		return nil
	}
	out := new(RouteDeleteHTTPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteDeleteHTTPHeaders) DeepCopyInto(out *RouteDeleteHTTPHeaders) {
	*out = *in
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = make([]RouteDeleteHTTPHeader, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteDeleteHTTPHeaders.
func (in *RouteDeleteHTTPHeaders) DeepCopy() *RouteDeleteHTTPHeaders {
	if in == nil {
		return nil
	}
	out := new(RouteDeleteHTTPHeaders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteIngress) DeepCopyInto(out *RouteIngress) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RouteIngressCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteIngress.
func (in *RouteIngress) DeepCopy() *RouteIngress {
	if in == nil {
		return nil
	}
	out := new(RouteIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteIngressCondition) DeepCopyInto(out *RouteIngressCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteIngressCondition.
func (in *RouteIngressCondition) DeepCopy() *RouteIngressCondition {
	if in == nil {
		return nil
	}
	out := new(RouteIngressCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteList) DeepCopyInto(out *RouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteList.
func (in *RouteList) DeepCopy() *RouteList {
	if in == nil {
		return nil
	}
	out := new(RouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePort) DeepCopyInto(out *RoutePort) {
	*out = *in
	out.TargetPort = in.TargetPort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePort.
func (in *RoutePort) DeepCopy() *RoutePort {
	if in == nil {
		return nil
	}
	out := new(RoutePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSetHTTPHeader) DeepCopyInto(out *RouteSetHTTPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSetHTTPHeader.
func (in *RouteSetHTTPHeader) DeepCopy() *RouteSetHTTPHeader {
	if in == nil {
		return nil
	}
	out := new(RouteSetHTTPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSetHTTPHeaders) DeepCopyInto(out *RouteSetHTTPHeaders) {
	*out = *in
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = make([]RouteSetHTTPHeader, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSetHTTPHeaders.
func (in *RouteSetHTTPHeaders) DeepCopy() *RouteSetHTTPHeaders {
	if in == nil {
		return nil
	}
	out := new(RouteSetHTTPHeaders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	in.To.DeepCopyInto(&out.To)
	if in.AlternateBackends != nil {
		in, out := &in.AlternateBackends, &out.AlternateBackends
		*out = make([]RouteTargetReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(RoutePort)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	in.HTTPSetHeaders.DeepCopyInto(&out.HTTPSetHeaders)
	in.HTTPDeleteHeaders.DeepCopyInto(&out.HTTPDeleteHeaders)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]RouteIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTargetReference) DeepCopyInto(out *RouteTargetReference) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTargetReference.
func (in *RouteTargetReference) DeepCopy() *RouteTargetReference {
	if in == nil {
		return nil
	}
	out := new(RouteTargetReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouterShard) DeepCopyInto(out *RouterShard) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouterShard.
func (in *RouterShard) DeepCopy() *RouterShard {
	if in == nil {
		return nil
	}
	out := new(RouterShard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}
