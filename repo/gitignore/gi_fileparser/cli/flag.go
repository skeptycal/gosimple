package cli

import (
	"flag"
	"io"
	"os"

	"github.com/skeptycal/gosimple/cli/envvars"
	"github.com/skeptycal/gosimple/cli/errorlogger"
)

// config defaults
var (
	DefaultDebugFlag     = false
	DefaultInFile        = "./gilist.txt"
	DefaultOutFile       = "../gitignore_gen.go"
	DefaultLogLevel      = errorlogger.ErrorLevel
	Log                  = errorlogger.New()
	defaultVerboseWriter = os.Stdout
	defaultDebugWriter   = os.Stderr
	discard              = io.Discard
	HOME                 = envvars.HOME
	PWD                  = envvars.PWD
)

// CLI flags and options
type cliOptions struct {
	DebugFlag   bool `default:"DefaultDebugFlag"`
	ForceFlag   bool
	VerboseFlag bool
	QuietFlag   bool
	InFile      string `default:"DefaultInFile"`
	OutFile     string `default:"DefaultOutFile"`

	verboseWriter io.Writer `default:"defaultVerboseWriter"`
	debugWriter   io.Writer `default:"defaultDebugWriter"`

	logLevel errorlogger.Level `default:"DefaultLogLevel"`

	additionalFlags []Option
}

var Options cliOptions = cliOptions{}

var Flag = flag.NewFlagSet("cli", flag.ExitOnError)

// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
//
// If a Value has an IsBoolFlag() bool method returning true,
// the command-line parser makes -name equivalent to -name=true
// rather than using the next command-line argument.
//
// Set is called once, in command line order, for each flag present.
// The flag package may call the String method with a zero-valued receiver,
// such as a nil pointer.
type Value interface {
	String() string
	Set(string) error
}

type Option struct {
	Flag  flag.Flag
	Short string
}

func init() {
	Flag.BoolVar(&Options.DebugFlag, "debug", false, "turn on debug mode")
	Flag.BoolVar(&Options.ForceFlag, "force", false, "force writing to file")
	Flag.BoolVar(&Options.VerboseFlag, "verbose", false, "turn on verbose mode")
	Flag.BoolVar(&Options.QuietFlag, "quiet", false, "turn on quiet mode")

	Flag.StringVar(&Options.InFile, "In", DefaultInFile, "name of input file")
	Flag.StringVar(&Options.OutFile, "Out", DefaultOutFile, "name of output file")

	Flag.Parse(os.Args[1:])

	Options.logLevel = DefaultLogLevel

	if Options.DebugFlag {
		Options.logLevel = errorlogger.DebugLevel
		Options.debugWriter = defaultDebugWriter
	}

	if Options.VerboseFlag {
		Options.logLevel = errorlogger.InfoLevel
		Options.verboseWriter = defaultVerboseWriter
	}

	// no output even if other options are set
	if Options.QuietFlag {
		Options.logLevel = errorlogger.FatalLevel
		Options.verboseWriter = discard
		Options.debugWriter = discard
	}

	Log.SetLevel(Options.logLevel)
}
