package main

import (
	"fmt"
	mt "github.com/mintalk/mintui"
)

func main() {
	screen := mt.NewScreen()
	defer screen.Clean()

	screen.SetLocale(mt.LC_ALL, "")

	mt.EraseScreen()

	var c int
	for c != 'q' {
		c = mt.GetChar()
		mt.PutChar(c)
		if c == 'f' {
			width, height := mt.TerminalSize()
			fmt.Printf("%d:%d\n", width, height)
		}
	}
}
