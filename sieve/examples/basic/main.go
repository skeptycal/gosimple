package main

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/skeptycal/gosimple/primes10k"
	"github.com/skeptycal/gosimple/sieve"
)

var (
	n            int   = 100                 // number of iterations to test
	checklist    []int = []int{}             // list of primes for verification
	maxcheck     int   = len(primes10k.List) // maximum "checkable" primes
	Eratosthenes       = sieve.Eratosthenes
	list               = Eratosthenes.List()
)

func init() {
	if len(list) <= maxcheck {
		checklist = primes10k.List[:len(list)]
	} else {
		checklist = primes10k.List
	}

	sieve.DbEcho("projected max check number:", sieve.MaxSearchValue(n))
	sieve.DbEcho("checklist len: ", len(checklist))
	sieve.DbEcho("test list len: ", len(list))
	sieve.DbEcho("sieveCount: ", sieve.PrimeCountSieve(n))

}

func main() {

	// }
	// func main() {

	sizes := make([]int, n)

	for i := 0; i < n; i++ {
		sizes[i] = sieve.PrimeCountSieve(i)
	}

	if !sort.IntsAreSorted(list) {
		sort.Ints(list)
	}

	if !sort.IntsAreSorted(checklist) {
		sort.Ints(checklist)
	}

	if reflect.DeepEqual(list, checklist) {
		fmt.Println("slices are equal")
	} else {
		fmt.Println("slices are not equal")
		for i := 0; i < len(list); i++ {
			if list[i] != checklist[i] {
				fmt.Printf("%v != %v", list[i], checklist[i])
			}
		}
	}

	// for i := 1; i < n; i = i + 2 {
	// go sieve(n)
	// }
}
