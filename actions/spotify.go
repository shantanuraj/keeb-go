package actions

import "fmt"

func PlaySpotifyPlaylist(playlist string) (bool, error) {
	fmt.Println("Playing:", playlist)
	var err error
	if _, err = tell(Spotify, fmt.Sprintf("play track \"%s\"", playlist)); err != nil {
		return false, err
	}
	return true, nil
}
