package cli

import (
	"fmt"
	"strings"
)

// DbEcho sends output based on DebugFlag setting
// and Log level >= 2.
// If the first argument is a string and contains
// at least one % symbol, it is used as a format
// string for a Printf version of this function.
func DbEcho(args ...any) (int, error) {
	if *DebugFlag {
		// is the first argument a format string ... best guess
		if v, ok := args[0].(string); ok && len(args) > 1 {
			if strings.Count(v, "%") > 0 {
				return DbEchof(v, args[1:])
			}
		}
		// Log.Debug(args...)

		return fmt.Fprint(debugWriter, args...)
	}
	return 0, nil
}

// DbEchof sends output based on DebugFlag setting
// and Log level >= 2.
// The first argument is a format string for a Printf
// version of the DbEcho function.
func DbEchof(format string, args ...any) (int, error) {
	// Log.Debugf(format, args...)
	if *DebugFlag {
		args = append(args, NewLine)
		return fmt.Fprintf(debugWriter, format, args...)
	}
	return 0, nil
}
