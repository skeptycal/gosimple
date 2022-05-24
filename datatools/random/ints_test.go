package random

import (
	"fmt"
	"math/big"
	"runtime"
	"strings"
	"testing"

	"github.com/skeptycal/gosimple/datatools/random/assert.go"
)

var (
	NewTest = assert.NewTestCase[int, int]
)

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
		{"0", NewBigInt(0), new(big.Int).SetInt64(0), false},
		{"-1", NewBigInt(-1), new(big.Int).SetInt64(1), false},
		{"1", NewBigInt(1), new(big.Int).SetInt64(1), false},
		{"255", NewBigInt(255), new(big.Int).SetInt64(255), false},
		{"1<<12", NewBigInt(1 << 12), new(big.Int).SetInt64(1 << 12), false},
		{"1<<30", NewBigInt(1 << 30), new(big.Int).SetInt64(1 << 30), false},
		{"1<<42", NewBigInt(1 << 42), new(big.Int).SetInt64(1 << 42), false},
		{"1<<60", NewBigInt(1 << 60), new(big.Int).SetInt64(1 << 60), false},
		{"1-1<<42", NewBigInt(1 - 1<<42), new(big.Int).SetInt64(1 << 42), false},
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

func TestInt(t *testing.T) {

	testFunc := Int[int, int]
	testAssertion := assert.Success[int]

	assert.NewTestSet("TestInt", []assert.Runner{
		assert.NewTestCase("0", 0, 0, testFunc, testAssertion, false),
		assert.NewTestCase("42", 42, 42, testFunc, testAssertion, false),
		assert.NewTestCase("255", 255, 255, testFunc, testAssertion, false),
		assert.NewTestCase("T", 0, 0, testFunc, testAssertion, false),
	}).Run(t)

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
