package actions

import (
	"fmt"
	"os/exec"
)

func JoinZoom(link string) error {
	fmt.Println("Joining zoom:", link)
	cmd := exec.Command("/usr/bin/open", link)
	_, err := cmd.CombinedOutput()
	return err
}
