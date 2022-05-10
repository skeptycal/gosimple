package dict

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
