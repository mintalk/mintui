package main

import "C"

// #include <termios.h>
// #include <unistd.h>
// #include <stdio.h>
// #include <wchar.h>
// #include <sys/select.h>
/*
int select_input() {
    fd_set readfds;

	struct timeval tv;

	tv.tv_sec = 0;
	tv.tv_usec = 0;

    FD_ZERO(&readfds);
    FD_SET(STDIN_FILENO, &readfds);

    return select(STDIN_FILENO + 1, &readfds, NULL, NULL, &tv);
}
*/
import "C"

func HasChar() bool {
	return C.select_input() > 0
}

func GetChar() int {
	char := C.getwchar()
	if char == C.WEOF {
		return 0
	}
	return int(char)
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
