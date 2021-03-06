package file

// a copy of the constants from the gofile repo

import (
	"bytes"
	"fmt"
	"os"

	"github.com/skeptycal/gosimple/cli"
	"github.com/skeptycal/gosimple/reallyunsafe"
)

const (
	NormalMode        os.FileMode = 0644
	DirMode           os.FileMode = 0755
	MinBufferSize                 = 16
	SmallBufferSize               = 64
	Chunk                         = 512
	DefaultBufferSize             = 1024
	DefaultBufSize                = 4096
	MaxInt                        = int(^uint(0) >> 1)
	MinRead                       = bytes.MinRead
)

const (
	PathSep      = os.PathSeparator
	ListSep      = os.PathListSeparator
	NewLine      = "\n"
	Tab          = "\t"
	NL      byte = '\n'
	TAB     byte = '\t'
	NUL     byte = 0
)

var (
	V = cli.Vprintln
	// P   = fmt.Println

	b2s = reallyunsafe.B2S
	s2b = reallyunsafe.S2B
)

type TimeZone int

// American Time Zones
const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

var tzNames = map[TimeZone]string{
	1: "EST",
	2: "CST",
	3: "MST",
	4: "PST",
}

func (tz TimeZone) String() string {
	if s, ok := tzNames[tz]; ok {
		return s
	}
	return fmt.Sprintf("GMT%+dh", tz)
}
