package random

import (
	"fmt"
	"math/big"
	"runtime"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/skeptycal/gosimple/types/constraints"
)

func helperMax(x int64) *big.Int { return big.NewInt(x) }

func Test_bigInt(t *testing.T) {
	tests := []struct {
		name    string
		max     *big.Int
		wantN   *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"nil", nil, new(big.Int), false},
		{"empty", &big.Int{}, new(big.Int), false},
		{"0", helperMax(0), new(big.Int).SetInt64(0), false},
		{"-1", helperMax(-1), new(big.Int).SetInt64(1), false},
		{"1", helperMax(1), new(big.Int).SetInt64(1), false},
		{"255", helperMax(255), new(big.Int).SetInt64(255), false},
		{"1<<12", helperMax(1 << 12), new(big.Int).SetInt64(1 << 12), false},
		{"1<<30", helperMax(1 << 30), new(big.Int).SetInt64(1 << 30), false},
		{"1<<42", helperMax(1 << 42), new(big.Int).SetInt64(1 << 42), false},
		{"1<<60", helperMax(1 << 60), new(big.Int).SetInt64(1 << 60), false},
		{"1-1<<42", helperMax(1 - 1<<42), new(big.Int).SetInt64(1 << 42), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := bigInt(tt.max)
			if err != nil != tt.wantErr {
				t.Errorf("bigInt(%v) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if gotN.Cmp(tt.wantN) == 1 {
				t.Errorf("bigInt(%v) = %v, want %v", tt.name, gotN, tt.wantN)
			}
		})
	}
}

type Runner interface {
	Run(t *testing.T) (string, error)
}

type testSet []Runner

type testcase[T Ints] struct {
	name string
	in   T
	want T
	// wantErr bool
	assert func(got, want T) bool
}

type testRunner Runner

func (tc *testcase[T]) Run(t *testing.T) (fn string, err error) {
	got := Int(tc.in)
	fn = fnName(1)
	// ExampleFrames()
	if ok := tc.assert(got, tc.want); !ok {
		err = fmt.Errorf("Int(%v) (%q) = %v, want %v", tc.name, fn, got, tc.want)
		t.Errorf(err.Error())
		return
	}
	return
}

func (ts *testSet) Run(t *testing.T) (fn string, wrap error) {
	for _, tt := range *ts {
		fn, err := tt.Run(t)
		wrap = errors.Wrap(err, fn)
	}
	return
}

func helpTestInt[T Ints](tc []testcase[T]) Runner {
	newts := make(testSet, len(tc))
	for i, tt := range tc {
		newts[i] = &tt
	}
	return &newts
}

func newRunner[T Ints](name string, in T, want T, assert func(got, want T) bool) Runner {
	return &testcase[T]{name, in, want, assert}
}
func AssertLT[T constraints.Ordered](got, want T) bool { return got < want }
func AssertGT[T constraints.Ordered](got, want T) bool { return got > want }
func AssertNE[T constraints.Ordered](got, want T) bool { return got != want }
func AssertEQ[T constraints.Ordered](got, want T) bool { return got == want }

func TestInt(t *testing.T) {
	ts := testSet{
		newRunner("0", 0, 0, AssertLT[int]),
		newRunner("42", 42, 42, AssertLT[int]),
		newRunner("255", 255, 255, AssertLT[int]),
		newRunner("T", 0, 0, AssertLT[int]),
	}

	ts.Run(t)
}

func trace() (file, name string, line int) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[3])
	file, line = f.FileLine(pc[1])
	name = f.Name()
	return
	// fmt.Printf("%s:%d %s\n", file, line, f.Name())
}

func fnName(n int) string {
	if n < 1 {
		n = 1
	}
	pc := make([]uintptr, n+1) // at least 1 entry needed
	// v := runtime.Callers(n, pc)
	f := runtime.CallersFrames(pc)
	counter := 0
	for {
		frame, ok := f.Next()
		{
			if counter == n {
				return frame.Function
			}
			counter++
		}
		if !ok {
			break
		}
	}
	return ""
}

func exampleExampleFrames() {
	c := func() {
		// Ask runtime.Callers for up to 10 PCs, including runtime.Callers itself.
		pc := make([]uintptr, 10)
		n := runtime.Callers(0, pc)
		if n == 0 {
			// No PCs available. This can happen if the first argument to
			// runtime.Callers is large.
			//
			// Return now to avoid processing the zero Frame that would
			// otherwise be returned by frames.Next below.
			return
		}

		pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
		frames := runtime.CallersFrames(pc)

		// Loop to get frames.
		// A fixed number of PCs can expand to an indefinite number of Frames.
		for {
			frame, more := frames.Next()

			// Process this frame.
			//
			// To keep this example's output stable
			// even if there are changes in the testing package,
			// stop unwinding when we leave package runtime.
			if !strings.Contains(frame.File, "runtime/") {
				break
			}
			fmt.Printf("- more:%v | %s\n", more, frame.Function)

			// Check whether there are more frames to process after this one.
			if !more {
				break
			}
		}
	}

	b := func() { c() }
	a := func() { b() }

	a()
	// Output:
	// - more:true | runtime.Callers
	// - more:true | runtime_test.ExampleFrames.func1
	// - more:true | runtime_test.ExampleFrames.func2
	// - more:true | runtime_test.ExampleFrames.func3
	// - more:true | runtime_test.ExampleFrames
}
