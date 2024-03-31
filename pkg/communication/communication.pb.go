// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: communication.proto

package communication

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

// Type of Message
type MessageType int32

const (
	// Usual communicating messages between user and device
	MessageType_Interaction MessageType = 0
	// Message sends by device for requesting linkage code
	MessageType_AcquiringCode MessageType = 1
	// Message sends by server for provide linkage code
	MessageType_LinkageCode MessageType = 2
	// Отсоединить пользователя от текущего устройства
	MessageType_UnlinkUser MessageType = 3
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "Interaction",
		1: "AcquiringCode",
		2: "LinkageCode",
		3: "UnlinkUser",
	}
	MessageType_value = map[string]int32{
		"Interaction":   0,
		"AcquiringCode": 1,
		"LinkageCode":   2,
		"UnlinkUser":    3,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_communication_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_communication_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_communication_proto_rawDescGZIP(), []int{0}
}

// Different keyboard styles
type KeyboardStyle int32

const (
	KeyboardStyle_Chat    KeyboardStyle = 0
	KeyboardStyle_Message KeyboardStyle = 1
)

// Enum value maps for KeyboardStyle.
var (
	KeyboardStyle_name = map[int32]string{
		0: "Chat",
		1: "Message",
	}
	KeyboardStyle_value = map[string]int32{
		"Chat":    0,
		"Message": 1,
	}
)

func (x KeyboardStyle) Enum() *KeyboardStyle {
	p := new(KeyboardStyle)
	*p = x
	return p
}

func (x KeyboardStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyboardStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_communication_proto_enumTypes[1].Descriptor()
}

func (KeyboardStyle) Type() protoreflect.EnumType {
	return &file_communication_proto_enumTypes[1]
}

func (x KeyboardStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyboardStyle.Descriptor instead.
func (KeyboardStyle) EnumDescriptor() ([]byte, []int) {
	return file_communication_proto_rawDescGZIP(), []int{1}
}

type Attachment_Type int32

const (
	Attachment_PhotoURL Attachment_Type = 0
	Attachment_Photo    Attachment_Type = 1
	Attachment_Video    Attachment_Type = 2
	Attachment_Document Attachment_Type = 3
	Attachment_Audio    Attachment_Type = 4
	Attachment_Voice    Attachment_Type = 5
)

// Enum value maps for Attachment_Type.
var (
	Attachment_Type_name = map[int32]string{
		0: "PhotoURL",
		1: "Photo",
		2: "Video",
		3: "Document",
		4: "Audio",
		5: "Voice",
	}
	Attachment_Type_value = map[string]int32{
		"PhotoURL": 0,
		"Photo":    1,
		"Video":    2,
		"Document": 3,
		"Audio":    4,
		"Voice":    5,
	}
)

func (x Attachment_Type) Enum() *Attachment_Type {
	p := new(Attachment_Type)
	*p = x
	return p
}

func (x Attachment_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Attachment_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_communication_proto_enumTypes[2].Descriptor()
}

func (Attachment_Type) Type() protoreflect.EnumType {
	return &file_communication_proto_enumTypes[2]
}

func (x Attachment_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Attachment_Type.Descriptor instead.
func (Attachment_Type) EnumDescriptor() ([]byte, []int) {
	return file_communication_proto_rawDescGZIP(), []int{2, 0}
}

// Telegram message from user to device
type UserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=communication.MessageType" json:"type,omitempty"`
	Text       string      `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	User       int32       `protobuf:"varint,4,opt,name=user,proto3" json:"user,omitempty"`
	Attachment *Attachment `protobuf:"bytes,5,opt,name=attachment,proto3,oneof" json:"attachment,omitempty"`
}

func (x *UserMessage) Reset() {
	*x = UserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMessage) ProtoMessage() {}

func (x *UserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_communication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMessage.ProtoReflect.Descriptor instead.
func (*UserMessage) Descriptor() ([]byte, []int) {
	return file_communication_proto_rawDescGZIP(), []int{0}
}

func (x *UserMessage) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_Interaction
}

func (x *UserMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UserMessage) GetUser() int32 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *UserMessage) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

// Button defines Telegram chat button, associated with message
type Button struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *Button) Reset() {
	*x = Button{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Button) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Button) ProtoMessage() {}

func (x *Button) ProtoReflect() protoreflect.Message {
	mi := &file_communication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Button.ProtoReflect.Descriptor instead.
func (*Button) Descriptor() ([]byte, []int) {
	return file_communication_proto_rawDescGZIP(), []int{1}
}

func (x *Button) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Button) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

// Attachment provide ability to attach media data to message
type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     Attachment_Type `protobuf:"varint,1,opt,name=type,proto3,enum=communication.Attachment_Type" json:"type,omitempty"`
	MimeType string          `protobuf:"bytes,2,opt,name=mimeType,proto3" json:"mimeType,omitempty"`
	Content  []byte          `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_communication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_communication_proto_rawDescGZIP(), []int{2}
}

func (x *Attachment) GetType() Attachment_Type {
	if x != nil {
		return x.Type
	}
	return Attachment_PhotoURL
}

func (x *Attachment) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Attachment) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// Message from Telegram bot to user
type BotMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          MessageType   `protobuf:"varint,1,opt,name=type,proto3,enum=communication.MessageType" json:"type,omitempty"`
	Text          string        `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Buttons       []*Button     `protobuf:"bytes,4,rep,name=buttons,proto3" json:"buttons,omitempty"`
	KeyboardStyle KeyboardStyle `protobuf:"varint,5,opt,name=keyboardStyle,proto3,enum=communication.KeyboardStyle" json:"keyboardStyle,omitempty"`
	Attachment    *Attachment   `protobuf:"bytes,6,opt,name=attachment,proto3,oneof" json:"attachment,omitempty"`
	User          int32         `protobuf:"varint,7,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *BotMessage) Reset() {
	*x = BotMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotMessage) ProtoMessage() {}

func (x *BotMessage) ProtoReflect() protoreflect.Message {
	mi := &file_communication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotMessage.ProtoReflect.Descriptor instead.
func (*BotMessage) Descriptor() ([]byte, []int) {
	return file_communication_proto_rawDescGZIP(), []int{3}
}

func (x *BotMessage) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_Interaction
}

func (x *BotMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *BotMessage) GetButtons() []*Button {
	if x != nil {
		return x.Buttons
	}
	return nil
}

func (x *BotMessage) GetKeyboardStyle() KeyboardStyle {
	if x != nil {
		return x.KeyboardStyle
	}
	return KeyboardStyle_Chat
}

func (x *BotMessage) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *BotMessage) GetUser() int32 {
	if x != nil {
		return x.User
	}
	return 0
}

var File_communication_proto protoreflect.FileDescriptor

var file_communication_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb4, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0a,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x06, 0x42,
	0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0xc6, 0x01, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x4e,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55,
	0x52, 0x4c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x10, 0x05, 0x22, 0xa8,
	0x02, 0x0a, 0x0a, 0x42, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x2f, 0x0a, 0x07, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52, 0x07, 0x62, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x73, 0x12, 0x42, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74,
	0x79, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x52, 0x0a, 0x0b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x63, 0x71,
	0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x4c, 0x69, 0x6e, 0x6b, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x10, 0x03, 0x2a, 0x26, 0x0a,
	0x0d, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x10, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x63, 0x6f, 0x6f, 0x6e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x6d, 0x73, 0x2d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communication_proto_rawDescOnce sync.Once
	file_communication_proto_rawDescData = file_communication_proto_rawDesc
)

func file_communication_proto_rawDescGZIP() []byte {
	file_communication_proto_rawDescOnce.Do(func() {
		file_communication_proto_rawDescData = protoimpl.X.CompressGZIP(file_communication_proto_rawDescData)
	})
	return file_communication_proto_rawDescData
}

var file_communication_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_communication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_communication_proto_goTypes = []interface{}{
	(MessageType)(0),     // 0: communication.MessageType
	(KeyboardStyle)(0),   // 1: communication.KeyboardStyle
	(Attachment_Type)(0), // 2: communication.Attachment.Type
	(*UserMessage)(nil),  // 3: communication.UserMessage
	(*Button)(nil),       // 4: communication.Button
	(*Attachment)(nil),   // 5: communication.Attachment
	(*BotMessage)(nil),   // 6: communication.BotMessage
}
var file_communication_proto_depIdxs = []int32{
	0, // 0: communication.UserMessage.type:type_name -> communication.MessageType
	5, // 1: communication.UserMessage.attachment:type_name -> communication.Attachment
	2, // 2: communication.Attachment.type:type_name -> communication.Attachment.Type
	0, // 3: communication.BotMessage.type:type_name -> communication.MessageType
	4, // 4: communication.BotMessage.buttons:type_name -> communication.Button
	1, // 5: communication.BotMessage.keyboardStyle:type_name -> communication.KeyboardStyle
	5, // 6: communication.BotMessage.attachment:type_name -> communication.Attachment
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_communication_proto_init() }
func file_communication_proto_init() {
	if File_communication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMessage); i {
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
		file_communication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Button); i {
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
		file_communication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
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
		file_communication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotMessage); i {
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
	file_communication_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_communication_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_communication_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_communication_proto_goTypes,
		DependencyIndexes: file_communication_proto_depIdxs,
		EnumInfos:         file_communication_proto_enumTypes,
		MessageInfos:      file_communication_proto_msgTypes,
	}.Build()
	File_communication_proto = out.File
	file_communication_proto_rawDesc = nil
	file_communication_proto_goTypes = nil
	file_communication_proto_depIdxs = nil
}
