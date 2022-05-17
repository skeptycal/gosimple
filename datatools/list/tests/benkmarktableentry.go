package tests

import "testing"

type (
	BenchmarkRunner interface {
		Name() string
		Run(b *testing.B) error
	}

	BenchmarkTableEntry[G any, W comparable] interface {
		BenchmarkRunner
		TableEntryInfo[G, W]
	}

	benchmarkTableEntry[G any, W comparable] struct {
		name    string
		fn      func(in ...G) W
		in      []G
		want    W
		wantErr bool
	}
)

func (bte *benchmarkTableEntry[G, W]) Name() string  { return bte.name }
func (bte *benchmarkTableEntry[G, W]) In() []G       { return bte.in }
func (bte *benchmarkTableEntry[G, W]) Want() W       { return *new(W) }
func (bte *benchmarkTableEntry[G, W]) Got() W        { return *new(W) }
func (bte *benchmarkTableEntry[G, W]) WantErr() bool { return false }
func (bte *benchmarkTableEntry[G, W]) Run(b *testing.B) error {
	retval := bte.fn(bte.in...)
	_ = retval /// TODO: global return, etc.

	return errNotImplemented
}
