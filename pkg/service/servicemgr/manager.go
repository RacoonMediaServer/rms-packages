package servicemgr

import (
	rms_backup "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-backup"
	rms_bot_client "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-bot-client"
	rms_bot_server "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-bot-server"
	rms_cctv "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-cctv"
	rms_library "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-library"
	rms_notes "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-notes"
	rms_notifier "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-notifier"
	rms_torrent "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-torrent"
	rms_users "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-users"
	"go-micro.dev/v4/client"
)

//go:generate ../rms-users/gen.sh
//go:generate ../rms-bot-server/gen.sh
//go:generate ../rms-library/gen.sh
//go:generate ../rms-torrent/gen.sh
//go:generate ../rms-bot-client/gen.sh
//go:generate ../rms-notifier/gen.sh
//go:generate ../rms-notes/gen.sh
//go:generate ../rms-cctv/gen.sh
//go:generate ../rms-backup/gen.sh

// ClientFactory can spawn microservices clients
type ClientFactory interface {
	Client() client.Client
	Name() string
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

// NewBotClient creates connection to rms-bot-client service
func (f ServiceFactory) NewBotClient() rms_bot_client.RmsBotClientService {
	return rms_bot_client.NewRmsBotClientService("rms-bot-client", f.constructor.Client())
}

// NewNotifier creates connection to rms-notifier service
func (f ServiceFactory) NewNotifier() rms_notifier.RmsNotifierService {
	return rms_notifier.NewRmsNotifierService("rms-notifier", f.constructor.Client())
}

// NewNotes creates connection to rms-notes service
func (f ServiceFactory) NewNotes() rms_notes.RmsNotesService {
	return rms_notes.NewRmsNotesService("rms-notes", f.constructor.Client())
}

// NewCctv creates connection to rms-cctv service
func (f ServiceFactory) NewCctv() rms_cctv.RmsCctvService {
	return rms_cctv.NewRmsCctvService("rms-cctv", f.constructor.Client())
}

// NewBackup creates connection to rms-backup service
func (f ServiceFactory) NewBackup() rms_backup.RmsBackupService {
	return rms_backup.NewRmsBackupService("rms-backup", f.constructor.Client())
}
