package miniansi

import (
	"fmt"
	"os"
	"strings"

	"github.com/skeptycal/gosimple/constraints"
)

var DEBUG bool = true // debug or DEV mode

const (
	InfoColor    = ansiPrefix + "1;34m"
	NoticeColor  = ansiPrefix + "1;36m"
	WarningColor = ansiPrefix + "1;33m"
	ErrorColor   = ansiPrefix + "1;31m"
	DebugColor   = ansiPrefix + "0;36m"
	dbcolor      = ansiPrefix + "1;31m" // ANSI dbecho code
	ResetColor   = ansiPrefix + "0m"    // ANSI reset code

	// semicolon delimited ANSI codes for %v
	// string to print for %s
	ansiFmt = "\033[%vm%s\033[0m"

	ansiSEP    = ";"
	ansiPrefix = "\033["
	ansiSuffix = "m"
)

type ansiConstraint interface {
	~string | []byte | constraints.Integer
}

type ansi[T ansiConstraint] struct {
	out T
}

// NewAnsi creates a new ansi color code string
// from components. Each argument is parsed and
// encoded and wrapped in an ANSI
func NewAnsi(in ...any) string {
	// func NewAnsi[T ansiConstraint](in ...any) string {
	// TODO: handle inappropriate types
	s := fmt.Sprint(in...)
	list := strings.Fields(s)
	return ansiPrefix + strings.Join(list, ansiSEP) + ansiSuffix
}

// func (a ansi[T]) String() string {
// 	return fmt.Sprintf(a.out)
// }

func ansiEncode(code any, s ...string) string {
	return fmt.Sprintf(ansiFmt, code, s)
}

// dbecho prints to os.Stdout if the global DEBUG is true
func DbEcho(args ...any) (n int, err error) {
	if DEBUG {
		s := fmt.Sprint(args...)
		return fmt.Fprintf(os.Stdout, "%s%s%s\n", dbcolor, s, ResetColor)
	}
	return
}
