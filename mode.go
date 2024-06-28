package main

import "fmt"

type ScreenMode uint

const (
	T_MCH_40x25      ScreenMode = 0
	T_CLR_40x25      ScreenMode = 1
	T_MCH_80x25      ScreenMode = 2
	T_CLR_80x25      ScreenMode = 3
	G_4CLR_320x200   ScreenMode = 4
	G_MCH_320x200    ScreenMode = 5
	G_MCH_640x200    ScreenMode = 6
	LINE_WRAP        ScreenMode = 7
	G_CLR_320x200    ScreenMode = 13
	G_16CLR_640x200  ScreenMode = 14
	G_2MCH_640x350   ScreenMode = 15
	G_16CLR_640x350  ScreenMode = 16
	G_2MCH_640x280   ScreenMode = 17
	G_16CLR_640x480  ScreenMode = 18
	G_256CLR_320x200 ScreenMode = 19
)

type PrivateMode uint

const (
	CURSOR_VISIBLE PrivateMode = 25
	SCREEN_SAVE    PrivateMode = 47
	ALT_BUFFER                 = 1049
)

func ScreenModeEnable(mode ScreenMode) {
	PutEsc(fmt.Sprintf("[=%dh", int(mode)))
}

func ScreenModeDisable(mode ScreenMode) {
	PutEsc(fmt.Sprintf("[=%dl", int(mode)))
}

func PrivateModeEnable(mode PrivateMode) {
	PutEsc(fmt.Sprintf("[?%dh", int(mode)))
}

func PrivateModeDisable(mode PrivateMode) {
	PutEsc(fmt.Sprintf("[?%dl", int(mode)))
}
