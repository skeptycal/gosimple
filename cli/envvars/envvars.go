package envvars

import (
	"github.com/skeptycal/gosimple/cli/goshell"
)

// conversion function shortcut aliases
var (
	home   string = EnvGet("$HOME")
	pwd    string = EnvGet("$PWD")
	shell  string = EnvGet("$SHELL")
	gopath string = EnvGet("$GOPATH")
)

// Getenv returns the value of the string while it
// replaces ${var} or $var in the string according
// to the values of the current environment variables.
// References to undefined variables are replaced by defaultValue.
func EnvGet(v string) string         { return goshell.Getenv(v, "") }
func EnvSet(key, value string) error { return goshell.Setenv(key, value) }
