// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/extension/service.proto

package extension

import (
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v33 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PostRouteModifyRequest sends a Route that was generated by Envoy Gateway along with context information to an extension so that the Route can be modified
type PostRouteModifyRequest struct {
	state            protoimpl.MessageState     `protogen:"open.v1"`
	Route            *v3.Route                  `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	PostRouteContext *PostRouteExtensionContext `protobuf:"bytes,2,opt,name=post_route_context,json=postRouteContext,proto3" json:"post_route_context,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *PostRouteModifyRequest) Reset() {
	*x = PostRouteModifyRequest{}
	mi := &file_proto_extension_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostRouteModifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRouteModifyRequest) ProtoMessage() {}

func (x *PostRouteModifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRouteModifyRequest.ProtoReflect.Descriptor instead.
func (*PostRouteModifyRequest) Descriptor() ([]byte, []int) {
	return file_proto_extension_service_proto_rawDescGZIP(), []int{0}
}

func (x *PostRouteModifyRequest) GetRoute() *v3.Route {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *PostRouteModifyRequest) GetPostRouteContext() *PostRouteExtensionContext {
	if x != nil {
		return x.PostRouteContext
	}
	return nil
}

// PostRouteModifyResponse is the expected response from an extension and contains a modified version of the Route that was sent
// If an extension returns a nil Route then it will not be modified
type PostRouteModifyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Route         *v3.Route              `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Clusters      []*v31.Cluster         `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostRouteModifyResponse) Reset() {
	*x = PostRouteModifyResponse{}
	mi := &file_proto_extension_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostRouteModifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRouteModifyResponse) ProtoMessage() {}

func (x *PostRouteModifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRouteModifyResponse.ProtoReflect.Descriptor instead.
func (*PostRouteModifyResponse) Descriptor() ([]byte, []int) {
	return file_proto_extension_service_proto_rawDescGZIP(), []int{1}
}

func (x *PostRouteModifyResponse) GetRoute() *v3.Route {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *PostRouteModifyResponse) GetClusters() []*v31.Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

// PostVirtualHostModifyRequest sends a VirtualHost that was generated by Envoy Gateway along with context information to an extension so that the VirtualHost can be modified
type PostVirtualHostModifyRequest struct {
	state                  protoimpl.MessageState           `protogen:"open.v1"`
	VirtualHost            *v3.VirtualHost                  `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost,proto3" json:"virtual_host,omitempty"`
	PostVirtualHostContext *PostVirtualHostExtensionContext `protobuf:"bytes,2,opt,name=post_virtual_host_context,json=postVirtualHostContext,proto3" json:"post_virtual_host_context,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *PostVirtualHostModifyRequest) Reset() {
	*x = PostVirtualHostModifyRequest{}
	mi := &file_proto_extension_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostVirtualHostModifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostVirtualHostModifyRequest) ProtoMessage() {}

func (x *PostVirtualHostModifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostVirtualHostModifyRequest.ProtoReflect.Descriptor instead.
func (*PostVirtualHostModifyRequest) Descriptor() ([]byte, []int) {
	return file_proto_extension_service_proto_rawDescGZIP(), []int{2}
}

func (x *PostVirtualHostModifyRequest) GetVirtualHost() *v3.VirtualHost {
	if x != nil {
		return x.VirtualHost
	}
	return nil
}

func (x *PostVirtualHostModifyRequest) GetPostVirtualHostContext() *PostVirtualHostExtensionContext {
	if x != nil {
		return x.PostVirtualHostContext
	}
	return nil
}

// PostVirtualHostModifyResponse is the expected response from an extension and contains a modified version of the VirtualHost that was sent
// If an extension returns a nil Virtual Host then it will not be modified
type PostVirtualHostModifyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VirtualHost   *v3.VirtualHost        `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost,proto3" json:"virtual_host,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostVirtualHostModifyResponse) Reset() {
	*x = PostVirtualHostModifyResponse{}
	mi := &file_proto_extension_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostVirtualHostModifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostVirtualHostModifyResponse) ProtoMessage() {}

func (x *PostVirtualHostModifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostVirtualHostModifyResponse.ProtoReflect.Descriptor instead.
func (*PostVirtualHostModifyResponse) Descriptor() ([]byte, []int) {
	return file_proto_extension_service_proto_rawDescGZIP(), []int{3}
}

func (x *PostVirtualHostModifyResponse) GetVirtualHost() *v3.VirtualHost {
	if x != nil {
		return x.VirtualHost
	}
	return nil
}

// PostVirtualHostModifyRequest sends a Listener that was generated by Envoy Gateway along with context information to an extension so that the Listener can be modified
type PostHTTPListenerModifyRequest struct {
	state               protoimpl.MessageState            `protogen:"open.v1"`
	Listener            *v32.Listener                     `protobuf:"bytes,1,opt,name=listener,proto3" json:"listener,omitempty"`
	PostListenerContext *PostHTTPListenerExtensionContext `protobuf:"bytes,2,opt,name=post_listener_context,json=postListenerContext,proto3" json:"post_listener_context,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *PostHTTPListenerModifyRequest) Reset() {
	*x = PostHTTPListenerModifyRequest{}
	mi := &file_proto_extension_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostHTTPListenerModifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostHTTPListenerModifyRequest) ProtoMessage() {}

func (x *PostHTTPListenerModifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostHTTPListenerModifyRequest.ProtoReflect.Descriptor instead.
func (*PostHTTPListenerModifyRequest) Descriptor() ([]byte, []int) {
	return file_proto_extension_service_proto_rawDescGZIP(), []int{4}
}

func (x *PostHTTPListenerModifyRequest) GetListener() *v32.Listener {
	if x != nil {
		return x.Listener
	}
	return nil
}

func (x *PostHTTPListenerModifyRequest) GetPostListenerContext() *PostHTTPListenerExtensionContext {
	if x != nil {
		return x.PostListenerContext
	}
	return nil
}

// PostHTTPListenerModifyResponse is the expected response from an extension and contains a modified version of the Listener that was sent
// If an extension returns a nil Listener then it will not be modified
type PostHTTPListenerModifyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Listener      *v32.Listener          `protobuf:"bytes,1,opt,name=listener,proto3" json:"listener,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostHTTPListenerModifyResponse) Reset() {
	*x = PostHTTPListenerModifyResponse{}
	mi := &file_proto_extension_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostHTTPListenerModifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostHTTPListenerModifyResponse) ProtoMessage() {}

func (x *PostHTTPListenerModifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostHTTPListenerModifyResponse.ProtoReflect.Descriptor instead.
func (*PostHTTPListenerModifyResponse) Descriptor() ([]byte, []int) {
	return file_proto_extension_service_proto_rawDescGZIP(), []int{5}
}

func (x *PostHTTPListenerModifyResponse) GetListener() *v32.Listener {
	if x != nil {
		return x.Listener
	}
	return nil
}

// PostTranslateModifyRequest currently sends only clusters and secrets to an extension.
// The extension is free to add/modify/remove the resources it received.
type PostTranslateModifyRequest struct {
	state                protoimpl.MessageState         `protogen:"open.v1"`
	PostTranslateContext *PostTranslateExtensionContext `protobuf:"bytes,1,opt,name=post_translate_context,json=postTranslateContext,proto3" json:"post_translate_context,omitempty"`
	Clusters             []*v31.Cluster                 `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	Secrets              []*v33.Secret                  `protobuf:"bytes,3,rep,name=secrets,proto3" json:"secrets,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *PostTranslateModifyRequest) Reset() {
	*x = PostTranslateModifyRequest{}
	mi := &file_proto_extension_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostTranslateModifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTranslateModifyRequest) ProtoMessage() {}

func (x *PostTranslateModifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTranslateModifyRequest.ProtoReflect.Descriptor instead.
func (*PostTranslateModifyRequest) Descriptor() ([]byte, []int) {
	return file_proto_extension_service_proto_rawDescGZIP(), []int{6}
}

func (x *PostTranslateModifyRequest) GetPostTranslateContext() *PostTranslateExtensionContext {
	if x != nil {
		return x.PostTranslateContext
	}
	return nil
}

func (x *PostTranslateModifyRequest) GetClusters() []*v31.Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *PostTranslateModifyRequest) GetSecrets() []*v33.Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

// PostTranslateModifyResponse is the expected response from an extension and contains
// the full list of xDS clusters and secrets to be used for the xDS config.
type PostTranslateModifyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Clusters      []*v31.Cluster         `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	Secrets       []*v33.Secret          `protobuf:"bytes,2,rep,name=secrets,proto3" json:"secrets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostTranslateModifyResponse) Reset() {
	*x = PostTranslateModifyResponse{}
	mi := &file_proto_extension_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostTranslateModifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTranslateModifyResponse) ProtoMessage() {}

func (x *PostTranslateModifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_extension_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTranslateModifyResponse.ProtoReflect.Descriptor instead.
func (*PostTranslateModifyResponse) Descriptor() ([]byte, []int) {
	return file_proto_extension_service_proto_rawDescGZIP(), []int{7}
}

func (x *PostTranslateModifyResponse) GetClusters() []*v31.Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *PostTranslateModifyResponse) GetSecrets() []*v33.Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

var File_proto_extension_service_proto protoreflect.FileDescriptor

var file_proto_extension_service_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a,
	0x16, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x5f, 0x0a, 0x12, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x10, 0x70, 0x6f, 0x73, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x8b, 0x01, 0x0a,
	0x17, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x08,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0xd9, 0x01, 0x0a, 0x1c, 0x50,
	0x6f, 0x73, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0c, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f,
	0x73, 0x74, 0x12, 0x72, 0x0a, 0x19, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x16,
	0x70, 0x6f, 0x73, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x66, 0x0a, 0x1d, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73,
	0x74, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x22, 0xcd,
	0x01, 0x0a, 0x1d, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3e, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x12, 0x6c, 0x0a, 0x15, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x54, 0x54,
	0x50, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x13, 0x70, 0x6f, 0x73, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x60,
	0x0a, 0x1e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x22, 0x94, 0x02, 0x0a, 0x1a, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x6b, 0x0a, 0x16, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x14, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3c, 0x0a, 0x08,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x07, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x1b, 0x50, 0x6f, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x32, 0xa5, 0x04, 0x0a, 0x15, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x74, 0x0a, 0x0f,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12,
	0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x86, 0x01, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12, 0x34, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x89, 0x01, 0x0a, 0x16,
	0x50, 0x6f, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x54, 0x54, 0x50, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12,
	0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_extension_service_proto_rawDescOnce sync.Once
	file_proto_extension_service_proto_rawDescData []byte
)

func file_proto_extension_service_proto_rawDescGZIP() []byte {
	file_proto_extension_service_proto_rawDescOnce.Do(func() {
		file_proto_extension_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_extension_service_proto_rawDesc), len(file_proto_extension_service_proto_rawDesc)))
	})
	return file_proto_extension_service_proto_rawDescData
}

var file_proto_extension_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_extension_service_proto_goTypes = []any{
	(*PostRouteModifyRequest)(nil),           // 0: envoygateway.extension.PostRouteModifyRequest
	(*PostRouteModifyResponse)(nil),          // 1: envoygateway.extension.PostRouteModifyResponse
	(*PostVirtualHostModifyRequest)(nil),     // 2: envoygateway.extension.PostVirtualHostModifyRequest
	(*PostVirtualHostModifyResponse)(nil),    // 3: envoygateway.extension.PostVirtualHostModifyResponse
	(*PostHTTPListenerModifyRequest)(nil),    // 4: envoygateway.extension.PostHTTPListenerModifyRequest
	(*PostHTTPListenerModifyResponse)(nil),   // 5: envoygateway.extension.PostHTTPListenerModifyResponse
	(*PostTranslateModifyRequest)(nil),       // 6: envoygateway.extension.PostTranslateModifyRequest
	(*PostTranslateModifyResponse)(nil),      // 7: envoygateway.extension.PostTranslateModifyResponse
	(*v3.Route)(nil),                         // 8: envoy.config.route.v3.Route
	(*PostRouteExtensionContext)(nil),        // 9: envoygateway.extension.PostRouteExtensionContext
	(*v31.Cluster)(nil),                      // 10: envoy.config.cluster.v3.Cluster
	(*v3.VirtualHost)(nil),                   // 11: envoy.config.route.v3.VirtualHost
	(*PostVirtualHostExtensionContext)(nil),  // 12: envoygateway.extension.PostVirtualHostExtensionContext
	(*v32.Listener)(nil),                     // 13: envoy.config.listener.v3.Listener
	(*PostHTTPListenerExtensionContext)(nil), // 14: envoygateway.extension.PostHTTPListenerExtensionContext
	(*PostTranslateExtensionContext)(nil),    // 15: envoygateway.extension.PostTranslateExtensionContext
	(*v33.Secret)(nil),                       // 16: envoy.extensions.transport_sockets.tls.v3.Secret
}
var file_proto_extension_service_proto_depIdxs = []int32{
	8,  // 0: envoygateway.extension.PostRouteModifyRequest.route:type_name -> envoy.config.route.v3.Route
	9,  // 1: envoygateway.extension.PostRouteModifyRequest.post_route_context:type_name -> envoygateway.extension.PostRouteExtensionContext
	8,  // 2: envoygateway.extension.PostRouteModifyResponse.route:type_name -> envoy.config.route.v3.Route
	10, // 3: envoygateway.extension.PostRouteModifyResponse.clusters:type_name -> envoy.config.cluster.v3.Cluster
	11, // 4: envoygateway.extension.PostVirtualHostModifyRequest.virtual_host:type_name -> envoy.config.route.v3.VirtualHost
	12, // 5: envoygateway.extension.PostVirtualHostModifyRequest.post_virtual_host_context:type_name -> envoygateway.extension.PostVirtualHostExtensionContext
	11, // 6: envoygateway.extension.PostVirtualHostModifyResponse.virtual_host:type_name -> envoy.config.route.v3.VirtualHost
	13, // 7: envoygateway.extension.PostHTTPListenerModifyRequest.listener:type_name -> envoy.config.listener.v3.Listener
	14, // 8: envoygateway.extension.PostHTTPListenerModifyRequest.post_listener_context:type_name -> envoygateway.extension.PostHTTPListenerExtensionContext
	13, // 9: envoygateway.extension.PostHTTPListenerModifyResponse.listener:type_name -> envoy.config.listener.v3.Listener
	15, // 10: envoygateway.extension.PostTranslateModifyRequest.post_translate_context:type_name -> envoygateway.extension.PostTranslateExtensionContext
	10, // 11: envoygateway.extension.PostTranslateModifyRequest.clusters:type_name -> envoy.config.cluster.v3.Cluster
	16, // 12: envoygateway.extension.PostTranslateModifyRequest.secrets:type_name -> envoy.extensions.transport_sockets.tls.v3.Secret
	10, // 13: envoygateway.extension.PostTranslateModifyResponse.clusters:type_name -> envoy.config.cluster.v3.Cluster
	16, // 14: envoygateway.extension.PostTranslateModifyResponse.secrets:type_name -> envoy.extensions.transport_sockets.tls.v3.Secret
	0,  // 15: envoygateway.extension.EnvoyGatewayExtension.PostRouteModify:input_type -> envoygateway.extension.PostRouteModifyRequest
	2,  // 16: envoygateway.extension.EnvoyGatewayExtension.PostVirtualHostModify:input_type -> envoygateway.extension.PostVirtualHostModifyRequest
	4,  // 17: envoygateway.extension.EnvoyGatewayExtension.PostHTTPListenerModify:input_type -> envoygateway.extension.PostHTTPListenerModifyRequest
	6,  // 18: envoygateway.extension.EnvoyGatewayExtension.PostTranslateModify:input_type -> envoygateway.extension.PostTranslateModifyRequest
	1,  // 19: envoygateway.extension.EnvoyGatewayExtension.PostRouteModify:output_type -> envoygateway.extension.PostRouteModifyResponse
	3,  // 20: envoygateway.extension.EnvoyGatewayExtension.PostVirtualHostModify:output_type -> envoygateway.extension.PostVirtualHostModifyResponse
	5,  // 21: envoygateway.extension.EnvoyGatewayExtension.PostHTTPListenerModify:output_type -> envoygateway.extension.PostHTTPListenerModifyResponse
	7,  // 22: envoygateway.extension.EnvoyGatewayExtension.PostTranslateModify:output_type -> envoygateway.extension.PostTranslateModifyResponse
	19, // [19:23] is the sub-list for method output_type
	15, // [15:19] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_proto_extension_service_proto_init() }
func file_proto_extension_service_proto_init() {
	if File_proto_extension_service_proto != nil {
		return
	}
	file_proto_extension_context_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_extension_service_proto_rawDesc), len(file_proto_extension_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_extension_service_proto_goTypes,
		DependencyIndexes: file_proto_extension_service_proto_depIdxs,
		MessageInfos:      file_proto_extension_service_proto_msgTypes,
	}.Build()
	File_proto_extension_service_proto = out.File
	file_proto_extension_service_proto_goTypes = nil
	file_proto_extension_service_proto_depIdxs = nil
}
