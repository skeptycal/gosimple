// The constants.go file contains type definitions and constants mainly from the ansi package.
// Reference: github.com/skeptycal/ansi

package cli

import (
	"io"
	"os"

	"github.com/skeptycal/gosimple/cli/envvars"
	"github.com/skeptycal/gosimple/cli/terminal"
)

var (

	// DEBUG flag to enable debug logging and features
	DEBUG = true

	// Terminal flag to enable CLI terminal display
	IsTerminal = terminal.IsTerminal(int(os.Stdout.Fd()))
)

// environment variables
var (

	// user home directory
	HOME = envvars.HOME

	// current working directory
	PWD = envvars.PWD

	// Initial column width of CLI terminal display
	COLUMNS = 80

	// Initial row height of CLI terminal display
	ROWS = 24
)

var (
	defaultWriter      io.Writer = newAnsiStdout()
	defaultErrorWriter io.Writer = newAnsiStderr()
	// Output             CLI       = New()
)

var (
	DbColor string = "\033[1;31m" // ANSI dbecho code

)

const (
	NewLine                  string = "\n" // Newline character
	Tab                      string = "\t" // Tab character
	Space                           = " "  // Space character
	DefaultScreenWidthString        = "80"
	DefaultHeadByteLength           = 79
	DefaultTailByteLength           = 20
	DefaultHeadLineLength           = 5
	DefaultTailLineLength           = 5
)
