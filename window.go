package main

type Window struct {
	line   uint
	column uint
	width  uint
	height uint
}

func (window *Window) MoveCursor(line, column uint) {
	CursorMove(window.line+line, window.column+column)
}

func (window *Window) Border() {
	for i := uint(1); i < window.width-1; i++ {
		window.MoveCursor(0, i)
		PutChar('─')
		window.MoveCursor(window.height-1, i)
		PutChar('─')
	}
	for i := uint(1); i < window.height-1; i++ {
		window.MoveCursor(i, 0)
		PutChar('│')
		window.MoveCursor(i, window.width-1)
		PutChar('│')
	}
	window.MoveCursor(0, 0)
	PutChar('┌')
	window.MoveCursor(0, window.width-1)
	PutChar('┐')
	window.MoveCursor(window.height-1, 0)
	PutChar('└')
	window.MoveCursor(window.height-1, window.width-1)
	PutChar('┘')
}
