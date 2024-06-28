package main

import (
	"strconv"
)

func main() {
	screen := NewScreen()
	defer screen.Clean()

	screen.SetLocale(LC_ALL, "")
	screen.SetLFlag(ECHO, false)

	EraseScreen()

	Color(CMR_BLINKING)
	PrivateModeDisable(CURSOR_VISIBLE)

	window := Window{0, 0, 10, 10}
	window.Border()
	window.CursorMove(1, 1)
	window.Print("Hello!")

	var c int
	for c != 'q' {
		window.Border()
		window.CursorMove(2, 1)
		window.Print(strconv.Itoa(c) + ";")
		//PutChar(c)
		if HasChar() {
			c = GetChar()
		} else {
			c = 0
		}
	}

	Color(CMR)
	Color(CM_BLINKING)
	PrivateModeEnable(CURSOR_VISIBLE)
}
