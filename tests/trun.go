package tests

import (
	"reflect"
	"testing"

	"github.com/skeptycal/gosimple/types/constraints"
)

type Ordered constraints.Ordered

func TRun[T comparable](t *testing.T, name, arg string, got, want T) {
	if !AssertEqual(got, want) && !AssertDeepEqual(got, want) {
		t.Errorf("%v(%v) = %v, want %v", name, arg, got, want)
	}
}

func TRunEqual[T comparable](t *testing.T, name, arg string, got, want T) {
	if !AssertEqual(got, want) {
		t.Errorf("%v(%v) = %v, want %v", name, arg, got, want)
	}
}

func TRunDeep[T any](t *testing.T, name, arg string, got, want T) {
	if !AssertDeepEqual(got, want) {
		t.Errorf("%v(%v) = %v, want %v", name, arg, got, want)
	}
}

func AssertLessThan[T Ordered](got, want T) bool {
	return got < want
}

func AssertGreaterThan[T Ordered](got, want T) bool {
	return got > want
}

func AssertContains[T Ordered, E ~[]T](needle T, haystack E) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func AssertNotContains[T Ordered, E ~[]T](needle T, haystack E) bool {
	for _, v := range haystack {
		if v == needle {
			return false
		}
	}
	return true
}

func AssertEqual[T comparable](got, want T) bool {
	return got == want
	// t.Errorf("AssertEqual: %v(%v) = %v, want %v", name, arg, got, want)
}

func AssertNotEqual[T comparable](got, want T) bool {
	return got != want
	// t.Errorf("AssertEqual: %v(%v) = %v, want %v", name, arg, got, want)
}

func AssertDeepEqual[T any](got, want T) bool {
	return reflect.DeepEqual(got, want)
	// t.Errorf("AssertDeepEqual: %v(%v) = %v, want %v", name, arg, got, want)
}

func AssertNotDeepEqual[T any](got, want T) bool {
	return !reflect.DeepEqual(got, want)
	// t.Errorf("AssertDeepEqual: %v(%v) = %v, want %v", name, arg, got, want)
}

// t.Run(tt.name, func(t *testing.T) {
// 	if got := NewBufferPool(tt.size); !reflect.DeepEqual(got, tt.want) {
// 		t.Errorf("%v() = %v, want %v", testname, got, tt.want)
// 	}
// })
