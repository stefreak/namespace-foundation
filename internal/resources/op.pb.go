// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/resources/op.proto

package resources

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	protos "namespacelabs.dev/foundation/internal/codegen/protos"
	schema "namespacelabs.dev/foundation/schema"
	runtime "namespacelabs.dev/foundation/schema/runtime"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OpWaitForProviderResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceInstanceId string                           `protobuf:"bytes,1,opt,name=resource_instance_id,json=resourceInstanceId,proto3" json:"resource_instance_id,omitempty"`
	Deployable         *runtime.Deployable              `protobuf:"bytes,2,opt,name=deployable,proto3" json:"deployable,omitempty"`
	ResourceClass      *schema.ResourceClass            `protobuf:"bytes,3,opt,name=resource_class,json=resourceClass,proto3" json:"resource_class,omitempty"`
	InstanceTypeSource *protos.FileDescriptorSetAndDeps `protobuf:"bytes,4,opt,name=instance_type_source,json=instanceTypeSource,proto3" json:"instance_type_source,omitempty"`
}

func (x *OpWaitForProviderResults) Reset() {
	*x = OpWaitForProviderResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_resources_op_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpWaitForProviderResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpWaitForProviderResults) ProtoMessage() {}

func (x *OpWaitForProviderResults) ProtoReflect() protoreflect.Message {
	mi := &file_internal_resources_op_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpWaitForProviderResults.ProtoReflect.Descriptor instead.
func (*OpWaitForProviderResults) Descriptor() ([]byte, []int) {
	return file_internal_resources_op_proto_rawDescGZIP(), []int{0}
}

func (x *OpWaitForProviderResults) GetResourceInstanceId() string {
	if x != nil {
		return x.ResourceInstanceId
	}
	return ""
}

func (x *OpWaitForProviderResults) GetDeployable() *runtime.Deployable {
	if x != nil {
		return x.Deployable
	}
	return nil
}

func (x *OpWaitForProviderResults) GetResourceClass() *schema.ResourceClass {
	if x != nil {
		return x.ResourceClass
	}
	return nil
}

func (x *OpWaitForProviderResults) GetInstanceTypeSource() *protos.FileDescriptorSetAndDeps {
	if x != nil {
		return x.InstanceTypeSource
	}
	return nil
}

var File_internal_resources_op_proto protoreflect.FileDescriptor

var file_internal_resources_op_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x64, 0x65, 0x73, 0x63, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcc, 0x02, 0x0a, 0x18, 0x4f, 0x70, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x14,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x45,
	0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x6e,
	0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x74, 0x41, 0x6e, 0x64, 0x44, 0x65, 0x70, 0x73, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x31,
	0x5a, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_resources_op_proto_rawDescOnce sync.Once
	file_internal_resources_op_proto_rawDescData = file_internal_resources_op_proto_rawDesc
)

func file_internal_resources_op_proto_rawDescGZIP() []byte {
	file_internal_resources_op_proto_rawDescOnce.Do(func() {
		file_internal_resources_op_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_resources_op_proto_rawDescData)
	})
	return file_internal_resources_op_proto_rawDescData
}

var file_internal_resources_op_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_resources_op_proto_goTypes = []interface{}{
	(*OpWaitForProviderResults)(nil),        // 0: foundation.internal.resources.OpWaitForProviderResults
	(*runtime.Deployable)(nil),              // 1: foundation.schema.runtime.Deployable
	(*schema.ResourceClass)(nil),            // 2: foundation.schema.ResourceClass
	(*protos.FileDescriptorSetAndDeps)(nil), // 3: foundation.workspace.source.protos.FileDescriptorSetAndDeps
}
var file_internal_resources_op_proto_depIdxs = []int32{
	1, // 0: foundation.internal.resources.OpWaitForProviderResults.deployable:type_name -> foundation.schema.runtime.Deployable
	2, // 1: foundation.internal.resources.OpWaitForProviderResults.resource_class:type_name -> foundation.schema.ResourceClass
	3, // 2: foundation.internal.resources.OpWaitForProviderResults.instance_type_source:type_name -> foundation.workspace.source.protos.FileDescriptorSetAndDeps
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_resources_op_proto_init() }
func file_internal_resources_op_proto_init() {
	if File_internal_resources_op_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_resources_op_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpWaitForProviderResults); i {
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
			RawDescriptor: file_internal_resources_op_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_resources_op_proto_goTypes,
		DependencyIndexes: file_internal_resources_op_proto_depIdxs,
		MessageInfos:      file_internal_resources_op_proto_msgTypes,
	}.Build()
	File_internal_resources_op_proto = out.File
	file_internal_resources_op_proto_rawDesc = nil
	file_internal_resources_op_proto_goTypes = nil
	file_internal_resources_op_proto_depIdxs = nil
}
