// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/resolvable.proto

package schema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Resolvable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value                              string                       `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ExperimentalFromSecret             string                       `protobuf:"bytes,2,opt,name=experimental_from_secret,json=experimentalFromSecret,proto3" json:"experimental_from_secret,omitempty"`                                         // Kubernetes specific.
	ExperimentalFromDownwardsFieldPath string                       `protobuf:"bytes,3,opt,name=experimental_from_downwards_field_path,json=experimentalFromDownwardsFieldPath,proto3" json:"experimental_from_downwards_field_path,omitempty"` // Kubernetes specific.
	FromSecretRef                      *PackageRef                  `protobuf:"bytes,4,opt,name=from_secret_ref,json=fromSecretRef,proto3" json:"from_secret_ref,omitempty"`
	FromServiceEndpoint                *ServiceRef                  `protobuf:"bytes,5,opt,name=from_service_endpoint,json=fromServiceEndpoint,proto3" json:"from_service_endpoint,omitempty"` // Inject the endpoint of a service in the stack.
	FromServiceIngress                 *ServiceRef                  `protobuf:"bytes,6,opt,name=from_service_ingress,json=fromServiceIngress,proto3" json:"from_service_ingress,omitempty"`    // Inject the ingress address of a service. This is available from outside the cluster.
	FromResourceField                  *ResourceConfigFieldSelector `protobuf:"bytes,7,opt,name=from_resource_field,json=fromResourceField,proto3" json:"from_resource_field,omitempty"`
	FromFieldSelector                  *FieldSelector               `protobuf:"bytes,8,opt,name=from_field_selector,json=fromFieldSelector,proto3" json:"from_field_selector,omitempty"`
}

func (x *Resolvable) Reset() {
	*x = Resolvable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resolvable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resolvable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resolvable) ProtoMessage() {}

func (x *Resolvable) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resolvable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resolvable.ProtoReflect.Descriptor instead.
func (*Resolvable) Descriptor() ([]byte, []int) {
	return file_schema_resolvable_proto_rawDescGZIP(), []int{0}
}

func (x *Resolvable) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Resolvable) GetExperimentalFromSecret() string {
	if x != nil {
		return x.ExperimentalFromSecret
	}
	return ""
}

func (x *Resolvable) GetExperimentalFromDownwardsFieldPath() string {
	if x != nil {
		return x.ExperimentalFromDownwardsFieldPath
	}
	return ""
}

func (x *Resolvable) GetFromSecretRef() *PackageRef {
	if x != nil {
		return x.FromSecretRef
	}
	return nil
}

func (x *Resolvable) GetFromServiceEndpoint() *ServiceRef {
	if x != nil {
		return x.FromServiceEndpoint
	}
	return nil
}

func (x *Resolvable) GetFromServiceIngress() *ServiceRef {
	if x != nil {
		return x.FromServiceIngress
	}
	return nil
}

func (x *Resolvable) GetFromResourceField() *ResourceConfigFieldSelector {
	if x != nil {
		return x.FromResourceField
	}
	return nil
}

func (x *Resolvable) GetFromFieldSelector() *FieldSelector {
	if x != nil {
		return x.FromFieldSelector
	}
	return nil
}

type NamedResolvable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *Resolvable `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NamedResolvable) Reset() {
	*x = NamedResolvable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resolvable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamedResolvable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedResolvable) ProtoMessage() {}

func (x *NamedResolvable) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resolvable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedResolvable.ProtoReflect.Descriptor instead.
func (*NamedResolvable) Descriptor() ([]byte, []int) {
	return file_schema_resolvable_proto_rawDescGZIP(), []int{1}
}

func (x *NamedResolvable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamedResolvable) GetValue() *Resolvable {
	if x != nil {
		return x.Value
	}
	return nil
}

type FieldSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance      *ResolvableSource `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	FieldSelector string            `protobuf:"bytes,2,opt,name=field_selector,json=fieldSelector,proto3" json:"field_selector,omitempty"`
}

func (x *FieldSelector) Reset() {
	*x = FieldSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resolvable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldSelector) ProtoMessage() {}

func (x *FieldSelector) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resolvable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldSelector.ProtoReflect.Descriptor instead.
func (*FieldSelector) Descriptor() ([]byte, []int) {
	return file_schema_resolvable_proto_rawDescGZIP(), []int{2}
}

func (x *FieldSelector) GetInstance() *ResolvableSource {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *FieldSelector) GetFieldSelector() string {
	if x != nil {
		return x.FieldSelector
	}
	return ""
}

type ResolvableSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server                       *PackageRef   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`                                                                                       // Resolves into a runtime.Server
	Service                      *ServiceRef   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`                                                                                     // Resolves into a runtime.Server.Service
	SelectInternalEndpointByKind string        `protobuf:"bytes,3,opt,name=select_internal_endpoint_by_kind,json=selectInternalEndpointByKind,proto3" json:"select_internal_endpoint_by_kind,omitempty"` // Resolves into a runtime.Service.InternalEndpoint
	UntypedJson                  *FileContents `protobuf:"bytes,4,opt,name=untyped_json,json=untypedJson,proto3" json:"untyped_json,omitempty"`
}

func (x *ResolvableSource) Reset() {
	*x = ResolvableSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resolvable_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolvableSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolvableSource) ProtoMessage() {}

func (x *ResolvableSource) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resolvable_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolvableSource.ProtoReflect.Descriptor instead.
func (*ResolvableSource) Descriptor() ([]byte, []int) {
	return file_schema_resolvable_proto_rawDescGZIP(), []int{3}
}

func (x *ResolvableSource) GetServer() *PackageRef {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *ResolvableSource) GetService() *ServiceRef {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ResolvableSource) GetSelectInternalEndpointByKind() string {
	if x != nil {
		return x.SelectInternalEndpointByKind
	}
	return ""
}

func (x *ResolvableSource) GetUntypedJson() *FileContents {
	if x != nil {
		return x.UntypedJson
	}
	return nil
}

type ResourceConfigFieldSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource      *PackageRef `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	FieldSelector string      `protobuf:"bytes,2,opt,name=field_selector,json=fieldSelector,proto3" json:"field_selector,omitempty"`
}

func (x *ResourceConfigFieldSelector) Reset() {
	*x = ResourceConfigFieldSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resolvable_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceConfigFieldSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceConfigFieldSelector) ProtoMessage() {}

func (x *ResourceConfigFieldSelector) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resolvable_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceConfigFieldSelector.ProtoReflect.Descriptor instead.
func (*ResourceConfigFieldSelector) Descriptor() ([]byte, []int) {
	return file_schema_resolvable_proto_rawDescGZIP(), []int{4}
}

func (x *ResourceConfigFieldSelector) GetResource() *PackageRef {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *ResourceConfigFieldSelector) GetFieldSelector() string {
	if x != nil {
		return x.FieldSelector
	}
	return ""
}

var File_schema_resolvable_proto protoreflect.FileDescriptor

var file_schema_resolvable_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x14, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x04,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x52, 0x0a, 0x26,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x77, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x22, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x6f,
	0x77, 0x6e, 0x77, 0x61, 0x72, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x45, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x51, 0x0a, 0x15, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x66, 0x52, 0x13, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x14, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x66, 0x52, 0x12, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x5e, 0x0a, 0x13, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x11, 0x66, 0x72, 0x6f, 0x6d, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x50, 0x0a, 0x13, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x11, 0x66, 0x72, 0x6f, 0x6d,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x5a, 0x0a,
	0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x77, 0x0a, 0x0d, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0x8e, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x37,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x66, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x20, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x42, 0x0a, 0x0c, 0x75, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0b, 0x75, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x4a,
	0x73, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x1b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x66, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x42, 0x25, 0x5a, 0x23, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_schema_resolvable_proto_rawDescOnce sync.Once
	file_schema_resolvable_proto_rawDescData = file_schema_resolvable_proto_rawDesc
)

func file_schema_resolvable_proto_rawDescGZIP() []byte {
	file_schema_resolvable_proto_rawDescOnce.Do(func() {
		file_schema_resolvable_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_resolvable_proto_rawDescData)
	})
	return file_schema_resolvable_proto_rawDescData
}

var file_schema_resolvable_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_schema_resolvable_proto_goTypes = []interface{}{
	(*Resolvable)(nil),                  // 0: foundation.schema.Resolvable
	(*NamedResolvable)(nil),             // 1: foundation.schema.NamedResolvable
	(*FieldSelector)(nil),               // 2: foundation.schema.FieldSelector
	(*ResolvableSource)(nil),            // 3: foundation.schema.ResolvableSource
	(*ResourceConfigFieldSelector)(nil), // 4: foundation.schema.ResourceConfigFieldSelector
	(*PackageRef)(nil),                  // 5: foundation.schema.PackageRef
	(*ServiceRef)(nil),                  // 6: foundation.schema.ServiceRef
	(*FileContents)(nil),                // 7: foundation.schema.FileContents
}
var file_schema_resolvable_proto_depIdxs = []int32{
	5,  // 0: foundation.schema.Resolvable.from_secret_ref:type_name -> foundation.schema.PackageRef
	6,  // 1: foundation.schema.Resolvable.from_service_endpoint:type_name -> foundation.schema.ServiceRef
	6,  // 2: foundation.schema.Resolvable.from_service_ingress:type_name -> foundation.schema.ServiceRef
	4,  // 3: foundation.schema.Resolvable.from_resource_field:type_name -> foundation.schema.ResourceConfigFieldSelector
	2,  // 4: foundation.schema.Resolvable.from_field_selector:type_name -> foundation.schema.FieldSelector
	0,  // 5: foundation.schema.NamedResolvable.value:type_name -> foundation.schema.Resolvable
	3,  // 6: foundation.schema.FieldSelector.instance:type_name -> foundation.schema.ResolvableSource
	5,  // 7: foundation.schema.ResolvableSource.server:type_name -> foundation.schema.PackageRef
	6,  // 8: foundation.schema.ResolvableSource.service:type_name -> foundation.schema.ServiceRef
	7,  // 9: foundation.schema.ResolvableSource.untyped_json:type_name -> foundation.schema.FileContents
	5,  // 10: foundation.schema.ResourceConfigFieldSelector.resource:type_name -> foundation.schema.PackageRef
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_schema_resolvable_proto_init() }
func file_schema_resolvable_proto_init() {
	if File_schema_resolvable_proto != nil {
		return
	}
	file_schema_package_proto_init()
	file_schema_filecontents_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_resolvable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resolvable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_resolvable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamedResolvable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_resolvable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_resolvable_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolvableSource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_resolvable_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceConfigFieldSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_resolvable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_resolvable_proto_goTypes,
		DependencyIndexes: file_schema_resolvable_proto_depIdxs,
		MessageInfos:      file_schema_resolvable_proto_msgTypes,
	}.Build()
	File_schema_resolvable_proto = out.File
	file_schema_resolvable_proto_rawDesc = nil
	file_schema_resolvable_proto_goTypes = nil
	file_schema_resolvable_proto_depIdxs = nil
}
