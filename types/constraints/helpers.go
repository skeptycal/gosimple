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
	Orderer[O any] interface {
		Comparabler[O]
		LT(other O) bool
		GT(other O) bool
		LE(other O) bool
		GE(other O) bool
	}

	Comparabler[O any] interface {
		EQ(other O) bool
		NE(other O) bool
	}

	// Sorter is the interface for sorting from the standard library sort package.
	Sorter interface {
		Len() int
		Less(i, j int)
		Swap(i, j int)
	}

	UserSorter[T Sorter] interface{ *T }

	userOrdered[T any] struct {
		userComparable[T]
		any
	}

	userComparable[T any] struct {
		any
	}
)
