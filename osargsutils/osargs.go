package osargs

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/skeptycal/goutil/repo/errorlogger"
)

var log = errorlogger.New()

// global OsArgs shortcuts ready to use.
var (
	OsArgs arger = osArgs{}.init()
	App          = OsArgs.App()
	Args         = OsArgs.Args()
	Me           = OsArgs.Base()
	Here         = OsArgs.Dir()
)

type arger interface {
	App() string
	Args() []string
	ArgString() string
	Dir() string
	Base() string
}

// osArgs hold the command-line arguments:
//
// app: cached program name
// args: cached command line arguments
// base: cached base file name of program
// dir: cached directory name of program
type osArgs struct {
	app  string
	args []string
	base string
	dir  string
}

func (o osArgs) init() osArgs {
	_ = o.App()
	_ = o.Args()
	_ = o.Base()
	_ = o.Dir()
	return o
}

// App returns the absolute path name for the executable
// that started the current process. EvalSymlinks is run
// on the resulting path to provide a stable result.
func (o osArgs) App() string {
	if o.app == "" {
		a, err := arg0()
		if err != nil {
			log.Errorf("could not determine command line app (os.args[0]): %v", err)
		}
		o.app = a
	}
	return o.app
}

// Args returns the command line arguments as
// a slice of strings.
func (o osArgs) Args() []string {
	if len(o.args) == 0 {
		o.args = args()
	}
	return o.args
}

// ArgString returns the command line arguments as
// a single space delimited string.
func (o osArgs) ArgString() string {
	return strings.Join(o.Args(), " ")
}

func (o osArgs) Base() string {
	if o.base == "" {
		o.base = filepath.Base(o.App())
	}
	return o.base
}

func (o osArgs) Dir() string {
	if o.dir == "" {
		o.dir = filepath.Dir(o.App())
	}
	return o.dir
}

// HereMe returns the folder (here) and basename (me) of
// the executable that started the current process.
func HereMe() (string, string) { return OsArgs.Dir(), OsArgs.Base() }

// Arg0 returns the absolute path name for the executable
// that started the current process. EvalSymlinks is run
// on the resulting path to provide a stable result.
//
// As of Go 1.8 (Released February 2017) the recommended
// way of doing this is with os.Executable. Prior to Go 1.8,
// and as a backup, os.Args[0] is used.
func arg0() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return filepath.EvalSymlinks(os.Args[0])
	}
	return filepath.EvalSymlinks(ex)
}

func args() []string {
	switch len(os.Args) {
	case 0:
		log.Errorf("os.Args was zero length.")
		return []string{}
	case 1:
		return []string{""}
	default:
		return os.Args[1:]
	}
}
