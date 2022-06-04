package cli

import (
	"fmt"

	"github.com/pkg/errors"
)

// DbEcho formats using the default formats for its
// operands and writes to debugWriter. Spaces are always
// added between operands and a newline is appended.
// This is equivalent to fmt.Println behavior, not
// fmt.Print behavior.
//
// It returns the number of bytes written and any
// write error encountered.
//
// DbEcho sends output based on DebugFlag setting
// and Log level >= 2.
// If the first argument is a string and contains
// at least one % symbol, it is used as a format
// string for a DbEchof() version of this function.
func DbEcho(args ...any) (n int, err error) {

	if skip, n, err := dbEchoCheck(args...); skip {
		return n, err
	}

	// is the first arg is format string - best guess =)
	if v, ok := args[0].(string); ok {
		return DbEchof(v, args[1:]...)
	}

	// Log.Debugln(args...)
	return fmt.Fprintln(debugWriter, args...)

}

// DbEchof formats according to a format specifier and
// writes to debugWriter. It returns the number of bytes
// written and any write error encountered.
//
// DbEcho sends output based on DebugFlag setting
// and Log level >= 2.
//
// As a convenience, if the final character in the
// format string is not a Newline, then one is added.
func DbEchof(format string, args ...any) (n int, err error) {
	if skip, n, err := dbEchoCheck(args...); skip {
		return n, err
	}

	// Log.Debugf(format, args...)
	n, err = fmt.Fprintf(debugWriter, format, args[1:])

	if format[len(format)-1] != '\n' {
		fmt.Fprint(debugWriter, NewLine)
	}
	return
}

// DbEchoNoLn formats using the default formats for its
// operands and writes to debugWriter. Spaces are added
// between operands when neither is a string. It returns
// the number of bytes written and any write error encountered.
//
// DbEchoNoLn sends output based on DebugFlag setting
// and Log level >= 2.
//
// Behavior is opposite of Print vs. Println in that
// the default DbEcho behavior is to add a Newline and
// the default DbEchoNoLn behavior is to leave all
// arguments unchanged.
// This function writes only the arguments; no
// formatting or updates are done; no newlines are added.
func DbEchoNoLn(args ...any) (n int, err error) {
	if skip, n, err := dbEchoCheck(args...); skip {
		return n, err
	}

	// Log.Debugf(format, args...)
	return fmt.Fprint(debugWriter, args...)
}

// dbEchoCheck handles situations where debug output
// is not enabled, no args are provided, or only one
// argument is provided. In such cases, the skip flag
// is set to true and the caller can safely return
// n and err.
// The skip flag will be false if debug output is enabled
// and two or more arguments are provided.
func dbEchoCheck(args ...any) (skip bool, n int, err error) {
	if !DEBUG || !*DebugFlag {
		return true, 0, nil
	}

	if len(args) == 0 {
		return true, 0, errNoArgumentsProvided
	}

	if len(args) == 1 {
		n, err = fmt.Fprintln(debugWriter, args[0])
		return true, n, err
	}

	return
}

var errNoArgumentsProvided = errors.New("DbEcho(): no arguments provided")
