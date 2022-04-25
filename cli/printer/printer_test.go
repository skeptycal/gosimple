package printer

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

/*  Sprint Benchmarks
/sprintf1-8         	1000000000	         0.0000024 ns/op	       0 B/op	       0 allocs/op
/sprintf2-8         	1000000000	         0.0000020 ns/op	       0 B/op	       0 allocs/op
/sprintf3-8         	1000000000	         0.0000023 ns/op	       0 B/op	       0 allocs/op
/sprintf4-8         	1000000000	         0.0000052 ns/op	       0 B/op	       0 allocs/op
/sprintf5-8         	1000000000	         0.0000023 ns/op	       0 B/op	       0 allocs/op
/sprintf6-8         	1000000000	         0.0000021 ns/op	       0 B/op	       0 allocs/op

* notes:
- returning string is much faster than returning (s string)
*/

var s = ""

func randString(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return string(b)
}
func BenchmarkSprint(b *testing.B) {

	maxsize := 12
	p := &print{io.Discard, io.Discard, nil, nil, nil}
	format := "%v(%v): %v"

	args := make([]any, maxsize)
	for i := 0; i < maxsize; i++ {
		args[i] = randString(1 << i)
	}

	benchmarks := []struct {
		name string
		fn   func(string, ...interface{}) string
	}{
		{"sprintf1", p.sprintf1},
		{"sprintf2", p.sprintf2},
		{"sprintf3", p.sprintf3},
		{"sprintf4", p.sprintf4},
		{"sprintf5", p.sprintf5},
		{"sprintf6", p.sprintf6},
	}
	for j := 1; j < maxsize; j++ {

		for _, bb := range benchmarks {
			for i := 0; i < b.N; i++ {
				name := fmt.Sprintf("%v(%d)", bb.name, j)
				b.Run(name, func(b *testing.B) {
					s = bb.fn(format, args...)
				})
			}
		}
	}
}
