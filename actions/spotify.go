package actions

import (
	"fmt"
	"log"
)

func PlaySpotifyTrack(track string) error {
	log.Println("Playing", track)
	var err error
	if _, err = tell(Spotify, fmt.Sprintf("play track \"%s\"", track)); err != nil {
		return err
	}
	return nil
}
