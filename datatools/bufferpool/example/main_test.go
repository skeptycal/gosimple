package main

import (
	"testing"
)

/*

* Using swimmer is ~23% faster than Get() + defer Put() for Int pool

/Benchmark_getput/swimmer-8         	 2073265	       561.9 ns/op	    1213 B/op	       3 allocs/op
/Benchmark_getput/getput-8          	 1660987	       729.1 ns/op	    2250 B/op	       4 allocs/op

* Using swimmer is ~35% faster than Get() + defer Put() for bytes.Buffer pool (1k buffer)

/Benchmark_getput/swimmer-8         	 1406859	       896.8 ns/op	    2309 B/op	       5 allocs/op
/Benchmark_getput/getput-8          	  877038	      1376 ns/op	    3353 B/op	       7 allocs/op

* With larger bytes.Buffer objects, using swimmer is only ~7% faster (32k buffer)
* to be fair, much of the time is spent on the "work" of generating random numbers

/Benchmark_getput/swimmer-8         	  287380	      4170 ns/op	   66818 B/op	       5 allocs/op
/Benchmark_getput/getput-8          	  266580	      4496 ns/op	   67844 B/op	       7 allocs/op

* It seems likely that these pools will favor smaller buffers, anyway ... so this is a good thing
* With 16 byte buffers, it is faster again (~30%)
/Benchmark_getput/swimmer-8         	 1737507	       693.6 ns/op	    1304 B/op	       5 allocs/op
/Benchmark_getput/getput-8          	 1217234	       983.9 ns/op	    2350 B/op	       7 allocs/op
*/

const bytesize = 1 << 24

func Benchmark_getput(b *testing.B) {
	b.Run("swimmer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			swimmer(bytesize)
		}
	})
	b.Run("getput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			getput(bytesize)
		}
	})
}

func Test_getput(t *testing.T) {

	tests := []struct {
		name string
	}{
		{"getput"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getput(bytesize)
		})
	}
}

func Test_swimmer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"swimmer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swimmer(bytesize)
		})
	}
}
