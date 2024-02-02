// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cctv-backend.proto

package cctv_backend

import (
	fmt "fmt"
	_ "github.com/RacoonMediaServer/rms-packages/pkg/video"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Stream service

func NewStreamEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Stream service

type StreamService interface {
	// Добавить URL на медиа поток
	AddStream(ctx context.Context, in *AddStreamRequest, opts ...client.CallOption) (*AddStreamResponse, error)
	// Получить ссылку на ретрансляцию потока
	GetStreamUri(ctx context.Context, in *GetStreamUriRequest, opts ...client.CallOption) (*GetUriResponse, error)
	// Удалить поток
	RemoveStream(ctx context.Context, in *RemoveStreamRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type streamService struct {
	c    client.Client
	name string
}

func NewStreamService(name string, c client.Client) StreamService {
	return &streamService{
		c:    c,
		name: name,
	}
}

func (c *streamService) AddStream(ctx context.Context, in *AddStreamRequest, opts ...client.CallOption) (*AddStreamResponse, error) {
	req := c.c.NewRequest(c.name, "Stream.AddStream", in)
	out := new(AddStreamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamService) GetStreamUri(ctx context.Context, in *GetStreamUriRequest, opts ...client.CallOption) (*GetUriResponse, error) {
	req := c.c.NewRequest(c.name, "Stream.GetStreamUri", in)
	out := new(GetUriResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamService) RemoveStream(ctx context.Context, in *RemoveStreamRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Stream.RemoveStream", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stream service

type StreamHandler interface {
	// Добавить URL на медиа поток
	AddStream(context.Context, *AddStreamRequest, *AddStreamResponse) error
	// Получить ссылку на ретрансляцию потока
	GetStreamUri(context.Context, *GetStreamUriRequest, *GetUriResponse) error
	// Удалить поток
	RemoveStream(context.Context, *RemoveStreamRequest, *emptypb.Empty) error
}

func RegisterStreamHandler(s server.Server, hdlr StreamHandler, opts ...server.HandlerOption) error {
	type stream interface {
		AddStream(ctx context.Context, in *AddStreamRequest, out *AddStreamResponse) error
		GetStreamUri(ctx context.Context, in *GetStreamUriRequest, out *GetUriResponse) error
		RemoveStream(ctx context.Context, in *RemoveStreamRequest, out *emptypb.Empty) error
	}
	type Stream struct {
		stream
	}
	h := &streamHandler{hdlr}
	return s.Handle(s.NewHandler(&Stream{h}, opts...))
}

type streamHandler struct {
	StreamHandler
}

func (h *streamHandler) AddStream(ctx context.Context, in *AddStreamRequest, out *AddStreamResponse) error {
	return h.StreamHandler.AddStream(ctx, in, out)
}

func (h *streamHandler) GetStreamUri(ctx context.Context, in *GetStreamUriRequest, out *GetUriResponse) error {
	return h.StreamHandler.GetStreamUri(ctx, in, out)
}

func (h *streamHandler) RemoveStream(ctx context.Context, in *RemoveStreamRequest, out *emptypb.Empty) error {
	return h.StreamHandler.RemoveStream(ctx, in, out)
}

// Api Endpoints for Recording service

func NewRecordingEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Recording service

type RecordingService interface {
	// Поставить поток на запись
	AddRecording(ctx context.Context, in *AddRecordingRequest, opts ...client.CallOption) (*AddRecordingResponse, error)
	// Получить ссылку на воспроизведение записи
	GetRecordingUri(ctx context.Context, in *GetRecordingUriRequest, opts ...client.CallOption) (*GetUriResponse, error)
	// Возобновить/остановить запись
	SetRecordingState(ctx context.Context, in *SetRecordingStateRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// Установить качество записи (прореживание)
	SetRecordingQuality(ctx context.Context, in *SetRecordingQualityRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// Удалить запись
	RemoveRecording(ctx context.Context, in *RemoveRecordingRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type recordingService struct {
	c    client.Client
	name string
}

func NewRecordingService(name string, c client.Client) RecordingService {
	return &recordingService{
		c:    c,
		name: name,
	}
}

func (c *recordingService) AddRecording(ctx context.Context, in *AddRecordingRequest, opts ...client.CallOption) (*AddRecordingResponse, error) {
	req := c.c.NewRequest(c.name, "Recording.AddRecording", in)
	out := new(AddRecordingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordingService) GetRecordingUri(ctx context.Context, in *GetRecordingUriRequest, opts ...client.CallOption) (*GetUriResponse, error) {
	req := c.c.NewRequest(c.name, "Recording.GetRecordingUri", in)
	out := new(GetUriResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordingService) SetRecordingState(ctx context.Context, in *SetRecordingStateRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Recording.SetRecordingState", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordingService) SetRecordingQuality(ctx context.Context, in *SetRecordingQualityRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Recording.SetRecordingQuality", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordingService) RemoveRecording(ctx context.Context, in *RemoveRecordingRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Recording.RemoveRecording", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Recording service

type RecordingHandler interface {
	// Поставить поток на запись
	AddRecording(context.Context, *AddRecordingRequest, *AddRecordingResponse) error
	// Получить ссылку на воспроизведение записи
	GetRecordingUri(context.Context, *GetRecordingUriRequest, *GetUriResponse) error
	// Возобновить/остановить запись
	SetRecordingState(context.Context, *SetRecordingStateRequest, *emptypb.Empty) error
	// Установить качество записи (прореживание)
	SetRecordingQuality(context.Context, *SetRecordingQualityRequest, *emptypb.Empty) error
	// Удалить запись
	RemoveRecording(context.Context, *RemoveRecordingRequest, *emptypb.Empty) error
}

func RegisterRecordingHandler(s server.Server, hdlr RecordingHandler, opts ...server.HandlerOption) error {
	type recording interface {
		AddRecording(ctx context.Context, in *AddRecordingRequest, out *AddRecordingResponse) error
		GetRecordingUri(ctx context.Context, in *GetRecordingUriRequest, out *GetUriResponse) error
		SetRecordingState(ctx context.Context, in *SetRecordingStateRequest, out *emptypb.Empty) error
		SetRecordingQuality(ctx context.Context, in *SetRecordingQualityRequest, out *emptypb.Empty) error
		RemoveRecording(ctx context.Context, in *RemoveRecordingRequest, out *emptypb.Empty) error
	}
	type Recording struct {
		recording
	}
	h := &recordingHandler{hdlr}
	return s.Handle(s.NewHandler(&Recording{h}, opts...))
}

type recordingHandler struct {
	RecordingHandler
}

func (h *recordingHandler) AddRecording(ctx context.Context, in *AddRecordingRequest, out *AddRecordingResponse) error {
	return h.RecordingHandler.AddRecording(ctx, in, out)
}

func (h *recordingHandler) GetRecordingUri(ctx context.Context, in *GetRecordingUriRequest, out *GetUriResponse) error {
	return h.RecordingHandler.GetRecordingUri(ctx, in, out)
}

func (h *recordingHandler) SetRecordingState(ctx context.Context, in *SetRecordingStateRequest, out *emptypb.Empty) error {
	return h.RecordingHandler.SetRecordingState(ctx, in, out)
}

func (h *recordingHandler) SetRecordingQuality(ctx context.Context, in *SetRecordingQualityRequest, out *emptypb.Empty) error {
	return h.RecordingHandler.SetRecordingQuality(ctx, in, out)
}

func (h *recordingHandler) RemoveRecording(ctx context.Context, in *RemoveRecordingRequest, out *emptypb.Empty) error {
	return h.RecordingHandler.RemoveRecording(ctx, in, out)
}

// Api Endpoints for System service

func NewSystemEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for System service

type SystemService interface {
	// Получить бекап настроек, БД и пр.
	GetBackup(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetBackupResponse, error)
}

type systemService struct {
	c    client.Client
	name string
}

func NewSystemService(name string, c client.Client) SystemService {
	return &systemService{
		c:    c,
		name: name,
	}
}

func (c *systemService) GetBackup(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetBackupResponse, error) {
	req := c.c.NewRequest(c.name, "System.GetBackup", in)
	out := new(GetBackupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for System service

type SystemHandler interface {
	// Получить бекап настроек, БД и пр.
	GetBackup(context.Context, *emptypb.Empty, *GetBackupResponse) error
}

func RegisterSystemHandler(s server.Server, hdlr SystemHandler, opts ...server.HandlerOption) error {
	type system interface {
		GetBackup(ctx context.Context, in *emptypb.Empty, out *GetBackupResponse) error
	}
	type System struct {
		system
	}
	h := &systemHandler{hdlr}
	return s.Handle(s.NewHandler(&System{h}, opts...))
}

type systemHandler struct {
	SystemHandler
}

func (h *systemHandler) GetBackup(ctx context.Context, in *emptypb.Empty, out *GetBackupResponse) error {
	return h.SystemHandler.GetBackup(ctx, in, out)
}