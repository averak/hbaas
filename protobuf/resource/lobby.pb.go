// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: resource/lobby.proto

package resource

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

type RoomType int32

const (
	RoomType_ROOM_TYPE_UNSPECIFIED RoomType = 0
	RoomType_ROOM_TYPE_PRIVATE     RoomType = 1
	RoomType_ROOM_TYPE_PUBLIC      RoomType = 2
)

// Enum value maps for RoomType.
var (
	RoomType_name = map[int32]string{
		0: "ROOM_TYPE_UNSPECIFIED",
		1: "ROOM_TYPE_PRIVATE",
		2: "ROOM_TYPE_PUBLIC",
	}
	RoomType_value = map[string]int32{
		"ROOM_TYPE_UNSPECIFIED": 0,
		"ROOM_TYPE_PRIVATE":     1,
		"ROOM_TYPE_PUBLIC":      2,
	}
)

func (x RoomType) Enum() *RoomType {
	p := new(RoomType)
	*p = x
	return p
}

func (x RoomType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomType) Descriptor() protoreflect.EnumDescriptor {
	return file_resource_lobby_proto_enumTypes[0].Descriptor()
}

func (RoomType) Type() protoreflect.EnumType {
	return &file_resource_lobby_proto_enumTypes[0]
}

func (x RoomType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomType.Descriptor instead.
func (RoomType) EnumDescriptor() ([]byte, []int) {
	return file_resource_lobby_proto_rawDescGZIP(), []int{0}
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId      string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	OwnerUserId string   `protobuf:"bytes,2,opt,name=owner_user_id,json=ownerUserId,proto3" json:"owner_user_id,omitempty"`
	Type        RoomType `protobuf:"varint,3,opt,name=type,proto3,enum=resource.RoomType" json:"type,omitempty"`
	MaxCapacity int64    `protobuf:"varint,4,opt,name=max_capacity,json=maxCapacity,proto3" json:"max_capacity,omitempty"`
	Secret      string   `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
	Details     []byte   `protobuf:"bytes,6,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_lobby_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_resource_lobby_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_resource_lobby_proto_rawDescGZIP(), []int{0}
}

func (x *Room) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Room) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *Room) GetType() RoomType {
	if x != nil {
		return x.Type
	}
	return RoomType_ROOM_TYPE_UNSPECIFIED
}

func (x *Room) GetMaxCapacity() int64 {
	if x != nil {
		return x.MaxCapacity
	}
	return 0
}

func (x *Room) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Room) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_resource_lobby_proto protoreflect.FileDescriptor

var file_resource_lobby_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x62, 0x62, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0xc0, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x2a, 0x52, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x15, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x4f,
	0x4f, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x02, 0x42, 0x85, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0a, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x6b, 0x2f, 0x68, 0x62, 0x61, 0x61, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0xca, 0x02, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xe2, 0x02, 0x14,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_lobby_proto_rawDescOnce sync.Once
	file_resource_lobby_proto_rawDescData = file_resource_lobby_proto_rawDesc
)

func file_resource_lobby_proto_rawDescGZIP() []byte {
	file_resource_lobby_proto_rawDescOnce.Do(func() {
		file_resource_lobby_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_lobby_proto_rawDescData)
	})
	return file_resource_lobby_proto_rawDescData
}

var file_resource_lobby_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resource_lobby_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resource_lobby_proto_goTypes = []any{
	(RoomType)(0), // 0: resource.RoomType
	(*Room)(nil),  // 1: resource.Room
}
var file_resource_lobby_proto_depIdxs = []int32{
	0, // 0: resource.Room.type:type_name -> resource.RoomType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resource_lobby_proto_init() }
func file_resource_lobby_proto_init() {
	if File_resource_lobby_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_lobby_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Room); i {
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
			RawDescriptor: file_resource_lobby_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_lobby_proto_goTypes,
		DependencyIndexes: file_resource_lobby_proto_depIdxs,
		EnumInfos:         file_resource_lobby_proto_enumTypes,
		MessageInfos:      file_resource_lobby_proto_msgTypes,
	}.Build()
	File_resource_lobby_proto = out.File
	file_resource_lobby_proto_rawDesc = nil
	file_resource_lobby_proto_goTypes = nil
	file_resource_lobby_proto_depIdxs = nil
}
