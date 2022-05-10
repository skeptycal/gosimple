package structures

import "errors"

// LinkedList is a data structure allowing for quick insertion and deletion.
// Access O(n)
// Search O(n)
//* Insert O(1) for known node location -or- O(n) search for node location
//* Delete O(1) for known node location -or- O(n) search for node location
type (
	LinkedList[E any] struct {
		head *node[E]
		tail *node[E]
		len  int
	}

	Node[E any] interface {
		Data() *E
		Next() Node[E]
		Prev() Node[E]

		Connect(first, second Node[E])
		// Remove(n Node[E])
		InsertBetween(first, second Node[E])

		SetData(data *E)
		SetNext(Node[E])
		SetPrev(Node[E])

		Clear()
	}

	node[E any] struct {
		// Node[E]
		data *E
		next Node[E]
		prev Node[E]
	}
)

// NewNode returns a new Node object containing the data
// and pointer arguments.
func NewNode[E any, N *node[E]](data *E, next Node[E]) Node[E] {
	return &node[E]{data: data, next: next, prev: nil}
}

// Data returns the Data contained in the Node.
func (n *node[E]) Data() *E { return n.data }

// Next returns a pointer to the next Node.
func (n *node[E]) Next() Node[E] { return n.next }

// Prev always returns nil for singly linked lists.
func (n *node[E]) Prev() Node[E] { return nil }

func (n *node[E]) Clear() {
	n.next = nil
	n.data = new(E)
}

func (n *node[E]) SetPrev(v Node[E]) { n.prev = v }
func (n *node[E]) SetNext(v Node[E]) { n.next = v }
func (n *node[E]) SetData(data *E)   { n.data = data }

func (n *node[E]) InsertBetween(first, second Node[E]) {
	n.SetNext(second)
	n.SetPrev(first)
	first.SetNext(n)
	second.SetPrev(n)
}

func (n *node[E]) Remove() {

}

func (n *node[E]) Connect(first, second Node[E]) {
	first.SetNext(second)
	second.SetPrev(first)
}

func (l *LinkedList[E]) InsertBefore(parent *node[E], new *node[E]) {
	next := parent.next
	if parent == l.head {
		l.head = new
	}
	new.next = next
	l.len++
}

func (l *LinkedList[E]) InsertAfter(parent *node[E], new *node[E]) {
	next := parent.next
	if parent == l.tail {
		next = nil
		l.tail = new
	}
	parent.next = new
	new.next = next
	l.len++
}

func (l *LinkedList[E]) Delete(n Node[E]) error {
	parent, err := l.iterFind(n)
	if err != nil {
		return err
	}
	parent.SetNext(n.Next())
	n.Clear()
	l.len--
	return nil
}

func (l *LinkedList[E]) Len() int       { return l.len }
func (l *LinkedList[E]) Head() Node[E]  { return l.head }
func (l *LinkedList[E]) Tail() Node[E]  { return l.tail }
func (l *LinkedList[E]) blank() Node[E] { return &node[E]{} }

var ErrNodeNotFound = errors.New("node not found")

func (l *LinkedList[E]) iterFind(n Node[E]) (Node[E], error) {
	for v := l.Head(); v != nil; v = v.Next() {
		if v.Next() == n {
			return v, nil
		}
	}
	return nil, ErrNodeNotFound
}
