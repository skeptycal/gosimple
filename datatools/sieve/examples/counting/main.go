package main

import (
	"fmt"

	"github.com/skeptycal/gosimple/sieve"
)

func main() {
	// test with known good algorithm:
	// list := sieve.SieveOfEratosthenes(100)
	// fmt.Println(list)
	// [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]

	// list2 := sieve.SieveCounting(100)
	// fmt.Println(list2)

	c := sieve.CountBy(3)
	d := make(chan int)
	for n := range c {
		d <- n
		fmt.Printf("%4d, ", n)
		if n > 100 {
			break
		}
	}

}
