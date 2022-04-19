package concurrent

import (
	"fmt"
	"testing"
)

func newGlobalVar[T any](v T) globalVar[T] { return globalVar[T]{v} }

var testsStringGen = []struct {
	name    string             // name of the test
	n       int                // length of string
	fn      func(n int) string // function producing string
	wantErr bool               // indicates an error is expected
}{
	// TODO: Add test cases.
	{"base test", 32, RandString, false},
}

type globalVar[T any] struct{ v T }

func (v globalVar[T]) String() string { return fmt.Sprintf("%v(%T)", v.v) }
func (v globalVar[T]) Set(t T)        { v.v = t }
func (v globalVar[T]) Get() T         { return v.v }

func BenchmarkRandString(b *testing.B) {

	// TODO test using map, sync.map, channels, etc ...

	// loop with local function call
	for _, bb := range testsStringGen {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bb.fn(bb.n)
			}
		})
	}

	g := newGlobalVar("")
	// return to global variable
	for _, bb := range testsStringGen {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				g.Set(bb.fn(bb.n))
			}
		})
	}

	var gmap = newGlobalVar(map[int]string{})
	// save in global list (or structure)
	for _, bb := range testsStringGen {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				g.Set(bb.fn(bb.n))
			}
		})
	}

}

func TestRandString(t *testing.T) {

	for _, tt := range testsStringGen {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandString(tt.n); len(got) != tt.n {
				t.Errorf("RandString(%v) length = %v, want %v", tt.n, got, tt.n)
			}
		})
	}
}
