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
		return actions.JoinZoom(os.Getenv("KEEB_ZOOM"))
	},
	F18: func(isShiftPressed bool) error {
		if isShiftPressed {
			// TODO implement
		}
		return actions.PlaySpotifyPlaylist(os.Getenv("KEEB_PLAYLIST"))
	},
	F19: func(isShiftPressed bool) error {
		if isShiftPressed {
			// TODO implement
		} else {
			// TODO implement
		}
		return nil
	},
}
