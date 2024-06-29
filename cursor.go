package mintui

import "fmt"

func CursorHome() {
	PutEsc("[H")
}

func CursorMove(x, y uint) {
	PutEsc(fmt.Sprintf("[%d;%dH", y+1, x+1))
	PutEsc(fmt.Sprintf("[%d;%df", y+1, x+1))
}

func CursorUp(dy uint) {
	PutEsc(fmt.Sprintf("[%dA", dy))
}

func CursorDown(dy uint) {
	PutEsc(fmt.Sprintf("[%dB", dy))
}

func CursorLeft(dx uint) {
	PutEsc(fmt.Sprintf("[%dC", dx))
}

func CursorRight(dx uint) {
	PutEsc(fmt.Sprintf("[%dD", dx))
}

func CursorPos() (x, y uint) {
	PutEsc("[6n")
	fmt.Scanf("\033[%d;%dR", &y, &x)
	return x - 1, y - 1
}
