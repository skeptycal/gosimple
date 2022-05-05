package list

import (
	"sort"
	"unsafe"

	"github.com/skeptycal/gosimple/types/constraints"
)

// Insert is a generic slice operation that inserts an item
// at position pos. If pos is invalid, item is appended to
// the end of the list.
func Insert[T comparable, E ~[]T](list E, item T, pos int) E {
	return insert(list, item, pos)
}

/// Below are other implementations used for benchmark testing.
func insert[T comparable, E ~[]T](list E, item T, pos int) E {
	if pos < 0 || pos >= len(list) {
		return Append(list, item)
	}

	temp := make(E, 0, len(list)+1) // preallocating is faster on most benchmarks

	copy(temp, list[:pos]) // copy is faster than append on most benchmarks
	copy(temp, E{item})    // temp = append(temp, s)
	copy(temp, list[pos+1:])

	list = nil // maybe helps GC? Doesn't cost much at all ...

	return temp
}

// insertNoZeroValue inserts an item at pos position. If pos is invalid,
// item is appended to the end of the list. If item is the zero value,
// the original list is returned unchanged.
func insertNoZeroValue[T comparable, E ~[]T](list E, item T, pos int) E {
	blank := *new(T)
	if item == blank {
		return list
	}
	if pos < 0 || pos >= len(list) {
		return append(list, item)
	}

	temp := make(E, 0, len(list)+1) // preallocating is faster on most benchmarks

	copy(temp, list[:pos]) // copy is faster than append on most benchmarks
	copy(temp, E{item})    // temp = append(temp, s)
	copy(temp, list[pos+1:])

	list = nil // maybe helps GC?

	return temp
}

// InsertSorted inserts s into the list in sorted order.
func InsertSorted[T constraints.Ordered, E ~[]T](list E, s T) E {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] > s {
			continue
		}
		return Insert(list, s, i)
	}
	return Append(list, s)
}

func insertSort[T constraints.Ordered, E ~[]T](haystack E, needle T) E {
	index := sort.Search(len(haystack), func(i int) bool { return haystack[i] > needle }) // or >= ??
	haystack = append(haystack, *new(T))
	copy(haystack[index+1:], haystack[index:])
	haystack[index] = needle
	return haystack
}

func insertSort2[T constraints.Ordered, E ~[]T](haystack E, needle T) E {
	newlist := make(E, 0, len(haystack)+1)
	index := iLoc(haystack, needle)
	// haystack = append(haystack, "")
	// copy(haystack[index+1:], haystack[index:])
	// haystack[index] = needle
	copy(newlist, haystack[:index])
	copy(newlist, E{needle})
	copy(newlist, haystack[index+1:])

	return haystack
}

func insertSortappend[T constraints.Ordered, E ~[]T](haystack E, needle T) E {
	// newlist := make(E, 0, len(haystack)+1)
	index := iLoc(haystack, needle)
	// haystack = append(haystack, "")
	// copy(haystack[index+1:], haystack[index:])
	// haystack[index] = needle
	newlist := append(haystack[:index], needle)
	newlist = append(newlist, haystack[index+1:]...)
	// copy(newlist, E{needle})
	// copy(newlist, haystack[index+1:])

	return newlist
}

func insertSortOneAppend[T constraints.Ordered, E ~[]T](haystack E, needle T) E {
	newlist := make(E, 0, len(haystack)+1)
	index := iLoc(haystack, needle)
	copy(newlist, haystack[:index])
	newlist = append(newlist, needle)
	copy(newlist, haystack[index+1:])

	return haystack
}

// toString converts []byte to string efficiently.
func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
