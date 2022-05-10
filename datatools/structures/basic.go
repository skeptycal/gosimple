package structures

import "github.com/skeptycal/gosimple/types/constraints"

/////////// Random Access

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

/////////// Sequential Access

// Stack is a LIFO sequential access data structure.
// Access O(n)
// Search O(n)
//* Insert O(1)
//* Delete O(1)
type Stack[I constraints.Integer, E any, S ~[]E] interface {
	Push(value E)
	Pop() E
	Peek() E
	Contains(value E) bool
	Size() I
}

// Queue is a FIFO sequential access data structure.
// Access O(n)
// Search O(n)
//* Insert O(1)
//* Delete O(1)
type Queue[I constraints.Integer, E any, S ~[]E] interface {
	Enqueue(value E)
	Dequeue() E
	Peek() E
	Contains(value E) bool
	Size() I
}

// LinkedLister is a data structure allowing for quick insertion and deletion.
// Access O(n)
// Search O(n)
//* Insert O(1) for known node location -or- O(n) search for node location
//* Delete O(1) for known node location -or- O(n) search for node location
type LinkedLister[E any] interface {
	Head() node[E]
	Tail() node[E]
	InsertBefore(parent *node[E], new *node[E])
	InsertAfter(parent *node[E], new *node[E])
	Delete(n *node[E]) error
	Len() int
}

// DoublyLinkedLister is a data structure allowing for quick insertion and deletion.
// Access O(n)
// Search O(n)
//* Insert O(1) for known dNode location -or- O(n) search for dNode location
//* Delete O(1) for known dNode location -or- O(n) search for dNode location
type DoublyLinkedLister[E any] interface {
	Head() node[E]
	Tail() node[E]
	InsertBefore(parent *node[E], new *node[E])
	InsertAfter(parent *node[E], new *node[E])
	Delete(n *node[E]) error
	Len() int
}

/////////// Dynamic Access

// Dict is an implementation of a Dictionary.
// A dictionary (also called maps or associative arrays) are a
// data structure where values are indexed with comparable keys.
// It is one of the most abstract of the basic data structures.
//
//* Access O(1) *on average (amortized)
//* Search O(1) *on average (amortized)
//* Insert O(1) *on average (amortized)
//* Delete O(1) *on average (amortized)
type Dict[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V) error
}
