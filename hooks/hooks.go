package hooks

import (
	"os"

	"sraj.me/keeb-go/actions"
)

type Keys = uint16
type KeyHandler = func(isShiftPressed bool) error

const (
	Shift Keys = 56
	F17   Keys = 64
	F18   Keys = 79
	F19   Keys = 80
)

var Handlers = map[Keys]KeyHandler{
	F17: func(_ bool) error {
		return actions.OpenURLWithMessage(os.Getenv("KEEB_ZOOM"), "Joining Zoom")
	},
	F18: func(isShiftPressed bool) error {
		playlist := os.Getenv("KEEB_PLAYLIST")
		if isShiftPressed {
			playlist = os.Getenv("KEEB_PLAYLIST_ALT")
		}
		return actions.PlaySpotifyTrack(playlist)
	},
	F19: func(isShiftPressed bool) error {
		if isShiftPressed {
			return actions.OpenSlack(os.Getenv("KEEB_SLACK_TEAM_ID"))
		}
		return actions.OpenStackOverflowWithClipboard()
	},
}
