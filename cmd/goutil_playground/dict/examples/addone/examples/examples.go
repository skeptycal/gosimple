package examples

import (
	"fmt"

	"github.com/skeptycal/goutil_playground/constraint"
)

func ExampleAddOne[T constraint.Number](n T) {
	i := n + 1
	fmt.Println(i)
}
