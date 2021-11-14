package actions

import (
	"strings"

	"golang.design/x/clipboard"
)

func OpenStackOverflowWithClipboard() error {
	clip := strings.TrimSpace(string(clipboard.Read(clipboard.FmtText)))
	url := "https://stackoverflow.com/"
	if clip != "" {
		url += "search?q=" + clip
	}
	return OpenURLWithMessage(url, "Opening StackOverflow")
}
