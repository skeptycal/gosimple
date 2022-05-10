package main

import (
	"fmt"

	"github.com/skeptycal/gosimple/datatools/structures/stack"
)

func main() {
	s := stack.New[int, []*int](10)

	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	fmt.Println("Len: ", s.Len())
	fmt.Println("Cap: ", s.Cap())
	fmt.Println(s)
}
