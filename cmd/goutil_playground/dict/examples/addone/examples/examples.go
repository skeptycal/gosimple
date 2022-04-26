package examples

import (
	"fmt"

	"github.com/skeptycal/gosimple/types/constraints"
)

func ExampleAddOne[T constraints.Number](n T) {
	i := n + 1
	fmt.Println(i)
}
