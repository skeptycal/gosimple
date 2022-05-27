// Package constraints defines a set of useful constraints to be used with type parameters.
package constraints

type (
	// Text defines a text constraint that permits
	// types that are often used to represent text streams.
	Text interface {
		~string | []byte
	}

	// Sortable defines a sortable constraint that allows
	// user defined sortable types.
	Sortable interface {
		Ordered | UserOrdered[Sorter]
	}

	// Number is a constraint that permits any real
	// or complex number type.
	// If future releases of Go add new predeclared
	// real or complex types, this constraint will
	// be modified to include them.
	Number interface {
		Integer | Float | Complex
	}

	// Real is a constraint that permits any real
	// number type.
	// If future releases of Go add new predeclared
	// integer or floating-point types, this constraint
	// will be modified to include them.
	Real interface {
		Integer | Float
	}

	// Real64 is a constraint that permits any real
	// number type 64 bits in size.
	// If future releases of Go add new predeclared
	// integer or floating-point types, this constraint
	// will be modified to include them.
	Real64 interface {
		Integer64 | ~float64
	}

	//// Go (soon to be) standard library constraints

	// From the go repo: https://cs.opensource.google/go/x/exp/+/7b9b53b0:constraints/constraints.go
	// Copyright 2021 The Go Authors. All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE file.

	// Comparable is an interface that is implemented
	// by all comparable types (booleans, numbers, strings,
	// pointers, channels, arrays of comparable types,
	// structs whose fields are all comparable types).
	// The Comparable interface may only be used as a
	// type parameter constraint, not as the type of a variable.
	// If future releases of Go add new comparable types,
	// this constraint will be modified to include them.
	Comparable interface{ comparable }

	// Ordered is a constraint that permits any
	// ordered type: any type that supports the
	// operators < <= >= >, as well as == and !=.
	// If future releases of Go add new ordered types,
	// this constraint will be modified to include them.
	Ordered interface {
		Integer | Float | ~string
	}

	// Integer is a constraint that permits any
	// integer type.
	// If future releases of Go add new predeclared
	// integer types, this constraint will be modified
	// to include them.
	Integer interface {
		Signed | Unsigned
	}

	// Integer is a constraint that permits any
	// integer type 64 bits in size.
	// If future releases of Go add new predeclared
	// integer types, this constraint will be modified
	// to include them.
	Integer64 interface {
		Signed64 | Unsigned64
	}

	// IntegerNo8Bit is a constraint that permits any
	// integer type larger than 8 bits.
	// If future releases of Go add new predeclared
	// integer types, this constraint will be modified
	// to include them.
	IntegerNo8Bit interface {
		SignedNo8Bit | UnsignedNo8Bit
	}

	// Float is a constraint that permits any
	// floating-point type.
	// If future releases of Go add new predeclared
	// floating-point types, this constraint will be
	// modified to include them.
	Float interface {
		~float32 | ~float64
	}

	// Complex is a constraint that permits any
	// complex numeric .
	// If future releases of Go add new predeclared
	// complex numeric types, this constraint will
	// be modified to include them.
	Complex interface {
		~complex64 | ~complex128
	}

	// Signed is a constraint that permits any
	// signed integer type.
	// If future releases of Go add new predeclared
	// signed integer types, this constraint will be
	// modified to include them.
	Signed interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}

	// SignedNo8Bit is a constraint that permits any
	// signed integer 64 bits in size.
	// If future releases of Go add new predeclared
	// signed integer types, this constraint will be
	// modified to include them.
	Signed64 interface {
		~int | ~int64
	}

	// SignedNo8Bit is a constraint that permits any
	// signed integer type larger than 8 bits.
	// If future releases of Go add new predeclared
	// signed integer types, this constraint will be
	// modified to include them.
	SignedNo8Bit interface {
		~int | ~int16 | ~int32 | ~int64
	}

	// Unsigned is a constraint that permits any
	// unsigned integer type.
	// If future releases of Go add new predeclared
	// unsigned integer types, this constraint will
	// be modified to include them.
	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}

	// Unsigned64 is a constraint that permits any
	// unsigned integer 64 bits in size.
	// If future releases of Go add new predeclared
	// unsigned integer types, this constraint will
	// be modified to include them.
	Unsigned64 interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}

	// UnsignedNo8Bit is a constraint that permits any
	// unsigned integer type larger than 8 bits.
	// If future releases of Go add new predeclared
	// unsigned integer types, this constraint will
	// be modified to include them.
	UnsignedNo8Bit interface {
		~uint | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
)
