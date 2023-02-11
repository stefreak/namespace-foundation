// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/runtime/config.proto

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

// The runtime configuration is generated from the deployment plan. It is
// injected into the server at runtime. Because RuntimeConfig is serialized
// as JSON, the field names have great significance and should be changed
// with care.
type RuntimeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environment *RuntimeConfig_Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	Current     *Server                    `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	// References to other servers that this server depends on.
	StackEntry []*Server `protobuf:"bytes,3,rep,name=stack_entry,json=stackEntry,proto3" json:"stack_entry,omitempty"`
}

func (x *RuntimeConfig) Reset() {
	*x = RuntimeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeConfig) ProtoMessage() {}

func (x *RuntimeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeConfig.ProtoReflect.Descriptor instead.
func (*RuntimeConfig) Descriptor() ([]byte, []int) {
	return file_schema_runtime_config_proto_rawDescGZIP(), []int{0}
}

func (x *RuntimeConfig) GetEnvironment() *RuntimeConfig_Environment {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *RuntimeConfig) GetCurrent() *Server {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *RuntimeConfig) GetStackEntry() []*Server {
	if x != nil {
		return x.StackEntry
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName      string                     `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	ModuleName       string                     `protobuf:"bytes,2,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	Port             []*Server_Port             `protobuf:"bytes,3,rep,name=port,proto3" json:"port,omitempty"`
	Service          []*Server_Service          `protobuf:"bytes,4,rep,name=service,proto3" json:"service,omitempty"`
	InternalEndpoint []*Server_InternalEndpoint `protobuf:"bytes,6,rep,name=internal_endpoint,json=internalEndpoint,proto3" json:"internal_endpoint,omitempty"`
	ImageRef         string                     `protobuf:"bytes,5,opt,name=image_ref,json=imageRef,proto3" json:"image_ref,omitempty"` // Only set for current.
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_schema_runtime_config_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *Server) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *Server) GetPort() []*Server_Port {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *Server) GetService() []*Server_Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Server) GetInternalEndpoint() []*Server_InternalEndpoint {
	if x != nil {
		return x.InternalEndpoint
	}
	return nil
}

func (x *Server) GetImageRef() string {
	if x != nil {
		return x.ImageRef
	}
	return ""
}

type BuildVCS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision    string `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
	CommitTime  string `protobuf:"bytes,2,opt,name=commit_time,json=commitTime,proto3" json:"commit_time,omitempty"`
	Uncommitted bool   `protobuf:"varint,3,opt,name=uncommitted,proto3" json:"uncommitted,omitempty"`
}

func (x *BuildVCS) Reset() {
	*x = BuildVCS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildVCS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildVCS) ProtoMessage() {}

func (x *BuildVCS) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildVCS.ProtoReflect.Descriptor instead.
func (*BuildVCS) Descriptor() ([]byte, []int) {
	return file_schema_runtime_config_proto_rawDescGZIP(), []int{2}
}

func (x *BuildVCS) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *BuildVCS) GetCommitTime() string {
	if x != nil {
		return x.CommitTime
	}
	return ""
}

func (x *BuildVCS) GetUncommitted() bool {
	if x != nil {
		return x.Uncommitted
	}
	return false
}

type RuntimeConfig_Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Empty if ephemeral is true.
	Purpose string `protobuf:"bytes,2,opt,name=purpose,proto3" json:"purpose,omitempty"`
	// Typically only set for tests. Signals that this environment is single-use and not meant to be user serviceable.
	Ephemeral bool `protobuf:"varint,3,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
}

func (x *RuntimeConfig_Environment) Reset() {
	*x = RuntimeConfig_Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeConfig_Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeConfig_Environment) ProtoMessage() {}

func (x *RuntimeConfig_Environment) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeConfig_Environment.ProtoReflect.Descriptor instead.
func (*RuntimeConfig_Environment) Descriptor() ([]byte, []int) {
	return file_schema_runtime_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RuntimeConfig_Environment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RuntimeConfig_Environment) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

func (x *RuntimeConfig_Environment) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

type Server_Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Server_Port) Reset() {
	*x = Server_Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Port) ProtoMessage() {}

func (x *Server_Port) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Port.ProtoReflect.Descriptor instead.
func (*Server_Port) Descriptor() ([]byte, []int) {
	return file_schema_runtime_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Server_Port) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Server_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner    string          `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`       // Package name.
	Name     string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`         // Scoped to package name.
	Endpoint string          `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"` // E.g. hostname:port.
	Ingress  *Server_Ingress `protobuf:"bytes,4,opt,name=ingress,proto3" json:"ingress,omitempty"`   // Computed ingress for this service.
}

func (x *Server_Service) Reset() {
	*x = Server_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Service) ProtoMessage() {}

func (x *Server_Service) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Service.ProtoReflect.Descriptor instead.
func (*Server_Service) Descriptor() ([]byte, []int) {
	return file_schema_runtime_config_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_Service) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Server_Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Server_Service) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Server_Service) GetIngress() *Server_Ingress {
	if x != nil {
		return x.Ingress
	}
	return nil
}

type Server_Ingress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain []*Server_Ingress_Domain `protobuf:"bytes,4,rep,name=domain,proto3" json:"domain,omitempty"`
}

func (x *Server_Ingress) Reset() {
	*x = Server_Ingress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Ingress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Ingress) ProtoMessage() {}

func (x *Server_Ingress) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Ingress.ProtoReflect.Descriptor instead.
func (*Server_Ingress) Descriptor() ([]byte, []int) {
	return file_schema_runtime_config_proto_rawDescGZIP(), []int{1, 2}
}

func (x *Server_Ingress) GetDomain() []*Server_Ingress_Domain {
	if x != nil {
		return x.Domain
	}
	return nil
}

type Server_InternalEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortName string   `protobuf:"bytes,1,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	Port     int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Kinds    []string `protobuf:"bytes,3,rep,name=kinds,proto3" json:"kinds,omitempty"`
}

func (x *Server_InternalEndpoint) Reset() {
	*x = Server_InternalEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_InternalEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_InternalEndpoint) ProtoMessage() {}

func (x *Server_InternalEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_InternalEndpoint.ProtoReflect.Descriptor instead.
func (*Server_InternalEndpoint) Descriptor() ([]byte, []int) {
	return file_schema_runtime_config_proto_rawDescGZIP(), []int{1, 3}
}

func (x *Server_InternalEndpoint) GetPortName() string {
	if x != nil {
		return x.PortName
	}
	return ""
}

func (x *Server_InternalEndpoint) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Server_InternalEndpoint) GetKinds() []string {
	if x != nil {
		return x.Kinds
	}
	return nil
}

type Server_Ingress_Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Includes the protocol, e.g. https://, and a port if needed.
	BaseUrl string `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
}

func (x *Server_Ingress_Domain) Reset() {
	*x = Server_Ingress_Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_runtime_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Ingress_Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Ingress_Domain) ProtoMessage() {}

func (x *Server_Ingress_Domain) ProtoReflect() protoreflect.Message {
	mi := &file_schema_runtime_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Ingress_Domain.ProtoReflect.Descriptor instead.
func (*Server_Ingress_Domain) Descriptor() ([]byte, []int) {
	return file_schema_runtime_config_proto_rawDescGZIP(), []int{1, 2, 0}
}

func (x *Server_Ingress_Domain) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

var File_schema_runtime_config_proto protoreflect.FileDescriptor

var file_schema_runtime_config_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xc3, 0x02, 0x0a, 0x0d, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x56, 0x0a, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x42, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x1a, 0x59, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x22, 0xe7,
	0x05, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f,
	0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x10, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x1a, 0x2e, 0x0a, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x94, 0x01, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x43,
	0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x1a, 0x78, 0x0a, 0x07, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x48,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x23, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x1a, 0x59, 0x0a,
	0x10, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x22, 0x69, 0x0a, 0x08, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x56, 0x43, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_runtime_config_proto_rawDescOnce sync.Once
	file_schema_runtime_config_proto_rawDescData = file_schema_runtime_config_proto_rawDesc
)

func file_schema_runtime_config_proto_rawDescGZIP() []byte {
	file_schema_runtime_config_proto_rawDescOnce.Do(func() {
		file_schema_runtime_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_runtime_config_proto_rawDescData)
	})
	return file_schema_runtime_config_proto_rawDescData
}

var file_schema_runtime_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_schema_runtime_config_proto_goTypes = []interface{}{
	(*RuntimeConfig)(nil),             // 0: foundation.schema.runtime.RuntimeConfig
	(*Server)(nil),                    // 1: foundation.schema.runtime.Server
	(*BuildVCS)(nil),                  // 2: foundation.schema.runtime.BuildVCS
	(*RuntimeConfig_Environment)(nil), // 3: foundation.schema.runtime.RuntimeConfig.Environment
	(*Server_Port)(nil),               // 4: foundation.schema.runtime.Server.Port
	(*Server_Service)(nil),            // 5: foundation.schema.runtime.Server.Service
	(*Server_Ingress)(nil),            // 6: foundation.schema.runtime.Server.Ingress
	(*Server_InternalEndpoint)(nil),   // 7: foundation.schema.runtime.Server.InternalEndpoint
	(*Server_Ingress_Domain)(nil),     // 8: foundation.schema.runtime.Server.Ingress.Domain
}
var file_schema_runtime_config_proto_depIdxs = []int32{
	3, // 0: foundation.schema.runtime.RuntimeConfig.environment:type_name -> foundation.schema.runtime.RuntimeConfig.Environment
	1, // 1: foundation.schema.runtime.RuntimeConfig.current:type_name -> foundation.schema.runtime.Server
	1, // 2: foundation.schema.runtime.RuntimeConfig.stack_entry:type_name -> foundation.schema.runtime.Server
	4, // 3: foundation.schema.runtime.Server.port:type_name -> foundation.schema.runtime.Server.Port
	5, // 4: foundation.schema.runtime.Server.service:type_name -> foundation.schema.runtime.Server.Service
	7, // 5: foundation.schema.runtime.Server.internal_endpoint:type_name -> foundation.schema.runtime.Server.InternalEndpoint
	6, // 6: foundation.schema.runtime.Server.Service.ingress:type_name -> foundation.schema.runtime.Server.Ingress
	8, // 7: foundation.schema.runtime.Server.Ingress.domain:type_name -> foundation.schema.runtime.Server.Ingress.Domain
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_schema_runtime_config_proto_init() }
func file_schema_runtime_config_proto_init() {
	if File_schema_runtime_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_runtime_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeConfig); i {
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
		file_schema_runtime_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_schema_runtime_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildVCS); i {
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
		file_schema_runtime_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeConfig_Environment); i {
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
		file_schema_runtime_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Port); i {
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
		file_schema_runtime_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Service); i {
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
		file_schema_runtime_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Ingress); i {
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
		file_schema_runtime_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_InternalEndpoint); i {
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
		file_schema_runtime_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Ingress_Domain); i {
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
			RawDescriptor: file_schema_runtime_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_runtime_config_proto_goTypes,
		DependencyIndexes: file_schema_runtime_config_proto_depIdxs,
		MessageInfos:      file_schema_runtime_config_proto_msgTypes,
	}.Build()
	File_schema_runtime_config_proto = out.File
	file_schema_runtime_config_proto_rawDesc = nil
	file_schema_runtime_config_proto_goTypes = nil
	file_schema_runtime_config_proto_depIdxs = nil
}
