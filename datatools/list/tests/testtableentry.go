package tests

import (
	"fmt"
	"testing"
)

// func NewEntry[In any, W comparable](name string, fn func(in ...In) W, in []In, want W, wantErr bool) *TestTableEntry[In, W] {
// 	return &TestTableEntry[In, W]{name, fn, in, want, wantErr}
// }

type (
	TestRunner interface {
		Run(t *testing.T) error
	}

	TestDataType[In any, W comparable] struct {
		NameFunc string
		Fn       func(In) W

		TestDataDetails[In, W]
	}

	TestDataDetails[In any, W comparable] struct {
		Name    string
		In      In
		Want    W
		WantErr bool
	}
)

func (tt *TestDataType[In, W]) Got() W { return tt.Fn(tt.In) }

func (tt *TestDataType[In, W]) Run(t *testing.T) (err error) {
	name := fmt.Sprintf("%s(%s): ", tt.NameFunc, tt.Name)
	t.Run(name, func(t *testing.T) {
		if tt.Got() != tt.Want {
			if !tt.WantErr {
				err = fmt.Errorf(fmtErrorfWithWantErr, name, tt.Got(), tt.Want, tt.WantErr, err)
			}
		}
	})
	return err
}

func (tt *TestDataType[In, W]) Benchmark(b *testing.B) (err error) {
	name := fmt.Sprintf("%s(%s): ", tt.NameFunc, tt.Name)
	b.Run(name, func(b *testing.B) {
		if tt.Got() != tt.Want {
			if !tt.WantErr {
				err = fmt.Errorf(fmtErrorfWithWantErr, name, tt.Got(), tt.Want, tt.WantErr, err)
			}
		}
	})
	return err
}

// TableEntryInfo[In any, W comparable] interface {
// 	In() []In
// 	Got() W
// 	Want() W
// 	WantErr() bool
// }

// TestTableEntryer[In any, W comparable] interface {
// 	TestRunner
// 	TableEntryInfo[In, W]
// }

// TestTableEntry describes a set of tests with
// with the same name and function and set of
// expected input and output values
// TestTableEntry[In any, W comparable] struct {
// 	FnName string
// 	Fn     func(in ...In) W
// 	Tests  []TestDataType[In, W]
// }

// testTableEntry[In any, W comparable] struct {
// 	name    string
// 	fn      func(in ...In) W
// 	in      []In
// 	want    W
// 	wantErr bool
// }
