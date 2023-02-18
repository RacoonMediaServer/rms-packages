package servicemgr

import (
	rms_bot_server "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-bot-server"
	rms_library "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-library"
	rms_torrent "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-torrent"
	rms_users "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-users"
	"go-micro.dev/v4/client"
)

//go:generate ../rms-users/gen.sh
//go:generate ../rms-bot-server/gen.sh
//go:generate ../rms-library/gen.sh
//go:generate ../rms-torrent/gen.sh

// ClientFactory can spawn microservices clients
type ClientFactory interface {
	Client() client.Client
}

// ServiceFactory can spawn connections to all RMS microservices
type ServiceFactory struct {
	constructor ClientFactory
}

// NewServiceFactory creates new ServiceFactory
func NewServiceFactory(f ClientFactory) ServiceFactory {
	return ServiceFactory{constructor: f}
}

// NewUsers create connection to rms-users service
func (f ServiceFactory) NewUsers() rms_users.RmsUsersService {
	return rms_users.NewRmsUsersService("rms-users", f.constructor.Client())
}

// NewBotServer creates connection to rms-bot-server service
func (f ServiceFactory) NewBotServer() rms_bot_server.RmsBotServerService {
	return rms_bot_server.NewRmsBotServerService("rms-bot-server", f.constructor.Client())
}

// NewLibrary creates connection to rms-library service
func (f ServiceFactory) NewLibrary() rms_library.RmsLibraryService {
	return rms_library.NewRmsLibraryService("rms-library", f.constructor.Client())
}

// NewTorrent creates connection to rms-library service
func (f ServiceFactory) NewTorrent() rms_torrent.RmsTorrentService {
	return rms_torrent.NewRmsTorrentService("rms-torrent", f.constructor.Client())
}
