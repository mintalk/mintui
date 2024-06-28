package mintui

type Window struct {
	line   uint
	column uint
	width  uint
	height uint
}

func NewWindow(line, column, width, height uint) Window {
	return Window{line, column, width, height}
}

func (window *Window) CursorMove(line, column uint) {
	CursorMove(window.line+line, window.column+column)
}

func (window *Window) Print(s string) {
	for i := 0; i < len(s); i++ {
		PutChar(int(s[i]))
	}
}

func (window *Window) Border(t, b, l, r, tl, tr, bl, br int) {
	for i := uint(1); i < window.width-1; i++ {
		window.CursorMove(0, i)
		PutChar(t)
		window.CursorMove(window.height-1, i)
		PutChar(b)
	}
	for i := uint(1); i < window.height-1; i++ {
		window.CursorMove(i, 0)
		PutChar(l)
		window.CursorMove(i, window.width-1)
		PutChar(r)
	}
	window.CursorMove(0, 0)
	PutChar(tl)
	window.CursorMove(0, window.width-1)
	PutChar(tr)
	window.CursorMove(window.height-1, 0)
	PutChar(bl)
	window.CursorMove(window.height-1, window.width-1)
	PutChar(br)
}

func (window *Window) BorderThin() {
	window.Border('─', '─', '│', '│', '┌', '┐', '└', '┘')
}

func (window *Window) BorderThick() {
	window.Border('━', '━', '┃', '┃', '┏', '┓', '┗', '┛')
}
