// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rms-library.proto

package rms_library

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

// Api Endpoints for Movies service

func NewMoviesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Movies service

type MoviesService interface {
	// поиск информации о фильмах и сериалах во внешних источниках
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchMovieResponse, error)
	// скачать указанный фильм или сериал в автоматическом режиме
	DownloadAuto(ctx context.Context, in *DownloadMovieAutoRequest, opts ...client.CallOption) (*DownloadMovieAutoResponse, error)
	// найти варианты раздачи для фильма или сериала
	FindTorrents(ctx context.Context, in *FindMovieTorrentsRequest, opts ...client.CallOption) (*FindTorrentsResponse, error)
	// скачать выбранную раздачу
	Download(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// загрузить торрент-файл фильма или сериала или просто ролика
	Upload(ctx context.Context, in *UploadMovieRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// получить список фильмов/сериалов и пути к их контенту
	List(ctx context.Context, in *GetMoviesRequest, opts ...client.CallOption) (*GetMoviesResponse, error)
	// получить инфу о фильме/сериале, присутствующем в библиотеке по его ID
	Get(ctx context.Context, in *GetMovieRequest, opts ...client.CallOption) (*GetMovieResponse, error)
	// удаление фильмов/сериалов
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// получить список доступных новых сезонов для скачивания
	GetTvSeriesUpdates(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetTvSeriesUpdatesResponse, error)
	// добавить медиа в библиотеку с целью скачивания позднее по запросу
	WatchLater(ctx context.Context, in *WatchLaterRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// посмотреть, что было добавлено в список на скачивание по запросу
	GetWatchList(ctx context.Context, in *GetMoviesRequest, opts ...client.CallOption) (*GetMoviesResponse, error)
}

type moviesService struct {
	c    client.Client
	name string
}

func NewMoviesService(name string, c client.Client) MoviesService {
	return &moviesService{
		c:    c,
		name: name,
	}
}

func (c *moviesService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.Search", in)
	out := new(SearchMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) DownloadAuto(ctx context.Context, in *DownloadMovieAutoRequest, opts ...client.CallOption) (*DownloadMovieAutoResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.DownloadAuto", in)
	out := new(DownloadMovieAutoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) FindTorrents(ctx context.Context, in *FindMovieTorrentsRequest, opts ...client.CallOption) (*FindTorrentsResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.FindTorrents", in)
	out := new(FindTorrentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) Download(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Movies.Download", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) Upload(ctx context.Context, in *UploadMovieRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Movies.Upload", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) List(ctx context.Context, in *GetMoviesRequest, opts ...client.CallOption) (*GetMoviesResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.List", in)
	out := new(GetMoviesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) Get(ctx context.Context, in *GetMovieRequest, opts ...client.CallOption) (*GetMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.Get", in)
	out := new(GetMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Movies.Delete", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) GetTvSeriesUpdates(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetTvSeriesUpdatesResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.GetTvSeriesUpdates", in)
	out := new(GetTvSeriesUpdatesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) WatchLater(ctx context.Context, in *WatchLaterRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Movies.WatchLater", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) GetWatchList(ctx context.Context, in *GetMoviesRequest, opts ...client.CallOption) (*GetMoviesResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.GetWatchList", in)
	out := new(GetMoviesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Movies service

type MoviesHandler interface {
	// поиск информации о фильмах и сериалах во внешних источниках
	Search(context.Context, *SearchRequest, *SearchMovieResponse) error
	// скачать указанный фильм или сериал в автоматическом режиме
	DownloadAuto(context.Context, *DownloadMovieAutoRequest, *DownloadMovieAutoResponse) error
	// найти варианты раздачи для фильма или сериала
	FindTorrents(context.Context, *FindMovieTorrentsRequest, *FindTorrentsResponse) error
	// скачать выбранную раздачу
	Download(context.Context, *DownloadTorrentRequest, *emptypb.Empty) error
	// загрузить торрент-файл фильма или сериала или просто ролика
	Upload(context.Context, *UploadMovieRequest, *emptypb.Empty) error
	// получить список фильмов/сериалов и пути к их контенту
	List(context.Context, *GetMoviesRequest, *GetMoviesResponse) error
	// получить инфу о фильме/сериале, присутствующем в библиотеке по его ID
	Get(context.Context, *GetMovieRequest, *GetMovieResponse) error
	// удаление фильмов/сериалов
	Delete(context.Context, *DeleteRequest, *emptypb.Empty) error
	// получить список доступных новых сезонов для скачивания
	GetTvSeriesUpdates(context.Context, *emptypb.Empty, *GetTvSeriesUpdatesResponse) error
	// добавить медиа в библиотеку с целью скачивания позднее по запросу
	WatchLater(context.Context, *WatchLaterRequest, *emptypb.Empty) error
	// посмотреть, что было добавлено в список на скачивание по запросу
	GetWatchList(context.Context, *GetMoviesRequest, *GetMoviesResponse) error
}

func RegisterMoviesHandler(s server.Server, hdlr MoviesHandler, opts ...server.HandlerOption) error {
	type movies interface {
		Search(ctx context.Context, in *SearchRequest, out *SearchMovieResponse) error
		DownloadAuto(ctx context.Context, in *DownloadMovieAutoRequest, out *DownloadMovieAutoResponse) error
		FindTorrents(ctx context.Context, in *FindMovieTorrentsRequest, out *FindTorrentsResponse) error
		Download(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error
		Upload(ctx context.Context, in *UploadMovieRequest, out *emptypb.Empty) error
		List(ctx context.Context, in *GetMoviesRequest, out *GetMoviesResponse) error
		Get(ctx context.Context, in *GetMovieRequest, out *GetMovieResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *emptypb.Empty) error
		GetTvSeriesUpdates(ctx context.Context, in *emptypb.Empty, out *GetTvSeriesUpdatesResponse) error
		WatchLater(ctx context.Context, in *WatchLaterRequest, out *emptypb.Empty) error
		GetWatchList(ctx context.Context, in *GetMoviesRequest, out *GetMoviesResponse) error
	}
	type Movies struct {
		movies
	}
	h := &moviesHandler{hdlr}
	return s.Handle(s.NewHandler(&Movies{h}, opts...))
}

type moviesHandler struct {
	MoviesHandler
}

func (h *moviesHandler) Search(ctx context.Context, in *SearchRequest, out *SearchMovieResponse) error {
	return h.MoviesHandler.Search(ctx, in, out)
}

func (h *moviesHandler) DownloadAuto(ctx context.Context, in *DownloadMovieAutoRequest, out *DownloadMovieAutoResponse) error {
	return h.MoviesHandler.DownloadAuto(ctx, in, out)
}

func (h *moviesHandler) FindTorrents(ctx context.Context, in *FindMovieTorrentsRequest, out *FindTorrentsResponse) error {
	return h.MoviesHandler.FindTorrents(ctx, in, out)
}

func (h *moviesHandler) Download(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error {
	return h.MoviesHandler.Download(ctx, in, out)
}

func (h *moviesHandler) Upload(ctx context.Context, in *UploadMovieRequest, out *emptypb.Empty) error {
	return h.MoviesHandler.Upload(ctx, in, out)
}

func (h *moviesHandler) List(ctx context.Context, in *GetMoviesRequest, out *GetMoviesResponse) error {
	return h.MoviesHandler.List(ctx, in, out)
}

func (h *moviesHandler) Get(ctx context.Context, in *GetMovieRequest, out *GetMovieResponse) error {
	return h.MoviesHandler.Get(ctx, in, out)
}

func (h *moviesHandler) Delete(ctx context.Context, in *DeleteRequest, out *emptypb.Empty) error {
	return h.MoviesHandler.Delete(ctx, in, out)
}

func (h *moviesHandler) GetTvSeriesUpdates(ctx context.Context, in *emptypb.Empty, out *GetTvSeriesUpdatesResponse) error {
	return h.MoviesHandler.GetTvSeriesUpdates(ctx, in, out)
}

func (h *moviesHandler) WatchLater(ctx context.Context, in *WatchLaterRequest, out *emptypb.Empty) error {
	return h.MoviesHandler.WatchLater(ctx, in, out)
}

func (h *moviesHandler) GetWatchList(ctx context.Context, in *GetMoviesRequest, out *GetMoviesResponse) error {
	return h.MoviesHandler.GetWatchList(ctx, in, out)
}

// Api Endpoints for Music service

func NewMusicEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Music service

type MusicService interface {
	// поиск информации о музыке во внешних источниках
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchMusicResponse, error)
	// скачать указанный альбом или дискографию в автоматическом режиме
	DownloadAuto(ctx context.Context, in *DownloadMusicAutoRequest, opts ...client.CallOption) (*DownloadMusicAutoResponse, error)
	// найти варианты раздачи для исполнителя или альбома
	FindTorrents(ctx context.Context, in *FindMusicTorrentsRequest, opts ...client.CallOption) (*FindTorrentsResponse, error)
	// скачать выбранную раздачу
	Download(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// загрузить торрент-файл альбома или дискографии
	Upload(ctx context.Context, in *UploadMusicRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// получить список исполнителей
	List(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListArtistsResponse, error)
	// удаление  исполнителей
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// получить список доступных новых сезонов для скачивания
	GetAlbumsUpdates(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetAlbumsUpdatesResponse, error)
	// добавить медиа в библиотеку с целью скачивания позднее по запросу
	ListenLater(ctx context.Context, in *ListenLaterRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// посмотреть, что было добавлено в список на скачивание по запросу
	GetListenList(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetListenListResponse, error)
}

type musicService struct {
	c    client.Client
	name string
}

func NewMusicService(name string, c client.Client) MusicService {
	return &musicService{
		c:    c,
		name: name,
	}
}

func (c *musicService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchMusicResponse, error) {
	req := c.c.NewRequest(c.name, "Music.Search", in)
	out := new(SearchMusicResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicService) DownloadAuto(ctx context.Context, in *DownloadMusicAutoRequest, opts ...client.CallOption) (*DownloadMusicAutoResponse, error) {
	req := c.c.NewRequest(c.name, "Music.DownloadAuto", in)
	out := new(DownloadMusicAutoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicService) FindTorrents(ctx context.Context, in *FindMusicTorrentsRequest, opts ...client.CallOption) (*FindTorrentsResponse, error) {
	req := c.c.NewRequest(c.name, "Music.FindTorrents", in)
	out := new(FindTorrentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicService) Download(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Music.Download", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicService) Upload(ctx context.Context, in *UploadMusicRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Music.Upload", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicService) List(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ListArtistsResponse, error) {
	req := c.c.NewRequest(c.name, "Music.List", in)
	out := new(ListArtistsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Music.Delete", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicService) GetAlbumsUpdates(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetAlbumsUpdatesResponse, error) {
	req := c.c.NewRequest(c.name, "Music.GetAlbumsUpdates", in)
	out := new(GetAlbumsUpdatesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicService) ListenLater(ctx context.Context, in *ListenLaterRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Music.ListenLater", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicService) GetListenList(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetListenListResponse, error) {
	req := c.c.NewRequest(c.name, "Music.GetListenList", in)
	out := new(GetListenListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Music service

type MusicHandler interface {
	// поиск информации о музыке во внешних источниках
	Search(context.Context, *SearchRequest, *SearchMusicResponse) error
	// скачать указанный альбом или дискографию в автоматическом режиме
	DownloadAuto(context.Context, *DownloadMusicAutoRequest, *DownloadMusicAutoResponse) error
	// найти варианты раздачи для исполнителя или альбома
	FindTorrents(context.Context, *FindMusicTorrentsRequest, *FindTorrentsResponse) error
	// скачать выбранную раздачу
	Download(context.Context, *DownloadTorrentRequest, *emptypb.Empty) error
	// загрузить торрент-файл альбома или дискографии
	Upload(context.Context, *UploadMusicRequest, *emptypb.Empty) error
	// получить список исполнителей
	List(context.Context, *emptypb.Empty, *ListArtistsResponse) error
	// удаление  исполнителей
	Delete(context.Context, *DeleteRequest, *emptypb.Empty) error
	// получить список доступных новых сезонов для скачивания
	GetAlbumsUpdates(context.Context, *emptypb.Empty, *GetAlbumsUpdatesResponse) error
	// добавить медиа в библиотеку с целью скачивания позднее по запросу
	ListenLater(context.Context, *ListenLaterRequest, *emptypb.Empty) error
	// посмотреть, что было добавлено в список на скачивание по запросу
	GetListenList(context.Context, *emptypb.Empty, *GetListenListResponse) error
}

func RegisterMusicHandler(s server.Server, hdlr MusicHandler, opts ...server.HandlerOption) error {
	type music interface {
		Search(ctx context.Context, in *SearchRequest, out *SearchMusicResponse) error
		DownloadAuto(ctx context.Context, in *DownloadMusicAutoRequest, out *DownloadMusicAutoResponse) error
		FindTorrents(ctx context.Context, in *FindMusicTorrentsRequest, out *FindTorrentsResponse) error
		Download(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error
		Upload(ctx context.Context, in *UploadMusicRequest, out *emptypb.Empty) error
		List(ctx context.Context, in *emptypb.Empty, out *ListArtistsResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *emptypb.Empty) error
		GetAlbumsUpdates(ctx context.Context, in *emptypb.Empty, out *GetAlbumsUpdatesResponse) error
		ListenLater(ctx context.Context, in *ListenLaterRequest, out *emptypb.Empty) error
		GetListenList(ctx context.Context, in *emptypb.Empty, out *GetListenListResponse) error
	}
	type Music struct {
		music
	}
	h := &musicHandler{hdlr}
	return s.Handle(s.NewHandler(&Music{h}, opts...))
}

type musicHandler struct {
	MusicHandler
}

func (h *musicHandler) Search(ctx context.Context, in *SearchRequest, out *SearchMusicResponse) error {
	return h.MusicHandler.Search(ctx, in, out)
}

func (h *musicHandler) DownloadAuto(ctx context.Context, in *DownloadMusicAutoRequest, out *DownloadMusicAutoResponse) error {
	return h.MusicHandler.DownloadAuto(ctx, in, out)
}

func (h *musicHandler) FindTorrents(ctx context.Context, in *FindMusicTorrentsRequest, out *FindTorrentsResponse) error {
	return h.MusicHandler.FindTorrents(ctx, in, out)
}

func (h *musicHandler) Download(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error {
	return h.MusicHandler.Download(ctx, in, out)
}

func (h *musicHandler) Upload(ctx context.Context, in *UploadMusicRequest, out *emptypb.Empty) error {
	return h.MusicHandler.Upload(ctx, in, out)
}

func (h *musicHandler) List(ctx context.Context, in *emptypb.Empty, out *ListArtistsResponse) error {
	return h.MusicHandler.List(ctx, in, out)
}

func (h *musicHandler) Delete(ctx context.Context, in *DeleteRequest, out *emptypb.Empty) error {
	return h.MusicHandler.Delete(ctx, in, out)
}

func (h *musicHandler) GetAlbumsUpdates(ctx context.Context, in *emptypb.Empty, out *GetAlbumsUpdatesResponse) error {
	return h.MusicHandler.GetAlbumsUpdates(ctx, in, out)
}

func (h *musicHandler) ListenLater(ctx context.Context, in *ListenLaterRequest, out *emptypb.Empty) error {
	return h.MusicHandler.ListenLater(ctx, in, out)
}

func (h *musicHandler) GetListenList(ctx context.Context, in *emptypb.Empty, out *GetListenListResponse) error {
	return h.MusicHandler.GetListenList(ctx, in, out)
}

// Api Endpoints for Torrents service

func NewTorrentsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Torrents service

type TorrentsService interface {
	// найти на торрентах просто какой то контент без привзяки к типу
	FindTorrents(ctx context.Context, in *FindTorrentsRequest, opts ...client.CallOption) (*FindTorrentsResponse, error)
	// скачать указанный торрент
	Download(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// сохранить указанный торрент для дальнейшего скачивания
	Save(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// получить список сохраненных элементов
	GetSavedList(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*FindTorrentsResponse, error)
	// выкачать сохраненный торрент
	DownloadSavedItem(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type torrentsService struct {
	c    client.Client
	name string
}

func NewTorrentsService(name string, c client.Client) TorrentsService {
	return &torrentsService{
		c:    c,
		name: name,
	}
}

func (c *torrentsService) FindTorrents(ctx context.Context, in *FindTorrentsRequest, opts ...client.CallOption) (*FindTorrentsResponse, error) {
	req := c.c.NewRequest(c.name, "Torrents.FindTorrents", in)
	out := new(FindTorrentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torrentsService) Download(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Torrents.Download", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torrentsService) Save(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Torrents.Save", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torrentsService) GetSavedList(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*FindTorrentsResponse, error) {
	req := c.c.NewRequest(c.name, "Torrents.GetSavedList", in)
	out := new(FindTorrentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torrentsService) DownloadSavedItem(ctx context.Context, in *DownloadTorrentRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Torrents.DownloadSavedItem", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Torrents service

type TorrentsHandler interface {
	// найти на торрентах просто какой то контент без привзяки к типу
	FindTorrents(context.Context, *FindTorrentsRequest, *FindTorrentsResponse) error
	// скачать указанный торрент
	Download(context.Context, *DownloadTorrentRequest, *emptypb.Empty) error
	// сохранить указанный торрент для дальнейшего скачивания
	Save(context.Context, *DownloadTorrentRequest, *emptypb.Empty) error
	// получить список сохраненных элементов
	GetSavedList(context.Context, *emptypb.Empty, *FindTorrentsResponse) error
	// выкачать сохраненный торрент
	DownloadSavedItem(context.Context, *DownloadTorrentRequest, *emptypb.Empty) error
}

func RegisterTorrentsHandler(s server.Server, hdlr TorrentsHandler, opts ...server.HandlerOption) error {
	type torrents interface {
		FindTorrents(ctx context.Context, in *FindTorrentsRequest, out *FindTorrentsResponse) error
		Download(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error
		Save(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error
		GetSavedList(ctx context.Context, in *emptypb.Empty, out *FindTorrentsResponse) error
		DownloadSavedItem(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error
	}
	type Torrents struct {
		torrents
	}
	h := &torrentsHandler{hdlr}
	return s.Handle(s.NewHandler(&Torrents{h}, opts...))
}

type torrentsHandler struct {
	TorrentsHandler
}

func (h *torrentsHandler) FindTorrents(ctx context.Context, in *FindTorrentsRequest, out *FindTorrentsResponse) error {
	return h.TorrentsHandler.FindTorrents(ctx, in, out)
}

func (h *torrentsHandler) Download(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error {
	return h.TorrentsHandler.Download(ctx, in, out)
}

func (h *torrentsHandler) Save(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error {
	return h.TorrentsHandler.Save(ctx, in, out)
}

func (h *torrentsHandler) GetSavedList(ctx context.Context, in *emptypb.Empty, out *FindTorrentsResponse) error {
	return h.TorrentsHandler.GetSavedList(ctx, in, out)
}

func (h *torrentsHandler) DownloadSavedItem(ctx context.Context, in *DownloadTorrentRequest, out *emptypb.Empty) error {
	return h.TorrentsHandler.DownloadSavedItem(ctx, in, out)
}
