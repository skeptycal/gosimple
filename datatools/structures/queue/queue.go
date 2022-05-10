package queue

import "github.com/skeptycal/gosimple/types/constraints"

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
