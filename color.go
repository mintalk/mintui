package mintui

import (
	"strconv"
)

type ColorMode uint

const (
	CMR               ColorMode = 0
	CM_BOLD           ColorMode = 1
	CM_DIM            ColorMode = 2
	CM_ITALIC         ColorMode = 3
	CM_UNDERLINE      ColorMode = 4
	CM_BLINKING       ColorMode = 5
	CM_INVERSE        ColorMode = 7
	CM_HIDDEN         ColorMode = 8
	CM_STRIKETHROUGH  ColorMode = 9
	CMR_STRENGTH      ColorMode = 22
	CMR_ITALIC        ColorMode = 23
	CMR_UNDERLINE     ColorMode = 24
	CMR_BLINKING      ColorMode = 25
	CMR_INVERSE       ColorMode = 27
	CMR_HIDDEN        ColorMode = 28
	CMR_STRIKETHROUGH ColorMode = 29

	CMF_BLACK   ColorMode = 30
	CMF_RED     ColorMode = 31
	CMF_GREEN   ColorMode = 32
	CMF_YELLOW  ColorMode = 33
	CMF_BLUE    ColorMode = 34
	CMF_MAGENTA ColorMode = 35
	CMF_CYAN    ColorMode = 36
	CMF_WHITE   ColorMode = 37
	CMF_DEFAULT ColorMode = 39
	CMB_BLACK   ColorMode = 40
	CMB_RED     ColorMode = 41
	CMB_GREEN   ColorMode = 42
	CMB_YELLOW  ColorMode = 43
	CMB_BLUE    ColorMode = 44
	CMB_MAGENTA ColorMode = 45
	CMB_CYAN    ColorMode = 46
	CMB_WHITE   ColorMode = 47
	CMB_DEFAULT ColorMode = 49

	CMF_B_BLACK   ColorMode = 90
	CMF_B_RED     ColorMode = 91
	CMF_B_GREEN   ColorMode = 92
	CMF_B_YELLOW  ColorMode = 93
	CMF_B_BLUE    ColorMode = 94
	CMF_B_MAGENTA ColorMode = 95
	CMF_B_CYAN    ColorMode = 96
	CMF_B_WHITE   ColorMode = 97
	CMB_B_BLACK   ColorMode = 100
	CMB_B_RED     ColorMode = 101
	CMB_B_GREEN   ColorMode = 102
	CMB_B_YELLOW  ColorMode = 103
	CMB_B_BLUE    ColorMode = 104
	CMB_B_MAGENTA ColorMode = 105
	CMB_B_CYAN    ColorMode = 106
	CMB_B_WHITE   ColorMode = 107
)

func Color(modes ...ColorMode) {
	colorString := ""
	for idx, mode := range modes {
		if idx != 0 {
			colorString += ";"
		}
		colorString += strconv.FormatUint(uint64(mode), 10)
	}
	PutEsc("[" + colorString + "m")
}

func Color256F(color uint8) {
	Color(38, 5, ColorMode(color))
}

func Color256B(color uint8) {
	Color(48, 5, ColorMode(color))
}

func ColorTrueF(r uint32, g uint32, b uint32) {
	Color(38, 2, ColorMode(r), ColorMode(g), ColorMode(b))
}

func ColorTrueB(r uint32, g uint32, b uint32) {
	Color(48, 2, ColorMode(r), ColorMode(g), ColorMode(b))
}
