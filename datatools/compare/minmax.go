package compare

import (
	. "github.com/skeptycal/gosimple/types/constraints"
)

// func LessThan[T Ordered](a, b T) bool {
// 	if v, ok := a.(Basic[T]); ok {
// 		return v.Less(b)
// 	}
// }

// Max returns the item that is greater as defined
// by the constraint type T. T must be an Ordered
// data type or
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// MinSet returns the minimum value in a set of values.
func MinSet[T Ordered](v ...T) T {
	list := sortOrdered[T]{v}
	list.Sort()
	return list.First()
}

// MaxSet returns the maximum value in a set of values.
func MaxSet[T Ordered](v ...T) T {
	list := sortOrdered[T]{v}
	list.Sort()
	return list.Last()
}

// Sort returns a sorted list from a list of values.
func Sort[T Ordered](v ...T) []T {
	list := sortOrdered[T]{v}
	list.Sort()
	return list.list
}

// MinMax returns the min and max of the two inputs
// in increasing order, i.e. the min is the first
// return value and the max is the second.
// If the inputs are equal, the the order is maintained.
func MinMax[T Ordered](a, b T) (T, T) {
	if a > b {
		return b, a
	}
	return a, b
}
