package atomic

import (
	"fmt"
	"testing"
)

var incTests = []struct {
	name    string
	start   int
	finish  int
	workers int
	want    int
	wantErr bool
}{
	{"3k/1worker", 0, 3000, 1, 3000, false},
	{"3k/4workers", 0, 3000, 4, 3000, false},
}

var incFuncs = []struct {
	name string
	fn   func(start, finish, workers int) int
}{
	{"concurrentInc", concurrentInc},
	{"concurrentIncFail", concurrentIncFail},
	{"loopInc", loopInc},
}

func TestConcurrentInc(t *testing.T) {
	for _, tt := range incTests {
		for _, ff := range incFuncs {
			name := fmt.Sprintf("%v(%v):", ff.name, tt.name)
			t.Run(name, func(t *testing.T) {
				x := ff.fn(tt.start, tt.finish, tt.workers)
				if x != tt.want {
					t.Errorf("incorrect sum: got %v, want %v", x, tt.want)
				}
			})
		}
	}
}

func TestLoopInc(t *testing.T) {

}
