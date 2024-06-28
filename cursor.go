package mintui

import "fmt"

func CursorHome() {
	PutEsc("[H")
}

func CursorMove(line, column uint) {
	PutEsc(fmt.Sprintf("[%d;%dH", line+1, column+1))
	PutEsc(fmt.Sprintf("[%d;%df", line+1, column+1))
}

func CursorUp(lines uint) {
	PutEsc(fmt.Sprintf("[%dA", lines))
}

func CursorDown(lines uint) {
	PutEsc(fmt.Sprintf("[%dB", lines))
}

func CursorLeft(columns uint) {
	PutEsc(fmt.Sprintf("[%dC", columns))
}

func CursorRight(columns uint) {
	PutEsc(fmt.Sprintf("[%dD", columns))
}
