package mintui

import "fmt"

func SetTitle(s string) {
	PutEsc(fmt.Sprintf("]0;%s\a", s))
}
