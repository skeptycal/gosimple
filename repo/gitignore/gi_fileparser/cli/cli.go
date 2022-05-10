// Package cli contains command line interface components that address common
// use cases in cli development and design.
//
// It has very few (standard library only) dependencies and is a simple drop-in
// addition to any cli toolkit.
package cli

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	defaultHeadByteLength    = 79
	defaultTailByteLength    = 20
	defaultHeadLineLength    = 5
	defaultTailLineLength    = 5
	defaultScreenWidthString = "80"
	Newline                  = "\n"
)

// fast conversion utilities
var (
	B2S = unsafeBytesToString
	S2B = unsafeStringToBytes
)

// StatCli returns the os.FileInfo from filename.
// In the Cli version, any error results in log.Fatal().
func StatCli(filename string) os.FileInfo {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fi
}

// GetDataCli gets the contents of filename and returns
// the string version.
// In the Cli version, any error results in log.Fatal().
func ReadFileCli(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return B2S(data)
}

// Atoi returns the integer representation of s. If any error
// occurs, 0 is returned.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// Columns returns the current number of columns in the terminal.
// If any error occurs, the default value is returned.
func Columns() int { return Atoi(Getenv("$COLUMNS", defaultScreenWidthString)) }

// Br prints a blank line to os.Stdout.
func Br() { fmt.Print(Newline) }

// Hr prints a screen-wide underline, or hard return.
func Hr() { fmt.Print(strings.Repeat("_", Columns())) }

// Cr prints a screen-wide character pattern.
func Cr(c string) { fmt.Print(strings.Repeat(c, Columns()/len(c))) }

// Head returns the first n elements of a sequence s.
// If n is longer than s, s is returned unchanged.
func HeadN[E any, S ~[]E](s S, n int) S {
	if n > len(s) {
		return s
	}
	return s[:n]
}

// Head returns the first n elements of a sequence.
// If n is longer than s, s is returned unchanged.
// The default value of n is used. If another value
// of n is needed, use HeadN(s S, n int).
func Head[E any, S ~[]E](s S) S {
	return HeadN(s, defaultHeadLineLength)
}

// Tail returns the last n elements of a sequence s.
// If n is longer than s, s is returned unchanged.
func TailN[E any, S ~[]E](s S, n int) S {
	if n > len(s) {
		return s
	}
	return s[len(s)-n:]
}

// Tail returns the last n elements of a sequence.
// If n is longer than s, s is returned unchanged.
// The default value of n is used. If another value
// of n is needed, use TailN(s S, n int).
func Tail[E any, S ~[]E](s S) S {
	return TailN(s, defaultTailLineLength)
}

// Getenv returns the value of the string while
// replaces ${var} or $var in the string according
// to the values of the current environment variables.
// References to undefined variables are replaced by
// defaultValue.
//  d := Getenv("${HOME}/.config")
//  fmt.Println(d)
//  // /Users/skeptycal/.config
func Getenv(envVarName string, defaultValue string) (retval string) {
	retval = os.ExpandEnv(envVarName)
	if retval == "" {
		return defaultValue
	}
	return
}
