// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: universe/db/maria/incluster/creds/provider.proto

package creds

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

type CredsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CredsRequest) Reset() {
	*x = CredsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_maria_incluster_creds_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredsRequest) ProtoMessage() {}

func (x *CredsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_maria_incluster_creds_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredsRequest.ProtoReflect.Descriptor instead.
func (*CredsRequest) Descriptor() ([]byte, []int) {
	return file_universe_db_maria_incluster_creds_provider_proto_rawDescGZIP(), []int{0}
}

type Creds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Creds) Reset() {
	*x = Creds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_maria_incluster_creds_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Creds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Creds) ProtoMessage() {}

func (x *Creds) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_maria_incluster_creds_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Creds.ProtoReflect.Descriptor instead.
func (*Creds) Descriptor() ([]byte, []int) {
	return file_universe_db_maria_incluster_creds_provider_proto_rawDescGZIP(), []int{1}
}

func (x *Creds) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_universe_db_maria_incluster_creds_provider_proto protoreflect.FileDescriptor

var file_universe_db_maria_incluster_creds_provider_proto_rawDesc = []byte{
	0x0a, 0x30, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x64, 0x62, 0x2f, 0x6d, 0x61,
	0x72, 0x69, 0x61, 0x2f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x72,
	0x65, 0x64, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x26, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x64, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x69, 0x61,
	0x2e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x05, 0x43, 0x72,
	0x65, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42,
	0x40, 0x5a, 0x3e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x64, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x69,
	0x61, 0x2f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x64,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_universe_db_maria_incluster_creds_provider_proto_rawDescOnce sync.Once
	file_universe_db_maria_incluster_creds_provider_proto_rawDescData = file_universe_db_maria_incluster_creds_provider_proto_rawDesc
)

func file_universe_db_maria_incluster_creds_provider_proto_rawDescGZIP() []byte {
	file_universe_db_maria_incluster_creds_provider_proto_rawDescOnce.Do(func() {
		file_universe_db_maria_incluster_creds_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_universe_db_maria_incluster_creds_provider_proto_rawDescData)
	})
	return file_universe_db_maria_incluster_creds_provider_proto_rawDescData
}

var file_universe_db_maria_incluster_creds_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_universe_db_maria_incluster_creds_provider_proto_goTypes = []interface{}{
	(*CredsRequest)(nil), // 0: foundation.universe.db.maria.incluster.CredsRequest
	(*Creds)(nil),        // 1: foundation.universe.db.maria.incluster.Creds
}
var file_universe_db_maria_incluster_creds_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_universe_db_maria_incluster_creds_provider_proto_init() }
func file_universe_db_maria_incluster_creds_provider_proto_init() {
	if File_universe_db_maria_incluster_creds_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_universe_db_maria_incluster_creds_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredsRequest); i {
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
		file_universe_db_maria_incluster_creds_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Creds); i {
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
			RawDescriptor: file_universe_db_maria_incluster_creds_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_universe_db_maria_incluster_creds_provider_proto_goTypes,
		DependencyIndexes: file_universe_db_maria_incluster_creds_provider_proto_depIdxs,
		MessageInfos:      file_universe_db_maria_incluster_creds_provider_proto_msgTypes,
	}.Build()
	File_universe_db_maria_incluster_creds_provider_proto = out.File
	file_universe_db_maria_incluster_creds_provider_proto_rawDesc = nil
	file_universe_db_maria_incluster_creds_provider_proto_goTypes = nil
	file_universe_db_maria_incluster_creds_provider_proto_depIdxs = nil
}
