// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/naming.proto

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

type Naming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithOrg                 string                         `protobuf:"bytes,1,opt,name=with_org,json=withOrg,proto3" json:"with_org,omitempty"`                                                   // Allocate with the following organization.
	AdditionalUserSpecified []*Naming_AdditionalDomainName `protobuf:"bytes,2,rep,name=additional_user_specified,json=additionalUserSpecified,proto3" json:"additional_user_specified,omitempty"` // If set, also configures the ingress with these domains.
	AdditionalTlsManaged    []*Naming_AdditionalDomainName `protobuf:"bytes,3,rep,name=additional_tls_managed,json=additionalTlsManaged,proto3" json:"additional_tls_managed,omitempty"`          // If set, both configures an ingress with these domains, and requests TLS certificate from Foundation.
	EnableNamespaceManaged  bool                           `protobuf:"varint,4,opt,name=enable_namespace_managed,json=enableNamespaceManaged,proto3" json:"enable_namespace_managed,omitempty"`   // If set, enables namespace managed domain names.
}

func (x *Naming) Reset() {
	*x = Naming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_naming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Naming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Naming) ProtoMessage() {}

func (x *Naming) ProtoReflect() protoreflect.Message {
	mi := &file_schema_naming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Naming.ProtoReflect.Descriptor instead.
func (*Naming) Descriptor() ([]byte, []int) {
	return file_schema_naming_proto_rawDescGZIP(), []int{0}
}

func (x *Naming) GetWithOrg() string {
	if x != nil {
		return x.WithOrg
	}
	return ""
}

func (x *Naming) GetAdditionalUserSpecified() []*Naming_AdditionalDomainName {
	if x != nil {
		return x.AdditionalUserSpecified
	}
	return nil
}

func (x *Naming) GetAdditionalTlsManaged() []*Naming_AdditionalDomainName {
	if x != nil {
		return x.AdditionalTlsManaged
	}
	return nil
}

func (x *Naming) GetEnableNamespaceManaged() bool {
	if x != nil {
		return x.EnableNamespaceManaged
	}
	return false
}

type ComputedNaming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source                   *Naming            `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	BaseDomain               string             `protobuf:"bytes,2,opt,name=base_domain,json=baseDomain,proto3" json:"base_domain,omitempty"`                                               // E.g. {org}.nscloud.dev
	TlsPassthroughBaseDomain string             `protobuf:"bytes,3,opt,name=tls_passthrough_base_domain,json=tlsPassthroughBaseDomain,proto3" json:"tls_passthrough_base_domain,omitempty"` // E.g. int-a.nscluster.cloud
	Managed                  Domain_ManagedType `protobuf:"varint,4,opt,name=managed,proto3,enum=foundation.schema.Domain_ManagedType" json:"managed,omitempty"`
	UpstreamTlsTermination   bool               `protobuf:"varint,5,opt,name=upstream_tls_termination,json=upstreamTlsTermination,proto3" json:"upstream_tls_termination,omitempty"` // If true, these addresses will be TLS terminated.
	DomainFragmentSuffix     string             `protobuf:"bytes,6,opt,name=domain_fragment_suffix,json=domainFragmentSuffix,proto3" json:"domain_fragment_suffix,omitempty"`        // If set, domain name generation will use this suffix rather than creating a new sub-domain.
	UseShortAlias            bool               `protobuf:"varint,7,opt,name=use_short_alias,json=useShortAlias,proto3" json:"use_short_alias,omitempty"`                            // Enable new endpoint alias generation.
	MainModuleName           string             `protobuf:"bytes,8,opt,name=main_module_name,json=mainModuleName,proto3" json:"main_module_name,omitempty"`
}

func (x *ComputedNaming) Reset() {
	*x = ComputedNaming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_naming_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedNaming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedNaming) ProtoMessage() {}

func (x *ComputedNaming) ProtoReflect() protoreflect.Message {
	mi := &file_schema_naming_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputedNaming.ProtoReflect.Descriptor instead.
func (*ComputedNaming) Descriptor() ([]byte, []int) {
	return file_schema_naming_proto_rawDescGZIP(), []int{1}
}

func (x *ComputedNaming) GetSource() *Naming {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ComputedNaming) GetBaseDomain() string {
	if x != nil {
		return x.BaseDomain
	}
	return ""
}

func (x *ComputedNaming) GetTlsPassthroughBaseDomain() string {
	if x != nil {
		return x.TlsPassthroughBaseDomain
	}
	return ""
}

func (x *ComputedNaming) GetManaged() Domain_ManagedType {
	if x != nil {
		return x.Managed
	}
	return Domain_MANAGED_UNKNOWN
}

func (x *ComputedNaming) GetUpstreamTlsTermination() bool {
	if x != nil {
		return x.UpstreamTlsTermination
	}
	return false
}

func (x *ComputedNaming) GetDomainFragmentSuffix() string {
	if x != nil {
		return x.DomainFragmentSuffix
	}
	return ""
}

func (x *ComputedNaming) GetUseShortAlias() bool {
	if x != nil {
		return x.UseShortAlias
	}
	return false
}

func (x *ComputedNaming) GetMainModuleName() string {
	if x != nil {
		return x.MainModuleName
	}
	return ""
}

type Naming_AdditionalDomainName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllocatedName string `protobuf:"bytes,1,opt,name=allocated_name,json=allocatedName,proto3" json:"allocated_name,omitempty"`
	Fqdn          string `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
}

func (x *Naming_AdditionalDomainName) Reset() {
	*x = Naming_AdditionalDomainName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_naming_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Naming_AdditionalDomainName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Naming_AdditionalDomainName) ProtoMessage() {}

func (x *Naming_AdditionalDomainName) ProtoReflect() protoreflect.Message {
	mi := &file_schema_naming_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Naming_AdditionalDomainName.ProtoReflect.Descriptor instead.
func (*Naming_AdditionalDomainName) Descriptor() ([]byte, []int) {
	return file_schema_naming_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Naming_AdditionalDomainName) GetAllocatedName() string {
	if x != nil {
		return x.AllocatedName
	}
	return ""
}

func (x *Naming_AdditionalDomainName) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

var File_schema_naming_proto protoreflect.FileDescriptor

var file_schema_naming_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6e, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03,
	0x0a, 0x06, 0x4e, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x69, 0x74, 0x68,
	0x5f, 0x6f, 0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x69, 0x74, 0x68,
	0x4f, 0x72, 0x67, 0x12, 0x6a, 0x0a, 0x19, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x17, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12,
	0x64, 0x0a, 0x16, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x6c,
	0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x14, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x6c, 0x73, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x1a,
	0x51, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71,
	0x64, 0x6e, 0x22, 0xa6, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x4e,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3d, 0x0a, 0x1b, 0x74, 0x6c, 0x73,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18,
	0x74, 0x6c, 0x73, 0x50, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x42, 0x61,
	0x73, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x75, 0x70, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x75, 0x70, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x6c, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x66, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x72, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65,
	0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x69,
	0x6e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_naming_proto_rawDescOnce sync.Once
	file_schema_naming_proto_rawDescData = file_schema_naming_proto_rawDesc
)

func file_schema_naming_proto_rawDescGZIP() []byte {
	file_schema_naming_proto_rawDescOnce.Do(func() {
		file_schema_naming_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_naming_proto_rawDescData)
	})
	return file_schema_naming_proto_rawDescData
}

var file_schema_naming_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schema_naming_proto_goTypes = []interface{}{
	(*Naming)(nil),                      // 0: foundation.schema.Naming
	(*ComputedNaming)(nil),              // 1: foundation.schema.ComputedNaming
	(*Naming_AdditionalDomainName)(nil), // 2: foundation.schema.Naming.AdditionalDomainName
	(Domain_ManagedType)(0),             // 3: foundation.schema.Domain.ManagedType
}
var file_schema_naming_proto_depIdxs = []int32{
	2, // 0: foundation.schema.Naming.additional_user_specified:type_name -> foundation.schema.Naming.AdditionalDomainName
	2, // 1: foundation.schema.Naming.additional_tls_managed:type_name -> foundation.schema.Naming.AdditionalDomainName
	0, // 2: foundation.schema.ComputedNaming.source:type_name -> foundation.schema.Naming
	3, // 3: foundation.schema.ComputedNaming.managed:type_name -> foundation.schema.Domain.ManagedType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_schema_naming_proto_init() }
func file_schema_naming_proto_init() {
	if File_schema_naming_proto != nil {
		return
	}
	file_schema_domain_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_naming_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Naming); i {
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
		file_schema_naming_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputedNaming); i {
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
		file_schema_naming_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Naming_AdditionalDomainName); i {
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
			RawDescriptor: file_schema_naming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_naming_proto_goTypes,
		DependencyIndexes: file_schema_naming_proto_depIdxs,
		MessageInfos:      file_schema_naming_proto_msgTypes,
	}.Build()
	File_schema_naming_proto = out.File
	file_schema_naming_proto_rawDesc = nil
	file_schema_naming_proto_goTypes = nil
	file_schema_naming_proto_depIdxs = nil
}
