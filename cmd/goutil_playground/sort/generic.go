package sort

import "github.com/skeptycal/gosimple/types/constraints"

// An implementation of the standard library sort.Interface
// that can be sorted by the routines in this package.
type Sorter[K constraints.Ordered, V any] interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	// More details can be found in the standard
	// library sort package documentation ...
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
