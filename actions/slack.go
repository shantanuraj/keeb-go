package actions

import (
	"fmt"
)

func OpenSlack(teamId string) error {
	return OpenURLWithMessage(fmt.Sprintf("slack://app?team=%s", teamId), "Opening Slack")
}
