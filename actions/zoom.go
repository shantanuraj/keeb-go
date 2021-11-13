package actions

import (
	"fmt"
	"os/exec"
)

func JoinZoom(link string) (bool, error) {
	fmt.Println("Joining zoom:", link)
	cmd := exec.Command("/usr/bin/open", link)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	return true, nil
}
