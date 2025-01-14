// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.1
// source: rms-torrent.proto

package rms_torrent

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

type Status int32

const (
	Status_Pending     Status = 0
	Status_Downloading Status = 1
	Status_Done        Status = 2
	Status_Failed      Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Pending",
		1: "Downloading",
		2: "Done",
		3: "Failed",
	}
	Status_value = map[string]int32{
		"Pending":     0,
		"Downloading": 1,
		"Done":        2,
		"Failed":      3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_rms_torrent_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_rms_torrent_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{0}
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Содержимое торрент-файла или магнет-ссылка
	What []byte `protobuf:"bytes,1,opt,name=what,proto3" json:"what,omitempty"`
	// Текстовое описание контента (нужно для отправки события)
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Флажок для игноирования ограничений на скорость скачивания
	Faster bool `protobuf:"varint,3,opt,name=faster,proto3" json:"faster,omitempty"`
	// Можно определить категорию (по сути - директория, куда будут сохраняться скачанные раздачи)
	Category string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	// Можно переопределить базовую директорию для сохранения данных
	ForceLocation *string `protobuf:"bytes,5,opt,name=forceLocation,proto3,oneof" json:"forceLocation,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_torrent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_torrent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadRequest) GetWhat() []byte {
	if x != nil {
		return x.What
	}
	return nil
}

func (x *DownloadRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DownloadRequest) GetFaster() bool {
	if x != nil {
		return x.Faster
	}
	return false
}

func (x *DownloadRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *DownloadRequest) GetForceLocation() string {
	if x != nil && x.ForceLocation != nil {
		return *x.ForceLocation
	}
	return ""
}

type DownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID торрента
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Пути к файлам, которые входят в раздачу
	Files []string `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	// Заголовок торрента
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_torrent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rms_torrent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse.ProtoReflect.Descriptor instead.
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DownloadResponse) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *DownloadResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetTorrentInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTorrentInfoRequest) Reset() {
	*x = GetTorrentInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_torrent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTorrentInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTorrentInfoRequest) ProtoMessage() {}

func (x *GetTorrentInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_torrent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTorrentInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTorrentInfoRequest) Descriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{2}
}

func (x *GetTorrentInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TorrentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор загрузки
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Название раздачи
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Состояние
	Status Status `protobuf:"varint,3,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	// Сколько % данных загружено
	Progress float32 `protobuf:"fixed32,4,opt,name=progress,proto3" json:"progress,omitempty"`
	// Сколько осталось времени до заверщения (time.Duration)
	RemainingTime int64 `protobuf:"varint,5,opt,name=remainingTime,proto3" json:"remainingTime,omitempty"`
	// Сколько Мб необходимо загрузить
	SizeMB uint64 `protobuf:"varint,6,opt,name=sizeMB,proto3" json:"sizeMB,omitempty"`
}

func (x *TorrentInfo) Reset() {
	*x = TorrentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_torrent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TorrentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TorrentInfo) ProtoMessage() {}

func (x *TorrentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rms_torrent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TorrentInfo.ProtoReflect.Descriptor instead.
func (*TorrentInfo) Descriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{3}
}

func (x *TorrentInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TorrentInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TorrentInfo) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Pending
}

func (x *TorrentInfo) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *TorrentInfo) GetRemainingTime() int64 {
	if x != nil {
		return x.RemainingTime
	}
	return 0
}

func (x *TorrentInfo) GetSizeMB() uint64 {
	if x != nil {
		return x.SizeMB
	}
	return 0
}

type GetTorrentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeDoneTorrents bool `protobuf:"varint,1,opt,name=includeDoneTorrents,proto3" json:"includeDoneTorrents,omitempty"`
}

func (x *GetTorrentsRequest) Reset() {
	*x = GetTorrentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_torrent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTorrentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTorrentsRequest) ProtoMessage() {}

func (x *GetTorrentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_torrent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTorrentsRequest.ProtoReflect.Descriptor instead.
func (*GetTorrentsRequest) Descriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{4}
}

func (x *GetTorrentsRequest) GetIncludeDoneTorrents() bool {
	if x != nil {
		return x.IncludeDoneTorrents
	}
	return false
}

type GetTorrentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Torrents []*TorrentInfo `protobuf:"bytes,1,rep,name=torrents,proto3" json:"torrents,omitempty"`
}

func (x *GetTorrentsResponse) Reset() {
	*x = GetTorrentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_torrent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTorrentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTorrentsResponse) ProtoMessage() {}

func (x *GetTorrentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rms_torrent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTorrentsResponse.ProtoReflect.Descriptor instead.
func (*GetTorrentsResponse) Descriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{5}
}

func (x *GetTorrentsResponse) GetTorrents() []*TorrentInfo {
	if x != nil {
		return x.Torrents
	}
	return nil
}

type RemoveTorrentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveTorrentRequest) Reset() {
	*x = RemoveTorrentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_torrent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTorrentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTorrentRequest) ProtoMessage() {}

func (x *RemoveTorrentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_torrent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTorrentRequest.ProtoReflect.Descriptor instead.
func (*RemoveTorrentRequest) Descriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveTorrentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpPriorityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpPriorityRequest) Reset() {
	*x = UpPriorityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_torrent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpPriorityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpPriorityRequest) ProtoMessage() {}

func (x *UpPriorityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_torrent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpPriorityRequest.ProtoReflect.Descriptor instead.
func (*UpPriorityRequest) Descriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{7}
}

func (x *UpPriorityRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TorrentSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DownloadLimit uint64 `protobuf:"varint,1,opt,name=DownloadLimit,proto3" json:"DownloadLimit,omitempty"`
	UploadLimit   uint64 `protobuf:"varint,2,opt,name=UploadLimit,proto3" json:"UploadLimit,omitempty"`
}

func (x *TorrentSettings) Reset() {
	*x = TorrentSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_torrent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TorrentSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TorrentSettings) ProtoMessage() {}

func (x *TorrentSettings) ProtoReflect() protoreflect.Message {
	mi := &file_rms_torrent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TorrentSettings.ProtoReflect.Descriptor instead.
func (*TorrentSettings) Descriptor() ([]byte, []int) {
	return file_rms_torrent_proto_rawDescGZIP(), []int{8}
}

func (x *TorrentSettings) GetDownloadLimit() uint64 {
	if x != nil {
		return x.DownloadLimit
	}
	return 0
}

func (x *TorrentSettings) GetUploadLimit() uint64 {
	if x != nil {
		return x.UploadLimit
	}
	return 0
}

var File_rms_torrent_proto protoreflect.FileDescriptor

var file_rms_torrent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x6d, 0x73, 0x2d, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb8, 0x01, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x77, 0x68, 0x61, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x29,
	0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x10, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x42, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73,
	0x69, 0x7a, 0x65, 0x4d, 0x42, 0x22, 0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3f, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x26,
	0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x55, 0x70, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x0f, 0x54,
	0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24,
	0x0a, 0x0d, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x2a, 0x3c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x10, 0x03, 0x32, 0x9b, 0x03, 0x0a, 0x0a, 0x52, 0x6d, 0x73, 0x54, 0x6f, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x10, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0a, 0x55, 0x70, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x53, 0x65, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x10, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x72, 0x6d, 0x73, 0x2d, 0x74, 0x6f, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rms_torrent_proto_rawDescOnce sync.Once
	file_rms_torrent_proto_rawDescData = file_rms_torrent_proto_rawDesc
)

func file_rms_torrent_proto_rawDescGZIP() []byte {
	file_rms_torrent_proto_rawDescOnce.Do(func() {
		file_rms_torrent_proto_rawDescData = protoimpl.X.CompressGZIP(file_rms_torrent_proto_rawDescData)
	})
	return file_rms_torrent_proto_rawDescData
}

var file_rms_torrent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rms_torrent_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rms_torrent_proto_goTypes = []interface{}{
	(Status)(0),                   // 0: Status
	(*DownloadRequest)(nil),       // 1: DownloadRequest
	(*DownloadResponse)(nil),      // 2: DownloadResponse
	(*GetTorrentInfoRequest)(nil), // 3: GetTorrentInfoRequest
	(*TorrentInfo)(nil),           // 4: TorrentInfo
	(*GetTorrentsRequest)(nil),    // 5: GetTorrentsRequest
	(*GetTorrentsResponse)(nil),   // 6: GetTorrentsResponse
	(*RemoveTorrentRequest)(nil),  // 7: RemoveTorrentRequest
	(*UpPriorityRequest)(nil),     // 8: UpPriorityRequest
	(*TorrentSettings)(nil),       // 9: TorrentSettings
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_rms_torrent_proto_depIdxs = []int32{
	0,  // 0: TorrentInfo.status:type_name -> Status
	4,  // 1: GetTorrentsResponse.torrents:type_name -> TorrentInfo
	1,  // 2: RmsTorrent.Download:input_type -> DownloadRequest
	3,  // 3: RmsTorrent.GetTorrentInfo:input_type -> GetTorrentInfoRequest
	5,  // 4: RmsTorrent.GetTorrents:input_type -> GetTorrentsRequest
	7,  // 5: RmsTorrent.RemoveTorrent:input_type -> RemoveTorrentRequest
	8,  // 6: RmsTorrent.UpPriority:input_type -> UpPriorityRequest
	10, // 7: RmsTorrent.GetSettings:input_type -> google.protobuf.Empty
	9,  // 8: RmsTorrent.SetSettings:input_type -> TorrentSettings
	2,  // 9: RmsTorrent.Download:output_type -> DownloadResponse
	4,  // 10: RmsTorrent.GetTorrentInfo:output_type -> TorrentInfo
	6,  // 11: RmsTorrent.GetTorrents:output_type -> GetTorrentsResponse
	10, // 12: RmsTorrent.RemoveTorrent:output_type -> google.protobuf.Empty
	10, // 13: RmsTorrent.UpPriority:output_type -> google.protobuf.Empty
	9,  // 14: RmsTorrent.GetSettings:output_type -> TorrentSettings
	10, // 15: RmsTorrent.SetSettings:output_type -> google.protobuf.Empty
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_rms_torrent_proto_init() }
func file_rms_torrent_proto_init() {
	if File_rms_torrent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rms_torrent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_rms_torrent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadResponse); i {
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
		file_rms_torrent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTorrentInfoRequest); i {
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
		file_rms_torrent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TorrentInfo); i {
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
		file_rms_torrent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTorrentsRequest); i {
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
		file_rms_torrent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTorrentsResponse); i {
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
		file_rms_torrent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTorrentRequest); i {
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
		file_rms_torrent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpPriorityRequest); i {
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
		file_rms_torrent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TorrentSettings); i {
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
	file_rms_torrent_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rms_torrent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rms_torrent_proto_goTypes,
		DependencyIndexes: file_rms_torrent_proto_depIdxs,
		EnumInfos:         file_rms_torrent_proto_enumTypes,
		MessageInfos:      file_rms_torrent_proto_msgTypes,
	}.Build()
	File_rms_torrent_proto = out.File
	file_rms_torrent_proto_rawDesc = nil
	file_rms_torrent_proto_goTypes = nil
	file_rms_torrent_proto_depIdxs = nil
}
