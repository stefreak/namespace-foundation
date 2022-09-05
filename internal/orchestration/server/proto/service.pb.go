// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/orchestration/server/proto/service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fnapi "namespacelabs.dev/foundation/internal/fnapi"
	aws "namespacelabs.dev/foundation/providers/aws"
	schema "namespacelabs.dev/foundation/schema"
	orchestration "namespacelabs.dev/foundation/schema/orchestration"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeployRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan *schema.DeployPlan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	Aws  *aws.Conf          `protobuf:"bytes,4,opt,name=aws,proto3" json:"aws,omitempty"`
	Auth *fnapi.UserAuth    `protobuf:"bytes,5,opt,name=auth,proto3" json:"auth,omitempty"` // Time-limited Namespace session.
}

func (x *DeployRequest) Reset() {
	*x = DeployRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orchestration_server_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployRequest) ProtoMessage() {}

func (x *DeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orchestration_server_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployRequest.ProtoReflect.Descriptor instead.
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return file_internal_orchestration_server_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *DeployRequest) GetPlan() *schema.DeployPlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *DeployRequest) GetAws() *aws.Conf {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *DeployRequest) GetAuth() *fnapi.UserAuth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type DeployResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Deployment to follow
}

func (x *DeployResponse) Reset() {
	*x = DeployResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orchestration_server_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployResponse) ProtoMessage() {}

func (x *DeployResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orchestration_server_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployResponse.ProtoReflect.Descriptor instead.
func (*DeployResponse) Descriptor() ([]byte, []int) {
	return file_internal_orchestration_server_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeployResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeploymentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Deployment to follow
}

func (x *DeploymentStatusRequest) Reset() {
	*x = DeploymentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orchestration_server_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentStatusRequest) ProtoMessage() {}

func (x *DeploymentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orchestration_server_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentStatusRequest.ProtoReflect.Descriptor instead.
func (*DeploymentStatusRequest) Descriptor() ([]byte, []int) {
	return file_internal_orchestration_server_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeploymentStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeploymentStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *orchestration.Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *DeploymentStatusResponse) Reset() {
	*x = DeploymentStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orchestration_server_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentStatusResponse) ProtoMessage() {}

func (x *DeploymentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orchestration_server_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentStatusResponse.ProtoReflect.Descriptor instead.
func (*DeploymentStatusResponse) Descriptor() ([]byte, []int) {
	return file_internal_orchestration_server_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeploymentStatusResponse) GetEvent() *orchestration.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_internal_orchestration_server_proto_service_proto protoreflect.FileDescriptor

var file_internal_orchestration_server_proto_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x66, 0x6e, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0d, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x70,
	0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x30,
	0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x03, 0x61, 0x77, 0x73,
	0x12, 0x30, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x66, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x20,
	0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x18, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x32, 0xd4, 0x01, 0x0a, 0x14,
	0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x20,
	0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c,
	0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x72, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_orchestration_server_proto_service_proto_rawDescOnce sync.Once
	file_internal_orchestration_server_proto_service_proto_rawDescData = file_internal_orchestration_server_proto_service_proto_rawDesc
)

func file_internal_orchestration_server_proto_service_proto_rawDescGZIP() []byte {
	file_internal_orchestration_server_proto_service_proto_rawDescOnce.Do(func() {
		file_internal_orchestration_server_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_orchestration_server_proto_service_proto_rawDescData)
	})
	return file_internal_orchestration_server_proto_service_proto_rawDescData
}

var file_internal_orchestration_server_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_orchestration_server_proto_service_proto_goTypes = []interface{}{
	(*DeployRequest)(nil),            // 0: nsl.orchestration.DeployRequest
	(*DeployResponse)(nil),           // 1: nsl.orchestration.DeployResponse
	(*DeploymentStatusRequest)(nil),  // 2: nsl.orchestration.DeploymentStatusRequest
	(*DeploymentStatusResponse)(nil), // 3: nsl.orchestration.DeploymentStatusResponse
	(*schema.DeployPlan)(nil),        // 4: foundation.schema.DeployPlan
	(*aws.Conf)(nil),                 // 5: foundation.providers.aws.Conf
	(*fnapi.UserAuth)(nil),           // 6: nsl.internal.fnapi.UserAuth
	(*orchestration.Event)(nil),      // 7: foundation.schema.orchestration.Event
}
var file_internal_orchestration_server_proto_service_proto_depIdxs = []int32{
	4, // 0: nsl.orchestration.DeployRequest.plan:type_name -> foundation.schema.DeployPlan
	5, // 1: nsl.orchestration.DeployRequest.aws:type_name -> foundation.providers.aws.Conf
	6, // 2: nsl.orchestration.DeployRequest.auth:type_name -> nsl.internal.fnapi.UserAuth
	7, // 3: nsl.orchestration.DeploymentStatusResponse.event:type_name -> foundation.schema.orchestration.Event
	0, // 4: nsl.orchestration.OrchestrationService.Deploy:input_type -> nsl.orchestration.DeployRequest
	2, // 5: nsl.orchestration.OrchestrationService.DeploymentStatus:input_type -> nsl.orchestration.DeploymentStatusRequest
	1, // 6: nsl.orchestration.OrchestrationService.Deploy:output_type -> nsl.orchestration.DeployResponse
	3, // 7: nsl.orchestration.OrchestrationService.DeploymentStatus:output_type -> nsl.orchestration.DeploymentStatusResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_orchestration_server_proto_service_proto_init() }
func file_internal_orchestration_server_proto_service_proto_init() {
	if File_internal_orchestration_server_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_orchestration_server_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployRequest); i {
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
		file_internal_orchestration_server_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployResponse); i {
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
		file_internal_orchestration_server_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentStatusRequest); i {
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
		file_internal_orchestration_server_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentStatusResponse); i {
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
			RawDescriptor: file_internal_orchestration_server_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_orchestration_server_proto_service_proto_goTypes,
		DependencyIndexes: file_internal_orchestration_server_proto_service_proto_depIdxs,
		MessageInfos:      file_internal_orchestration_server_proto_service_proto_msgTypes,
	}.Build()
	File_internal_orchestration_server_proto_service_proto = out.File
	file_internal_orchestration_server_proto_service_proto_rawDesc = nil
	file_internal_orchestration_server_proto_service_proto_goTypes = nil
	file_internal_orchestration_server_proto_service_proto_depIdxs = nil
}
