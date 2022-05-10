package structures

// DoubleLinkedList is a data structure allowing for quick insertion and deletion.
// Access O(n)
// Search O(n)
//* Insert O(1) for known dNode location -or- O(n) search for dNode location
//* Delete O(1) for known dNode location -or- O(n) search for dNode location
type (
	DoubleLinkedList[E any] struct {
		head Node[E]
		tail Node[E]
		len  int
	}

	dNode[E any] struct {
		node[E]
		// Node[E]
		// data *E
		// next Node[E]
		// prev Node[E]
	}
)

func NewDNode[E any](data *E, prev Node[E], next Node[E]) Node[E] {
	return &node[E]{data: data, next: next, prev: prev}
}

func (n *dNode[E]) Data() *E      { return n.data }
func (n *dNode[E]) Prev() Node[E] { return n.prev }
func (n *dNode[E]) Next() Node[E] { return n.next }

func (n *dNode[E]) Clear() {
	n.prev = nil
	// n.Node.Clear()
	n.next = nil
	n.data = new(E)
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
