package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"

	"github.com/skeptycal/gorepos/constraints"
	"github.com/skeptycal/gorepos/miniansi"
	"github.com/skeptycal/gorepos/primes10k"
)

var dbecho = miniansi.DbEcho

type Option[T constraints.Ordered] struct{ value T }

var (
	None Option[int] = Option[int]{0}
)

type Opt[T any] interface {
	any
}
type primeTest[T constraints.Ordered] struct {
	name string // algorithm or idea used
	n    int    // max search space
	list []T    // resulting list of primes between [1 .. n]
}

// type primeTester interface {
// 	constraints.Ordered
// 	Len() int
// }
// func Len[T primeTester](x T) int {
// 	return x.Len()
// }

func (p *primeTest[T]) Len() int           { return len(p.list) }
func (p *primeTest[T]) Less(i, j int) bool { return p.list[i] < p.list[j] }
func (p *primeTest[T]) Swap(i, j int)      { p.list[i], p.list[j] = p.list[j], p.list[i] }

var (
	n            int   = 100                 // number of iterations to test
	checklist    []int = []int{}             // list of primes for verification
	maxcheck     int   = len(primes10k.List) // maximum "checkable" primes
	Eratosthenes       = primeTest[int]{"Eratosthenes", n, SieveOfEratosthenes(n)}
	list               = Eratosthenes.list
)

func init() {
	if len(list) <= maxcheck {
		checklist = primes10k.List[:len(list)]
	} else {
		checklist = primes10k.List
	}

	dbecho("projected max check number:", size(n))
	dbecho("checklist len: ", len(checklist))
	dbecho("test list len: ", len(list))
	dbecho("sieveCount: ", PrimeCountSieve(n))

}

func main() {

	// }
	// func main() {

	sizes := make([]int, n)

	for i := 0; i < n; i++ {
		sizes[i] = PrimeCountSieve(i)
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

// Example of Type Chaining
// Reference: https://benjiv.com/golang-generics-introduction/
func ToChan[U ~[]T, T any](t U) <-chan T {
	c := make(chan T)

	// ...

	return c
}

func Drain[E any](ch <-chan E) {

}

func Merge[E any](chs ...<-chan E) <-chan E {

	// TODO not implemented
	return chs[0]
}

func Identity[T any](v T) T { return v }

func sievecheck(n int) []int {
	size := int(math.Sqrt(float64(n)))
	ret := make([]int, size)
	for i := 0; i < n; i++ {
		// TODO stuff

	}
	return ret
}

// IsPrime returns true if value is prime. This is the simplest
// (and slowest) implementation.
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// IsPrimeSqrt returns true if value is prime. This is an improvement
// over IsPrime.
func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// size is an estimate of the highest value that
// will contain the square root of n.
func size(n int) int {
	return int(math.Sqrt(float64(n)))
}

// SieveOfEratosthenes is the simplest implementation
// of the famous sieve.
func SieveOfEratosthenes(value int) []int {
	ret := make([]int, 0, size(value))
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if !f[i] {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if !f[i] {
			ret = append(ret, i)
		}
	}
	return ret
}

// PrimeCountSieve returns the number of primes in
// the interval [0 .. n] using the sieve of
// Eratosthenes.
func PrimeCountSieve(n int) int {
	count := 0
	f := make([]bool, n)
	for i := 2; i <= size(n); i++ {
		if !f[i] {
			for j := i * i; j < n; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !f[i] {
			count++
		}
	}
	return count
}

// basic uses the IsPrime function to count
// the number of primes in the interval
// [0 .. max], prints the list, and prints
// the control list from the sieve of
// Eratosthenes.
func basic(max int) {
	list := make([]int, max)
	for i := 1; i <= max; i++ {
		if IsPrime(i) {
			list = append(list, i)
		}
	}
	fmt.Print(listStringer("", list, Option.Both))
	fmt.Println("")
	for i := 1; i <= max; i++ {
		if IsPrimeSqrt(i) {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println("")
	SieveOfEratosthenes(max)
}

func listStringer[T constraints.Ordered](list []T) string {
	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteString(fmt.Sprintf("length of list (type: %T): %v\n", list, len(list)))
	for _, v := range list {
		sb.WriteString(fmt.Sprintf(" %v\n", v))
	}
	return sb.String()
}

func sieve(n int) {
	for i := 1; i < n; i = i + 2 { // skip multiples of 2
		for j := i; j < n; j = j + 3 {

		}
	}

}

func sieveRecursiveLoops(n int) {
	for i := 1; i < n; i = i + 2 { // skip multiples of 2
		for j := i; j < n; j = j + 3 {

		}
	}

}

// TODO not implemented
func recloop(c, n int) int {
	return 0
}

func sieveBoolMap(n int) {
	m := make(map[int]bool, n)
	var newmap map[int]bool = make(map[int]bool, 50)
	stop := int(math.Sqrt(float64(n)))

	counter := 2
	for i := 0; i < stop; i++ {
		for j := 1; j < counter; j++ {
			m[i] = false
		}
		i++
		// skip multiples of counter
		m[i] = true

		// TODO below is only for debugging and testing
		if len(m) == 49 {
			newmap = m
		}
		if len(newmap) > 40 {
			printprogress(newmap)
		} else {
			printprogress(m)
		}
	}
}

func printprogress(m map[int]bool) {
	s := ""
	var kk = 0
	for k, v := range m {
		if v {
			s += "1"
		} else {
			s += "0"
		}
		kk = k
	}
	fmt.Println(kk, " : ", s)
}

func count(count, start, end int, ch chan int) chan int {
	ret := make(chan int)
	go func() {
		for j := start; j < end; j++ {
			n := <-ch // get next number
			for i := 0; i < count-1; i++ {
				n++
				ret <- n
			}
			n++ // skip i == count (would be n mod count == 0)
		}
	}()
	return ret
}
