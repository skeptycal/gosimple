package compare

import (
	"sort"

	"github.com/skeptycal/gosimple/types/constraints"
)

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type sortOrdered[T constraints.Ordered] struct {
	list []T
}

func (l sortOrdered[T]) Len() int {
	return len(l.list)
}
func (l sortOrdered[T]) Less(i, j int) bool {
	return l.list[i] < l.list[j]
}

func (l sortOrdered[T]) Swap(i, j int) {
	l.list[i], l.list[j] = l.list[j], l.list[i]
}

// Sort sorts the list if necessary.
func (l sortOrdered[T]) Sort() {
	if !sort.IsSorted(l) {
		sort.Sort(l)
	}
}

func (l sortOrdered[T]) First() T {
	return l.list[0]
}

func (l sortOrdered[T]) Last() T {
	return l.list[l.Len()-1]
}

// MinSet returns the minimum value in a set of values.
func MinSet[T constraints.Ordered](v ...T) T {
	list := sortOrdered[T]{v}
	list.Sort()
	return list.First()
}

// MaxSet returns the maximum value in a set of values.
func MaxSet[T constraints.Ordered](v ...T) T {
	list := sortOrdered[T]{v}
	list.Sort()
	return list.Last()
}

// Sort returns a sorted list from a list of values.
func Sort[T constraints.Ordered](v ...T) []T {
	list := sortOrdered[T]{v}
	list.Sort()
	return list.list
}

// MinMax returns the min and max of the two inputs
// in increasing order, i.e. the min is the first
// return value and the max is the second.
// If the inputs are equal, the the order is maintained.
func MinMax[T constraints.Ordered](a, b T) (T, T) {
	if a > b {
		return b, a
	}
	return a, b
}
