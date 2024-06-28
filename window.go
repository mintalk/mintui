package main

type Window struct {
	line   uint
	column uint
	width  uint
	height uint
}

func (window *Window) CursorMove(line, column uint) {
	CursorMove(window.line+line, window.column+column)
}

func (window *Window) Print(s string) {
	for i := 0; i < len(s); i++ {
		PutChar(int(s[i]))
	}
}

func (window *Window) Border() {
	for i := uint(1); i < window.width-1; i++ {
		window.CursorMove(0, i)
		PutChar('─')
		window.CursorMove(window.height-1, i)
		PutChar('─')
	}
	for i := uint(1); i < window.height-1; i++ {
		window.CursorMove(i, 0)
		PutChar('│')
		window.CursorMove(i, window.width-1)
		PutChar('│')
	}
	window.CursorMove(0, 0)
	PutChar('┌')
	window.CursorMove(0, window.width-1)
	PutChar('┐')
	window.CursorMove(window.height-1, 0)
	PutChar('└')
	window.CursorMove(window.height-1, window.width-1)
	PutChar('┘')
}
