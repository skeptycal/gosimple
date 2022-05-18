package main

import (
	"fmt"

	"github.com/skeptycal/gosimple/datatools/list"
	"github.com/skeptycal/gosimple/types/constraints"
)

func listExample[T constraints.Ordered, E ~[]T](v E) {
	b := list.New(v)
	fmt.Println("list len: ", b.Len())
	fmt.Println("list cap: ", b.Cap())
	fmt.Println(b)
	fmt.Println()
}

func main() {
	listExample([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	listExample([]byte("hello world"))
	listExample([]string{"h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d"})
	listExample([]rune{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'})
	listExample([]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'})
	listExample([]float64{1, 2, 3, 4, 5, 6, 7, 8})

}
