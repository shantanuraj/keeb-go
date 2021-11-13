package actions

import (
	"os/exec"

	"golang.design/x/clipboard"
)

func OpenStackOverflowWithClipboard() error {
	clip := string(clipboard.Read(clipboard.FmtText))
	url := "https://stackoverflow.com/"
	if clip != "" {
		url += "search?q=" + clip
	}
	cmd := exec.Command("/usr/bin/open", url)
	_, err := cmd.CombinedOutput()
	return err
}
