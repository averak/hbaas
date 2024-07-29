// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: cloud_pubsub/pubsub.proto

package cloud_pubsub

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

type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED        EventType = 0
	EventType_EVENT_TYPE_BAAS_USER_DELETION EventType = 1
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNSPECIFIED",
		1: "EVENT_TYPE_BAAS_USER_DELETION",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED":        0,
		"EVENT_TYPE_BAAS_USER_DELETION": 1,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_pubsub_pubsub_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_cloud_pubsub_pubsub_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_pubsub_pubsub_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType EventType `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=cloud_pubsub.EventType" json:"event_type,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*Message_BaasUserDeletion
	Payload isMessage_Payload `protobuf_oneof:"payload"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_pubsub_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_pubsub_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_cloud_pubsub_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_EVENT_TYPE_UNSPECIFIED
}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Message) GetBaasUserDeletion() *BaasUserDeletion {
	if x, ok := x.GetPayload().(*Message_BaasUserDeletion); ok {
		return x.BaasUserDeletion
	}
	return nil
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_BaasUserDeletion struct {
	BaasUserDeletion *BaasUserDeletion `protobuf:"bytes,2,opt,name=baas_user_deletion,json=baasUserDeletion,proto3,oneof"`
}

func (*Message_BaasUserDeletion) isMessage_Payload() {}

type BaasUserDeletion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaasUserId string `protobuf:"bytes,1,opt,name=baas_user_id,json=baasUserId,proto3" json:"baas_user_id,omitempty"`
}

func (x *BaasUserDeletion) Reset() {
	*x = BaasUserDeletion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_pubsub_pubsub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaasUserDeletion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaasUserDeletion) ProtoMessage() {}

func (x *BaasUserDeletion) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_pubsub_pubsub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaasUserDeletion.ProtoReflect.Descriptor instead.
func (*BaasUserDeletion) Descriptor() ([]byte, []int) {
	return file_cloud_pubsub_pubsub_proto_rawDescGZIP(), []int{1}
}

func (x *BaasUserDeletion) GetBaasUserId() string {
	if x != nil {
		return x.BaasUserId
	}
	return ""
}

var File_cloud_pubsub_pubsub_proto protoreflect.FileDescriptor

var file_cloud_pubsub_pubsub_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x22, 0x9c, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a,
	0x12, 0x62, 0x61, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x42, 0x61, 0x61, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x62, 0x61, 0x61,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x34, 0x0a, 0x10, 0x42, 0x61, 0x61, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c,
	0x62, 0x61, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x61, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x4a,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x41, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0x9a, 0x01, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x42,
	0x0b, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x6b, 0x2f, 0x68, 0x62, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0xa2, 0x02, 0x03,
	0x43, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x75, 0x62, 0x73, 0x75,
	0x62, 0xca, 0x02, 0x0b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0xe2,
	0x02, 0x17, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_pubsub_pubsub_proto_rawDescOnce sync.Once
	file_cloud_pubsub_pubsub_proto_rawDescData = file_cloud_pubsub_pubsub_proto_rawDesc
)

func file_cloud_pubsub_pubsub_proto_rawDescGZIP() []byte {
	file_cloud_pubsub_pubsub_proto_rawDescOnce.Do(func() {
		file_cloud_pubsub_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_pubsub_pubsub_proto_rawDescData)
	})
	return file_cloud_pubsub_pubsub_proto_rawDescData
}

var file_cloud_pubsub_pubsub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_pubsub_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_pubsub_pubsub_proto_goTypes = []any{
	(EventType)(0),           // 0: cloud_pubsub.EventType
	(*Message)(nil),          // 1: cloud_pubsub.Message
	(*BaasUserDeletion)(nil), // 2: cloud_pubsub.BaasUserDeletion
}
var file_cloud_pubsub_pubsub_proto_depIdxs = []int32{
	0, // 0: cloud_pubsub.Message.event_type:type_name -> cloud_pubsub.EventType
	2, // 1: cloud_pubsub.Message.baas_user_deletion:type_name -> cloud_pubsub.BaasUserDeletion
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_pubsub_pubsub_proto_init() }
func file_cloud_pubsub_pubsub_proto_init() {
	if File_cloud_pubsub_pubsub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_pubsub_pubsub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Message); i {
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
		file_cloud_pubsub_pubsub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BaasUserDeletion); i {
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
	file_cloud_pubsub_pubsub_proto_msgTypes[0].OneofWrappers = []any{
		(*Message_BaasUserDeletion)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_pubsub_pubsub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_pubsub_pubsub_proto_goTypes,
		DependencyIndexes: file_cloud_pubsub_pubsub_proto_depIdxs,
		EnumInfos:         file_cloud_pubsub_pubsub_proto_enumTypes,
		MessageInfos:      file_cloud_pubsub_pubsub_proto_msgTypes,
	}.Build()
	File_cloud_pubsub_pubsub_proto = out.File
	file_cloud_pubsub_pubsub_proto_rawDesc = nil
	file_cloud_pubsub_pubsub_proto_goTypes = nil
	file_cloud_pubsub_pubsub_proto_depIdxs = nil
}
