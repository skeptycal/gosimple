package tests

import "testing"

type (
	// BenchmarkRunner interface {
	// 	Name() string
	// 	Run(b *testing.B) error
	// }

	// BenchmarkTableEntry[G any, W comparable] interface {
	// 	BenchmarkRunner
	// 	TableEntryInfo[G, W]
	// }

	BenchmarkTableEntry[G any, W comparable] struct {
		Name    string
		Fn      func(in ...G) W
		In      []G
		Want    W
		WantErr bool
	}
)

// func (bte *BenchmarkTableEntry[G, W]) Name() string  { return bte.name }
// func (bte *BenchmarkTableEntry[G, W]) In() []G       { return bte.in }
// func (bte *BenchmarkTableEntry[G, W]) Want() W       { return *new(W) }
// func (bte *BenchmarkTableEntry[G, W]) Got() W        { return *new(W) }
// func (bte *BenchmarkTableEntry[G, W]) WantErr() bool { return false }
func (bte *BenchmarkTableEntry[G, W]) Run(b *testing.B) error {
	var retval W
	for i := 0; i < b.N; i++ {
		retval = bte.Fn(bte.In...)
	}
	_ = retval /// TODO: global return, etc.

	return errNotImplemented
}
