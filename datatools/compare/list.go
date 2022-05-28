package compare

import "sort"

type sortOrdered[T Ordered] struct {
	list []T
}

func (l sortOrdered[T]) Len() int {
	return len(l.list)
}
func (l sortOrdered[T]) Less(i, j int) bool {
	return l.list[i] < l.list[j]
}

func (l sortOrdered[T]) Swap(i, j int) {
	l.list[i], l.list[j] = l.list[j], l.list[i]
}

// Sort sorts the list if necessary.
func (l sortOrdered[T]) Sort() {
	if !sort.IsSorted(l) {
		sort.Sort(l)
	}
}

func (l sortOrdered[T]) First() T {
	return l.list[0]
}

func (l sortOrdered[T]) Last() T {
	return l.list[l.Len()-1]
}
