// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/private_kvs.proto

package api

import (
	_ "github.com/averak/hbaas/protobuf/custom_option"
	resource "github.com/averak/hbaas/protobuf/resource"
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

type PrivateKVSServiceGetETagV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrivateKVSServiceGetETagV1Request) Reset() {
	*x = PrivateKVSServiceGetETagV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_private_kvs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKVSServiceGetETagV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKVSServiceGetETagV1Request) ProtoMessage() {}

func (x *PrivateKVSServiceGetETagV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_private_kvs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKVSServiceGetETagV1Request.ProtoReflect.Descriptor instead.
func (*PrivateKVSServiceGetETagV1Request) Descriptor() ([]byte, []int) {
	return file_api_private_kvs_proto_rawDescGZIP(), []int{0}
}

type PrivateKVSServiceGetETagV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etag string `protobuf:"bytes,1,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *PrivateKVSServiceGetETagV1Response) Reset() {
	*x = PrivateKVSServiceGetETagV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_private_kvs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKVSServiceGetETagV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKVSServiceGetETagV1Response) ProtoMessage() {}

func (x *PrivateKVSServiceGetETagV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_private_kvs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKVSServiceGetETagV1Response.ProtoReflect.Descriptor instead.
func (*PrivateKVSServiceGetETagV1Response) Descriptor() ([]byte, []int) {
	return file_api_private_kvs_proto_rawDescGZIP(), []int{1}
}

func (x *PrivateKVSServiceGetETagV1Response) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

type PrivateKVSServiceGetV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Criteria []*resource.KVSCriterion `protobuf:"bytes,1,rep,name=criteria,proto3" json:"criteria,omitempty"`
}

func (x *PrivateKVSServiceGetV1Request) Reset() {
	*x = PrivateKVSServiceGetV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_private_kvs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKVSServiceGetV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKVSServiceGetV1Request) ProtoMessage() {}

func (x *PrivateKVSServiceGetV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_private_kvs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKVSServiceGetV1Request.ProtoReflect.Descriptor instead.
func (*PrivateKVSServiceGetV1Request) Descriptor() ([]byte, []int) {
	return file_api_private_kvs_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateKVSServiceGetV1Request) GetCriteria() []*resource.KVSCriterion {
	if x != nil {
		return x.Criteria
	}
	return nil
}

type PrivateKVSServiceGetV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*resource.KVSEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	Etag    string               `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *PrivateKVSServiceGetV1Response) Reset() {
	*x = PrivateKVSServiceGetV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_private_kvs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKVSServiceGetV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKVSServiceGetV1Response) ProtoMessage() {}

func (x *PrivateKVSServiceGetV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_private_kvs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKVSServiceGetV1Response.ProtoReflect.Descriptor instead.
func (*PrivateKVSServiceGetV1Response) Descriptor() ([]byte, []int) {
	return file_api_private_kvs_proto_rawDescGZIP(), []int{3}
}

func (x *PrivateKVSServiceGetV1Response) GetEntries() []*resource.KVSEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *PrivateKVSServiceGetV1Response) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

type PrivateKVSServiceSetV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*resource.KVSEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	// 同時更新の競合を楽観ロックで防ぐためのバージョン管理情報です。
	// 最新の ETag を指定する必要があります。
	Etag string `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *PrivateKVSServiceSetV1Request) Reset() {
	*x = PrivateKVSServiceSetV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_private_kvs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKVSServiceSetV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKVSServiceSetV1Request) ProtoMessage() {}

func (x *PrivateKVSServiceSetV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_private_kvs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKVSServiceSetV1Request.ProtoReflect.Descriptor instead.
func (*PrivateKVSServiceSetV1Request) Descriptor() ([]byte, []int) {
	return file_api_private_kvs_proto_rawDescGZIP(), []int{4}
}

func (x *PrivateKVSServiceSetV1Request) GetEntries() []*resource.KVSEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *PrivateKVSServiceSetV1Request) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

type PrivateKVSServiceSetV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etag string `protobuf:"bytes,1,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *PrivateKVSServiceSetV1Response) Reset() {
	*x = PrivateKVSServiceSetV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_private_kvs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKVSServiceSetV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKVSServiceSetV1Response) ProtoMessage() {}

func (x *PrivateKVSServiceSetV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_private_kvs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKVSServiceSetV1Response.ProtoReflect.Descriptor instead.
func (*PrivateKVSServiceSetV1Response) Descriptor() ([]byte, []int) {
	return file_api_private_kvs_proto_rawDescGZIP(), []int{5}
}

func (x *PrivateKVSServiceSetV1Response) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

var File_api_private_kvs_proto protoreflect.FileDescriptor

var file_api_private_kvs_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x76,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x21, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6b, 0x76, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x21, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x56,
	0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x45, 0x54, 0x61, 0x67, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x22, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74,
	0x45, 0x54, 0x61, 0x67, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74,
	0x61, 0x67, 0x22, 0x53, 0x0a, 0x1d, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x56, 0x53,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x4b, 0x56, 0x53, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x63,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x22, 0x62, 0x0a, 0x1e, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4b, 0x56, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0x61, 0x0a, 0x1d, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4b, 0x56, 0x53, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74,
	0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0x34,
	0x0a, 0x1e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x65, 0x74, 0x61, 0x67, 0x32, 0xfc, 0x02, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x45, 0x54, 0x61, 0x67, 0x56, 0x31, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47,
	0x65, 0x74, 0x45, 0x54, 0x61, 0x67, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x56, 0x53,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x45, 0x54, 0x61, 0x67, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x56,
	0x31, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xb6, 0x01, 0x0a, 0x05, 0x53,
	0x65, 0x74, 0x56, 0x31, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64, 0x82,
	0xb5, 0x18, 0x60, 0x0a, 0x3a, 0x08, 0xd0, 0x0f, 0x10, 0x04, 0x1a, 0x33, 0x54, 0x68, 0x65, 0x20,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x62, 0x79, 0x74, 0x65, 0x73, 0x20, 0x6d, 0x75, 0x73, 0x74,
	0x20, 0x62, 0x65, 0x20, 0x6c, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x61, 0x6e, 0x20, 0x6f, 0x72,
	0x20, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x31, 0x4b, 0x69, 0x42, 0x2e, 0x0a,
	0x22, 0x08, 0xd2, 0x0f, 0x10, 0x04, 0x1a, 0x1b, 0x54, 0x68, 0x65, 0x20, 0x45, 0x54, 0x61, 0x67,
	0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x75, 0x70, 0x2d, 0x74, 0x6f, 0x2d, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x42, 0x6c, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0f,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x76, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x6b, 0x2f, 0x68, 0x62, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x03,
	0x41, 0x70, 0x69, 0xca, 0x02, 0x03, 0x41, 0x70, 0x69, 0xe2, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x03, 0x41, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_private_kvs_proto_rawDescOnce sync.Once
	file_api_private_kvs_proto_rawDescData = file_api_private_kvs_proto_rawDesc
)

func file_api_private_kvs_proto_rawDescGZIP() []byte {
	file_api_private_kvs_proto_rawDescOnce.Do(func() {
		file_api_private_kvs_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_private_kvs_proto_rawDescData)
	})
	return file_api_private_kvs_proto_rawDescData
}

var file_api_private_kvs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_private_kvs_proto_goTypes = []any{
	(*PrivateKVSServiceGetETagV1Request)(nil),  // 0: api.PrivateKVSServiceGetETagV1Request
	(*PrivateKVSServiceGetETagV1Response)(nil), // 1: api.PrivateKVSServiceGetETagV1Response
	(*PrivateKVSServiceGetV1Request)(nil),      // 2: api.PrivateKVSServiceGetV1Request
	(*PrivateKVSServiceGetV1Response)(nil),     // 3: api.PrivateKVSServiceGetV1Response
	(*PrivateKVSServiceSetV1Request)(nil),      // 4: api.PrivateKVSServiceSetV1Request
	(*PrivateKVSServiceSetV1Response)(nil),     // 5: api.PrivateKVSServiceSetV1Response
	(*resource.KVSCriterion)(nil),              // 6: resource.KVSCriterion
	(*resource.KVSEntry)(nil),                  // 7: resource.KVSEntry
}
var file_api_private_kvs_proto_depIdxs = []int32{
	6, // 0: api.PrivateKVSServiceGetV1Request.criteria:type_name -> resource.KVSCriterion
	7, // 1: api.PrivateKVSServiceGetV1Response.entries:type_name -> resource.KVSEntry
	7, // 2: api.PrivateKVSServiceSetV1Request.entries:type_name -> resource.KVSEntry
	0, // 3: api.PrivateKVSService.GetETagV1:input_type -> api.PrivateKVSServiceGetETagV1Request
	2, // 4: api.PrivateKVSService.GetV1:input_type -> api.PrivateKVSServiceGetV1Request
	4, // 5: api.PrivateKVSService.SetV1:input_type -> api.PrivateKVSServiceSetV1Request
	1, // 6: api.PrivateKVSService.GetETagV1:output_type -> api.PrivateKVSServiceGetETagV1Response
	3, // 7: api.PrivateKVSService.GetV1:output_type -> api.PrivateKVSServiceGetV1Response
	5, // 8: api.PrivateKVSService.SetV1:output_type -> api.PrivateKVSServiceSetV1Response
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_private_kvs_proto_init() }
func file_api_private_kvs_proto_init() {
	if File_api_private_kvs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_private_kvs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateKVSServiceGetETagV1Request); i {
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
		file_api_private_kvs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateKVSServiceGetETagV1Response); i {
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
		file_api_private_kvs_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateKVSServiceGetV1Request); i {
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
		file_api_private_kvs_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateKVSServiceGetV1Response); i {
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
		file_api_private_kvs_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateKVSServiceSetV1Request); i {
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
		file_api_private_kvs_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateKVSServiceSetV1Response); i {
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
			RawDescriptor: file_api_private_kvs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_private_kvs_proto_goTypes,
		DependencyIndexes: file_api_private_kvs_proto_depIdxs,
		MessageInfos:      file_api_private_kvs_proto_msgTypes,
	}.Build()
	File_api_private_kvs_proto = out.File
	file_api_private_kvs_proto_rawDesc = nil
	file_api_private_kvs_proto_goTypes = nil
	file_api_private_kvs_proto_depIdxs = nil
}
