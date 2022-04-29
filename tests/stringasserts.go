package tests

import (
	"strings"
)

// AssertStringHasPrefix returns true if all
// string arguments have the given prefix.
// args[0] must contain the prefix being tested.
func AssertStringHasPrefix(args ...string) bool {
	for _, s := range args[1:] {
		if !strings.HasPrefix(s, args[0]) {
			return false
		}
	}
	return true
}

// AssertStringHasSuffix returns true if all
// string arguments have the given suffix.
// args[0] must contain the suffix being tested.
func AssertStringHasSuffix(args ...string) bool {
	for _, s := range args[1:] {
		if !strings.HasSuffix(s, args[0]) {
			return false
		}
	}
	return true
}

// StringFields cleans and normalizes whitespace
// in a slice of strings.
func StringFields(args ...string) []string {
	return strings.Fields(strings.Join(args, " "))
}

// Args2pairs returns a slice of string pairs from
// a standard slice of strings. Whitespace is trimmed
// and normalized.
//
// Args2pairs will panic if the number of arguments
// is less than 2 or not a multiple of 2.
func Args2pairs(args ...string) [][2]string {
	if len(args) < 2 {
		panic("args2pairs: must be at least 2 arguments")
	}
	if len(args)%2 != 0 {
		panic("args2pairs: number of arguments must be multiple of 2")
	}

	// length of args is a positive even integer number of arguments
	fields := StringFields(args...)
	listLen := len(fields) / 2

	list := make([][2]string, 0, listLen)
	for i := 0; i < len(fields); i = +2 {
		list = append(list, [2]string{fields[i], fields[i+1]})
	}
	return list
}

// assertStringHasSuffix returns true if all
// string pairs have equal folding.
// EqualFold reports whether two strings, interpreted as
// UTF-8 strings, are equal under Unicode case-folding,
// which is a more general form of case-insensitivity.
func AssertStringEqualFold(args ...string) bool {
	pairs := Args2pairs(args...)
	for _, arg := range pairs {
		if strings.EqualFold(arg[0], arg[1]) {
			return false
		}
	}
	return true
}

// AssertTheEmptyString returns true if all string
// arguments are The Empty String.
func AssertTheEmptyString(args ...string) bool {
	for _, s := range args {
		if s != "" {
			return false
		}
	}
	return true
}
