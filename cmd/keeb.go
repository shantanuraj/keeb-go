package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
	"sraj.me/keeb-go/hooks"
)

func main() {
	isShiftPressed := false
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		if ev.Rawcode == hooks.Shift && ev.Kind == hook.KeyHold {
			isShiftPressed = true
		} else if ev.Rawcode == hooks.Shift && ev.Kind == hook.KeyUp {
			isShiftPressed = false
		}

		if ev.Kind != hook.KeyUp {
			continue
		}
		switch ev.Rawcode {
		case hooks.F17:
			fallthrough
		case hooks.F18:
			fallthrough
		case hooks.F19:
			if err := hooks.Handlers[ev.Rawcode](isShiftPressed); err != nil {
				fmt.Println(err)
			}
		}
	}
}
