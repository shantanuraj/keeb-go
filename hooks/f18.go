package hooks

import (
	"fmt"
	"os"

	"sraj.me/keeb-go/actions"
)

func HandleF18(isShiftPressed bool) {
	fmt.Println(actions.JoinZoom(os.Getenv("KEEB_ZOOM")))
}
