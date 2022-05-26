package textfiles

import (
	"unicode"

	"github.com/skeptycal/gosimple/types/convert"
)

const (
	newline = "\n"
	space   = " "
	tab     = "\t"
)

var (
	S2B             = convert.UnsafeStringToBytes
	B2S             = convert.UnsafeBytesToString
	linesep         = S2B(newline)
	utf8Replacement = []byte(string(unicode.ReplacementChar))
)
