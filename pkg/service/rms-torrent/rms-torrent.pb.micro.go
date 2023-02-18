// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rms-torrent.proto

package rms_torrent

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

// Api Endpoints for RmsTorrent service

func NewRmsTorrentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RmsTorrent service

type RmsTorrentService interface {
	// Скачать торрент
	Download(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (*DownloadResponse, error)
	GetTorrentInfo(ctx context.Context, in *GetTorrentInfoRequest, opts ...client.CallOption) (*TorrentInfo, error)
	GetTorrents(ctx context.Context, in *GetTorrentsRequest, opts ...client.CallOption) (*GetTorrentsResponse, error)
	RemoveTorrent(ctx context.Context, in *RemoveTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	UpPriority(ctx context.Context, in *UpPriorityRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type rmsTorrentService struct {
	c    client.Client
	name string
}

func NewRmsTorrentService(name string, c client.Client) RmsTorrentService {
	return &rmsTorrentService{
		c:    c,
		name: name,
	}
}

func (c *rmsTorrentService) Download(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (*DownloadResponse, error) {
	req := c.c.NewRequest(c.name, "RmsTorrent.Download", in)
	out := new(DownloadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsTorrentService) GetTorrentInfo(ctx context.Context, in *GetTorrentInfoRequest, opts ...client.CallOption) (*TorrentInfo, error) {
	req := c.c.NewRequest(c.name, "RmsTorrent.GetTorrentInfo", in)
	out := new(TorrentInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsTorrentService) GetTorrents(ctx context.Context, in *GetTorrentsRequest, opts ...client.CallOption) (*GetTorrentsResponse, error) {
	req := c.c.NewRequest(c.name, "RmsTorrent.GetTorrents", in)
	out := new(GetTorrentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsTorrentService) RemoveTorrent(ctx context.Context, in *RemoveTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "RmsTorrent.RemoveTorrent", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rmsTorrentService) UpPriority(ctx context.Context, in *UpPriorityRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "RmsTorrent.UpPriority", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RmsTorrent service

type RmsTorrentHandler interface {
	// Скачать торрент
	Download(context.Context, *DownloadRequest, *DownloadResponse) error
	GetTorrentInfo(context.Context, *GetTorrentInfoRequest, *TorrentInfo) error
	GetTorrents(context.Context, *GetTorrentsRequest, *GetTorrentsResponse) error
	RemoveTorrent(context.Context, *RemoveTorrentRequest, *emptypb.Empty) error
	UpPriority(context.Context, *UpPriorityRequest, *emptypb.Empty) error
}

func RegisterRmsTorrentHandler(s server.Server, hdlr RmsTorrentHandler, opts ...server.HandlerOption) error {
	type rmsTorrent interface {
		Download(ctx context.Context, in *DownloadRequest, out *DownloadResponse) error
		GetTorrentInfo(ctx context.Context, in *GetTorrentInfoRequest, out *TorrentInfo) error
		GetTorrents(ctx context.Context, in *GetTorrentsRequest, out *GetTorrentsResponse) error
		RemoveTorrent(ctx context.Context, in *RemoveTorrentRequest, out *emptypb.Empty) error
		UpPriority(ctx context.Context, in *UpPriorityRequest, out *emptypb.Empty) error
	}
	type RmsTorrent struct {
		rmsTorrent
	}
	h := &rmsTorrentHandler{hdlr}
	return s.Handle(s.NewHandler(&RmsTorrent{h}, opts...))
}

type rmsTorrentHandler struct {
	RmsTorrentHandler
}

func (h *rmsTorrentHandler) Download(ctx context.Context, in *DownloadRequest, out *DownloadResponse) error {
	return h.RmsTorrentHandler.Download(ctx, in, out)
}

func (h *rmsTorrentHandler) GetTorrentInfo(ctx context.Context, in *GetTorrentInfoRequest, out *TorrentInfo) error {
	return h.RmsTorrentHandler.GetTorrentInfo(ctx, in, out)
}

func (h *rmsTorrentHandler) GetTorrents(ctx context.Context, in *GetTorrentsRequest, out *GetTorrentsResponse) error {
	return h.RmsTorrentHandler.GetTorrents(ctx, in, out)
}

func (h *rmsTorrentHandler) RemoveTorrent(ctx context.Context, in *RemoveTorrentRequest, out *emptypb.Empty) error {
	return h.RmsTorrentHandler.RemoveTorrent(ctx, in, out)
}

func (h *rmsTorrentHandler) UpPriority(ctx context.Context, in *UpPriorityRequest, out *emptypb.Empty) error {
	return h.RmsTorrentHandler.UpPriority(ctx, in, out)
}
