package main

func main() {
	screen := NewScreen()
	defer screen.Clean()

	screen.SetLocale(LC_ALL, "")
	screen.SetLFlag(ECHO, false)

	EraseScreen()

	Color(CMR_BLINKING)
	PrivateModeDisable(CURSOR_VISIBLE)

	window := Window{0, 0, 10, 10}

	var c int
	for c != 'q' {
		PutChar(c)
		window.Border()
		c = GetChar()
	}

	Color(CMR)
	Color(CM_BLINKING)
	PrivateModeEnable(CURSOR_VISIBLE)
}
