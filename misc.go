package mintui

import "fmt"

func SetTitle(s string) {
	PutEsc(fmt.Sprintf("]0;%s\a", s))
}

func TerminalSize() (uint, uint) {
	beforeLine, beforeColumn := CursorPos()
	CursorMove(65535, 65535)
	width, height := CursorPos()
	CursorMove(beforeLine, beforeColumn)
	return width + 1, height + 1
}
