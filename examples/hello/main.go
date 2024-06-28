package main

import (
	mt "github.com/mintalk/mintui"
	"strconv"
)

func main() {
	screen := mt.NewScreen()
	defer screen.Clean()

	screen.SetLocale(mt.LC_ALL, "")
	screen.SetLFlag(mt.ECHO, false)

	mt.EraseScreen()

	mt.Color(mt.CMR_BLINKING)
	mt.PrivateModeDisable(mt.SHOW_CURSOR)
	mt.Color256F(40)
	mt.ColorTrueB(100, 0, 255)

	window := mt.NewWindow(0, 0, 20, 10)
	window.CursorMove(1, 1)
	window.Print("Hello! E to exit")

	var c int
	for c != 'e' {
		window.BorderThick()
		window.CursorMove(2, 1)
		window.Print(strconv.Itoa(c) + ";")
		//PutChar(c)
		if mt.HasChar() {
			c = mt.GetChar()
			mt.SetTitle(strconv.Itoa(c))
		} else {
			c = 0
		}
	}

	mt.PrivateModeEnable(mt.SHOW_CURSOR)
	mt.Color(mt.CMR)
}
