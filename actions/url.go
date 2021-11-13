package actions

import (
	"fmt"
	"os/exec"
	"runtime"
)

func openURL(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("/usr/bin/open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}
	return cmd.Start()
}

func OpenURLWithMessage(url string, message string) error {
	fmt.Println(message, url)
	return openURL(url)
}
