package sieve

import (
	"math"
	"unsafe"

	"github.com/skeptycal/gosimple/types/constraints"
)

type Floater64 interface {
	ToFloat64() float64
}

type Real constraints.Real

type PrimeTest struct {
	name string // algorithm or idea used
	n    int    // max search space
	list []int  // resulting list of primes between [1 .. n]
}

func (p *PrimeTest) List() []int {
	if len(p.list) == 0 {
		p.Generate()
	}
	return p.list
}

func (p *PrimeTest) Generate() {
	list := make([]int, 0, PrimeCountLog(p.n))
	p.list = append(list, SieveOfEratosthenes(p.n)...)
}

// type primeTester interface {
// 	constraints.Ordered
// 	Len() int
// }
// func Len[T primeTester](x T) int {
// 	return x.Len()
// }

func (p *PrimeTest) Len() int           { return len(p.list) }
func (p *PrimeTest) Less(i, j int) bool { return p.list[i] < p.list[j] }
func (p *PrimeTest) Swap(i, j int)      { p.list[i], p.list[j] = p.list[j], p.list[i] }

// basic uses the IsPrime function to count
// the number of primes in the interval
// [0 .. max], prints the list, and prints
// the control list from the sieve of
// // Eratosthenes.
// func basic(max int) {
// 	list := make([]int, max)
// 	for i := 1; i <= max; i++ {
// 		if IsPrime(i) {
// 			list = append(list, i)
// 		}
// 	}
// 	fmt.Print(listStringer("", list))
// 	fmt.Println("")
// 	for i := 1; i <= max; i++ {
// 		if IsPrimeSqrt(i) {
// 			fmt.Printf("%v ", i)
// 		}
// 	}
// 	fmt.Println("")
// 	SieveOfEratosthenes(max)
// }

func PrimeCountLog[T Real](n T) T {

	return T(float64(n) / math.Log(float64(n)))
}

// func real2float[T Real](n T) float64 {
// 	return *(*float64)(unsafe.Pointer(&n))
// }

func ToFloat64[T Real](n T) float64 {
	return *(*float64)(unsafe.Pointer(&n))
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

// MaxSearchValue is an estimate of the highest value that
// will contain the square root of n.
func MaxSearchValue(n int) int {
	return int(math.Sqrt(float64(n)))
}

// PrimeCountSieve returns the number of primes in
// the interval [0 .. n] using the sieve of
// Eratosthenes.
func PrimeCountSieve(n int) int {
	count := 0
	f := make([]bool, n)
	for i := 2; i <= MaxSearchValue(n); i++ {
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
