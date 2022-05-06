package goshell

import (
	"os"
	"strings"
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
		return defaultValue
	}
	return
}

// Setenv sets the value of the environment variable named
// by the key. It returns an error, if any.
func Setenv(key, value string) error {
	return os.Setenv(key, value)
}

// Environ returns a copy of strings representing the
// environment, in the form "key=value".
func Environ() []string {
	return os.Environ()
}

// EnvMap returns a mapping of string keys to string
// values representing the environment, in the form
// "m[key] == value", e.g.
//  home := m["HOME"]
//  fmt.Println(home)
//  // /Users/skeptycal/
func EnvMap() map[string]string {
	list := os.Environ()
	m := make(map[string]string, len(list))
	for _, item := range os.Environ() {
		pair := strings.SplitN(item, "=", 2)
		m[pair[0]] = pair[1]
	}
	return m
}

/// Alternate functions for testing or possible later export

func getenv(key string) string {
	return os.Getenv(key)
}

func lookup(key string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return "!EnvVarError"
}
