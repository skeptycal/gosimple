package envvars

import (
	"github.com/skeptycal/gosimple/cli/goshell"
)

// common shell environment variables
var (
	HOME   string = Getenv("$HOME")
	PWD    string = Getenv("$PWD")
	SHELL  string = Getenv("$SHELL")
	GOPATH string = Getenv("$GOPATH")
)

// Getenv returns the value of the string while it
// replaces ${var} or $var in the string according
// to the values of the current environment variables.
// References to undefined variables are replaced by
// a the empty string.
func Getenv(v string) string { return goshell.Getenv(v, "") }

func EnvSet(key, value string) error { return goshell.Setenv(key, value) }
