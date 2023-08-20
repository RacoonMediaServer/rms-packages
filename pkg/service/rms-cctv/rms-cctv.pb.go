// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: rms-cctv.proto

package rms_cctv

import (
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

// Режим записи камеры в архив
type RecordingMode int32

const (
	// Не писать
	RecordingMode_Never RecordingMode = 0
	// Всегда писать
	RecordingMode_Always RecordingMode = 1
	// Пишем, только если сработало событие
	RecordingMode_ByEvents RecordingMode = 2
	// Пишем с прореживанием, когда сработало событие - все кадры
	RecordingMode_Optimal RecordingMode = 3
)

// Enum value maps for RecordingMode.
var (
	RecordingMode_name = map[int32]string{
		0: "Never",
		1: "Always",
		2: "ByEvents",
		3: "Optimal",
	}
	RecordingMode_value = map[string]int32{
		"Never":    0,
		"Always":   1,
		"ByEvents": 2,
		"Optimal":  3,
	}
)

func (x RecordingMode) Enum() *RecordingMode {
	p := new(RecordingMode)
	*p = x
	return p
}

func (x RecordingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_rms_cctv_proto_enumTypes[0].Descriptor()
}

func (RecordingMode) Type() protoreflect.EnumType {
	return &file_rms_cctv_proto_enumTypes[0]
}

func (x RecordingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordingMode.Descriptor instead.
func (RecordingMode) EnumDescriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{0}
}

type VideoTransport int32

const (
	VideoTransport_RTSP               VideoTransport = 0
	VideoTransport_HTTP_HLS_MPEGTS    VideoTransport = 1
	VideoTransport_HTTP_HLS_MP4       VideoTransport = 2
	VideoTransport_HTTP_HLS_ONE_CHUNK VideoTransport = 3
	VideoTransport_HTTP_MP4           VideoTransport = 4
	VideoTransport_HTTP_FLV           VideoTransport = 5
	VideoTransport_HTTP_MPEGTS        VideoTransport = 6
	VideoTransport_HTTP_DASH          VideoTransport = 7
)

// Enum value maps for VideoTransport.
var (
	VideoTransport_name = map[int32]string{
		0: "RTSP",
		1: "HTTP_HLS_MPEGTS",
		2: "HTTP_HLS_MP4",
		3: "HTTP_HLS_ONE_CHUNK",
		4: "HTTP_MP4",
		5: "HTTP_FLV",
		6: "HTTP_MPEGTS",
		7: "HTTP_DASH",
	}
	VideoTransport_value = map[string]int32{
		"RTSP":               0,
		"HTTP_HLS_MPEGTS":    1,
		"HTTP_HLS_MP4":       2,
		"HTTP_HLS_ONE_CHUNK": 3,
		"HTTP_MP4":           4,
		"HTTP_FLV":           5,
		"HTTP_MPEGTS":        6,
		"HTTP_DASH":          7,
	}
)

func (x VideoTransport) Enum() *VideoTransport {
	p := new(VideoTransport)
	*p = x
	return p
}

func (x VideoTransport) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VideoTransport) Descriptor() protoreflect.EnumDescriptor {
	return file_rms_cctv_proto_enumTypes[1].Descriptor()
}

func (VideoTransport) Type() protoreflect.EnumType {
	return &file_rms_cctv_proto_enumTypes[1]
}

func (x VideoTransport) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VideoTransport.Descriptor instead.
func (VideoTransport) EnumDescriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{1}
}

type CctvSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventNotifyThresholdIntervalSec uint32 `protobuf:"varint,1,opt,name=EventNotifyThresholdIntervalSec,proto3" json:"EventNotifyThresholdIntervalSec,omitempty"`
	OneEventDefaultDurationSec      uint32 `protobuf:"varint,2,opt,name=OneEventDefaultDurationSec,proto3" json:"OneEventDefaultDurationSec,omitempty"`
}

func (x *CctvSettings) Reset() {
	*x = CctvSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CctvSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CctvSettings) ProtoMessage() {}

func (x *CctvSettings) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CctvSettings.ProtoReflect.Descriptor instead.
func (*CctvSettings) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{0}
}

func (x *CctvSettings) GetEventNotifyThresholdIntervalSec() uint32 {
	if x != nil {
		return x.EventNotifyThresholdIntervalSec
	}
	return 0
}

func (x *CctvSettings) GetOneEventDefaultDurationSec() uint32 {
	if x != nil {
		return x.OneEventDefaultDurationSec
	}
	return 0
}

type Camera struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Внутренний идентификатор камеры
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Видимое имя камеры
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// URL для доступа к камере (HTTP)
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// Имя пользователя
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	// Пароль
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	// Сколько дней хранить архив
	KeepDays uint32 `protobuf:"varint,6,opt,name=keepDays,proto3" json:"keepDays,omitempty"`
	// Режим записи
	Mode RecordingMode `protobuf:"varint,7,opt,name=mode,proto3,enum=RecordingMode" json:"mode,omitempty"`
	// Расписание уведомлений (в формате iCal)
	Schedule string `protobuf:"bytes,8,opt,name=schedule,proto3" json:"schedule,omitempty"`
}

func (x *Camera) Reset() {
	*x = Camera{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Camera) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Camera) ProtoMessage() {}

func (x *Camera) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Camera.ProtoReflect.Descriptor instead.
func (*Camera) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{1}
}

func (x *Camera) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Camera) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Camera) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Camera) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Camera) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Camera) GetKeepDays() uint32 {
	if x != nil {
		return x.KeepDays
	}
	return 0
}

func (x *Camera) GetMode() RecordingMode {
	if x != nil {
		return x.Mode
	}
	return RecordingMode_Never
}

func (x *Camera) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

type GetCamerasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Список камер
	Cameras []*Camera `protobuf:"bytes,1,rep,name=cameras,proto3" json:"cameras,omitempty"`
}

func (x *GetCamerasResponse) Reset() {
	*x = GetCamerasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCamerasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCamerasResponse) ProtoMessage() {}

func (x *GetCamerasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCamerasResponse.ProtoReflect.Descriptor instead.
func (*GetCamerasResponse) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{2}
}

func (x *GetCamerasResponse) GetCameras() []*Camera {
	if x != nil {
		return x.Cameras
	}
	return nil
}

type AddCameraResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID камеры
	CameraId uint32 `protobuf:"varint,1,opt,name=cameraId,proto3" json:"cameraId,omitempty"`
}

func (x *AddCameraResponse) Reset() {
	*x = AddCameraResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCameraResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCameraResponse) ProtoMessage() {}

func (x *AddCameraResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCameraResponse.ProtoReflect.Descriptor instead.
func (*AddCameraResponse) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{3}
}

func (x *AddCameraResponse) GetCameraId() uint32 {
	if x != nil {
		return x.CameraId
	}
	return 0
}

type DeleteCameraRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID камеры
	CameraId uint32 `protobuf:"varint,1,opt,name=cameraId,proto3" json:"cameraId,omitempty"`
}

func (x *DeleteCameraRequest) Reset() {
	*x = DeleteCameraRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCameraRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCameraRequest) ProtoMessage() {}

func (x *DeleteCameraRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCameraRequest.ProtoReflect.Descriptor instead.
func (*DeleteCameraRequest) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCameraRequest) GetCameraId() uint32 {
	if x != nil {
		return x.CameraId
	}
	return 0
}

type GetLiveUriRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID камеры
	CameraId uint32 `protobuf:"varint,1,opt,name=cameraId,proto3" json:"cameraId,omitempty"`
	// Протокол для формирования ссылки
	Transport VideoTransport `protobuf:"varint,2,opt,name=transport,proto3,enum=VideoTransport" json:"transport,omitempty"`
}

func (x *GetLiveUriRequest) Reset() {
	*x = GetLiveUriRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLiveUriRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLiveUriRequest) ProtoMessage() {}

func (x *GetLiveUriRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLiveUriRequest.ProtoReflect.Descriptor instead.
func (*GetLiveUriRequest) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{5}
}

func (x *GetLiveUriRequest) GetCameraId() uint32 {
	if x != nil {
		return x.CameraId
	}
	return 0
}

func (x *GetLiveUriRequest) GetTransport() VideoTransport {
	if x != nil {
		return x.Transport
	}
	return VideoTransport_RTSP
}

type GetLiveUriResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URI для воспроизведения
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *GetLiveUriResponse) Reset() {
	*x = GetLiveUriResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLiveUriResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLiveUriResponse) ProtoMessage() {}

func (x *GetLiveUriResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLiveUriResponse.ProtoReflect.Descriptor instead.
func (*GetLiveUriResponse) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{6}
}

func (x *GetLiveUriResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type GetReplayUriRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID камеры
	CameraId uint32 `protobuf:"varint,1,opt,name=cameraId,proto3" json:"cameraId,omitempty"`
	// Протокол для формирования ссылки
	Transport VideoTransport `protobuf:"varint,2,opt,name=transport,proto3,enum=VideoTransport" json:"transport,omitempty"`
	// С какого момента начать показывать запись
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetReplayUriRequest) Reset() {
	*x = GetReplayUriRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReplayUriRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReplayUriRequest) ProtoMessage() {}

func (x *GetReplayUriRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReplayUriRequest.ProtoReflect.Descriptor instead.
func (*GetReplayUriRequest) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{7}
}

func (x *GetReplayUriRequest) GetCameraId() uint32 {
	if x != nil {
		return x.CameraId
	}
	return 0
}

func (x *GetReplayUriRequest) GetTransport() VideoTransport {
	if x != nil {
		return x.Transport
	}
	return VideoTransport_RTSP
}

func (x *GetReplayUriRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type GetSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID камеры
	CameraId uint32 `protobuf:"varint,1,opt,name=cameraId,proto3" json:"cameraId,omitempty"`
}

func (x *GetSnapshotRequest) Reset() {
	*x = GetSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnapshotRequest) ProtoMessage() {}

func (x *GetSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnapshotRequest.ProtoReflect.Descriptor instead.
func (*GetSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{8}
}

func (x *GetSnapshotRequest) GetCameraId() uint32 {
	if x != nil {
		return x.CameraId
	}
	return 0
}

type GetSnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JPEG-изображение
	Snapshot []byte `protobuf:"bytes,1,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (x *GetSnapshotResponse) Reset() {
	*x = GetSnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_cctv_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnapshotResponse) ProtoMessage() {}

func (x *GetSnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rms_cctv_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnapshotResponse.ProtoReflect.Descriptor instead.
func (*GetSnapshotResponse) Descriptor() ([]byte, []int) {
	return file_rms_cctv_proto_rawDescGZIP(), []int{9}
}

func (x *GetSnapshotResponse) GetSnapshot() []byte {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

var File_rms_cctv_proto protoreflect.FileDescriptor

var file_rms_cctv_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x6d, 0x73, 0x2d, 0x63, 0x63, 0x74, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98,
	0x01, 0x0a, 0x0c, 0x43, 0x63, 0x74, 0x76, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x48, 0x0a, 0x1f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53,
	0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x12, 0x3e, 0x0a, 0x1a, 0x4f, 0x6e, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x4f,
	0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x22, 0xca, 0x01, 0x0a, 0x06, 0x43, 0x61,
	0x6d, 0x65, 0x72, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65,
	0x65, 0x70, 0x44, 0x61, 0x79, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6b, 0x65,
	0x65, 0x70, 0x44, 0x61, 0x79, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x37, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d,
	0x65, 0x72, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x07,
	0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x52, 0x07, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x73, 0x22,
	0x2f, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64,
	0x22, 0x31, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x72,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0x26, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x72,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22, 0x9a, 0x01, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2a, 0x41, 0x0a,
	0x0d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x4e, 0x65, 0x76, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x6c, 0x77,
	0x61, 0x79, 0x73, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x10, 0x03,
	0x2a, 0x95, 0x01, 0x0a, 0x0e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x54, 0x53, 0x50, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x48, 0x4c, 0x53, 0x5f, 0x4d, 0x50, 0x45, 0x47, 0x54, 0x53,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x48, 0x4c, 0x53, 0x5f, 0x4d,
	0x50, 0x34, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x48, 0x4c, 0x53,
	0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08,
	0x48, 0x54, 0x54, 0x50, 0x5f, 0x4d, 0x50, 0x34, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x54,
	0x54, 0x50, 0x5f, 0x46, 0x4c, 0x56, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x54, 0x54, 0x50,
	0x5f, 0x4d, 0x50, 0x45, 0x47, 0x54, 0x53, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x54, 0x54,
	0x50, 0x5f, 0x44, 0x41, 0x53, 0x48, 0x10, 0x07, 0x32, 0xf6, 0x03, 0x0a, 0x07, 0x52, 0x6d, 0x73,
	0x43, 0x63, 0x74, 0x76, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x43, 0x63,
	0x74, 0x76, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x53, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0d, 0x2e, 0x43, 0x63, 0x74, 0x76,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x41,
	0x64, 0x64, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x12, 0x07, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x1a, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43,
	0x61, 0x6d, 0x65, 0x72, 0x61, 0x12, 0x07, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x6d, 0x65, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55,
	0x72, 0x69, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x72, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x55, 0x72, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x69, 0x12, 0x14, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x72, 0x6d, 0x73, 0x2d, 0x63, 0x63, 0x74, 0x76,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rms_cctv_proto_rawDescOnce sync.Once
	file_rms_cctv_proto_rawDescData = file_rms_cctv_proto_rawDesc
)

func file_rms_cctv_proto_rawDescGZIP() []byte {
	file_rms_cctv_proto_rawDescOnce.Do(func() {
		file_rms_cctv_proto_rawDescData = protoimpl.X.CompressGZIP(file_rms_cctv_proto_rawDescData)
	})
	return file_rms_cctv_proto_rawDescData
}

var file_rms_cctv_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_rms_cctv_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rms_cctv_proto_goTypes = []interface{}{
	(RecordingMode)(0),            // 0: RecordingMode
	(VideoTransport)(0),           // 1: VideoTransport
	(*CctvSettings)(nil),          // 2: CctvSettings
	(*Camera)(nil),                // 3: Camera
	(*GetCamerasResponse)(nil),    // 4: GetCamerasResponse
	(*AddCameraResponse)(nil),     // 5: AddCameraResponse
	(*DeleteCameraRequest)(nil),   // 6: DeleteCameraRequest
	(*GetLiveUriRequest)(nil),     // 7: GetLiveUriRequest
	(*GetLiveUriResponse)(nil),    // 8: GetLiveUriResponse
	(*GetReplayUriRequest)(nil),   // 9: GetReplayUriRequest
	(*GetSnapshotRequest)(nil),    // 10: GetSnapshotRequest
	(*GetSnapshotResponse)(nil),   // 11: GetSnapshotResponse
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 13: google.protobuf.Empty
}
var file_rms_cctv_proto_depIdxs = []int32{
	0,  // 0: Camera.mode:type_name -> RecordingMode
	3,  // 1: GetCamerasResponse.cameras:type_name -> Camera
	1,  // 2: GetLiveUriRequest.transport:type_name -> VideoTransport
	1,  // 3: GetReplayUriRequest.transport:type_name -> VideoTransport
	12, // 4: GetReplayUriRequest.timestamp:type_name -> google.protobuf.Timestamp
	13, // 5: RmsCctv.GetSettings:input_type -> google.protobuf.Empty
	2,  // 6: RmsCctv.SetSettings:input_type -> CctvSettings
	13, // 7: RmsCctv.GetCameras:input_type -> google.protobuf.Empty
	3,  // 8: RmsCctv.AddCamera:input_type -> Camera
	3,  // 9: RmsCctv.ModifyCamera:input_type -> Camera
	6,  // 10: RmsCctv.DeleteCamera:input_type -> DeleteCameraRequest
	7,  // 11: RmsCctv.GetLiveUri:input_type -> GetLiveUriRequest
	9,  // 12: RmsCctv.GetReplayUri:input_type -> GetReplayUriRequest
	10, // 13: RmsCctv.GetSnapshot:input_type -> GetSnapshotRequest
	2,  // 14: RmsCctv.GetSettings:output_type -> CctvSettings
	13, // 15: RmsCctv.SetSettings:output_type -> google.protobuf.Empty
	4,  // 16: RmsCctv.GetCameras:output_type -> GetCamerasResponse
	5,  // 17: RmsCctv.AddCamera:output_type -> AddCameraResponse
	13, // 18: RmsCctv.ModifyCamera:output_type -> google.protobuf.Empty
	13, // 19: RmsCctv.DeleteCamera:output_type -> google.protobuf.Empty
	8,  // 20: RmsCctv.GetLiveUri:output_type -> GetLiveUriResponse
	9,  // 21: RmsCctv.GetReplayUri:output_type -> GetReplayUriRequest
	11, // 22: RmsCctv.GetSnapshot:output_type -> GetSnapshotResponse
	14, // [14:23] is the sub-list for method output_type
	5,  // [5:14] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_rms_cctv_proto_init() }
func file_rms_cctv_proto_init() {
	if File_rms_cctv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rms_cctv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CctvSettings); i {
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
		file_rms_cctv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Camera); i {
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
		file_rms_cctv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCamerasResponse); i {
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
		file_rms_cctv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCameraResponse); i {
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
		file_rms_cctv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCameraRequest); i {
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
		file_rms_cctv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLiveUriRequest); i {
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
		file_rms_cctv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLiveUriResponse); i {
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
		file_rms_cctv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReplayUriRequest); i {
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
		file_rms_cctv_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnapshotRequest); i {
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
		file_rms_cctv_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnapshotResponse); i {
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
			RawDescriptor: file_rms_cctv_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rms_cctv_proto_goTypes,
		DependencyIndexes: file_rms_cctv_proto_depIdxs,
		EnumInfos:         file_rms_cctv_proto_enumTypes,
		MessageInfos:      file_rms_cctv_proto_msgTypes,
	}.Build()
	File_rms_cctv_proto = out.File
	file_rms_cctv_proto_rawDesc = nil
	file_rms_cctv_proto_goTypes = nil
	file_rms_cctv_proto_depIdxs = nil
}
