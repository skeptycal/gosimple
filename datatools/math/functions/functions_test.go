package functions

import (
	"fmt"
	"testing"

	"github.com/skeptycal/gosimple/datatools/list/tests"
)

func TestFactorialr(t *testing.T) {
	testdata := []tests.TestDataDetails[int, int]{
		{"0", 0, 1, false},
		{"1", 1, 1, true},
		{"2", 2, 2, false},
		{"3", 3, 6, false},
		{"4", 4, 24, false},
		{"5", 5, 120, false},
		{"42", 42, 7538058755741581312, false},
		{"-1", -1, 1, false},
		{"-5", -5, 1, true},
	}

	funclist := []struct {
		NameFunc string
		Fn       func(n int) int
	}{
		{NameFunc: "factorialr", Fn: factorial},
		{NameFunc: "factorialRecursive", Fn: factorialRecursive},
		{NameFunc: "factorialLoop", Fn: factorialLoop},
		{NameFunc: "factorialPtr", Fn: factorialPtr},
		{NameFunc: "factorialIteration", Fn: factorialIteration},
		{NameFunc: "FactorialClosure", Fn: FactorialClosure()},
		{NameFunc: "factorialIteration2", Fn: factorialIteration2},
		{NameFunc: "FactorialMemoization", Fn: FactorialMemoization},
	}

	for _, ff := range funclist {
		tests.MakeTestRunner(ff.NameFunc, ff.Fn, testdata).Run(t)
	}

	for _, tt := range testdata {
		checkMathFact(t, tt.In, tt.Want)
	}
}

func BenchmarkFactorial(b *testing.B) {
	/*

		* Benchmark results:

		/factorial_-_factorialRecursive(255):_-8       	  476232	      2518 ns/op	       0 B/op	       0 allocs/op
		/factorial_-_factorialLoop(255):_-8            	 3680716	       326.2 ns/op	       0 B/op	       0 allocs/op
		/factorial_-_factorialPtr(255):_-8             	 1000000	      1031 ns/op	       0 B/op	       0 allocs/op
		/factorial_-_factorialIteration(255):_-8       	 3659563	       327.7 ns/op	       0 B/op	       0 allocs/op
		/factorial_-_FactorialClosure(255):_-8         	271471024	         4.421 ns/op	   0 B/op	       0 allocs/op
		/factorial_-_factorialIteration2(255):_-8      	 3656493	       326.0 ns/op	       0 B/op	       0 allocs/op
		/factorial_-_FactorialMemoization(255):_-8     	  422281	      2842 ns/op	       0 B/op	       0 allocs/op

	*/

	inputs := []tests.BenchmarkInput[int, int]{
		{Name: "0", In: 0},
		// {Name: "1", In: 1},
		// {Name: "2", In: 2},
		// {Name: "3", In: 3},
		// {Name: "4", In: 4},
		{Name: "5", In: 5},
		// {Name: "15", In: 15},
		// {Name: "25", In: 25},
		{Name: "42", In: 42},
		{Name: "255", In: 255},
		// {Name: "-1", In: -1},
		// {Name: "-5", In: -5},
	}
	funcs := []tests.BenchmarkFunc[int, int]{
		{NameFunc: "factorial", Fn: factorial},
		{NameFunc: "factorialRecursive", Fn: factorialRecursive},
		{NameFunc: "factorialLoop", Fn: factorialLoop},
		{NameFunc: "factorialPtr", Fn: factorialPtr},
		{NameFunc: "factorialIteration", Fn: factorialIteration},
		{NameFunc: "FactorialClosure", Fn: FactorialClosure()},
		{NameFunc: "factorialIteration2", Fn: factorialIteration2},
		{NameFunc: "FactorialMemoization", Fn: FactorialMemoization},
	}

	tests.MakeBenchmarkRunner("factorial", false, false, funcs, inputs).Run(b)

}

func checkMathFact(t *testing.T, in, want int) (ok bool) {
	name := fmt.Sprintf("math.factorial(%v)", in)
	t.Run(name, func(t *testing.T) {
		if got := Factorial(in); got != want {
			t.Errorf("factorial(%v) = %v, want %v", in, got, want)
			ok = false
		} else {
			ok = true
		}
	})
	return
}
