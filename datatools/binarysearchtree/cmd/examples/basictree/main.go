package main

import (
	"fmt"
	"math/rand"

	"github.com/skeptycal/gosimple/datatools/binarysearchtree"
)

func main() {

	max := 20

	b := binarysearchtree.New[int, int]()

	for i := 0; i < max; i++ {
		j := rand.Intn(255)
		b.Insert(j, j)
	}
	format := " "
	b.InOrderTraverse(func(v int) { fmt.Print(v, format) })
	fmt.Println()
	b.PreOrderTraverse(func(v int) { fmt.Print(v, format) })
	fmt.Println()
	b.PostOrderTraverse(func(v int) { fmt.Print(v, format) })
	fmt.Println()
	binarysearchtree.InOrder(b.Root())

}
