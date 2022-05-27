package constraints

/// User Defined types ... some of the discussion:
// Lesser defines a type-parameterized interface[1] with one method, Less, which returns a boolean for whether the caller, of type T, is less than some other instance of T. This is blatantly stolen from Robert Griesemer's talk at Gophercon 2020[2] about the type parameters proposal.
//
// Probably more controversially, this library also defines a wrapper called Basic over the built-in numerical types, exposing the underlying < operator through this Less method. The reasoning for this follows.
//
// theInterface is an interface that wraps the Less method.
//
// Less compares a caller of type T to some other variable of type T,
// returning true if the caller is the lesser of the two values, and false
// otherwise. If the two values are equal it returns false.
//
// Ian Lance Taylor made (what I think is) a great suggestion here, arguing for constraining containers to any and passing a comparison function to the constructor. This is a good idea and subjectively feels more idiomatic than what I have done here. [3]
//
// [1]: https://github.com/lelysses/lesser
//
// [2]: https://www.youtube.com/watch?v=TborQFPY2IM&ab_channel=GopherAcademy
//
// [3]: https://github.com/golang/go/issues/47632#issuecomment-897168431

type (
	Or interface {
	}

	Orderer[T any] interface {
		Less(other T) bool
		Greater(y T) bool
		LE(y T) bool
		GE(y T) bool
		Equal(y T) bool
		NotEqual(y T) bool
	}

	// theInterface[T any] interface{ Less(other T) bool }

	// Sorter is the interface for sorting from the standard library sort package.
	Sorter interface {
		Len() int
		Less(i, j int)
		Swap(i, j int)
	}

	UserOrdered[T Sorter] interface{ *T }
)

// Basic is a parameterized type that abstracts over
// the entire class of Ordered types (the set of Go
// built-in types which respond to < <= >= > == !=``````````````````````````````````````````````````````````2222222222222222222222222222222222222222222
// operators) and exposes this behavior via a Less
// method so that they fall under the Orderer constraint.
type Basic[O Ordered] struct{ Val O }

// Less implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is less than that of the parameter; otherwise returns
// false.
func (x Basic[O]) Less(y Basic[O]) bool { return x.Val < y.Val }

// Greater implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is greater than that of the parameter; otherwise returns
// false.
func (x Basic[O]) Greater(y Basic[O]) bool { return x.Val > y.Val }

// LE implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is less than or equal to that of the parameter;
// otherwise returns false.
func (x Basic[O]) LE(y Basic[O]) bool { return x.Val <= y.Val }

// GE implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is greater than or equal to that of the parameter;
// otherwise returns false.
func (x Basic[O]) GE(y Basic[O]) bool { return x.Val >= y.Val }

// Equal implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is equal to that of the parameter; otherwise returns
// false.
func (x Basic[O]) Equal(y Basic[O]) bool { return x.Val == y.Val }

// NotEqual implements Orderer[Basic[O]] for Basic[O]. Returns true
// if the value of the caller is not equal to that of the parameter;
// otherwise returns false.
func (x Basic[O]) NotEqual(y Basic[O]) bool { return x.Val != y.Val }
