// Package cli contains command line interface components that address common
// use cases in cli development and design.
//
// It has very few (standard library only) dependencies and is a simple drop-in
// addition to any cli toolkit.
package cli

import (
	"io"

	"github.com/skeptycal/gosimple/cli/envvars"
	"github.com/skeptycal/gosimple/cli/errorlogger"
)

const (
	DefaultHeadByteLength    = 79
	DefaultTailByteLength    = 20
	DefaultHeadLineLength    = 5
	DefaultTailLineLength    = 5
	DefaultScreenWidthString = "80"
	Newline                  = "\n"
	Tab                      = "\t"
	Space                    = " "
)

var (
	Log     = errorlogger.New()
	log     = Log // sometimes my habit for lowercase is too strong ...
	er      = Log.Err
	discard = io.Discard
	HOME    = envvars.HOME
	PWD     = envvars.PWD
)

// fast conversion utilities
var (
	B2S = unsafeBytesToString
	S2B = unsafeStringToBytes
)
