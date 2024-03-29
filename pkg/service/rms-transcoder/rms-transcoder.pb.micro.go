// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rms-transcoder.proto

package rms_transcoder

import (
	fmt "fmt"
	_ "github.com/RacoonMediaServer/rms-packages/pkg/media"
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

// Api Endpoints for Profiles service

func NewProfilesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Profiles service

type ProfilesService interface {
	// Добавить профиль с параметрами транскодирования
	AddProfile(ctx context.Context, in *Profile, opts ...client.CallOption) (*emptypb.Empty, error)
	// Получить все профили
	GetProfiles(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetProfilesResponse, error)
	// Обновить профиль
	UpdateProfile(ctx context.Context, in *Profile, opts ...client.CallOption) (*emptypb.Empty, error)
	// Удалить профиль
	RemoveProfile(ctx context.Context, in *RemoveProfileRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type profilesService struct {
	c    client.Client
	name string
}

func NewProfilesService(name string, c client.Client) ProfilesService {
	return &profilesService{
		c:    c,
		name: name,
	}
}

func (c *profilesService) AddProfile(ctx context.Context, in *Profile, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Profiles.AddProfile", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesService) GetProfiles(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetProfilesResponse, error) {
	req := c.c.NewRequest(c.name, "Profiles.GetProfiles", in)
	out := new(GetProfilesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesService) UpdateProfile(ctx context.Context, in *Profile, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Profiles.UpdateProfile", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesService) RemoveProfile(ctx context.Context, in *RemoveProfileRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Profiles.RemoveProfile", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profiles service

type ProfilesHandler interface {
	// Добавить профиль с параметрами транскодирования
	AddProfile(context.Context, *Profile, *emptypb.Empty) error
	// Получить все профили
	GetProfiles(context.Context, *emptypb.Empty, *GetProfilesResponse) error
	// Обновить профиль
	UpdateProfile(context.Context, *Profile, *emptypb.Empty) error
	// Удалить профиль
	RemoveProfile(context.Context, *RemoveProfileRequest, *emptypb.Empty) error
}

func RegisterProfilesHandler(s server.Server, hdlr ProfilesHandler, opts ...server.HandlerOption) error {
	type profiles interface {
		AddProfile(ctx context.Context, in *Profile, out *emptypb.Empty) error
		GetProfiles(ctx context.Context, in *emptypb.Empty, out *GetProfilesResponse) error
		UpdateProfile(ctx context.Context, in *Profile, out *emptypb.Empty) error
		RemoveProfile(ctx context.Context, in *RemoveProfileRequest, out *emptypb.Empty) error
	}
	type Profiles struct {
		profiles
	}
	h := &profilesHandler{hdlr}
	return s.Handle(s.NewHandler(&Profiles{h}, opts...))
}

type profilesHandler struct {
	ProfilesHandler
}

func (h *profilesHandler) AddProfile(ctx context.Context, in *Profile, out *emptypb.Empty) error {
	return h.ProfilesHandler.AddProfile(ctx, in, out)
}

func (h *profilesHandler) GetProfiles(ctx context.Context, in *emptypb.Empty, out *GetProfilesResponse) error {
	return h.ProfilesHandler.GetProfiles(ctx, in, out)
}

func (h *profilesHandler) UpdateProfile(ctx context.Context, in *Profile, out *emptypb.Empty) error {
	return h.ProfilesHandler.UpdateProfile(ctx, in, out)
}

func (h *profilesHandler) RemoveProfile(ctx context.Context, in *RemoveProfileRequest, out *emptypb.Empty) error {
	return h.ProfilesHandler.RemoveProfile(ctx, in, out)
}

// Api Endpoints for Transcoder service

func NewTranscoderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Transcoder service

type TranscoderService interface {
	// Добавить задачу транскодирования
	AddJob(ctx context.Context, in *AddJobRequest, opts ...client.CallOption) (*AddJobResponse, error)
	// Получить статус задачи
	GetJob(ctx context.Context, in *GetJobRequest, opts ...client.CallOption) (*GetJobResponse, error)
	// Отменить задачу (должно вызываться для каждой запущенной задачи)
	CancelJob(ctx context.Context, in *CancelJobRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type transcoderService struct {
	c    client.Client
	name string
}

func NewTranscoderService(name string, c client.Client) TranscoderService {
	return &transcoderService{
		c:    c,
		name: name,
	}
}

func (c *transcoderService) AddJob(ctx context.Context, in *AddJobRequest, opts ...client.CallOption) (*AddJobResponse, error) {
	req := c.c.NewRequest(c.name, "Transcoder.AddJob", in)
	out := new(AddJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transcoderService) GetJob(ctx context.Context, in *GetJobRequest, opts ...client.CallOption) (*GetJobResponse, error) {
	req := c.c.NewRequest(c.name, "Transcoder.GetJob", in)
	out := new(GetJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transcoderService) CancelJob(ctx context.Context, in *CancelJobRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Transcoder.CancelJob", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Transcoder service

type TranscoderHandler interface {
	// Добавить задачу транскодирования
	AddJob(context.Context, *AddJobRequest, *AddJobResponse) error
	// Получить статус задачи
	GetJob(context.Context, *GetJobRequest, *GetJobResponse) error
	// Отменить задачу (должно вызываться для каждой запущенной задачи)
	CancelJob(context.Context, *CancelJobRequest, *emptypb.Empty) error
}

func RegisterTranscoderHandler(s server.Server, hdlr TranscoderHandler, opts ...server.HandlerOption) error {
	type transcoder interface {
		AddJob(ctx context.Context, in *AddJobRequest, out *AddJobResponse) error
		GetJob(ctx context.Context, in *GetJobRequest, out *GetJobResponse) error
		CancelJob(ctx context.Context, in *CancelJobRequest, out *emptypb.Empty) error
	}
	type Transcoder struct {
		transcoder
	}
	h := &transcoderHandler{hdlr}
	return s.Handle(s.NewHandler(&Transcoder{h}, opts...))
}

type transcoderHandler struct {
	TranscoderHandler
}

func (h *transcoderHandler) AddJob(ctx context.Context, in *AddJobRequest, out *AddJobResponse) error {
	return h.TranscoderHandler.AddJob(ctx, in, out)
}

func (h *transcoderHandler) GetJob(ctx context.Context, in *GetJobRequest, out *GetJobResponse) error {
	return h.TranscoderHandler.GetJob(ctx, in, out)
}

func (h *transcoderHandler) CancelJob(ctx context.Context, in *CancelJobRequest, out *emptypb.Empty) error {
	return h.TranscoderHandler.CancelJob(ctx, in, out)
}
