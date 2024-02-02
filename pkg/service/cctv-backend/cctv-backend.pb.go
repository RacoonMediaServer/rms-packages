// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: cctv-backend.proto

package cctv_backend

import (
	video "github.com/RacoonMediaServer/rms-packages/pkg/video"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Номер камеры в системы (может быть нужен для генерации ID)
	CameraNo uint32 `protobuf:"varint,1,opt,name=cameraNo,proto3" json:"cameraNo,omitempty"`
	// Порекомендовать текстовый ID для потока
	AdviceId string `protobuf:"bytes,2,opt,name=adviceId,proto3" json:"adviceId,omitempty"`
	// URL птока
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *AddStreamRequest) Reset() {
	*x = AddStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStreamRequest) ProtoMessage() {}

func (x *AddStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStreamRequest.ProtoReflect.Descriptor instead.
func (*AddStreamRequest) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{0}
}

func (x *AddStreamRequest) GetCameraNo() uint32 {
	if x != nil {
		return x.CameraNo
	}
	return 0
}

func (x *AddStreamRequest) GetAdviceId() string {
	if x != nil {
		return x.AdviceId
	}
	return ""
}

func (x *AddStreamRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type AddStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор созданного потока
	StreamId string `protobuf:"bytes,1,opt,name=streamId,proto3" json:"streamId,omitempty"`
}

func (x *AddStreamResponse) Reset() {
	*x = AddStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStreamResponse) ProtoMessage() {}

func (x *AddStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStreamResponse.ProtoReflect.Descriptor instead.
func (*AddStreamResponse) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{1}
}

func (x *AddStreamResponse) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

type RemoveStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор потока
	StreamId string `protobuf:"bytes,1,opt,name=streamId,proto3" json:"streamId,omitempty"`
}

func (x *RemoveStreamRequest) Reset() {
	*x = RemoveStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveStreamRequest) ProtoMessage() {}

func (x *RemoveStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveStreamRequest.ProtoReflect.Descriptor instead.
func (*RemoveStreamRequest) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveStreamRequest) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

type GetStreamUriRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор потока
	StreamId string `protobuf:"bytes,1,opt,name=streamId,proto3" json:"streamId,omitempty"`
	// Транспорт для получения потока
	Transport *video.Transport `protobuf:"varint,2,opt,name=transport,proto3,enum=video.Transport,oneof" json:"transport,omitempty"`
}

func (x *GetStreamUriRequest) Reset() {
	*x = GetStreamUriRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamUriRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamUriRequest) ProtoMessage() {}

func (x *GetStreamUriRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamUriRequest.ProtoReflect.Descriptor instead.
func (*GetStreamUriRequest) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{3}
}

func (x *GetStreamUriRequest) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

func (x *GetStreamUriRequest) GetTransport() video.Transport {
	if x != nil && x.Transport != nil {
		return *x.Transport
	}
	return video.Transport(0)
}

type GetUriResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL на поток
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Реальный транспорт, для которого выдана ссылка
	Transport video.Transport `protobuf:"varint,2,opt,name=transport,proto3,enum=video.Transport" json:"transport,omitempty"`
}

func (x *GetUriResponse) Reset() {
	*x = GetUriResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUriResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUriResponse) ProtoMessage() {}

func (x *GetUriResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUriResponse.ProtoReflect.Descriptor instead.
func (*GetUriResponse) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{4}
}

func (x *GetUriResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *GetUriResponse) GetTransport() video.Transport {
	if x != nil {
		return x.Transport
	}
	return video.Transport(0)
}

type AddRecordingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор потока
	StreamId string `protobuf:"bytes,1,opt,name=streamId,proto3" json:"streamId,omitempty"`
	// Сколько дней хранить архив
	RotationDays uint32 `protobuf:"varint,2,opt,name=rotationDays,proto3" json:"rotationDays,omitempty"`
}

func (x *AddRecordingRequest) Reset() {
	*x = AddRecordingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecordingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecordingRequest) ProtoMessage() {}

func (x *AddRecordingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecordingRequest.ProtoReflect.Descriptor instead.
func (*AddRecordingRequest) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{5}
}

func (x *AddRecordingRequest) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

func (x *AddRecordingRequest) GetRotationDays() uint32 {
	if x != nil {
		return x.RotationDays
	}
	return 0
}

type AddRecordingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID записи
	RecordingId string `protobuf:"bytes,1,opt,name=recordingId,proto3" json:"recordingId,omitempty"`
}

func (x *AddRecordingResponse) Reset() {
	*x = AddRecordingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecordingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecordingResponse) ProtoMessage() {}

func (x *AddRecordingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecordingResponse.ProtoReflect.Descriptor instead.
func (*AddRecordingResponse) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{6}
}

func (x *AddRecordingResponse) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

type GetRecordingUriRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор потока
	StreamId string `protobuf:"bytes,1,opt,name=streamId,proto3" json:"streamId,omitempty"`
	// Транспорт для получения потока
	Transport *video.Transport `protobuf:"varint,2,opt,name=transport,proto3,enum=video.Transport,oneof" json:"transport,omitempty"`
	// С какого времени играть запись
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetRecordingUriRequest) Reset() {
	*x = GetRecordingUriRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordingUriRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordingUriRequest) ProtoMessage() {}

func (x *GetRecordingUriRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordingUriRequest.ProtoReflect.Descriptor instead.
func (*GetRecordingUriRequest) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{7}
}

func (x *GetRecordingUriRequest) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

func (x *GetRecordingUriRequest) GetTransport() video.Transport {
	if x != nil && x.Transport != nil {
		return *x.Transport
	}
	return video.Transport(0)
}

func (x *GetRecordingUriRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type SetRecordingStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID записи
	RecordingId string `protobuf:"bytes,1,opt,name=recordingId,proto3" json:"recordingId,omitempty"`
	// Зпись на паузе или нет
	Pause bool `protobuf:"varint,2,opt,name=pause,proto3" json:"pause,omitempty"`
}

func (x *SetRecordingStateRequest) Reset() {
	*x = SetRecordingStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRecordingStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRecordingStateRequest) ProtoMessage() {}

func (x *SetRecordingStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRecordingStateRequest.ProtoReflect.Descriptor instead.
func (*SetRecordingStateRequest) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{8}
}

func (x *SetRecordingStateRequest) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *SetRecordingStateRequest) GetPause() bool {
	if x != nil {
		return x.Pause
	}
	return false
}

type SetRecordingQualityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID записи
	RecordingId string `protobuf:"bytes,1,opt,name=recordingId,proto3" json:"recordingId,omitempty"`
	// Качество записи (0 - как есть, 1 - прорежено до 1 опорного кдра/GOP, 2 - до 1 опорного кадра из 2х GOP)
	Quality uint32 `protobuf:"varint,2,opt,name=quality,proto3" json:"quality,omitempty"`
}

func (x *SetRecordingQualityRequest) Reset() {
	*x = SetRecordingQualityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRecordingQualityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRecordingQualityRequest) ProtoMessage() {}

func (x *SetRecordingQualityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRecordingQualityRequest.ProtoReflect.Descriptor instead.
func (*SetRecordingQualityRequest) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{9}
}

func (x *SetRecordingQualityRequest) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *SetRecordingQualityRequest) GetQuality() uint32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

type RemoveRecordingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID записи
	RecordingId string `protobuf:"bytes,1,opt,name=recordingId,proto3" json:"recordingId,omitempty"`
	// Сохранить ли записанные на диск данные
	SaveRecordedData bool `protobuf:"varint,2,opt,name=saveRecordedData,proto3" json:"saveRecordedData,omitempty"`
}

func (x *RemoveRecordingRequest) Reset() {
	*x = RemoveRecordingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRecordingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRecordingRequest) ProtoMessage() {}

func (x *RemoveRecordingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRecordingRequest.ProtoReflect.Descriptor instead.
func (*RemoveRecordingRequest) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveRecordingRequest) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *RemoveRecordingRequest) GetSaveRecordedData() bool {
	if x != nil {
		return x.SaveRecordedData
	}
	return false
}

type GetBackupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Данные бекапа
	RawData []byte `protobuf:"bytes,1,opt,name=rawData,proto3" json:"rawData,omitempty"`
	// MIME тип
	MimeType string `protobuf:"bytes,2,opt,name=mimeType,proto3" json:"mimeType,omitempty"`
}

func (x *GetBackupResponse) Reset() {
	*x = GetBackupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cctv_backend_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBackupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBackupResponse) ProtoMessage() {}

func (x *GetBackupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cctv_backend_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBackupResponse.ProtoReflect.Descriptor instead.
func (*GetBackupResponse) Descriptor() ([]byte, []int) {
	return file_cctv_backend_proto_rawDescGZIP(), []int{11}
}

func (x *GetBackupResponse) GetRawData() []byte {
	if x != nil {
		return x.RawData
	}
	return nil
}

func (x *GetBackupResponse) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

var File_cctv_backend_proto protoreflect.FileDescriptor

var file_cctv_backend_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x63, 0x74, 0x76, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5c, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x4e, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x4e, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x64, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x64, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2f, 0x0a,
	0x11, 0x41, 0x64, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x31,
	0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x22, 0x74, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x72,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x00, 0x52, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x72,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x2e, 0x0a, 0x09, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x55, 0x0a, 0x13, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x79, 0x73, 0x22, 0x38, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0xb1, 0x01, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x00, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x52, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x75, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x1a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x66,
	0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x61,
	0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x32, 0xb1, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x32, 0x0a, 0x09,
	0x41, 0x64, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x41,
	0x64, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x72, 0x69,
	0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x72, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x72, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xdd, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x55, 0x72, 0x69, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x55, 0x72, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x72, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x11, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x2e, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x42, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x41, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2e, 0x2f, 0x63,
	0x63, 0x74, 0x76, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cctv_backend_proto_rawDescOnce sync.Once
	file_cctv_backend_proto_rawDescData = file_cctv_backend_proto_rawDesc
)

func file_cctv_backend_proto_rawDescGZIP() []byte {
	file_cctv_backend_proto_rawDescOnce.Do(func() {
		file_cctv_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_cctv_backend_proto_rawDescData)
	})
	return file_cctv_backend_proto_rawDescData
}

var file_cctv_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_cctv_backend_proto_goTypes = []interface{}{
	(*AddStreamRequest)(nil),           // 0: AddStreamRequest
	(*AddStreamResponse)(nil),          // 1: AddStreamResponse
	(*RemoveStreamRequest)(nil),        // 2: RemoveStreamRequest
	(*GetStreamUriRequest)(nil),        // 3: GetStreamUriRequest
	(*GetUriResponse)(nil),             // 4: GetUriResponse
	(*AddRecordingRequest)(nil),        // 5: AddRecordingRequest
	(*AddRecordingResponse)(nil),       // 6: AddRecordingResponse
	(*GetRecordingUriRequest)(nil),     // 7: GetRecordingUriRequest
	(*SetRecordingStateRequest)(nil),   // 8: SetRecordingStateRequest
	(*SetRecordingQualityRequest)(nil), // 9: SetRecordingQualityRequest
	(*RemoveRecordingRequest)(nil),     // 10: RemoveRecordingRequest
	(*GetBackupResponse)(nil),          // 11: GetBackupResponse
	(video.Transport)(0),               // 12: video.Transport
	(*timestamppb.Timestamp)(nil),      // 13: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 14: google.protobuf.Empty
}
var file_cctv_backend_proto_depIdxs = []int32{
	12, // 0: GetStreamUriRequest.transport:type_name -> video.Transport
	12, // 1: GetUriResponse.transport:type_name -> video.Transport
	12, // 2: GetRecordingUriRequest.transport:type_name -> video.Transport
	13, // 3: GetRecordingUriRequest.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 4: Stream.AddStream:input_type -> AddStreamRequest
	3,  // 5: Stream.GetStreamUri:input_type -> GetStreamUriRequest
	2,  // 6: Stream.RemoveStream:input_type -> RemoveStreamRequest
	5,  // 7: Recording.AddRecording:input_type -> AddRecordingRequest
	7,  // 8: Recording.GetRecordingUri:input_type -> GetRecordingUriRequest
	8,  // 9: Recording.SetRecordingState:input_type -> SetRecordingStateRequest
	9,  // 10: Recording.SetRecordingQuality:input_type -> SetRecordingQualityRequest
	10, // 11: Recording.RemoveRecording:input_type -> RemoveRecordingRequest
	14, // 12: System.GetBackup:input_type -> google.protobuf.Empty
	1,  // 13: Stream.AddStream:output_type -> AddStreamResponse
	4,  // 14: Stream.GetStreamUri:output_type -> GetUriResponse
	14, // 15: Stream.RemoveStream:output_type -> google.protobuf.Empty
	6,  // 16: Recording.AddRecording:output_type -> AddRecordingResponse
	4,  // 17: Recording.GetRecordingUri:output_type -> GetUriResponse
	14, // 18: Recording.SetRecordingState:output_type -> google.protobuf.Empty
	14, // 19: Recording.SetRecordingQuality:output_type -> google.protobuf.Empty
	14, // 20: Recording.RemoveRecording:output_type -> google.protobuf.Empty
	11, // 21: System.GetBackup:output_type -> GetBackupResponse
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_cctv_backend_proto_init() }
func file_cctv_backend_proto_init() {
	if File_cctv_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cctv_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStreamRequest); i {
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
		file_cctv_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStreamResponse); i {
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
		file_cctv_backend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveStreamRequest); i {
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
		file_cctv_backend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamUriRequest); i {
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
		file_cctv_backend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUriResponse); i {
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
		file_cctv_backend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecordingRequest); i {
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
		file_cctv_backend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecordingResponse); i {
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
		file_cctv_backend_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordingUriRequest); i {
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
		file_cctv_backend_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRecordingStateRequest); i {
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
		file_cctv_backend_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRecordingQualityRequest); i {
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
		file_cctv_backend_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRecordingRequest); i {
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
		file_cctv_backend_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBackupResponse); i {
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
	file_cctv_backend_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_cctv_backend_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cctv_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cctv_backend_proto_goTypes,
		DependencyIndexes: file_cctv_backend_proto_depIdxs,
		MessageInfos:      file_cctv_backend_proto_msgTypes,
	}.Build()
	File_cctv_backend_proto = out.File
	file_cctv_backend_proto_rawDesc = nil
	file_cctv_backend_proto_goTypes = nil
	file_cctv_backend_proto_depIdxs = nil
}
