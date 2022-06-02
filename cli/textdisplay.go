package cli

import (
	"fmt"
	"strconv"
	"strings"
)

// Atoi returns the integer representation of s. If any error
// occurs, 0 is returned.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if textErr(err, "Atoi()") != nil {
		return 0
	}
	return i
}

// Columns returns the current number of columns in the terminal.
// If any error occurs, the default value is returned.
func Columns() int { return Atoi(Getenv("$COLUMNS", DefaultScreenWidthString)) }

// Br prints a blank line to os.Stdout.
func Br() { fmt.Print(NewLine) }

// Hr prints a screen-wide underline, or hard return.
func Hr() { fmt.Print(strings.Repeat("_", Columns())) }

// Cr prints a screen-wide character pattern.
func Cr(c string) { fmt.Print(strings.Repeat(c, Columns()/len(c))) }
