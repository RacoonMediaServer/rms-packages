// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rms-notes.proto

package rms_notes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Заголовок заметки
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Текст заметки
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *AddNoteRequest) Reset() {
	*x = AddNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_notes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteRequest) ProtoMessage() {}

func (x *AddNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_notes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoteRequest.ProtoReflect.Descriptor instead.
func (*AddNoteRequest) Descriptor() ([]byte, []int) {
	return file_rms_notes_proto_rawDescGZIP(), []int{0}
}

func (x *AddNoteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddNoteRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type AddTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Текст задачи
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Дата задачи в формате yyyy-mm-dd
	DueDate *string `protobuf:"bytes,2,opt,name=dueDate,proto3,oneof" json:"dueDate,omitempty"`
}

func (x *AddTaskRequest) Reset() {
	*x = AddTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_notes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskRequest) ProtoMessage() {}

func (x *AddTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_notes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskRequest.ProtoReflect.Descriptor instead.
func (*AddTaskRequest) Descriptor() ([]byte, []int) {
	return file_rms_notes_proto_rawDescGZIP(), []int{1}
}

func (x *AddTaskRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AddTaskRequest) GetDueDate() string {
	if x != nil && x.DueDate != nil {
		return *x.DueDate
	}
	return ""
}

type SnoozeTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID задачи
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Дата задачи в формате yyyy-mm-dd
	DueDate *string `protobuf:"bytes,2,opt,name=dueDate,proto3,oneof" json:"dueDate,omitempty"`
}

func (x *SnoozeTaskRequest) Reset() {
	*x = SnoozeTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_notes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnoozeTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnoozeTaskRequest) ProtoMessage() {}

func (x *SnoozeTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_notes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnoozeTaskRequest.ProtoReflect.Descriptor instead.
func (*SnoozeTaskRequest) Descriptor() ([]byte, []int) {
	return file_rms_notes_proto_rawDescGZIP(), []int{2}
}

func (x *SnoozeTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SnoozeTaskRequest) GetDueDate() string {
	if x != nil && x.DueDate != nil {
		return *x.DueDate
	}
	return ""
}

type DoneTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID задачи
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DoneTaskRequest) Reset() {
	*x = DoneTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_notes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoneTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoneTaskRequest) ProtoMessage() {}

func (x *DoneTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_notes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoneTaskRequest.ProtoReflect.Descriptor instead.
func (*DoneTaskRequest) Descriptor() ([]byte, []int) {
	return file_rms_notes_proto_rawDescGZIP(), []int{3}
}

func (x *DoneTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NotesSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Путь к директории Obsidian
	Directory string `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	// Директория, в которую добавлять новые заметки
	NotesDirectory string `protobuf:"bytes,2,opt,name=notesDirectory,proto3" json:"notesDirectory,omitempty"`
	// Файл, в который добавляются задачи
	TasksFile string `protobuf:"bytes,3,opt,name=tasksFile,proto3" json:"tasksFile,omitempty"`
}

func (x *NotesSettings) Reset() {
	*x = NotesSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_notes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotesSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotesSettings) ProtoMessage() {}

func (x *NotesSettings) ProtoReflect() protoreflect.Message {
	mi := &file_rms_notes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotesSettings.ProtoReflect.Descriptor instead.
func (*NotesSettings) Descriptor() ([]byte, []int) {
	return file_rms_notes_proto_rawDescGZIP(), []int{4}
}

func (x *NotesSettings) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

func (x *NotesSettings) GetNotesDirectory() string {
	if x != nil {
		return x.NotesDirectory
	}
	return ""
}

func (x *NotesSettings) GetTasksFile() string {
	if x != nil {
		return x.TasksFile
	}
	return ""
}

var File_rms_notes_proto protoreflect.FileDescriptor

var file_rms_notes_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x6d, 0x73, 0x2d, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4f, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x1d, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x4e, 0x0a, 0x11, 0x53,
	0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x44,
	0x6f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x73,
	0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x0a,
	0x0e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x46, 0x69,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x46,
	0x69, 0x6c, 0x65, 0x32, 0xd0, 0x02, 0x0a, 0x08, 0x52, 0x6d, 0x73, 0x4e, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x32, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x41, 0x64,
	0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x0f, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0a, 0x53, 0x6e, 0x6f, 0x6f,
	0x7a, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x2e, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x34, 0x0a, 0x08, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x10,
	0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0e, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x35, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e,
	0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x72, 0x6d, 0x73,
	0x2d, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rms_notes_proto_rawDescOnce sync.Once
	file_rms_notes_proto_rawDescData = file_rms_notes_proto_rawDesc
)

func file_rms_notes_proto_rawDescGZIP() []byte {
	file_rms_notes_proto_rawDescOnce.Do(func() {
		file_rms_notes_proto_rawDescData = protoimpl.X.CompressGZIP(file_rms_notes_proto_rawDescData)
	})
	return file_rms_notes_proto_rawDescData
}

var file_rms_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rms_notes_proto_goTypes = []interface{}{
	(*AddNoteRequest)(nil),    // 0: AddNoteRequest
	(*AddTaskRequest)(nil),    // 1: AddTaskRequest
	(*SnoozeTaskRequest)(nil), // 2: SnoozeTaskRequest
	(*DoneTaskRequest)(nil),   // 3: DoneTaskRequest
	(*NotesSettings)(nil),     // 4: NotesSettings
	(*emptypb.Empty)(nil),     // 5: google.protobuf.Empty
}
var file_rms_notes_proto_depIdxs = []int32{
	0, // 0: RmsNotes.AddNote:input_type -> AddNoteRequest
	1, // 1: RmsNotes.AddTask:input_type -> AddTaskRequest
	2, // 2: RmsNotes.SnoozeTask:input_type -> SnoozeTaskRequest
	3, // 3: RmsNotes.DoneTask:input_type -> DoneTaskRequest
	5, // 4: RmsNotes.GetSettings:input_type -> google.protobuf.Empty
	4, // 5: RmsNotes.SetSettings:input_type -> NotesSettings
	5, // 6: RmsNotes.AddNote:output_type -> google.protobuf.Empty
	5, // 7: RmsNotes.AddTask:output_type -> google.protobuf.Empty
	5, // 8: RmsNotes.SnoozeTask:output_type -> google.protobuf.Empty
	5, // 9: RmsNotes.DoneTask:output_type -> google.protobuf.Empty
	4, // 10: RmsNotes.GetSettings:output_type -> NotesSettings
	5, // 11: RmsNotes.SetSettings:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rms_notes_proto_init() }
func file_rms_notes_proto_init() {
	if File_rms_notes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rms_notes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNoteRequest); i {
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
		file_rms_notes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskRequest); i {
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
		file_rms_notes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnoozeTaskRequest); i {
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
		file_rms_notes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoneTaskRequest); i {
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
		file_rms_notes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotesSettings); i {
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
	file_rms_notes_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_rms_notes_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rms_notes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rms_notes_proto_goTypes,
		DependencyIndexes: file_rms_notes_proto_depIdxs,
		MessageInfos:      file_rms_notes_proto_msgTypes,
	}.Build()
	File_rms_notes_proto = out.File
	file_rms_notes_proto_rawDesc = nil
	file_rms_notes_proto_goTypes = nil
	file_rms_notes_proto_depIdxs = nil
}
