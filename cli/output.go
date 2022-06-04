package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skeptycal/gosimple/cli/terminal"
)

var t = terminal.GetWinSize()
var Col = t.Col

// Atoi returns the integer representation of s. If any error
// occurs, 0 is returned.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if textErr(err, "Atoi()") != nil {
		return 0
	}
	return i
}

// Br prints a blank line to os.Stdout.
func Br() { fmt.Print(NewLine) }

// Hr prints a screen-wide header underline using
// the headerChar string contained in the package
// options.
// The screen width is updated by the Columns() function.
func Hr() { fmt.Println(headerString()) }

// Fr prints a screen-wide footer underline using
// the footerChar string contained in the package
// options.
// The screen width is updated by the Columns() function.
func Fr() { fmt.Println(footerString()) }

// Cr prints a screen-wide character pattern.
func Cr(c string) { fmt.Print(strings.Repeat(c, Col()/len(c))) }

func Box(args ...any) (n int, err error) {
	Hr()
	n, err = fmt.Print(args...)
	Fr()
	return
}

func headerString() string {
	return strings.Repeat(headerChar, Col()/len(headerChar))
}

func footerString() string {
	return strings.Repeat(footerChar, Col()/len(footerChar))
}
