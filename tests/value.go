package tests

import (
	"fmt"
	"reflect"
)

// Contains returns true if haystack contains needle.
func Contains[T comparable, E ~[]T](needle T, haystack E) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// Len returns the length of any arbitrary argument.
// The method used depends on the type of argument and
// it may not give the most expected results:
//
// for types Array, Pointer to array, Slice, Map, String,
// and Channel, the builtin len() function is used and
// will return the number of elements in the collection
// as described in the standard library:
// 	Array: the number of elements in v.
// 	Pointer to array: the number of elements in *v (even if v is nil).
// 	Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
// 	String: the number of bytes in v.
// 	Channel: the number of elements queued (unread) in the channel buffer;
// 	if v is nil, len(v) is zero.
//
// For boolean values, the length is 1.
// For pointer types, the value returned is based on the
// underlying value.
// For other types, it returns the length of the string
// representation of the argument using the fmt package.
func Len[T any](elems T) int {
	v := reflect.ValueOf(elems)
	switch v.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		// Len returns v's length. It panics if v's Kind is not Array, Chan, Map, Slice, or String.
		return v.Len()
	case reflect.Bool:
		// bool values are length 1 by definition without formatting
		return 1
	case reflect.Pointer:
		return Len(reflect.Indirect(v))
	default:
		return len(fmt.Sprintf("%v", elems))
	}
}
