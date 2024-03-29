// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rms-backup.proto

package rms_backup

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Api Endpoints for RmsBackup service

func NewRmsBackupEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RmsBackup service

type RmsBackupService interface {
	// Запустить бекап вручную
	LaunchBackup(ctx context.Context, in *LaunchBackupRequest, opts ...client.CallOption) (*LaunchBackupResponse, error)
	// Получить статус процедуры создания бекапа
	GetBackupStatus(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetBackupStatusResponse, error)
	// Получить список существующих бекапов
	GetBackups(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetBackupsResponse, error)
	// Удаление бекапа
	RemoveBackup(ctx context.Context, in *RemoveBackupRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// Получить настройки резервного копирования
	GetBackupSettings(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*BackupSettings, error)
	// Установить настройки резеврного копирования
	SetBackupSettings(ctx context.Context, in *BackupSettings, opts ...client.CallOption) (*emptypb.Empty, error)
}

type rmsBackupService struct {
	c    client.Client
	name string
}

func NewRmsBackupService(name string, c client.Client) RmsBackupService {
	return &rmsBackupService{
		c:    c,
		name: name,
	}
}

func (c *rmsBackupService) LaunchBackup(ctx context.Context, in *LaunchBackupRequest, opts ...client.CallOption) (*LaunchBackupResponse, error) {
	req := c.c.NewRequest(c.name, "RmsBackup.LaunchBackup", in)
	out := new(LaunchBackupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsBackupService) GetBackupStatus(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetBackupStatusResponse, error) {
	req := c.c.NewRequest(c.name, "RmsBackup.GetBackupStatus", in)
	out := new(GetBackupStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsBackupService) GetBackups(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetBackupsResponse, error) {
	req := c.c.NewRequest(c.name, "RmsBackup.GetBackups", in)
	out := new(GetBackupsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsBackupService) RemoveBackup(ctx context.Context, in *RemoveBackupRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "RmsBackup.RemoveBackup", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsBackupService) GetBackupSettings(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*BackupSettings, error) {
	req := c.c.NewRequest(c.name, "RmsBackup.GetBackupSettings", in)
	out := new(BackupSettings)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsBackupService) SetBackupSettings(ctx context.Context, in *BackupSettings, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "RmsBackup.SetBackupSettings", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RmsBackup service

type RmsBackupHandler interface {
	// Запустить бекап вручную
	LaunchBackup(context.Context, *LaunchBackupRequest, *LaunchBackupResponse) error
	// Получить статус процедуры создания бекапа
	GetBackupStatus(context.Context, *emptypb.Empty, *GetBackupStatusResponse) error
	// Получить список существующих бекапов
	GetBackups(context.Context, *emptypb.Empty, *GetBackupsResponse) error
	// Удаление бекапа
	RemoveBackup(context.Context, *RemoveBackupRequest, *emptypb.Empty) error
	// Получить настройки резервного копирования
	GetBackupSettings(context.Context, *emptypb.Empty, *BackupSettings) error
	// Установить настройки резеврного копирования
	SetBackupSettings(context.Context, *BackupSettings, *emptypb.Empty) error
}

func RegisterRmsBackupHandler(s server.Server, hdlr RmsBackupHandler, opts ...server.HandlerOption) error {
	type rmsBackup interface {
		LaunchBackup(ctx context.Context, in *LaunchBackupRequest, out *LaunchBackupResponse) error
		GetBackupStatus(ctx context.Context, in *emptypb.Empty, out *GetBackupStatusResponse) error
		GetBackups(ctx context.Context, in *emptypb.Empty, out *GetBackupsResponse) error
		RemoveBackup(ctx context.Context, in *RemoveBackupRequest, out *emptypb.Empty) error
		GetBackupSettings(ctx context.Context, in *emptypb.Empty, out *BackupSettings) error
		SetBackupSettings(ctx context.Context, in *BackupSettings, out *emptypb.Empty) error
	}
	type RmsBackup struct {
		rmsBackup
	}
	h := &rmsBackupHandler{hdlr}
	return s.Handle(s.NewHandler(&RmsBackup{h}, opts...))
}

type rmsBackupHandler struct {
	RmsBackupHandler
}

func (h *rmsBackupHandler) LaunchBackup(ctx context.Context, in *LaunchBackupRequest, out *LaunchBackupResponse) error {
	return h.RmsBackupHandler.LaunchBackup(ctx, in, out)
}

func (h *rmsBackupHandler) GetBackupStatus(ctx context.Context, in *emptypb.Empty, out *GetBackupStatusResponse) error {
	return h.RmsBackupHandler.GetBackupStatus(ctx, in, out)
}

func (h *rmsBackupHandler) GetBackups(ctx context.Context, in *emptypb.Empty, out *GetBackupsResponse) error {
	return h.RmsBackupHandler.GetBackups(ctx, in, out)
}

func (h *rmsBackupHandler) RemoveBackup(ctx context.Context, in *RemoveBackupRequest, out *emptypb.Empty) error {
	return h.RmsBackupHandler.RemoveBackup(ctx, in, out)
}

func (h *rmsBackupHandler) GetBackupSettings(ctx context.Context, in *emptypb.Empty, out *BackupSettings) error {
	return h.RmsBackupHandler.GetBackupSettings(ctx, in, out)
}

func (h *rmsBackupHandler) SetBackupSettings(ctx context.Context, in *BackupSettings, out *emptypb.Empty) error {
	return h.RmsBackupHandler.SetBackupSettings(ctx, in, out)
}
