package hooks

import (
	"fmt"
	"os"
)

func HandleF17() {
	playSpotifyPlaylist(os.Getenv("KEEB_PLAYLIST"))
}

func playSpotifyPlaylist(playlist string) (bool, error) {
	fmt.Println("Playing:", playlist)
	var err error
	if _, err = tell(Spotify, fmt.Sprintf("play track \"%s\"", playlist)); err != nil {
		return false, err
	}
	return true, nil
}
