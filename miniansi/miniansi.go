package miniansi

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/skeptycal/gosimple/printer"
)

var DEBUG bool = true // debug or DEV mode

func (c *control[T]) SetDebug(debug bool) {
	c.debug = debug
}

func (c *control[T]) SetEnabled(enabled bool) {
	c.enabled = enabled
}

type Ansier interface {
	String() string
	printer.Fprinter

	// test methods ...
	Fprintf1(w io.Writer, format string, args ...interface{}) (n int, err error)
	Fprintf2(w io.Writer, format string, args ...interface{}) (n int, err error)
	Fprintf3(w io.Writer, format string, args ...interface{}) (n int, err error)
}

type ansi[T AnsiConstraint] struct {
	value T
	out   string
	bOut  []byte
}

func (c *ansi[T]) String() string {
	return fmt.Sprint(c.out)
}

const fprintFormat = "%v%v%v"

func (c *ansi[T]) Fprint(w io.Writer, format string, args ...interface{}) (int, error) {
	return c.Fprintf3(w, format, args...)
}
func (c *ansi[T]) Fprintln(w io.Writer, format string, args ...interface{}) (int, error) {
	return c.Fprintf3(w, format, args...)
}

func (c *ansi[T]) Fprintf(w io.Writer, format string, args ...interface{}) (int, error) {
	return c.Fprintf3(w, format, args...)
}

func (c *ansi[T]) Fprintf1(w io.Writer, format string, args ...interface{}) (n int, err error) {
	if w == nil {
		w = os.Stdout
	}

	s := fmt.Sprintf(format, args...)
	return fmt.Fprintf(w, fprintFormat, c.out, s, ResetColor)
}

func (c *ansi[T]) Fprintf2(w io.Writer, format string, args ...interface{}) (n int, err error) {
	format = "%v" + format
	a := make([]interface{}, 1, len(args)+1)
	a[0] = c.out
	a = append(a, args...)
	return fmt.Fprintf(w, format, a...)
}

func (c *ansi[T]) Fprintf3(w io.Writer, format string, args ...interface{}) (n int, err error) {
	w.Write(c.bOut)
	n, err = fmt.Fprintf(w, format, args...)
	w.Write(bReset)
	return
}

func NewAnsi[T AnsiConstraint](in T) *ansi[T] {
	s := fmt.Sprint(in)
	return &ansi[T]{
		value: in,
		out:   s,
		bOut:  []byte(s),
	}
}

// NewAnsi creates a new ansi color code string
// from components. Each argument is parsed and
// encoded and wrapped in an ANSI
func NewAnsiString(in ...any) string {
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
		return fmt.Fprintf(os.Stdout, "%s%s%s\n", DbColor, s, ResetColor)
	}
	return
}
