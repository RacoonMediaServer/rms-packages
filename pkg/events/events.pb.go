// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.1
// source: events.proto

package events

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

type AlertKind int32

const (
	AlertKind_MotionDetected    AlertKind = 0
	AlertKind_TamperDetected    AlertKind = 1
	AlertKind_CrossLineDetected AlertKind = 2
	AlertKind_GuestDetected     AlertKind = 3
	AlertKind_IntrusionDetected AlertKind = 4
)

// Enum value maps for AlertKind.
var (
	AlertKind_name = map[int32]string{
		0: "MotionDetected",
		1: "TamperDetected",
		2: "CrossLineDetected",
		3: "GuestDetected",
		4: "IntrusionDetected",
	}
	AlertKind_value = map[string]int32{
		"MotionDetected":    0,
		"TamperDetected":    1,
		"CrossLineDetected": 2,
		"GuestDetected":     3,
		"IntrusionDetected": 4,
	}
)

func (x AlertKind) Enum() *AlertKind {
	p := new(AlertKind)
	*p = x
	return p
}

func (x AlertKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertKind) Descriptor() protoreflect.EnumDescriptor {
	return file_events_proto_enumTypes[0].Descriptor()
}

func (AlertKind) Type() protoreflect.EnumType {
	return &file_events_proto_enumTypes[0]
}

func (x AlertKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertKind.Descriptor instead.
func (AlertKind) EnumDescriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

type Notification_Kind int32

const (
	// Torrent done
	Notification_DownloadComplete Notification_Kind = 0
	// Something wrong with torrent
	Notification_DownloadFailed Notification_Kind = 1
	// Transcoding of some content has done
	Notification_TranscodingDone Notification_Kind = 2
	// Error raises during transcoding
	Notification_TranscodingFailed Notification_Kind = 3
	// Torrent removed by user
	Notification_TorrentRemoved Notification_Kind = 4
	// Backup operation complete
	Notification_BackupComplete Notification_Kind = 5
)

// Enum value maps for Notification_Kind.
var (
	Notification_Kind_name = map[int32]string{
		0: "DownloadComplete",
		1: "DownloadFailed",
		2: "TranscodingDone",
		3: "TranscodingFailed",
		4: "TorrentRemoved",
		5: "BackupComplete",
	}
	Notification_Kind_value = map[string]int32{
		"DownloadComplete":  0,
		"DownloadFailed":    1,
		"TranscodingDone":   2,
		"TranscodingFailed": 3,
		"TorrentRemoved":    4,
		"BackupComplete":    5,
	}
)

func (x Notification_Kind) Enum() *Notification_Kind {
	p := new(Notification_Kind)
	*p = x
	return p
}

func (x Notification_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Notification_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_events_proto_enumTypes[1].Descriptor()
}

func (Notification_Kind) Type() protoreflect.EnumType {
	return &file_events_proto_enumTypes[1]
}

func (x Notification_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Notification_Kind.Descriptor instead.
func (Notification_Kind) EnumDescriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0, 0}
}

type Malfunction_System int32

const (
	Malfunction_Cameras  Malfunction_System = 0
	Malfunction_Archive  Malfunction_System = 1
	Malfunction_Services Malfunction_System = 2
	Malfunction_Media    Malfunction_System = 3
)

// Enum value maps for Malfunction_System.
var (
	Malfunction_System_name = map[int32]string{
		0: "Cameras",
		1: "Archive",
		2: "Services",
		3: "Media",
	}
	Malfunction_System_value = map[string]int32{
		"Cameras":  0,
		"Archive":  1,
		"Services": 2,
		"Media":    3,
	}
)

func (x Malfunction_System) Enum() *Malfunction_System {
	p := new(Malfunction_System)
	*p = x
	return p
}

func (x Malfunction_System) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Malfunction_System) Descriptor() protoreflect.EnumDescriptor {
	return file_events_proto_enumTypes[2].Descriptor()
}

func (Malfunction_System) Type() protoreflect.EnumType {
	return &file_events_proto_enumTypes[2]
}

func (x Malfunction_System) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Malfunction_System.Descriptor instead.
func (Malfunction_System) EnumDescriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1, 0}
}

type Malfunction_Code int32

const (
	Malfunction_Unknown        Malfunction_Code = 0
	Malfunction_CannotAccess   Malfunction_Code = 1
	Malfunction_NotEnoughSpace Malfunction_Code = 2
	Malfunction_TaskIsHung     Malfunction_Code = 3
	Malfunction_ActionFailed   Malfunction_Code = 4
)

// Enum value maps for Malfunction_Code.
var (
	Malfunction_Code_name = map[int32]string{
		0: "Unknown",
		1: "CannotAccess",
		2: "NotEnoughSpace",
		3: "TaskIsHung",
		4: "ActionFailed",
	}
	Malfunction_Code_value = map[string]int32{
		"Unknown":        0,
		"CannotAccess":   1,
		"NotEnoughSpace": 2,
		"TaskIsHung":     3,
		"ActionFailed":   4,
	}
)

func (x Malfunction_Code) Enum() *Malfunction_Code {
	p := new(Malfunction_Code)
	*p = x
	return p
}

func (x Malfunction_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Malfunction_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_events_proto_enumTypes[3].Descriptor()
}

func (Malfunction_Code) Type() protoreflect.EnumType {
	return &file_events_proto_enumTypes[3]
}

func (x Malfunction_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Malfunction_Code.Descriptor instead.
func (Malfunction_Code) EnumDescriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1, 1}
}

// User information notifications
type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender        string            `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	Kind          Notification_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=events.Notification_Kind" json:"kind,omitempty"`
	TorrentID     *string           `protobuf:"bytes,2,opt,name=torrentID,proto3,oneof" json:"torrentID,omitempty"`
	MediaID       *string           `protobuf:"bytes,3,opt,name=mediaID,proto3,oneof" json:"mediaID,omitempty"`
	ItemTitle     *string           `protobuf:"bytes,4,opt,name=itemTitle,proto3,oneof" json:"itemTitle,omitempty"`
	VideoLocation *string           `protobuf:"bytes,5,opt,name=videoLocation,proto3,oneof" json:"videoLocation,omitempty"`
	SizeMB        *uint32           `protobuf:"varint,7,opt,name=sizeMB,proto3,oneof" json:"sizeMB,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Notification) GetKind() Notification_Kind {
	if x != nil {
		return x.Kind
	}
	return Notification_DownloadComplete
}

func (x *Notification) GetTorrentID() string {
	if x != nil && x.TorrentID != nil {
		return *x.TorrentID
	}
	return ""
}

func (x *Notification) GetMediaID() string {
	if x != nil && x.MediaID != nil {
		return *x.MediaID
	}
	return ""
}

func (x *Notification) GetItemTitle() string {
	if x != nil && x.ItemTitle != nil {
		return *x.ItemTitle
	}
	return ""
}

func (x *Notification) GetVideoLocation() string {
	if x != nil && x.VideoLocation != nil {
		return *x.VideoLocation
	}
	return ""
}

func (x *Notification) GetSizeMB() uint32 {
	if x != nil && x.SizeMB != nil {
		return *x.SizeMB
	}
	return 0
}

// Notification about malfunctions and any problems in the entire system
type Malfunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender     string             `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	Timestamp  int64              `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Error      string             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	System     Malfunction_System `protobuf:"varint,3,opt,name=system,proto3,enum=events.Malfunction_System" json:"system,omitempty"`
	Code       Malfunction_Code   `protobuf:"varint,4,opt,name=code,proto3,enum=events.Malfunction_Code" json:"code,omitempty"`
	StackTrace string             `protobuf:"bytes,5,opt,name=stackTrace,proto3" json:"stackTrace,omitempty"`
}

func (x *Malfunction) Reset() {
	*x = Malfunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Malfunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Malfunction) ProtoMessage() {}

func (x *Malfunction) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Malfunction.ProtoReflect.Descriptor instead.
func (*Malfunction) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *Malfunction) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Malfunction) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Malfunction) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Malfunction) GetSystem() Malfunction_System {
	if x != nil {
		return x.System
	}
	return Malfunction_Cameras
}

func (x *Malfunction) GetCode() Malfunction_Code {
	if x != nil {
		return x.Code
	}
	return Malfunction_Unknown
}

func (x *Malfunction) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

// CCTV security alerts
type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender         string    `protobuf:"bytes,8,opt,name=sender,proto3" json:"sender,omitempty"`
	Timestamp      int64     `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Camera         string    `protobuf:"bytes,2,opt,name=camera,proto3" json:"camera,omitempty"`
	Kind           AlertKind `protobuf:"varint,3,opt,name=kind,proto3,enum=events.AlertKind" json:"kind,omitempty"`
	Image          []byte    `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	ImageMimeType  string    `protobuf:"bytes,5,opt,name=imageMimeType,proto3" json:"imageMimeType,omitempty"`
	DurationSec    uint64    `protobuf:"varint,6,opt,name=durationSec,proto3" json:"durationSec,omitempty"`       // длительность события (мб предполагаемая)
	ExtraOffsetSec uint64    `protobuf:"varint,7,opt,name=extraOffsetSec,proto3" json:"extraOffsetSec,omitempty"` // сколько еще захватить секунд с обоих концов
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *Alert) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Alert) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Alert) GetCamera() string {
	if x != nil {
		return x.Camera
	}
	return ""
}

func (x *Alert) GetKind() AlertKind {
	if x != nil {
		return x.Kind
	}
	return AlertKind_MotionDetected
}

func (x *Alert) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Alert) GetImageMimeType() string {
	if x != nil {
		return x.ImageMimeType
	}
	return ""
}

func (x *Alert) GetDurationSec() uint64 {
	if x != nil {
		return x.DurationSec
	}
	return 0
}

func (x *Alert) GetExtraOffsetSec() uint64 {
	if x != nil {
		return x.ExtraOffsetSec
	}
	return 0
}

// Security incidents
type Incident struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   string             `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Cameras  []string           `protobuf:"bytes,2,rep,name=cameras,proto3" json:"cameras,omitempty"`
	Begin    int64              `protobuf:"varint,3,opt,name=begin,proto3" json:"begin,omitempty"`
	Duration int64              `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	End      int64              `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	Records  []*Incident_Record `protobuf:"bytes,6,rep,name=Records,proto3" json:"Records,omitempty"`
}

func (x *Incident) Reset() {
	*x = Incident{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Incident) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Incident) ProtoMessage() {}

func (x *Incident) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Incident.ProtoReflect.Descriptor instead.
func (*Incident) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{3}
}

func (x *Incident) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Incident) GetCameras() []string {
	if x != nil {
		return x.Cameras
	}
	return nil
}

func (x *Incident) GetBegin() int64 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *Incident) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Incident) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Incident) GetRecords() []*Incident_Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type Incident_Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Camera           string    `protobuf:"bytes,1,opt,name=camera,proto3" json:"camera,omitempty"`
	Kind             AlertKind `protobuf:"varint,2,opt,name=kind,proto3,enum=events.AlertKind" json:"kind,omitempty"`
	Begin            int64     `protobuf:"varint,3,opt,name=begin,proto3" json:"begin,omitempty"`
	Duration         int64     `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	End              int64     `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	SnapshotPath     string    `protobuf:"bytes,6,opt,name=snapshotPath,proto3" json:"snapshotPath,omitempty"`
	SnapshotMimeType string    `protobuf:"bytes,7,opt,name=snapshotMimeType,proto3" json:"snapshotMimeType,omitempty"`
	VideoPath        string    `protobuf:"bytes,8,opt,name=videoPath,proto3" json:"videoPath,omitempty"`
	VideoURL         string    `protobuf:"bytes,9,opt,name=videoURL,proto3" json:"videoURL,omitempty"`
}

func (x *Incident_Record) Reset() {
	*x = Incident_Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Incident_Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Incident_Record) ProtoMessage() {}

func (x *Incident_Record) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Incident_Record.ProtoReflect.Descriptor instead.
func (*Incident_Record) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Incident_Record) GetCamera() string {
	if x != nil {
		return x.Camera
	}
	return ""
}

func (x *Incident_Record) GetKind() AlertKind {
	if x != nil {
		return x.Kind
	}
	return AlertKind_MotionDetected
}

func (x *Incident_Record) GetBegin() int64 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *Incident_Record) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Incident_Record) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Incident_Record) GetSnapshotPath() string {
	if x != nil {
		return x.SnapshotPath
	}
	return ""
}

func (x *Incident_Record) GetSnapshotMimeType() string {
	if x != nil {
		return x.SnapshotMimeType
	}
	return ""
}

func (x *Incident_Record) GetVideoPath() string {
	if x != nil {
		return x.VideoPath
	}
	return ""
}

func (x *Incident_Record) GetVideoURL() string {
	if x != nil {
		return x.VideoURL
	}
	return ""
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xce, 0x03, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x2d, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x21,
	0x0a, 0x09, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0d, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x42, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04,
	0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x42, 0x88, 0x01, 0x01, 0x22, 0x84, 0x01, 0x0a, 0x04,
	0x4b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x6f, 0x6e,
	0x65, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x6f,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x04, 0x12, 0x12,
	0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x10, 0x05, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x44, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x42, 0x22, 0xf5, 0x02, 0x0a, 0x0b, 0x4d, 0x61, 0x6c, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4d, 0x61, 0x6c,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x2c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4d,
	0x61, 0x6c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x22, 0x3b, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x73, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x10, 0x03, 0x22, 0x5b, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x6f, 0x74,
	0x45, 0x6e, 0x6f, 0x75, 0x67, 0x68, 0x53, 0x70, 0x61, 0x63, 0x65, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x73, 0x48, 0x75, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x10, 0x0a,
	0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x04, 0x22,
	0x82, 0x02, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x12, 0x25, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x69, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x12, 0x26, 0x0a, 0x0e,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x53, 0x65, 0x63, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x53, 0x65, 0x63, 0x22, 0xcb, 0x03, 0x0a, 0x08, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x6d,
	0x65, 0x72, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x95, 0x02, 0x0a, 0x06, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x12, 0x25, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x10,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x50, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55,
	0x52, 0x4c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55,
	0x52, 0x4c, 0x2a, 0x74, 0x0a, 0x09, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x12, 0x0a, 0x0e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x61, 0x6d, 0x70, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x72, 0x6f, 0x73, 0x73,
	0x4c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x47, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10,
	0x03, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x72, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x04, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x63, 0x6f, 0x6f, 0x6e, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x6d, 0x73, 0x2d, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_events_proto_goTypes = []interface{}{
	(AlertKind)(0),          // 0: events.AlertKind
	(Notification_Kind)(0),  // 1: events.Notification.Kind
	(Malfunction_System)(0), // 2: events.Malfunction.System
	(Malfunction_Code)(0),   // 3: events.Malfunction.Code
	(*Notification)(nil),    // 4: events.Notification
	(*Malfunction)(nil),     // 5: events.Malfunction
	(*Alert)(nil),           // 6: events.Alert
	(*Incident)(nil),        // 7: events.Incident
	(*Incident_Record)(nil), // 8: events.Incident.Record
}
var file_events_proto_depIdxs = []int32{
	1, // 0: events.Notification.kind:type_name -> events.Notification.Kind
	2, // 1: events.Malfunction.system:type_name -> events.Malfunction.System
	3, // 2: events.Malfunction.code:type_name -> events.Malfunction.Code
	0, // 3: events.Alert.kind:type_name -> events.AlertKind
	8, // 4: events.Incident.Records:type_name -> events.Incident.Record
	0, // 5: events.Incident.Record.kind:type_name -> events.AlertKind
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Malfunction); i {
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
		file_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Incident); i {
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
		file_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Incident_Record); i {
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
	file_events_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		EnumInfos:         file_events_proto_enumTypes,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}
