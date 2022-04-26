package concurrent

import "fmt"

var (
	DefaultSliceCheck float64 = 0.20
	DefaultSliceAlloc float64 = 2.00
)

func NewList[T any](n int) Slc[T] {
	return make(Slc[T], 0, n)
}

type Slc[T any] []T // sorted in alphabetical order ...

func (l Slc[T]) Less(i, j int) bool { return fmt.Sprint(l[i]) < fmt.Sprint(l[j]) }
func (l Slc[T]) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l Slc[T]) Len() int           { return len(l) }
func (l *Slc[T]) Append(v T) {
	CheckSize(l)

	list := *l
	list[list.Len()] = v
}

// CheckSize accepts a pointer to a slice and
// checks to see if the len of a slice is
// within 20% of the capacity of the slice and
// reallocates if necessary.
//
// If the capacity is still sufficient, it returns
// the initial pointer unchanged.
//
// If the capacity is insufficient, it allocates
// a new slice, sets the old slice equal to the new
// slice, and returns a pointer to the new slice.
//
// The default 20% warning and 200% allocation
// can be overridden by setting the config
// variables DefaultSliceCheck and
// DefaultSliceAlloc, respectively.
func CheckSize[T any](plist *Slc[T]) *Slc[T] {
	list := *plist

	// if the list size is withing 20% of the capacity ...
	if cap(list) > int(float64(len(list))*DefaultSliceCheck) {
		return plist
	}

	// ... double the list size
	var newlist = make(Slc[T], 0, adjustInt(len(list), DefaultSliceAlloc))
	copy(newlist, list)
	plist = &newlist
	return &newlist
}

// adjustInt adjusts the integer by a floating
// point percentage and returns a new integer.
// This will have negligible effect on small
// integers and/or small percentages.
func adjustInt(i int, percent float64) int {
	return int(float64(i) * percent)
}
