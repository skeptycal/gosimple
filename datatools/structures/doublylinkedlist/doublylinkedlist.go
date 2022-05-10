package doublylinkedlist

import "errors"

// DoublyLinkedLister is a data structure allowing for quick insertion and deletion.
// Access O(n)
// Search O(n)
//* Insert O(1) for known node location -or- O(n) search for node location
//* Delete O(1) for known node location -or- O(n) search for node location
type DoublyLinkedLister[E any] interface {
	Head() node[E]
	Tail() node[E]
	InsertBefore(parent *node[E], new *node[E])
	InsertAfter(parent *node[E], new *node[E])
	Delete(n *node[E]) error
	Len() int
}

// DoubleLinkedList is a data structure allowing for quick insertion and deletion.
// Access O(n)
// Search O(n)
//* Insert O(1) for known node location -or- O(n) search for node location
//* Delete O(1) for known node location -or- O(n) search for node location
type (
	DoubleLinkedList[E any] struct {
		head Node[E]
		tail Node[E]
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
		data *E
		next Node[E]
		prev Node[E]
	}
)

func NewDNode[E any](data *E, prev Node[E], next Node[E]) Node[E] {
	return &node[E]{data: data, next: next, prev: prev}
}

// Clear sets all node values to their zero values.
func (n *node[E]) Clear() {
	n.prev = nil
	n.next = nil
	n.data = new(E)
}

// Data returns the Data contained in the Node.
func (n *node[E]) Data() *E { return n.data }

// Next returns a pointer to the next Node.
func (n *node[E]) Next() Node[E] { return n.next }

// Next returns a pointer to the previous Node.
func (n *node[E]) Prev() Node[E] { return n.prev }

func (n *node[E]) SetPrev(v Node[E]) { n.prev = v }
func (n *node[E]) SetNext(v Node[E]) { n.next = v }
func (n *node[E]) SetData(data *E)   { n.data = data }

func (n *node[E]) InsertBetween(first, second Node[E]) {
	n.SetNext(second)
	n.SetPrev(first)
	first.SetNext(n)
	second.SetPrev(n)
}

func (l *DoubleLinkedList[E]) InsertBefore(first Node[E], new Node[E]) {
	if first == l.head {
		l.head = new
	}

	new.SetNext(first)
	new.SetPrev(first.Prev())
	first.SetPrev(new)
	l.len++
}

func (l *DoubleLinkedList[E]) InsertAfter(second Node[E], new Node[E]) {
	if second == l.tail {
		l.tail = new
	}

	new.SetPrev(second)
	new.SetNext(second.Next())
	second.SetNext(new)
	l.len++
}

func (n *node[E]) Connect(first, second Node[E]) {
	first.SetNext(second)
	second.SetPrev(first)
}

func (l *DoubleLinkedList[E]) Delete(n Node[E]) error {
	defer n.Clear()

	// if Head
	if n.Prev() == nil {
		l.head = n.Next()
	}

	// if Tail
	if n.Next() == nil {
		l.tail = n.Prev()
	}

	n.Next().SetPrev(n.Prev())
	n.Prev().SetNext(n.Next())

	l.len--
	return nil
}

func (l *DoubleLinkedList[E]) Len() int       { return l.len }
func (l *DoubleLinkedList[E]) Head() Node[E]  { return l.head }
func (l *DoubleLinkedList[E]) Tail() Node[E]  { return l.tail }
func (l *DoubleLinkedList[E]) blank() Node[E] { return new(node[E]) }

// var ErrDNodeNotFound = errors.New("DNode not found")

func (l *DoubleLinkedList[E]) findFromHead(n Node[E]) (Node[E], error) {
	for v := l.head; v != nil; v = v.Next() {
		if v.Next() == n {
			return v, nil
		}
	}
	return nil, ErrNodeNotFound
}

func (l *DoubleLinkedList[E]) iterHead() (out chan Node[E]) {
	go func() {
		for v := l.head; v != nil; v = v.Next() {
			out <- v
		}
	}()
	return out
}

func (l *DoubleLinkedList[E]) iterTail() (out chan Node[E]) {
	go func() {
		for v := l.tail; v != nil; v = v.Prev() {
			out <- v
		}
	}()
	return out
}

func (l *DoubleLinkedList[E]) findFromTail(n Node[E]) (Node[E], error) {
	for v := l.tail; v != nil; v = v.Prev() {
		if v.Prev() == n {
			return v, nil
		}
	}
	return nil, ErrNodeNotFound
}

var ErrNodeNotFound = errors.New("node not found")
