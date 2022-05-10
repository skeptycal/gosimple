package array

import "github.com/skeptycal/gosimple/types/constraints"

// Array is a random access data structure with a fixed size.
//* Access O(1)
// Search O(n)
// Insert O(n)
// Delete O(n)
type Array[I constraints.Integer, E any, S ~[]E] interface {
	Insert(value E) error
	InsertAt(index I, value E) error
	Remove(value E) error
	RemoveAt(index I) error
	Get(value E) (index I, err error)
	Set(value E, new E) (index I, err error)
	GetAt(index int) (E, error)
	SetAt(index I, value E) error
	Clear() error
	List() S
}
