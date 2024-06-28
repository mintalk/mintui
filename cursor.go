package main

import "fmt"

func CursorHome() {
	PutEsc("[H")
}

func CursorMove(line, column uint) {
	PutEsc(fmt.Sprintf("[%d;%dH", line, column))
	PutEsc(fmt.Sprintf("[%d;%df", line, column))
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
