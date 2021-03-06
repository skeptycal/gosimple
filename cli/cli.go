// Package cli contains utility functions for dealing with cli commands
// within CLI applications written in Go.
//
// The main component is the CLI interface that implements cli features
// through the Terminal struct.
// Reference: github.com/skeptycal/cli/terminal
//
// The ansi color terminal support is provided by the ansi package.
// Reference: github.com/skeptycal/ansi
package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/skeptycal/gosimple/cli/ansi"
	"github.com/skeptycal/gosimple/cli/terminal"
)

const defaultDevMode = true

func init() {

	// TODO: causes seg fault in VsCode terminal window, e.g.
	/* time="2022-06-01T12:39:29-05:00" level=error msg="GetWinsize: operation not supported on socket"
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal SIGSEGV: segmentation violation code=0x2 addr=0x2 pc=0x104f87b44]
	*/
	winSize := terminal.GetWinSize()
	if winSize != nil {
		COLUMNS = winSize.Col()
		ROWS = winSize.Row()
	}
}

func checkColor() bool {
	// TODO check if color is supported - bring this code over ...
	return true
}

// newAnsiStdout returns stdout which converts escape sequences
// to Windows API calls on Windows environment.
func newAnsiStdout() io.Writer { return os.Stdout }

// newAnsiStderr returns stdout which converts escape sequences
// to Windows API calls on Windows environment.
func newAnsiStderr() io.Writer { return os.Stderr }

// Printer implements the common printer interface elements
// Print, Printf, and Println
type Printer interface {
	Print(args ...interface{}) (n int, err error)
	Printf(format string, args ...interface{}) (n int, err error)
	Println(args ...interface{}) (n int, err error)
}

// CLI implements an ANSI compatible terminal interface.
type CLI interface {
	io.Writer
	io.StringWriter
	fmt.Stringer
	Printer
	terminal.CLIControls
	SetColor(color ansi.AnsiColor)
	Reset() (n int, err error)
}

// // New returns a new ANSI compatible terminal interface based on
// // os.Stdout with ANSI support enabled by default.
// func New() CLI { return NewFromWriter(defaultWriter) }

// // New returns a new ANSI compatible terminal interface based on
// // os.Stderr with ANSI support enabled by default.
// func NewStderr() CLI { return NewFromWriter(defaultErrorWriter) }

// // New returns a new ANSI compatible terminal interface based on
// // the given io.Writer with ANSI support enabled by default.
// func NewFromWriter(w io.Writer) CLI {
// 	devMode := defaultDevMode
// 	if w == nil {
// 		w = defaultWriter
// 	}

// 	t := &Terminal{
// 		w:        w,
// 		useColor: checkColor(),
// 		devMode:  devMode,
// 	}

// 	if t.useColor {
// 		t.on = t.colorWrite
// 	} else {
// 		t.on = t.noOp
// 	}

// 	return t
// }
