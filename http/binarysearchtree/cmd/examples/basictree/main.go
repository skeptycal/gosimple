package main

import (
	"fmt"
	"math/rand"

	"github.com/skeptycal/gosimple/http/binarysearchtree"
)

func main() {

	max := 100

	b := binarysearchtree.New[int, int]()

	for i := 0; i < max; i++ {
		b.Insert(i, rand.Intn(255))
	}

	b.InOrderTraverse(func(v int) { fmt.Print(v, ", ") })
	fmt.Println()
	b.PreOrderTraverse(func(v int) { fmt.Print(v, ", ") })
	fmt.Println()
	b.PostOrderTraverse(func(v int) { fmt.Print(v, ", ") })
	fmt.Println()
	binarysearchtree.InOrder(b.Root())

}
