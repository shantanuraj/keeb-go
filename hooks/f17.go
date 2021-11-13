package hooks

import (
	"os"

	"sraj.me/keeb-go/actions"
)

func HandleF17() {
	actions.PlaySpotifyPlaylist(os.Getenv("KEEB_PLAYLIST"))
}
