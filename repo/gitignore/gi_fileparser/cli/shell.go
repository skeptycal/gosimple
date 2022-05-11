package cli

import (
	"fmt"
	"os"
)

// Getenv returns the value of the string while
// replaces ${var} or $var in the string according
// to the values of the current environment variables.
// References to undefined variables are replaced by
// defaultValue.
//  d := Getenv("${HOME}/.config")
//  fmt.Println(d)
//  // /Users/skeptycal/.config
func Getenv(envVarName string, defaultValue string) (retval string) {
	retval = os.ExpandEnv(envVarName)
	if retval == "" {
		osErr(os.ErrInvalid, fmt.Errorf("Getenv(%q) error: %q (using default value: %q", envVarName, retval, defaultValue).Error())
		return defaultValue
	}
	return
}
