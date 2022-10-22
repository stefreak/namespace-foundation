// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/secret.proto

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

type SecretSpec_GenerateSpec_Format int32

const (
	SecretSpec_GenerateSpec_FORMAT_UNKNOWN SecretSpec_GenerateSpec_Format = 0 // Defaults to base64.
	SecretSpec_GenerateSpec_FORMAT_BASE64  SecretSpec_GenerateSpec_Format = 1
	SecretSpec_GenerateSpec_FORMAT_BASE32  SecretSpec_GenerateSpec_Format = 2
)

// Enum value maps for SecretSpec_GenerateSpec_Format.
var (
	SecretSpec_GenerateSpec_Format_name = map[int32]string{
		0: "FORMAT_UNKNOWN",
		1: "FORMAT_BASE64",
		2: "FORMAT_BASE32",
	}
	SecretSpec_GenerateSpec_Format_value = map[string]int32{
		"FORMAT_UNKNOWN": 0,
		"FORMAT_BASE64":  1,
		"FORMAT_BASE32":  2,
	}
)

func (x SecretSpec_GenerateSpec_Format) Enum() *SecretSpec_GenerateSpec_Format {
	p := new(SecretSpec_GenerateSpec_Format)
	*p = x
	return p
}

func (x SecretSpec_GenerateSpec_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecretSpec_GenerateSpec_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_secret_proto_enumTypes[0].Descriptor()
}

func (SecretSpec_GenerateSpec_Format) Type() protoreflect.EnumType {
	return &file_schema_secret_proto_enumTypes[0]
}

func (x SecretSpec_GenerateSpec_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecretSpec_GenerateSpec_Format.Descriptor instead.
func (SecretSpec_GenerateSpec_Format) EnumDescriptor() ([]byte, []int) {
	return file_schema_secret_proto_rawDescGZIP(), []int{0, 0, 0}
}

type SecretSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Generate    *SecretSpec_GenerateSpec `protobuf:"bytes,5,opt,name=generate,proto3" json:"generate,omitempty"`
}

func (x *SecretSpec) Reset() {
	*x = SecretSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretSpec) ProtoMessage() {}

func (x *SecretSpec) ProtoReflect() protoreflect.Message {
	mi := &file_schema_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretSpec.ProtoReflect.Descriptor instead.
func (*SecretSpec) Descriptor() ([]byte, []int) {
	return file_schema_secret_proto_rawDescGZIP(), []int{0}
}

func (x *SecretSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecretSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SecretSpec) GetGenerate() *SecretSpec_GenerateSpec {
	if x != nil {
		return x.Generate
	}
	return nil
}

type SecretSpec_GenerateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueId        string                         `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"` // If not set, will default to a value derived from the secret reference.
	RandomByteCount int32                          `protobuf:"varint,2,opt,name=random_byte_count,json=randomByteCount,proto3" json:"random_byte_count,omitempty"`
	Format          SecretSpec_GenerateSpec_Format `protobuf:"varint,3,opt,name=format,proto3,enum=foundation.schema.SecretSpec_GenerateSpec_Format" json:"format,omitempty"`
}

func (x *SecretSpec_GenerateSpec) Reset() {
	*x = SecretSpec_GenerateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretSpec_GenerateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretSpec_GenerateSpec) ProtoMessage() {}

func (x *SecretSpec_GenerateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_schema_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretSpec_GenerateSpec.ProtoReflect.Descriptor instead.
func (*SecretSpec_GenerateSpec) Descriptor() ([]byte, []int) {
	return file_schema_secret_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SecretSpec_GenerateSpec) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *SecretSpec_GenerateSpec) GetRandomByteCount() int32 {
	if x != nil {
		return x.RandomByteCount
	}
	return 0
}

func (x *SecretSpec_GenerateSpec) GetFormat() SecretSpec_GenerateSpec_Format {
	if x != nil {
		return x.Format
	}
	return SecretSpec_GenerateSpec_FORMAT_UNKNOWN
}

var File_schema_secret_proto protoreflect.FileDescriptor

var file_schema_secret_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0xff, 0x02, 0x0a, 0x0a, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a,
	0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x1a, 0xe6, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x42, 0x79, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x49, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x31, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x42, 0x0a, 0x06, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x36, 0x34, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46,
	0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x33, 0x32, 0x10, 0x02, 0x4a, 0x04,
	0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x42, 0x25, 0x5a, 0x23, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_secret_proto_rawDescOnce sync.Once
	file_schema_secret_proto_rawDescData = file_schema_secret_proto_rawDesc
)

func file_schema_secret_proto_rawDescGZIP() []byte {
	file_schema_secret_proto_rawDescOnce.Do(func() {
		file_schema_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_secret_proto_rawDescData)
	})
	return file_schema_secret_proto_rawDescData
}

var file_schema_secret_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_schema_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_schema_secret_proto_goTypes = []interface{}{
	(SecretSpec_GenerateSpec_Format)(0), // 0: foundation.schema.SecretSpec.GenerateSpec.Format
	(*SecretSpec)(nil),                  // 1: foundation.schema.SecretSpec
	(*SecretSpec_GenerateSpec)(nil),     // 2: foundation.schema.SecretSpec.GenerateSpec
}
var file_schema_secret_proto_depIdxs = []int32{
	2, // 0: foundation.schema.SecretSpec.generate:type_name -> foundation.schema.SecretSpec.GenerateSpec
	0, // 1: foundation.schema.SecretSpec.GenerateSpec.format:type_name -> foundation.schema.SecretSpec.GenerateSpec.Format
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_schema_secret_proto_init() }
func file_schema_secret_proto_init() {
	if File_schema_secret_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretSpec); i {
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
		file_schema_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretSpec_GenerateSpec); i {
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
			RawDescriptor: file_schema_secret_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_secret_proto_goTypes,
		DependencyIndexes: file_schema_secret_proto_depIdxs,
		EnumInfos:         file_schema_secret_proto_enumTypes,
		MessageInfos:      file_schema_secret_proto_msgTypes,
	}.Build()
	File_schema_secret_proto = out.File
	file_schema_secret_proto_rawDesc = nil
	file_schema_secret_proto_goTypes = nil
	file_schema_secret_proto_depIdxs = nil
}
