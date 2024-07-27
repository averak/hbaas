// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/global_kvs.proto

package api

import (
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

type GlobalKVSServiceGetV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Criteria []*resource.KVSCriterion `protobuf:"bytes,1,rep,name=criteria,proto3" json:"criteria,omitempty"`
}

func (x *GlobalKVSServiceGetV1Request) Reset() {
	*x = GlobalKVSServiceGetV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_global_kvs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalKVSServiceGetV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalKVSServiceGetV1Request) ProtoMessage() {}

func (x *GlobalKVSServiceGetV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_kvs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalKVSServiceGetV1Request.ProtoReflect.Descriptor instead.
func (*GlobalKVSServiceGetV1Request) Descriptor() ([]byte, []int) {
	return file_api_global_kvs_proto_rawDescGZIP(), []int{0}
}

func (x *GlobalKVSServiceGetV1Request) GetCriteria() []*resource.KVSCriterion {
	if x != nil {
		return x.Criteria
	}
	return nil
}

type GlobalKVSServiceGetV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*resource.KVSEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *GlobalKVSServiceGetV1Response) Reset() {
	*x = GlobalKVSServiceGetV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_global_kvs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalKVSServiceGetV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalKVSServiceGetV1Response) ProtoMessage() {}

func (x *GlobalKVSServiceGetV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_kvs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalKVSServiceGetV1Response.ProtoReflect.Descriptor instead.
func (*GlobalKVSServiceGetV1Response) Descriptor() ([]byte, []int) {
	return file_api_global_kvs_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalKVSServiceGetV1Response) GetEntries() []*resource.KVSEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type GlobalKVSServiceSetV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*resource.KVSEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *GlobalKVSServiceSetV1Request) Reset() {
	*x = GlobalKVSServiceSetV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_global_kvs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalKVSServiceSetV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalKVSServiceSetV1Request) ProtoMessage() {}

func (x *GlobalKVSServiceSetV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_kvs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalKVSServiceSetV1Request.ProtoReflect.Descriptor instead.
func (*GlobalKVSServiceSetV1Request) Descriptor() ([]byte, []int) {
	return file_api_global_kvs_proto_rawDescGZIP(), []int{2}
}

func (x *GlobalKVSServiceSetV1Request) GetEntries() []*resource.KVSEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type GlobalKVSServiceSetV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GlobalKVSServiceSetV1Response) Reset() {
	*x = GlobalKVSServiceSetV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_global_kvs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalKVSServiceSetV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalKVSServiceSetV1Response) ProtoMessage() {}

func (x *GlobalKVSServiceSetV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_global_kvs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalKVSServiceSetV1Response.ProtoReflect.Descriptor instead.
func (*GlobalKVSServiceSetV1Response) Descriptor() ([]byte, []int) {
	return file_api_global_kvs_proto_rawDescGZIP(), []int{3}
}

var File_api_global_kvs_proto protoreflect.FileDescriptor

var file_api_global_kvs_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6b, 0x76, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x12, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6b, 0x76, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x52, 0x0a, 0x1c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4b, 0x56, 0x53,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x22, 0x4d, 0x0a, 0x1d, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b, 0x56, 0x53,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x4b, 0x56, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x4c, 0x0a, 0x1c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b, 0x56, 0x53, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4b,
	0x56, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x1f, 0x0a, 0x1d, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xb2, 0x01, 0x0a, 0x10, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b, 0x56, 0x53, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x56, 0x31, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b, 0x56, 0x53, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b,
	0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x53, 0x65, 0x74, 0x56, 0x31, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b, 0x56, 0x53, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b,
	0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6b, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x42, 0x0e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4b, 0x76, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x6b, 0x2f, 0x68, 0x62, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa,
	0x02, 0x03, 0x41, 0x70, 0x69, 0xca, 0x02, 0x03, 0x41, 0x70, 0x69, 0xe2, 0x02, 0x0f, 0x41, 0x70,
	0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x03,
	0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_global_kvs_proto_rawDescOnce sync.Once
	file_api_global_kvs_proto_rawDescData = file_api_global_kvs_proto_rawDesc
)

func file_api_global_kvs_proto_rawDescGZIP() []byte {
	file_api_global_kvs_proto_rawDescOnce.Do(func() {
		file_api_global_kvs_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_global_kvs_proto_rawDescData)
	})
	return file_api_global_kvs_proto_rawDescData
}

var file_api_global_kvs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_global_kvs_proto_goTypes = []any{
	(*GlobalKVSServiceGetV1Request)(nil),  // 0: api.GlobalKVSServiceGetV1Request
	(*GlobalKVSServiceGetV1Response)(nil), // 1: api.GlobalKVSServiceGetV1Response
	(*GlobalKVSServiceSetV1Request)(nil),  // 2: api.GlobalKVSServiceSetV1Request
	(*GlobalKVSServiceSetV1Response)(nil), // 3: api.GlobalKVSServiceSetV1Response
	(*resource.KVSCriterion)(nil),         // 4: resource.KVSCriterion
	(*resource.KVSEntry)(nil),             // 5: resource.KVSEntry
}
var file_api_global_kvs_proto_depIdxs = []int32{
	4, // 0: api.GlobalKVSServiceGetV1Request.criteria:type_name -> resource.KVSCriterion
	5, // 1: api.GlobalKVSServiceGetV1Response.entries:type_name -> resource.KVSEntry
	5, // 2: api.GlobalKVSServiceSetV1Request.entries:type_name -> resource.KVSEntry
	0, // 3: api.GlobalKVSService.GetV1:input_type -> api.GlobalKVSServiceGetV1Request
	2, // 4: api.GlobalKVSService.SetV1:input_type -> api.GlobalKVSServiceSetV1Request
	1, // 5: api.GlobalKVSService.GetV1:output_type -> api.GlobalKVSServiceGetV1Response
	3, // 6: api.GlobalKVSService.SetV1:output_type -> api.GlobalKVSServiceSetV1Response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_global_kvs_proto_init() }
func file_api_global_kvs_proto_init() {
	if File_api_global_kvs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_global_kvs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GlobalKVSServiceGetV1Request); i {
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
		file_api_global_kvs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GlobalKVSServiceGetV1Response); i {
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
		file_api_global_kvs_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GlobalKVSServiceSetV1Request); i {
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
		file_api_global_kvs_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GlobalKVSServiceSetV1Response); i {
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
			RawDescriptor: file_api_global_kvs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_global_kvs_proto_goTypes,
		DependencyIndexes: file_api_global_kvs_proto_depIdxs,
		MessageInfos:      file_api_global_kvs_proto_msgTypes,
	}.Build()
	File_api_global_kvs_proto = out.File
	file_api_global_kvs_proto_rawDesc = nil
	file_api_global_kvs_proto_goTypes = nil
	file_api_global_kvs_proto_depIdxs = nil
}