package actions

import "fmt"

func PlaySpotifyTrack(track string) error {
	fmt.Println("Playing:", track)
	var err error
	if _, err = tell(Spotify, fmt.Sprintf("play track \"%s\"", track)); err != nil {
		return err
	}
	return nil
}
