package sieve

import (
	"fmt"
	"math"
	"strings"

	"github.com/skeptycal/gosimple/constraints"
	"github.com/skeptycal/gosimple/miniansi"
)

var DbEcho = miniansi.DbEcho

type Option[T constraints.Ordered] struct{ value T }

var (
	None Option[int] = Option[int]{0}
)

type Opt[T any] interface {
	any
}

var Eratosthenes = PrimeTest{"Eratosthenes", MaxSearchValue(1000000), SieveOfEratosthenes(MaxSearchValue(1000000))}

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

// SieveOfEratosthenes is the simplest implementation
// of the famous sieve.
func SieveOfEratosthenes(value int) []int {
	ret := make([]int, 0, MaxSearchValue(value))
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

func listStringer[T constraints.Ordered](list []T) string {
	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteString(fmt.Sprintf("length of list (type: %T): %v\n", list, len(list)))
	for _, v := range list {
		sb.WriteString(fmt.Sprintf(" %v\n", v))
	}
	return sb.String()
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
