package list

import (
	"sort"

	"github.com/skeptycal/gosimple/types/constraints"
)

func Remove[T constraints.Ordered, E ~[]T](slice E, pos int) E {
	return removeOrderedGeneric(slice, pos)
}

// removeWhileIter iterates through the list and removes any
// items that match value while maintaining the order of the items.
func removeWhileIter[T comparable, E ~[]T](slice E, value T) E {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == value {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// removeOrderedGeneric is a generic function that removes an item from
// an ordered list of items without modifying the order.
func removeOrderedGeneric[T constraints.Ordered, E ~[]T](slice E, pos int) E {
	return append(slice[:pos], slice[pos+1:]...)
}

func removeOrderedInts(slice []int, pos int) []int {
	return append(slice[:pos], slice[pos+1:]...)
}

func removeOrderedInterface(slice []any, pos int) []any {
	return append(slice[:pos], slice[pos+1:]...)
}

func removeInterfaceWrapper(slice []int, pos int) (retval []int) {
	flag := 0
	retval = make([]int, len(slice))
	for i, v := range slice {
		if i == pos {
			flag = 1
			continue
		}
		retval[i-flag] = v
	}
	return
}

func removeOrderedWithCheckGeneric[T constraints.Ordered, E ~[]T](slice E, pos int) E {
	if !sort.SliceIsSorted(slice, func(i, j int) bool { return slice[i] < slice[j] }) {
		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	}
	return append(slice[:pos], slice[pos+1:]...)
}

// removeUnOrderedGeneric is a generic function that removes an item from
// an ordered list of items without maintaining the order.
func removeUnOrderedGeneric[T constraints.Ordered, E ~[]T](slice E, n int) E {
	slice[n] = slice[len(slice)-1]
	// s[len(s)-1] = "" // clear the last element before removing it ... maybe helps with GC?
	return slice[:len(slice)-1]
}

func removeUnOrderedInts(slice []int, n int) []int {
	slice[n] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func removeUnOrderedWithClear[T constraints.Ordered, E ~[]T](slice E, n int) E {
	blank := new(T)
	slice[n] = slice[len(slice)-1]
	slice[len(slice)-1] = *blank // clear the last element with zero value before removing it ... maybe helps with GC?
	return slice[:len(slice)-1]
}
