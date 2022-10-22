// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/artifacts/oci/types.proto

package oci

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

type RawDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repository  string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	MediaType   string `protobuf:"bytes,2,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	RawManifest []byte `protobuf:"bytes,3,opt,name=raw_manifest,json=rawManifest,proto3" json:"raw_manifest,omitempty"`
	RawConfig   []byte `protobuf:"bytes,4,opt,name=raw_config,json=rawConfig,proto3" json:"raw_config,omitempty"` // Only available when the descriptor points at an image.
}

func (x *RawDescriptor) Reset() {
	*x = RawDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_artifacts_oci_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawDescriptor) ProtoMessage() {}

func (x *RawDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_internal_artifacts_oci_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawDescriptor.ProtoReflect.Descriptor instead.
func (*RawDescriptor) Descriptor() ([]byte, []int) {
	return file_internal_artifacts_oci_types_proto_rawDescGZIP(), []int{0}
}

func (x *RawDescriptor) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *RawDescriptor) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *RawDescriptor) GetRawManifest() []byte {
	if x != nil {
		return x.RawManifest
	}
	return nil
}

func (x *RawDescriptor) GetRawConfig() []byte {
	if x != nil {
		return x.RawConfig
	}
	return nil
}

var File_internal_artifacts_oci_types_proto protoreflect.FileDescriptor

var file_internal_artifacts_oci_types_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x2f, 0x6f, 0x63, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x6f, 0x63, 0x69, 0x22, 0x90, 0x01, 0x0a, 0x0d, 0x52, 0x61, 0x77, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x61, 0x77, 0x5f,
	0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x72, 0x61, 0x77, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x61, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x35, 0x5a, 0x33, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x6f, 0x63,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_artifacts_oci_types_proto_rawDescOnce sync.Once
	file_internal_artifacts_oci_types_proto_rawDescData = file_internal_artifacts_oci_types_proto_rawDesc
)

func file_internal_artifacts_oci_types_proto_rawDescGZIP() []byte {
	file_internal_artifacts_oci_types_proto_rawDescOnce.Do(func() {
		file_internal_artifacts_oci_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_artifacts_oci_types_proto_rawDescData)
	})
	return file_internal_artifacts_oci_types_proto_rawDescData
}

var file_internal_artifacts_oci_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_artifacts_oci_types_proto_goTypes = []interface{}{
	(*RawDescriptor)(nil), // 0: foundation.internal.artifacts.oci.RawDescriptor
}
var file_internal_artifacts_oci_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_artifacts_oci_types_proto_init() }
func file_internal_artifacts_oci_types_proto_init() {
	if File_internal_artifacts_oci_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_artifacts_oci_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawDescriptor); i {
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
			RawDescriptor: file_internal_artifacts_oci_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_artifacts_oci_types_proto_goTypes,
		DependencyIndexes: file_internal_artifacts_oci_types_proto_depIdxs,
		MessageInfos:      file_internal_artifacts_oci_types_proto_msgTypes,
	}.Build()
	File_internal_artifacts_oci_types_proto = out.File
	file_internal_artifacts_oci_types_proto_rawDesc = nil
	file_internal_artifacts_oci_types_proto_goTypes = nil
	file_internal_artifacts_oci_types_proto_depIdxs = nil
}
