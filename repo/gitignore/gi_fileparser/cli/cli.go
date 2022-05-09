package cli

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/skeptycal/gosimple/cli/envvars"
	"github.com/skeptycal/gosimple/cli/errorlogger"
	"github.com/skeptycal/gosimple/types/convert"
)

const (
	DefaultInFile   = "./gilist.txt"
	DefaultOutFile  = "../gitignore_gen.go"
	defaultLogLevel = errorlogger.ErrorLevel
)

// conversion utilities
var (
	B2S = convert.UnsafeBytesToString
	S2B = convert.UnsafeStringToBytes
)

// config defaults
var (
	Log                  = errorlogger.New()
	defaultVerboseWriter = os.Stdout
	defaultDebugWriter   = os.Stderr
	discard              = io.Discard
	HOME                 = envvars.HOME
	PWD                  = envvars.PWD
)

// CLI flags and options
type cliOptions struct {
	DebugFlag   bool
	ForceFlag   bool
	VerboseFlag bool
	QuietFlag   bool
	FieldsFlag  bool
	LinesFlag   bool
	InFile      string `default:"DefaultInFile"`
	OutFile     string `default:"DefaultOutFile"`

	verboseWriter io.Writer
	debugWriter   io.Writer

	logLevel errorlogger.Level `default:"1"`

	additionalFlags []Option
}

type Option struct {
	Name         string
	DefaultValue any
	Value        any
	// typ          reflect.Type
	// kind         reflect.Kind
}

var Options cliOptions = cliOptions{}

func init() {
	flag.BoolVar(&Options.DebugFlag, "debug", false, "turn on debug mode")
	flag.BoolVar(&Options.ForceFlag, "force", false, "force writing to file")
	flag.BoolVar(&Options.VerboseFlag, "verbose", false, "turn on verbose mode")
	flag.BoolVar(&Options.QuietFlag, "quiet", false, "turn on quiet mode")

	flag.StringVar(&Options.InFile, "In", DefaultInFile, "name of input file")
	flag.StringVar(&Options.OutFile, "Out", DefaultOutFile, "name of output file")

	flag.Parse()

	Options.logLevel = defaultLogLevel

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

// Vprint sends output based on VerboseFlag setting
// and Log level >= 4.
// If the first argument is a string and contains
// at least one % symbol, it is used as a format
// string for a Printf version of this function.
func Vprint(args ...any) (int, error) {
	if v, ok := args[0].(string); ok {
		if strings.Count(v, "%") > 0 {
			return Vprintf(v, args[1:])
		}
	}
	// Log.Info(args...)
	if Options.VerboseFlag {
		return fmt.Fprint(Options.verboseWriter, args...)
	}
	return 0, nil
}

// Vprintf sends output based on VerboseFlag setting
// and Log level >= 4.
// The first argument is a format string for a Printf
// version of the Vprint function.
func Vprintf(format string, args ...any) (int, error) {
	// Log.Infof(format, args...)
	if Options.VerboseFlag {
		return fmt.Fprintf(Options.verboseWriter, format, args...)
	}
	return 0, nil
}

// DbEcho sends output based on DebugFlag setting
// and Log level >= 2.
// If the first argument is a string and contains
// at least one % symbol, it is used as a format
// string for a Printf version of this function.
func DbEcho(args ...any) (int, error) {
	if v, ok := args[0].(string); ok && len(args) > 1 {
		if strings.Count(v, "%") > 0 {
			return DbEchof(v, args[1:])
		}
	}
	// Log.Debug(args...)
	if Options.DebugFlag {
		return fmt.Fprint(Options.debugWriter, args...)
	}
	return 0, nil
}

// DbEchof sends output based on DebugFlag setting
// and Log level >= 2.
// The first argument is a format string for a Printf
// version of the DbEcho function.
func DbEchof(format string, args ...any) (int, error) {
	// Log.Debugf(format, args...)
	if Options.DebugFlag {
		return fmt.Fprintf(Options.debugWriter, format, args...)
	}
	return 0, nil
}

// StatCli returns the os.FileInfo from filename.
// In the Cli version, any error results in log.Fatal().
func StatCli(filename string) os.FileInfo {
	fiIn, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fiIn
}

// GetDataCli gets the bytes from filename and returns
// the string version.
// In the Cli version, any error results in log.Fatal().
func GetDataCli(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return B2S(data)
}
