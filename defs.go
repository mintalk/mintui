package mintui

// #include <termios.h>
// #include <unistd.h>
// #include <stdio.h>
// #include <locale.h>
import "C"

type LocaleCategory C.int

const (
	LC_COLLATE  LocaleCategory = C.LC_COLLATE
	LC_CTYPE    LocaleCategory = C.LC_CTYPE
	LC_MONETARY LocaleCategory = C.LC_MONETARY
	LC_NUMERIC  LocaleCategory = C.LC_NUMERIC
	LC_TIME     LocaleCategory = C.LC_TIME
	LC_MESSAGES LocaleCategory = C.LC_MESSAGES
	LC_ALL      LocaleCategory = C.LC_ALL
)

type iflag C.uint

const (
	IGNBRK  iflag = C.IGNBRK
	BRKINT  iflag = C.BRKINT
	IGNPAR  iflag = C.IGNPAR
	PARMRK  iflag = C.PARMRK
	INPCK   iflag = C.INPCK
	ISTRIP  iflag = C.ISTRIP
	INLCR   iflag = C.INLCR
	IGNCR   iflag = C.IGNCR
	ICRNL   iflag = C.ICRNL
	IUCLC   iflag = C.IUCLC
	IXON    iflag = C.IXON
	IXANY   iflag = C.IXANY
	IXOFF   iflag = C.IXOFF
	IMAXBEL iflag = C.IMAXBEL
	IUTF8   iflag = C.IUTF8
)

type oflag C.uint

const (
	OPOST  oflag = C.OPOST
	OLCUC  oflag = C.OLCUC
	ONLCR  oflag = C.ONLCR
	OCRNL  oflag = C.OCRNL
	ONOCR  oflag = C.ONOCR
	ONLRET oflag = C.ONLRET
	OFILL  oflag = C.OFILL
	OFDEL  oflag = C.OFDEL
	NLDLY  oflag = C.NLDLY
	CRDLY  oflag = C.CRDLY
	TABDLY oflag = C.TABDLY
	BSDLY  oflag = C.BSDLY
	VTDLY  oflag = C.VTDLY
	FFDLY  oflag = C.FFDLY
)

type cflag C.uint

const (
	CBAUD   cflag = C.CBAUD
	CBAUDEX cflag = C.CBAUDEX
	CSIZE   cflag = C.CSIZE
	CSTOPB  cflag = C.CSTOPB
	CREAD   cflag = C.CREAD
	PARENB  cflag = C.PARENB
	PARODD  cflag = C.PARODD
	HUPCL   cflag = C.HUPCL
	CLOCAL  cflag = C.CLOCAL
	// LOBLK cflag = C.LOBLK
	CIBAUD  cflag = C.CIBAUD
	CMSPAR  cflag = C.CMSPAR
	CRTSCTS cflag = C.CRTSCTS
)

type lflag C.uint

const (
	ISIG    lflag = C.ISIG
	ICANON  lflag = C.ICANON
	XCASE   lflag = C.XCASE
	ECHO    lflag = C.ECHO
	ECHOE   lflag = C.ECHOE
	ECHONL  lflag = C.ECHONL
	ECHOCTL lflag = C.ECHOCTL
	ECHOPRT lflag = C.ECHOPRT
	ECHOKE  lflag = C.ECHOKE
	// DEFECHO lflag = C.DEFECHO
	FLUSHO lflag = C.FLUSHO
	NOFLSH lflag = C.NOFLSH
	TOSTOP lflag = C.TOSTOP
	PENDIN lflag = C.PENDIN
	IEXTEN lflag = C.IEXTEN
)

type cc C.uint

const (
	VDISCARD cc = C.VDISCARD
	// VDSUSP cc = C.VDSUSP
	VEOF     cc = C.VEOF
	VEOL     cc = C.VEOL
	VEOL2    cc = C.VEOL2
	VERASE   cc = C.VERASE
	VINTR    cc = C.VINTR
	VKILL    cc = C.VKILL
	VLNEXT   cc = C.VLNEXT
	VMIN     cc = C.VMIN
	VQUIT    cc = C.VQUIT
	VREPRINT cc = C.VREPRINT
	VSTART   cc = C.VSTART
	// VSTATUS cc = C.VSTATUS
	VSTOP cc = C.VSTOP
	VSUSP cc = C.VSUSP
	// VSWTCH cc = C.VSWTCH
	VTIME   cc = C.VTIME
	VWERASE cc = C.VWERASE
)
