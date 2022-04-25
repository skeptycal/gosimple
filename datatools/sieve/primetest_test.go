package sieve

import (
	"fmt"
	"testing"
)

type testItem[T Real] struct {
	n T
}

func (t *testItem[T]) ToFloat64() float64 { return float64(t.n) }

func (t *testItem[T]) ToInt() int { return int(t.n) }

func (t *testItem[T]) String() string {
	return fmt.Sprint(t.n)
}

func NewItem[T Real](n T) Floater64 {
	ti := &testItem[T]{n}
	return ti
}

var rtfTests = []struct {
	name string
	n    Floater64
	want float64
}{
	// TODO: Add test cases.
	{"int 42", NewItem(42), 42},
	{"float 42", NewItem(42.0), 42.0},
	{"float 0.01", NewItem(0.01), 0.01},
	{"int -1", NewItem(-1), -1.0},
}

func rtfTest(t *testing.T) {
	for _, tt := range rtfTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.ToFloat64(); got != tt.want {
				t.Errorf("real2float() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_real2float(t *testing.T) {
	rtfTest(t)
}
