// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rms-notifier.proto

package rms_notifier

import (
	fmt "fmt"
	_ "github.com/RacoonMediaServer/rms-packages/pkg/communication"
	_ "github.com/RacoonMediaServer/rms-packages/pkg/events"
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

// Api Endpoints for RmsNotifier service

func NewRmsNotifierEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RmsNotifier service

type RmsNotifierService interface {
	// получить текущие настройки
	GetSettings(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*NotifierSettings, error)
	// обновить настройки
	SetSettings(ctx context.Context, in *NotifierSettings, opts ...client.CallOption) (*emptypb.Empty, error)
	// получить события пришедшые за период времени
	GetJournalEvents(ctx context.Context, in *GetJournalEventsRequest, opts ...client.CallOption) (*GetJournalEventsResponse, error)
	// получить события пришедшые за период времени в виде контента
	GetJournal(ctx context.Context, in *GetJournalEventsRequest, opts ...client.CallOption) (*GetJournalResponse, error)
}

type rmsNotifierService struct {
	c    client.Client
	name string
}

func NewRmsNotifierService(name string, c client.Client) RmsNotifierService {
	return &rmsNotifierService{
		c:    c,
		name: name,
	}
}

func (c *rmsNotifierService) GetSettings(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*NotifierSettings, error) {
	req := c.c.NewRequest(c.name, "RmsNotifier.GetSettings", in)
	out := new(NotifierSettings)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsNotifierService) SetSettings(ctx context.Context, in *NotifierSettings, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "RmsNotifier.SetSettings", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsNotifierService) GetJournalEvents(ctx context.Context, in *GetJournalEventsRequest, opts ...client.CallOption) (*GetJournalEventsResponse, error) {
	req := c.c.NewRequest(c.name, "RmsNotifier.GetJournalEvents", in)
	out := new(GetJournalEventsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsNotifierService) GetJournal(ctx context.Context, in *GetJournalEventsRequest, opts ...client.CallOption) (*GetJournalResponse, error) {
	req := c.c.NewRequest(c.name, "RmsNotifier.GetJournal", in)
	out := new(GetJournalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RmsNotifier service

type RmsNotifierHandler interface {
	// получить текущие настройки
	GetSettings(context.Context, *emptypb.Empty, *NotifierSettings) error
	// обновить настройки
	SetSettings(context.Context, *NotifierSettings, *emptypb.Empty) error
	// получить события пришедшые за период времени
	GetJournalEvents(context.Context, *GetJournalEventsRequest, *GetJournalEventsResponse) error
	// получить события пришедшые за период времени в виде контента
	GetJournal(context.Context, *GetJournalEventsRequest, *GetJournalResponse) error
}

func RegisterRmsNotifierHandler(s server.Server, hdlr RmsNotifierHandler, opts ...server.HandlerOption) error {
	type rmsNotifier interface {
		GetSettings(ctx context.Context, in *emptypb.Empty, out *NotifierSettings) error
		SetSettings(ctx context.Context, in *NotifierSettings, out *emptypb.Empty) error
		GetJournalEvents(ctx context.Context, in *GetJournalEventsRequest, out *GetJournalEventsResponse) error
		GetJournal(ctx context.Context, in *GetJournalEventsRequest, out *GetJournalResponse) error
	}
	type RmsNotifier struct {
		rmsNotifier
	}
	h := &rmsNotifierHandler{hdlr}
	return s.Handle(s.NewHandler(&RmsNotifier{h}, opts...))
}

type rmsNotifierHandler struct {
	RmsNotifierHandler
}

func (h *rmsNotifierHandler) GetSettings(ctx context.Context, in *emptypb.Empty, out *NotifierSettings) error {
	return h.RmsNotifierHandler.GetSettings(ctx, in, out)
}

func (h *rmsNotifierHandler) SetSettings(ctx context.Context, in *NotifierSettings, out *emptypb.Empty) error {
	return h.RmsNotifierHandler.SetSettings(ctx, in, out)
}

func (h *rmsNotifierHandler) GetJournalEvents(ctx context.Context, in *GetJournalEventsRequest, out *GetJournalEventsResponse) error {
	return h.RmsNotifierHandler.GetJournalEvents(ctx, in, out)
}

func (h *rmsNotifierHandler) GetJournal(ctx context.Context, in *GetJournalEventsRequest, out *GetJournalResponse) error {
	return h.RmsNotifierHandler.GetJournal(ctx, in, out)
}
