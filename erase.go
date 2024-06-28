package mintui

func EraseDisplay() {
	PutEsc("[J")
}

func EraseToEOS() {
	PutEsc("[0J")
}

func EraseToSOS() {
	PutEsc("[1J")
}

func EraseScreen() {
	PutEsc("[2J")
}

func EraseToEOL() {
	PutEsc("[0K")
}

func EraseToSOL() {
	PutEsc("[1K")
}

func EraseLine() {
	PutEsc("[2K")
}
