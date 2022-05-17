package main

import (
	"fmt"

	"github.com/skeptycal/gosimple/sort"
)

func main() {
	slice := sort.GenerateIntSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sorted := sort.NewQuickSortArray("[]int", slice).Sort()
	fmt.Println("\n--- Sorted ---\n\n", sorted, "\n")
}
