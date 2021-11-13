package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

type Keys = uint16

const (
	Shift Keys = 56
	F17   Keys = 64
	F18   Keys = 79
	F19   Keys = 80
)

func main() {
	isShiftPressed := false
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		if ev.Rawcode == Shift && ev.Kind == hook.KeyHold {
			isShiftPressed = true
		} else if ev.Rawcode == Shift && ev.Kind == hook.KeyUp {
			isShiftPressed = false
		}

		if ev.Kind != hook.KeyUp {
			continue
		}
		switch ev.Rawcode {
		case F17:
			fallthrough
		case F18:
			fallthrough
		case F19:
			handler(ev.Rawcode, isShiftPressed)
		}
	}
}

func handler(code Keys, isShiftPressed bool) {
	switch code {
	case F17:
		fmt.Println("f17")
	case F18:
		fmt.Println("f18: ", isShiftPressed)
	case F19:
		fmt.Println("f19: ", isShiftPressed)
	}
}
