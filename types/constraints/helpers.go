package constraints

/// User Defined types ... some of the discussion:

// Lesser defines a type-parameterized interface[1] with one method, Less, which returns a boolean for whether the caller, of type T, is less than some other instance of T. This is blatantly stolen from Robert Griesemer's talk at Gophercon 2020[2] about the type parameters proposal.
//
// Probably more controversially, this library also defines a wrapper called Basic over the built-in numerical types, exposing the underlying < operator through this Less method. The reasoning for this follows.
type Lesser[E any] interface {
	Less(E) bool
}

// Interface is an interface that wraps the Less method.
//
// Less compares a caller of type T to some other variable of type T,
// returning true if the caller is the lesser of the two values, and false
// otherwise. If the two values are equal it returns false.
type Interface[T any] interface{ Less(other T) bool }

// Basic is a parameterized type that abstracts over the entire class of
// Ordered types (the set of Go built-in types which respond to the <
// operator), and exposes this behavior via a Less method so that they
// fall under the lesser.Interface constraint.
type Basic[N Ordered] struct{ Val N }

// Less implements Interface[Basic[N]] for Basic[N]. Returns true if the value
// of the caller is less than that of the parameter; otherwise returns
// false.
func (x Basic[N]) Less(y Basic[N]) bool { return x.Val < y.Val }

// Ian Lance Taylor made (what I think is) a great suggestion here, arguing for constraining containers to any and passing a comparison function to the constructor. This is a good idea and subjectively feels more idiomatic than what I have done here. [3]
//
// [1]: https://github.com/lelysses/lesser
//
// [2]: https://www.youtube.com/watch?v=TborQFPY2IM&ab_channel=GopherAcademy
//
// [3]: https://github.com/golang/go/issues/47632#issuecomment-897168431

// Sorter is the interface for sorting from the standard library sort package.
type Sorter interface {
	Len() int
	Less(i, j int)
	Swap(i, j int)
}

type UserOrdered[T Sorter] interface{ *T }
