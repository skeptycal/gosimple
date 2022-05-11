package cli

import (
	"flag"
	"io"
	"os"

	"github.com/skeptycal/gosimple/cli/errorlogger"
)

// Flags is a flag.FlagSet that can be inherited or modified.
var Flags = flag.NewFlagSet("cli", flag.ExitOnError)

////////// Flags utilities

func isSet(name string) bool        { return Flags.Lookup(name) != nil }
func flagString(name string) string { return Flags.Lookup(name).Value.String() }

func Flag(flag *flag.Flag) any {
	if v, ok := flag.Value.(Getter); ok {
		return v.Get()
	}
	return nil
}

////////// Flags defaults
var (
	DefaultDebugFlag     = false
	DefaultInFile        = "./gilist.txt"
	DefaultOutFile       = "../gitignore_gen.go"
	DefaultConfigFile    = ".gosimple.conf"
	defaultLogLevel      = errorlogger.ErrorLevel
	defaultVerboseWriter = os.Stdout
	defaultDebugWriter   = os.Stderr
)

////////// cli Flags
var (
	DebugFlag   = Flags.Bool("debug", false, "turn on debug mode")
	ForceFlag   = Flags.Bool("force", false, "force writing to file")
	VerboseFlag = Flags.Bool("verbose", false, "turn on verbose mode")
	QuietFlag   = Flags.Bool("quiet", false, "turn on quiet mode")
	InFile      = Flags.String("in", DefaultInFile, "name of input file")
	OutFile     = Flags.String("out", DefaultOutFile, "name of output file")
	ConfigFile  = Flags.String("config", DefaultConfigFile, "name of config file")
	LogLevel    = &logLevelFlag{defaultLogLevel}

	debugWriter   io.Writer = discard
	verboseWriter io.Writer = discard
)

type logLevelFlag struct {
	level errorlogger.Level
}

func init() {
	Flags.Var(LogLevel, "loglevel", "level of logging feedback")
	Flags.Parse(os.Args[1:])

	if *VerboseFlag {
		LogLevel.Set("Info")
		verboseWriter = defaultVerboseWriter
	}

	// Debug level output ... overrides Verbose logging level.
	if *DebugFlag {
		LogLevel.Set("Debug")
		debugWriter = defaultDebugWriter
	}

	// no output even if other options are set
	// overrides Verbose and Debug logging level.
	if flagString("QuietFlag") == "true" {
		LogLevel.Set("Fatal")
		debugWriter = discard
		verboseWriter = discard
	}
}

////////// from standard library flag package

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
//
// Reference: standard library flag package
type Value interface {
	String() string
	Set(string) error
}

type Getter interface {
	Value
	Get() any
}

type hasIsBoolFlag interface {
	IsBoolFlag()
}
