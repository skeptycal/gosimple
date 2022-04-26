package list

// import (
// 	"sort"

// 	"github.com/skeptycal/gosimple/types/constraints"
// )

// type (
// 	Number     constraints.Number
// 	Ordered    constraints.Ordered
// 	Stringable any
// )

// type Lister[O Ordered] interface {
// 	sort.Interface
// 	BubbleSort()
// 	IsSorted() bool
// }

// type List[O Ordered] struct {
// 	list []O
// }

// func (s List[O]) Len() int           { return len(s.list) }
// func (s List[O]) Less(i, j int) bool { return s.list[i] < s.list[j] }
// func (s List[O]) Swap(i, j int)      { s.list[i], s.list[j] = s.list[j], s.list[i] }
// func (s List[O]) IsSorted() bool     { return IsSorted(s) }

// // BubbleSort sorts a slice of Ordered items in place.
// func (s List[O]) BubbleSort() {
// 	BubbleSort(s.list)
// }

// // IsSorted reports whether data is sorted.
// // func IsSorted(data Sort) bool {
// func IsSorted[T Ordered](data List[T]) bool {
// 	n := data.Len()
// 	for i := n - 1; i > 0; i-- {
// 		if data.Less(i, i-1) {
// 			return false
// 		}
// 	}
// 	return true
// }
