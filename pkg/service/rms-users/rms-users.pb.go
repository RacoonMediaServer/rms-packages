// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rms-users.proto

package rms_users

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

type GetPermissionsResponse_Permissions int32

const (
	// поиск через rms-media-discovery
	GetPermissionsResponse_Search GetPermissionsResponse_Permissions = 0
	// подключение к боту
	GetPermissionsResponse_ConnectingToTheBot GetPermissionsResponse_Permissions = 1
	// управление аккаунтами внешних систем
	GetPermissionsResponse_AccountManagement GetPermissionsResponse_Permissions = 2
)

// Enum value maps for GetPermissionsResponse_Permissions.
var (
	GetPermissionsResponse_Permissions_name = map[int32]string{
		0: "Search",
		1: "ConnectingToTheBot",
		2: "AccountManagement",
	}
	GetPermissionsResponse_Permissions_value = map[string]int32{
		"Search":             0,
		"ConnectingToTheBot": 1,
		"AccountManagement":  2,
	}
)

func (x GetPermissionsResponse_Permissions) Enum() *GetPermissionsResponse_Permissions {
	p := new(GetPermissionsResponse_Permissions)
	*p = x
	return p
}

func (x GetPermissionsResponse_Permissions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPermissionsResponse_Permissions) Descriptor() protoreflect.EnumDescriptor {
	return file_rms_users_proto_enumTypes[0].Descriptor()
}

func (GetPermissionsResponse_Permissions) Type() protoreflect.EnumType {
	return &file_rms_users_proto_enumTypes[0]
}

func (x GetPermissionsResponse_Permissions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPermissionsResponse_Permissions.Descriptor instead.
func (GetPermissionsResponse_Permissions) EnumDescriptor() ([]byte, []int) {
	return file_rms_users_proto_rawDescGZIP(), []int{1, 0}
}

type GetPermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// содержимое X-Token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetPermissionsRequest) Reset() {
	*x = GetPermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionsRequest) ProtoMessage() {}

func (x *GetPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionsRequest.ProtoReflect.Descriptor instead.
func (*GetPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_rms_users_proto_rawDescGZIP(), []int{0}
}

func (x *GetPermissionsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Perms []GetPermissionsResponse_Permissions `protobuf:"varint,1,rep,packed,name=perms,proto3,enum=GetPermissionsResponse_Permissions" json:"perms,omitempty"`
}

func (x *GetPermissionsResponse) Reset() {
	*x = GetPermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionsResponse) ProtoMessage() {}

func (x *GetPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rms_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionsResponse.ProtoReflect.Descriptor instead.
func (*GetPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_rms_users_proto_rawDescGZIP(), []int{1}
}

func (x *GetPermissionsResponse) GetPerms() []GetPermissionsResponse_Permissions {
	if x != nil {
		return x.Perms
	}
	return nil
}

var File_rms_users_proto protoreflect.FileDescriptor

var file_rms_users_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x6d, 0x73, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x9d, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x70,
	0x65, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x05, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x22, 0x48, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x6f, 0x54, 0x68, 0x65, 0x42, 0x6f, 0x74, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02,
	0x32, 0x4d, 0x0a, 0x08, 0x52, 0x6d, 0x73, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x72, 0x6d, 0x73, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rms_users_proto_rawDescOnce sync.Once
	file_rms_users_proto_rawDescData = file_rms_users_proto_rawDesc
)

func file_rms_users_proto_rawDescGZIP() []byte {
	file_rms_users_proto_rawDescOnce.Do(func() {
		file_rms_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_rms_users_proto_rawDescData)
	})
	return file_rms_users_proto_rawDescData
}

var file_rms_users_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rms_users_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rms_users_proto_goTypes = []interface{}{
	(GetPermissionsResponse_Permissions)(0), // 0: GetPermissionsResponse.Permissions
	(*GetPermissionsRequest)(nil),           // 1: GetPermissionsRequest
	(*GetPermissionsResponse)(nil),          // 2: GetPermissionsResponse
}
var file_rms_users_proto_depIdxs = []int32{
	0, // 0: GetPermissionsResponse.perms:type_name -> GetPermissionsResponse.Permissions
	1, // 1: RmsUsers.GetPermissions:input_type -> GetPermissionsRequest
	2, // 2: RmsUsers.GetPermissions:output_type -> GetPermissionsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rms_users_proto_init() }
func file_rms_users_proto_init() {
	if File_rms_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rms_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPermissionsRequest); i {
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
		file_rms_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPermissionsResponse); i {
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
			RawDescriptor: file_rms_users_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rms_users_proto_goTypes,
		DependencyIndexes: file_rms_users_proto_depIdxs,
		EnumInfos:         file_rms_users_proto_enumTypes,
		MessageInfos:      file_rms_users_proto_msgTypes,
	}.Build()
	File_rms_users_proto = out.File
	file_rms_users_proto_rawDesc = nil
	file_rms_users_proto_goTypes = nil
	file_rms_users_proto_depIdxs = nil
}