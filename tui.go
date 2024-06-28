package main

// #include <termios.h>
// #include <unistd.h>
// #include <stdio.h>
// #include <wchar.h>
import "C"

func GetChar() int {
	return int(C.getwchar())
}

func PutChar(char int) {
	C.putwchar(C.wchar_t(char))
}

func PutEsc(s string) {
	PutChar(27) // ESC
	for i := 0; i < len(s); i++ {
		PutChar(int(s[i]))
	}
}
