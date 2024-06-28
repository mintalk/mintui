package mintui

// #include <termios.h>
// #include <unistd.h>
// #include <stdio.h>
// #include <locale.h>
import "C"

type Screen struct {
	oldt C.struct_termios
	newt C.struct_termios
}

func NewScreen() *Screen {
	screen := new(Screen)

	C.tcgetattr(C.STDIN_FILENO, &screen.oldt)
	screen.newt = screen.oldt

	screen.SetLFlag(ICANON, false)
	screen.SetLFlag(ECHO, false)

	return screen
}

func (screen *Screen) Clean() {
	C.tcsetattr(C.STDIN_FILENO, C.TCSANOW, &screen.oldt)
}

func (screen *Screen) SetLocale(category LocaleCategory, locale string) {
	C.setlocale(C.int(category), C.CString(locale))
}

func (screen *Screen) SetIFlag(flag iflag, state bool) {
	if state {
		screen.newt.c_iflag |= C.uint(flag)
	} else {
		screen.newt.c_iflag &= ^C.uint(flag)
	}
	C.tcsetattr(C.STDIN_FILENO, C.TCSANOW, &screen.newt)
}

func (screen *Screen) SetOFlag(flag iflag, state bool) {
	if state {
		screen.newt.c_oflag |= C.uint(flag)
	} else {
		screen.newt.c_oflag &= ^C.uint(flag)
	}
	C.tcsetattr(C.STDIN_FILENO, C.TCSANOW, &screen.newt)
}

func (screen *Screen) SetCFlag(flag cflag, state bool) {
	if state {
		screen.newt.c_cflag |= C.uint(flag)
	} else {
		screen.newt.c_cflag &= ^C.uint(flag)
	}
	C.tcsetattr(C.STDIN_FILENO, C.TCSANOW, &screen.newt)
}

func (screen *Screen) SetLFlag(flag lflag, state bool) {
	if state {
		screen.newt.c_lflag |= C.uint(flag)
	} else {
		screen.newt.c_lflag &= ^C.uint(flag)
	}
	C.tcsetattr(C.STDIN_FILENO, C.TCSANOW, &screen.newt)
}

func (screen *Screen) SetCC(char cc, value uint) {
	screen.newt.c_cc[char] = C.cc_t(value)
	C.tcsetattr(C.STDIN_FILENO, C.TCSANOW, &screen.newt)
}
