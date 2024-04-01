// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: rms-speech.proto

package rms_speech

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

type GetRecognitionStatusResponse_Status int32

const (
	GetRecognitionStatusResponse_Pending    GetRecognitionStatusResponse_Status = 0
	GetRecognitionStatusResponse_Processing GetRecognitionStatusResponse_Status = 1
	GetRecognitionStatusResponse_Failed     GetRecognitionStatusResponse_Status = 3
	GetRecognitionStatusResponse_Done       GetRecognitionStatusResponse_Status = 4
)

// Enum value maps for GetRecognitionStatusResponse_Status.
var (
	GetRecognitionStatusResponse_Status_name = map[int32]string{
		0: "Pending",
		1: "Processing",
		3: "Failed",
		4: "Done",
	}
	GetRecognitionStatusResponse_Status_value = map[string]int32{
		"Pending":    0,
		"Processing": 1,
		"Failed":     3,
		"Done":       4,
	}
)

func (x GetRecognitionStatusResponse_Status) Enum() *GetRecognitionStatusResponse_Status {
	p := new(GetRecognitionStatusResponse_Status)
	*p = x
	return p
}

func (x GetRecognitionStatusResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetRecognitionStatusResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_rms_speech_proto_enumTypes[0].Descriptor()
}

func (GetRecognitionStatusResponse_Status) Type() protoreflect.EnumType {
	return &file_rms_speech_proto_enumTypes[0]
}

func (x GetRecognitionStatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetRecognitionStatusResponse_Status.Descriptor instead.
func (GetRecognitionStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_rms_speech_proto_rawDescGZIP(), []int{3, 0}
}

type StartRecognitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Аудио для распознавания
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Content-Type
	ContentType string `protobuf:"bytes,2,opt,name=contentType,proto3" json:"contentType,omitempty"`
	// Таймаут длительности процесса
	TimeoutSec uint32 `protobuf:"varint,3,opt,name=timeoutSec,proto3" json:"timeoutSec,omitempty"`
}

func (x *StartRecognitionRequest) Reset() {
	*x = StartRecognitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_speech_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRecognitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRecognitionRequest) ProtoMessage() {}

func (x *StartRecognitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_speech_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRecognitionRequest.ProtoReflect.Descriptor instead.
func (*StartRecognitionRequest) Descriptor() ([]byte, []int) {
	return file_rms_speech_proto_rawDescGZIP(), []int{0}
}

func (x *StartRecognitionRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StartRecognitionRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *StartRecognitionRequest) GetTimeoutSec() uint32 {
	if x != nil {
		return x.TimeoutSec
	}
	return 0
}

type StartRecognitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID созданной задачи
	JobId string `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
}

func (x *StartRecognitionResponse) Reset() {
	*x = StartRecognitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_speech_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRecognitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRecognitionResponse) ProtoMessage() {}

func (x *StartRecognitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rms_speech_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRecognitionResponse.ProtoReflect.Descriptor instead.
func (*StartRecognitionResponse) Descriptor() ([]byte, []int) {
	return file_rms_speech_proto_rawDescGZIP(), []int{1}
}

func (x *StartRecognitionResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type GetRecognitionStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
}

func (x *GetRecognitionStatusRequest) Reset() {
	*x = GetRecognitionStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_speech_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecognitionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecognitionStatusRequest) ProtoMessage() {}

func (x *GetRecognitionStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_speech_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecognitionStatusRequest.ProtoReflect.Descriptor instead.
func (*GetRecognitionStatusRequest) Descriptor() ([]byte, []int) {
	return file_rms_speech_proto_rawDescGZIP(), []int{2}
}

func (x *GetRecognitionStatusRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type GetRecognitionStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Состояние задачи
	Status GetRecognitionStatusResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=GetRecognitionStatusResponse_Status" json:"status,omitempty"`
	// Распознанный текст, если статус == Done
	RecognizedText string `protobuf:"bytes,2,opt,name=recognizedText,proto3" json:"recognizedText,omitempty"`
}

func (x *GetRecognitionStatusResponse) Reset() {
	*x = GetRecognitionStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_speech_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecognitionStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecognitionStatusResponse) ProtoMessage() {}

func (x *GetRecognitionStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rms_speech_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecognitionStatusResponse.ProtoReflect.Descriptor instead.
func (*GetRecognitionStatusResponse) Descriptor() ([]byte, []int) {
	return file_rms_speech_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecognitionStatusResponse) GetStatus() GetRecognitionStatusResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetRecognitionStatusResponse_Pending
}

func (x *GetRecognitionStatusResponse) GetRecognizedText() string {
	if x != nil {
		return x.RecognizedText
	}
	return ""
}

type StopRecognitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
}

func (x *StopRecognitionRequest) Reset() {
	*x = StopRecognitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rms_speech_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRecognitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRecognitionRequest) ProtoMessage() {}

func (x *StopRecognitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rms_speech_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRecognitionRequest.ProtoReflect.Descriptor instead.
func (*StopRecognitionRequest) Descriptor() ([]byte, []int) {
	return file_rms_speech_proto_rawDescGZIP(), []int{4}
}

func (x *StopRecognitionRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

var File_rms_speech_proto protoreflect.FileDescriptor

var file_rms_speech_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x6d, 0x73, 0x2d, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6f, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63,
	0x22, 0x30, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62,
	0x49, 0x64, 0x22, 0x33, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x7a, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x22, 0x3b,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x04, 0x22, 0x2e, 0x0a, 0x16, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x32, 0xea, 0x01, 0x0a, 0x06,
	0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x12, 0x47, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2e, 0x2f, 0x72,
	0x6d, 0x73, 0x2d, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rms_speech_proto_rawDescOnce sync.Once
	file_rms_speech_proto_rawDescData = file_rms_speech_proto_rawDesc
)

func file_rms_speech_proto_rawDescGZIP() []byte {
	file_rms_speech_proto_rawDescOnce.Do(func() {
		file_rms_speech_proto_rawDescData = protoimpl.X.CompressGZIP(file_rms_speech_proto_rawDescData)
	})
	return file_rms_speech_proto_rawDescData
}

var file_rms_speech_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rms_speech_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rms_speech_proto_goTypes = []interface{}{
	(GetRecognitionStatusResponse_Status)(0), // 0: GetRecognitionStatusResponse.Status
	(*StartRecognitionRequest)(nil),          // 1: StartRecognitionRequest
	(*StartRecognitionResponse)(nil),         // 2: StartRecognitionResponse
	(*GetRecognitionStatusRequest)(nil),      // 3: GetRecognitionStatusRequest
	(*GetRecognitionStatusResponse)(nil),     // 4: GetRecognitionStatusResponse
	(*StopRecognitionRequest)(nil),           // 5: StopRecognitionRequest
	(*emptypb.Empty)(nil),                    // 6: google.protobuf.Empty
}
var file_rms_speech_proto_depIdxs = []int32{
	0, // 0: GetRecognitionStatusResponse.status:type_name -> GetRecognitionStatusResponse.Status
	1, // 1: Speech.StartRecognition:input_type -> StartRecognitionRequest
	3, // 2: Speech.GetRecognitionStatus:input_type -> GetRecognitionStatusRequest
	5, // 3: Speech.StopRecognition:input_type -> StopRecognitionRequest
	2, // 4: Speech.StartRecognition:output_type -> StartRecognitionResponse
	4, // 5: Speech.GetRecognitionStatus:output_type -> GetRecognitionStatusResponse
	6, // 6: Speech.StopRecognition:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rms_speech_proto_init() }
func file_rms_speech_proto_init() {
	if File_rms_speech_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rms_speech_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRecognitionRequest); i {
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
		file_rms_speech_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRecognitionResponse); i {
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
		file_rms_speech_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecognitionStatusRequest); i {
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
		file_rms_speech_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecognitionStatusResponse); i {
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
		file_rms_speech_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRecognitionRequest); i {
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
			RawDescriptor: file_rms_speech_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rms_speech_proto_goTypes,
		DependencyIndexes: file_rms_speech_proto_depIdxs,
		EnumInfos:         file_rms_speech_proto_enumTypes,
		MessageInfos:      file_rms_speech_proto_msgTypes,
	}.Build()
	File_rms_speech_proto = out.File
	file_rms_speech_proto_rawDesc = nil
	file_rms_speech_proto_goTypes = nil
	file_rms_speech_proto_depIdxs = nil
}