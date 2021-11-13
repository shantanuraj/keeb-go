package actions

import (
	"golang.design/x/clipboard"
)

func OpenStackOverflowWithClipboard() error {
	clip := string(clipboard.Read(clipboard.FmtText))
	url := "https://stackoverflow.com/"
	if clip != "" {
		url += "search?q=" + clip
	}
	return OpenURLWithMessage(url, "Opening StackOverflow")
}
