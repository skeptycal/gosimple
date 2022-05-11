package cli

import (
	"fmt"
	"strings"
)

/// tools for Verbose flag and mode

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
	if *VerboseFlag {
		return fmt.Fprint(verboseWriter, args...)
	}
	return 0, nil
}

// Vprintln sends output based on VerboseFlag setting
// and Log level >= 4.
// A trailing newline is appended to the output.
//
// If the first argument is a string and contains
// at least one % symbol, it is used as a format
// string for a Printf version of this function.
func Vprintln(args ...any) (int, error) {
	args = append(args, Newline)
	return Vprint(args...)
}

// Vprintf sends output based on VerboseFlag setting
// and Log level >= 4.
// The first argument is a format string for a Printf
// version of the Vprint function.
func Vprintf(format string, args ...any) (int, error) {
	// Log.Infof(format, args...)
	if *VerboseFlag {
		return fmt.Fprintf(verboseWriter, format, args...)
	}
	return 0, nil
}
