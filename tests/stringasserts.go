package tests

import (
	"strings"
)

// AssertStringHasPrefix returns true if all
// string arguments have the given prefix.
// args[0] must contain the prefix being tested.
func AssertStringHasPrefix(args ...string) bool {
	if !CheckPairs(args...) {
		return false
	}
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
	if !CheckPairs(args...) {
		return false
	}
	for _, s := range args[1:] {
		if !strings.HasSuffix(s, args[0]) {
			return false
		}
	}
	return true
}

// AssertStringEqualFold returns true if all
// string pairs have equal folding.
//
// EqualFold reports whether two strings, interpreted as
// UTF-8 strings, are equal under Unicode case-folding,
// which is a more general form of case-insensitivity.
//
// Only pairs of args are considered; number of args
// must be a positive integer multiple of two.
func AssertStringEqualFold(args ...string) bool {
	if !CheckPairs(args...) {
		return false
	}
	pairs := Args2Pairs(args...)
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
