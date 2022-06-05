// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package ansi

import (
	"bytes"
	"fmt"
	"strings"
)

// NewColor returns a new ansi color string function.
//
// Several methods were profiled and the most efficient
// was aliased to NewColor. YMMV; change the alias as
// needed for your environment.
var NewColor = newColorConcat

const (
	fmtANSI            string = FmtANSI // format string for simple ANSI encoding ( "\x1b[%dm" )
	fa                 string = "\x1b[%dm"
	ansiEncodeBasicFMT string = "\x1b[%v;%v;%vm"
	ansiPrefix         string = "\033["
	ansiSuffix         string = "m"
	ansiSep            string = ";"
	fg                 string = "38;5;"
	bg                 string = "48;5;"
	fg2                string = ansiSep + fg
	bg2                string = ansiSep + bg
	esc                byte   = '\x1b'
	ansiPrefixByte1    byte   = esc
	ansiPrefixByte2    byte   = '['
	ansiSepByte        byte   = ';'
	ansiSuffixByte     byte   = 'm'

	Reset          string = "\033[0m"  // ANSI reset code
	ResetColor     string = "\033[32m" // Reset to default color
	ResetLineConst string = "\r\033[K" // Return cursor to start of line and clean it
	SetBold        string = "\033[1m"  // ANSI bold
	SetInverse     string = "\033[4m"  // ANSI inverse
)

var (
	ansiPrefixByte []byte = []byte(ansiPrefix)
	bAnsiPrefix    []byte = []byte(ansiPrefix)
	ResetBytes     []byte = []byte(Reset)
	BoldBytes      []byte = []byte(SetBold)
	InverseBytes   []byte = []byte(SetInverse)
)

var (
	defaultAnsiCode = NewColor("2", "0", "1")
)

func ansiEncodeBasic(fg, bg, ef string) string {
	return fmt.Sprintf(ansiEncodeBasicFMT, fg, bg, ef)
}

func newColorConcat(foreground, background, effect string) string {
	//     "\x1b[        %d      ;	  38;5;       %d      ;       48;5;     %d         m"
	return ansiPrefix + effect + ansiSep + fg + foreground + ansiSep + bg + background + ansiSuffix
}

func newColorConcat2(foreground, background, effect string) string {
	//     "\x1b[     %d    ;38;5;     %d      ;48;5;     %d         m"
	return ansiPrefix + effect + fg2 + foreground + bg2 + background + ansiSuffix
}

const ansitemplate = "\x1b[000;38;5;000;48;5;000m"
const z = '0'

var a [24]byte

func NewColorMake(foreground, background, effect string) string {
	//     "\x1b[     000    ;38;5;      000      ;48;5;     000        m"
	//      0,1      2,3,4    5-10     11,12,13   14-19    20,21,22    23

	var b = a[:0]

	b = []byte(ansitemplate)

	fmt.Printf("a: %q\n", a)
	fmt.Printf("b: %q\n", b)

	fz := 3 - len(foreground)
	bz := 3 - len(background)
	ez := 3 - len(effect)

	for i := 0; i < fz; i++ {
		b[2+i] = '0'
	}

	fmt.Println("fz: ", fz)
	fmt.Println("bz: ", bz)
	fmt.Println("ez: ", ez)

	foreground = strings.Repeat("0", fz) + foreground
	background = strings.Repeat("0", bz) + background
	effect = strings.Repeat("0", ez) + effect

	fmt.Println("foreground: ", foreground)
	fmt.Println("background: ", background)
	fmt.Println("effect: ", effect)

	return ""
}

func newColorSB(foreground, background, effect string) string {
	sb := strings.Builder{}
	sb.WriteString(ansiPrefix)
	sb.WriteString(effect)
	sb.WriteString(ansiSep)
	sb.WriteString(fg)
	sb.WriteString(foreground)
	sb.WriteString(ansiSep)
	sb.WriteString(bg)
	sb.WriteString(background)
	sb.WriteString(ansiSuffix)
	return sb.String()
}

func newColorBB(foreground, background, effect string) string {
	sb := bytes.Buffer{}
	sb.WriteString(ansiPrefix)
	sb.WriteString(effect)
	sb.WriteString(ansiSep)
	sb.WriteString(fg)
	sb.WriteString(foreground)
	sb.WriteString(ansiSep)
	sb.WriteString(bg)
	sb.WriteString(background)
	sb.WriteString(ansiSuffix)
	return sb.String()
}

func newColorJoin(foreground, background, effect string) string {
	return strings.Join(
		[]string{ansiPrefix + effect, fg + foreground, bg + background + ansiSuffix},
		ansiSep,
	)
}

func newColorSprintf(foreground, background, effect string) string {
	return fmt.Sprintf(ansiEncodeBasicFMT, effect, foreground, background)
}
