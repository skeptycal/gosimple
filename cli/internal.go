package cli

import (
	"github.com/skeptycal/gosimple/cli/ansi"
	"github.com/skeptycal/gosimple/unsafe"
)

var globalReturn = ""

const (
	fmtANSI         string = ansi.FmtANSI // format string for simple ANSI encoding ( "\x1b[%dm" )
	fa              string = "\x1b[%dm"
	ansiPrefix      string = "\033["
	esc             byte   = '\x1b'
	ansiPrefixByte1 byte   = esc
	ansiPrefixByte2 byte   = '['
	ansiSuffix      string = "m"
	ansiSep         string = ";"
	ansiSepByte     byte   = ';'
	ansiSuffixByte  byte   = 'm'
)

var (
	ansiPrefixByte = []byte(ansiPrefix)
)

var (

	// alias for local logging
	// sometimes my habit for lowercase is too strong ...
	log          = Log
	headerChar   = "*"
	footerChar   = "-"
	headerBorder = headerString()
	footerBorder = footerString()
)

// fast conversion utilities
var (
	B2S = unsafe.B2S
	S2B = unsafe.S2B
)
