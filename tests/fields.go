package tests

import (
	"fmt"
	"strings"
)

// CheckPairs returns true if the given args
// slice contains an even number of elements.
func CheckPairs[T any](args ...T) bool {
	return len(args) > 1 && len(args)%2 == 0
}

// Fields2String returns a list of string
// representations of a variadic list of args.
func Fields2String[T any](args ...T) []string {
	s := fmt.Sprint(args)
	return strings.Fields(s)
}

// ToFields cleans and normalizes whitespace
// in a slice of strings.
func ToFields(elems ...string) []string {
	return strings.Fields(strings.Join(elems, " "))
}

// Join concatenates the string representations of a variadic
// list of elements to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join[T any, E ~[]T](elems E, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return fmt.Sprint(elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += Len(elems[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(fmt.Sprintf("%v", elems[0]))
	// b.WriteString(elems[0])
	for _, v := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%v", v))
		// b.WriteString(v)
	}
	return b.String()
}

// Args2Pairs returns a slice of string pairs from
// a standard slice of strings. Whitespace is trimmed
// and normalized.
//
// Args2Pairs will panic if the number of arguments
// is less than 2 or not a multiple of 2.
func Args2Pairs(args ...string) [][2]string {
	if len(args) < 2 {
		panic("args2pairs: must be at least 2 arguments")
	}
	if len(args)%2 != 0 {
		panic("args2pairs: number of arguments must be multiple of 2")
	}

	// length of args is a positive even integer number of arguments
	fields := ToFields(args...)
	listLen := len(fields) / 2

	list := make([][2]string, 0, listLen)
	for i := 0; i < len(fields); i = +2 {
		list = append(list, [2]string{fields[i], fields[i+1]})
	}
	return list
}
