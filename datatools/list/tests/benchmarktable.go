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
) BenchmarkRunner {
	return BenchmarkTable[In, W]{name, nil}.AddSet(name, fn, testdata)
}

type (
	BenchmarkTable[In any, W comparable] struct {
		Name   string
		Return bool
		Global bool
		Funcs  []BenchmarkFunc[In, W]
		Inputs []BenchmarkInput[In, W]
	}

	BenchmarkRunner interface {
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

var globalSink any

func (tbl *BenchmarkTable[In, W]) Run(b *testing.B) {
	var retval W
	if tbl.Return {
		if tbl.Global {
			retval = globalSink.(W)
		} else {
			retval = *new(W)
		}
	}

	for _, ff := range tbl.Funcs {
		b.ResetTimer()
		for _, bb := range tbl.Inputs {
			name := fmt.Sprintf("%s - %s(%s): ", tbl.Name, ff.NameFunc, bb.Name)
			if tbl.Return {
				fn = BMReturn[In, W]
			}
			b.Run(name, func(b *testing.B) {

				for i := 0; i < b.N; i++ {
					retval = ff.Fn(bb.In)
				}

			})
			_ = bb.Run(b) // TODO: global return, etc
		}
	}
}

func BMReturn[In any, W comparable](in In, fn func(In) W) W {
	return fn(in)
}

func BMNoReturn[In any, W comparable](in In, fn func(In) W) {
	fn(in)
}
func (tbl *BenchmarkTable[In, W]) AddSet(name string, fn func(In) W, entries []TestDataDetails[In, W]) BenchmarkRunner {
	if len(entries) == 0 {
		return nil
	}
	for _, entry := range entries {
		tt := TestDataType[In, W]{name, fn, entry}
		tbl.Add(tt)
	}
	return tbl
}

func (tbl *BenchmarkTable[In, W]) Add(entry TestDataType[In, W]) {
	tbl.Tests = append(tbl.Tests, entry)
}
