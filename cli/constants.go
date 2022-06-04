// The constants.go file contains type definitions and constants mainly from the ansi package.
// Reference: github.com/skeptycal/ansi

package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/skeptycal/gosimple/cli/ansi"
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
	DbColor     string = "\033[1;31m" // ANSI dbecho code
	bAnsiPrefix []byte = []byte(ansiPrefix)
	SetBold     string = BasicEncode(fmt.Sprint(ansi.Bold)) // ANSI bold
)

const (
	Reset                    string = "\033[0m"  // ANSI reset code
	ResetColor               string = "\033[32m" // Reset to default color
	ResetLineConst           string = "\r\033[K" // Return cursor to start of line and clean it
	SetInverse               string = "\033[4m"  // ANSI inverse
	NewLine                  string = "\n"       // Newline character
	Tab                      string = "\t"       // Tab character
	Space                           = " "        // Space character
	DefaultScreenWidthString        = "80"
	DefaultHeadByteLength           = 79
	DefaultTailByteLength           = 20
	DefaultHeadLineLength           = 5
	DefaultTailLineLength           = 5
)

// List of possible colors (3 bit ANSI)
const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

// Ansi Control Codes provide ASCII representation of nonprintable characters.
const (
	NUL byte = iota //  0 = Null
	SOH             //	1 = Start of Heading
	STX             //	2 = Start of Text
	ETX             //	3 = End of Text
	EOT             //	4 = End of Transmission
	ENQ             //	5 = Enquiry
	ACK             //	6 = Acknowledge
	BEL             //	7 = Bell
	BS              //	8 = Backspace
	TAB             //	9 = Horizontal Tab
	LF              //	10 = Line Feed
	VT              //	11 = Vertical Tab
	FF              //	12 = Form Feed
	CR              //	13 = Carriage Return
	SO              //	14 = Shift Out
	SI              //	15 = Shift In
	DLE             //	16 = Data Link Escape
	DC1             //	17 = Device Control 1
	DC2             //	18 = Device Control 2
	DC3             //	19 = Device Control 3
	DC4             //	20 = Device Control 4
	NAK             //	21 = Negative Acknowledgement
	SYN             //	22 = Synchronous Idle
	ETB             //	23 = End of Transmission Block
	CAN             //	24 = Cancel
	EM              //	25 = End of Medium
	SUB             //	26 = Substitute
	ESC             //	27 = Escape
	FS              //	28 = File Separator
	GS              //	29 = Group Separator
	RS              //	30 = Record Separator
	US              //	31 = Unit Separator
)

var (
	ResetBytes   []byte = []byte(Reset)
	InverseBytes []byte = []byte(BasicEncode(string(ansi.Inverse)))
)
