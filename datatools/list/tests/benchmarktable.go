package tests

import (
	"testing"
)

type (
	BenchmarkTable[G any, W comparable, S ~[]BenchmarkTableEntry[G, W]] interface {
		BenchmarkRunner
		Benchmarks() S
	}

	benchmarkTable[G any, W comparable, S ~[]BenchmarkTableEntry[G, W]] struct {
		name  string
		tests S
	}
)

func (tbl *benchmarkTable[G, W, S]) Name() string  { return tbl.name }
func (tbl *benchmarkTable[G, W, S]) Benchmarks() S { return tbl.tests }
func (tbl *benchmarkTable[G, W, S]) Run(b *testing.B) error {
	for _, bb := range tbl.tests {
		name := tbl.Name() + "(" + bb.Name() + ")"
		b.ResetTimer()
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bb.Run(b)
			}
		})
		_ = bb.Run(b) // TODO: global return, etc
	}
	return errNotImplemented
}
