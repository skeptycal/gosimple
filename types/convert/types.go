package convert

import "github.com/skeptycal/gosimple/constraints"

type (

	// Number is a constraint that permits any real or complex number type.
	Number constraints.Number

	// Real is a constraint that permits any real number type.
	Real constraints.Real

	// Ordered is a constraint that permits any ordered type: any type that supports the operators < <= >= >.
	Ordered constraints.Ordered

	Stringable any
)
