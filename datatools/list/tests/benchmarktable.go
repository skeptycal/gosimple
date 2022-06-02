package tests

import (
	"fmt"
	"testing"
)

// MakeBenchmarkRunner returns an interface that
// can be Run() to perform all benchmarks comparing
// several implementations of varying forms of an
// algorithm.
//
// The goal, or incentive, for each can be setup to
// be speed, accuracy, throughput, accuracy, memory
// usage, etc.
//
// The function takes a name, func, and slice of data
// and adds them to a table.
// The entire set is returned as a Runner ready for
// immediate use in benchmarking.
//
// The anonymous table is not accessible for any other
// functionality.
//
func MakeBenchmarkRunner[In any, W comparable](
	name string,
	Return, Global bool,
	funcs []BenchmarkFunc[In, W],
	inputs []BenchmarkInput[In, W],
) Runner {
	tbl := BenchmarkTable[In, W]{Name: name, Return: Return, Global: Global}
	tbl.AddFuncs(funcs)
	tbl.AddInputs(inputs)
	return &tbl
}

type (
	BenchmarkTable[In any, W comparable] struct {
		Name   string
		Return bool
		Global bool
		Funcs  []BenchmarkFunc[In, W]
		Inputs []BenchmarkInput[In, W]
	}

	Runner interface {
		Run(b *testing.B)
	}

	BenchmarkFunc[In any, W comparable] struct {
		NameFunc string
		Fn       func(In) W
	}

	BenchmarkInput[In any, W comparable] struct {
		Name string
		In   In
	}
)

type globalSinkType[T any] struct{ v T }

func (g *globalSinkType[T]) Set(v T) { g.v = v }
func (*globalSinkType[T]) Noop(v T)  {}

// type retfunc[In any, W comparable] func(In any, fn func(in In) W) W

func (tbl *BenchmarkTable[In, W]) Run(b *testing.B) {
	var retval func(v W)
	var localret W
	_ = localret
	if tbl.Return {
		if tbl.Global {
			var rettype globalSinkType[W]
			retval = rettype.Set
		} else {
			retval = func(v W) { localret = v }
		}
	} else {
		retval = func(v W) {}
	}

	// b.ResetTimer()
	for _, bb := range tbl.Inputs {
		for _, ff := range tbl.Funcs {
			name := fmt.Sprintf("%s - %s(%s): ", tbl.Name, ff.NameFunc, bb.Name)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					retval(ff.Fn(bb.In))
				}
			})
		}
	}
}

func (tbl *BenchmarkTable[In, W]) AddFuncs(funcs []BenchmarkFunc[In, W]) {
	tbl.Funcs = append(tbl.Funcs, funcs...)
}

func (tbl *BenchmarkTable[In, W]) AddInputs(inputs []BenchmarkInput[In, W]) {
	tbl.Inputs = append(tbl.Inputs, inputs...)
}

func Ret[In any, W comparable, PT interface {
	*In
	M(in In) W
}](in In) W {
	p := PT(new(In))
	return p.M(in) // calling method on non-nil pointer
	// Reference: https://stackoverflow.com/questions/69573113/how-can-i-instantiate-a-non-nil-pointer-of-type-argument-with-generic-go
}
