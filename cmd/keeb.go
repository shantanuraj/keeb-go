package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
	"sraj.me/keeb-go/hooks"
)

func main() {
	isShiftPressed := false
	events := hook.Start()
	defer hook.End()

	for event := range events {
		if event.Rawcode == hooks.Shift {
			isShiftPressed = event.Kind != hook.KeyUp
		}

		if event.Kind != hook.KeyUp {
			continue
		}

		handler, ok := hooks.Handlers[event.Rawcode]
		if ok {
			if err := handler(isShiftPressed); err != nil {
				fmt.Println(err)
			}
		}
	}
}
