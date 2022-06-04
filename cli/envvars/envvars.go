package envvars

import (
	"os"
	"strconv"
)

// common shell environment variables
var (
	HOME   string = Getenv("$HOME")
	PWD    string = Getenv("$PWD")
	SHELL  string = Getenv("$SHELL")
	GOPATH string = Getenv("$GOPATH")
)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Getenv returns the value of the string while it
// replaces ${var} or $var in the string according
// to the values of the current environment variables.
// References to undefined variables are replaced by
// a the empty string.
func Getenv(v string) string { return os.ExpandEnv(v) }

func EnvSet(key, value string) error { return os.Setenv(key, value) }
