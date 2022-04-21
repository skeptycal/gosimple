package cli

import "github.com/skeptycal/gosimple/ansi"

var globalReturn = ""

const (
	fmtANSI         string = ansi.FmtANSI // format string for simple ANSI encoding ( "\x1b[%dm" )
	fa                     = "\x1b[%dm"
	ansiPrefix      string = "\033["
	esc             byte   = '\x1b'
	ansiPrefixByte1 byte   = esc
	ansiPrefixByte2 byte   = '['
	ansiSuffix      string = "m"
	ansiSep         string = ";"
	ansiSepByte     byte   = ';'
	ansiSuffixByte  byte   = 'm'
)
