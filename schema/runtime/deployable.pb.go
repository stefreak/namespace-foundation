// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/runtime/deployable.proto

package runtime

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

type Deployable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName     string `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	Id              string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DeployableClass string `protobuf:"bytes,4,opt,name=deployable_class,json=deployableClass,proto3" json:"deployable_class,omitempty"`
}

func (x *Deployable) Reset() {
	*x = Deployable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_deployable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployable) ProtoMessage() {}

func (x *Deployable) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_deployable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployable.ProtoReflect.Descriptor instead.
func (*Deployable) Descriptor() ([]byte, []int) {
	return file_schema_runtime_deployable_proto_rawDescGZIP(), []int{0}
}

func (x *Deployable) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *Deployable) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Deployable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deployable) GetDeployableClass() string {
	if x != nil {
		return x.DeployableClass
	}
	return ""
}

var File_schema_runtime_deployable_proto protoreflect.FileDescriptor

var file_schema_runtime_deployable_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x7e, 0x0a, 0x0a,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x2d, 0x5a, 0x2b,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65,
	0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_schema_runtime_deployable_proto_rawDescOnce sync.Once
	file_schema_runtime_deployable_proto_rawDescData = file_schema_runtime_deployable_proto_rawDesc
)

func file_schema_runtime_deployable_proto_rawDescGZIP() []byte {
	file_schema_runtime_deployable_proto_rawDescOnce.Do(func() {
		file_schema_runtime_deployable_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_runtime_deployable_proto_rawDescData)
	})
	return file_schema_runtime_deployable_proto_rawDescData
}

var file_schema_runtime_deployable_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_runtime_deployable_proto_goTypes = []interface{}{
	(*Deployable)(nil), // 0: foundation.schema.runtime.Deployable
}
var file_schema_runtime_deployable_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_runtime_deployable_proto_init() }
func file_schema_runtime_deployable_proto_init() {
	if File_schema_runtime_deployable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_runtime_deployable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployable); i {
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
			RawDescriptor: file_schema_runtime_deployable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_runtime_deployable_proto_goTypes,
		DependencyIndexes: file_schema_runtime_deployable_proto_depIdxs,
		MessageInfos:      file_schema_runtime_deployable_proto_msgTypes,
	}.Build()
	File_schema_runtime_deployable_proto = out.File
	file_schema_runtime_deployable_proto_rawDesc = nil
	file_schema_runtime_deployable_proto_goTypes = nil
	file_schema_runtime_deployable_proto_depIdxs = nil
}
