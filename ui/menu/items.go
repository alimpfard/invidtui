package menu

import (
	"github.com/darkhz/invidtui/cmd"
	"github.com/darkhz/invidtui/ui/app"
)

// Items describes the menu items.
var Items = &app.MenuData{
	Items: map[cmd.KeyContext][]cmd.Key{
		cmd.KeyContextApp: {
			cmd.KeyDashboard,
			cmd.KeyCancel,
			cmd.KeySuspend,
			cmd.KeyDownloadView,
			cmd.KeyDownloadOptions,
			cmd.KeyInstancesList,
			cmd.KeyQuit,
		},
		cmd.KeyContextStart: {
			cmd.KeyQuery,
		},
		cmd.KeyContextFiles: {
			cmd.KeyFilebrowserDirForward,
			cmd.KeyFilebrowserDirBack,
			cmd.KeyFilebrowserToggleHidden,
			cmd.KeyFilebrowserNewFolder,
			cmd.KeyFilebrowserRename,
			cmd.KeyClose,
		},
		cmd.KeyContextPlaylist: {
			cmd.KeyComments,
			cmd.KeyLink,
			cmd.KeyAdd,
			cmd.KeyRemove,
			cmd.KeyLoadMore,
			cmd.KeyPlaylistSave,
			cmd.KeyDownloadOptions,
			cmd.KeyClose,
		},
		cmd.KeyContextComments: {
			cmd.KeyCommentReplies,
			cmd.KeyClose,
		},
		cmd.KeyContextDownloads: {
			cmd.KeyDownloadOptionSelect,
			cmd.KeyDownloadChangeDir,
			cmd.KeyDownloadCancel,
			cmd.KeyClose,
		},
		cmd.KeyContextSearch: {
			cmd.KeySearchStart,
			cmd.KeyQuery,
			cmd.KeyLoadMore,
			cmd.KeySearchSwitchMode,
			cmd.KeySearchSuggestions,
			cmd.KeySearchParameters,
			cmd.KeyComments,
			cmd.KeyLink,
			cmd.KeyPlaylist,
			cmd.KeyChannelVideos,
			cmd.KeyChannelPlaylists,
			cmd.KeyAdd,
			cmd.KeyDownloadOptions,
		},
		cmd.KeyContextChannel: {
			cmd.KeySwitchTab,
			cmd.KeyLoadMore,
			cmd.KeyQuery,
			cmd.KeyPlaylist,
			cmd.KeyAdd,
			cmd.KeyComments,
			cmd.KeyLink,
			cmd.KeyDownloadOptions,
			cmd.KeyClose,
		},
		cmd.KeyContextDashboard: {
			cmd.KeySwitchTab,
			cmd.KeyDashboardReload,
			cmd.KeyLoadMore,
			cmd.KeyAdd,
			cmd.KeyComments,
			cmd.KeyPlaylist,
			cmd.KeyDashboardCreatePlaylist,
			cmd.KeyDashboardEditPlaylist,
			cmd.KeyChannelVideos,
			cmd.KeyChannelPlaylists,
			cmd.KeyRemove,
			cmd.KeyClose,
		},
		cmd.KeyContextPlayer: {
			cmd.KeyPlayerOpenPlaylist,
			cmd.KeyQueue,
			cmd.KeyPlayerHistory,
			cmd.KeyPlayerInfo,
			cmd.KeyPlayerInfoChangeQuality,
			cmd.KeyPlayerQueueAudio,
			cmd.KeyPlayerQueueVideo,
			cmd.KeyPlayerPlayAudio,
			cmd.KeyPlayerPlayVideo,
			cmd.KeyAudioURL,
			cmd.KeyVideoURL,
		},
		cmd.KeyContextQueue: {
			cmd.KeyQueuePlayMove,
			cmd.KeyQueueSave,
			cmd.KeyQueueAppend,
			cmd.KeyPlayerQueueAudio,
			cmd.KeyPlayerQueueVideo,
			cmd.KeyQueueDelete,
			cmd.KeyQueueMove,
			cmd.KeyQueueCancel,
			cmd.KeyClose,
		},
		cmd.KeyContextHistory: {
			cmd.KeyQuery,
			cmd.KeyChannelVideos,
			cmd.KeyChannelPlaylists,
			cmd.KeyClose,
		},
	},
	Visible: map[cmd.Key]func(menuType string) bool{
		cmd.KeyDownloadChangeDir:       downloadView,
		cmd.KeyDownloadView:            downloadView,
		cmd.KeyDownloadOptions:         downloadOptions,
		cmd.KeyComments:                isVideo,
		cmd.KeyLink:                    isVideo,
		cmd.KeyDownloadCancel:          downloadViewVisible,
		cmd.KeyAdd:                     add,
		cmd.KeyRemove:                  remove,
		cmd.KeyPlaylist:                isPlaylist,
		cmd.KeyChannelVideos:           isVideoOrChannel,
		cmd.KeyChannelPlaylists:        isVideoOrChannel,
		cmd.KeyQuery:                   query,
		cmd.KeySearchStart:             searchInputFocused,
		cmd.KeySearchSwitchMode:        searchInputFocused,
		cmd.KeySearchSuggestions:       searchInputFocused,
		cmd.KeySearchParameters:        searchInputFocused,
		cmd.KeyDashboardReload:         isDashboardFocused,
		cmd.KeyDashboardCreatePlaylist: createPlaylist,
		cmd.KeyDashboardEditPlaylist:   editPlaylist,
		cmd.KeyQueue:                   playerQueue,
		cmd.KeyPlayerInfo:              isPlaying,
		cmd.KeyPlayerInfoChangeQuality: infoShown,
		cmd.KeyPlayerQueueAudio:        queueMedia,
		cmd.KeyPlayerQueueVideo:        queueMedia,
		cmd.KeyPlayerPlayAudio:         isVideo,
		cmd.KeyPlayerPlayVideo:         isVideo,
	},
}
