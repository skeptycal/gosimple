package testes

import "reflect"

func AssertLessThan[T Ordered](got, want T) bool {
	return got < want
}

func AssertGreaterThan[T Ordered](got, want T) bool {
	return got > want
}

func AssertContains[T comparable, E ~[]T](needle T, haystack E) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func AssertNotContains[T comparable, E ~[]T](needle T, haystack E) bool {
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
