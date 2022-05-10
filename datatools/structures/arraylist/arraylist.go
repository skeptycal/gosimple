package arraylist

import "github.com/skeptycal/gosimple/types/constraints"

// ArrayList is a random access data structure with a variable size.
//* Access O(1)
// Search O(n)
// Insert O(n)
// Delete O(n)
type ArrayList[I constraints.Integer, E any, S ~[]E] interface {
	Insert(value E) error
	InsertAt(index I, value E) error
	Remove(value E) error
	RemoveAt(index I) error
	GetAt(index int) (E, error)
	Get(value E) (index I, err error)
	Set(value E, new E) (index I, err error)
	SetAt(index I, value E) error
	Clear() error
	ToArray() (S, error)
}
