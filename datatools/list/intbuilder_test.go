package list_test

import (
	"bytes"
	"math/rand"
	"testing"
	"unsafe"

	"github.com/skeptycal/gosimple/datatools/list"
)

/* Benchmark Results

/1Write_NoGrow(int)-8         	 9007520	       131.9 ns/op	     896 B/op	       1 allocs/op
/3Write_NoGrow(int)-8         	 1266499	       824.0 ns/op	    6784 B/op	       3 allocs/op
/3Write_Grow(int)-8           	 2912071	       405.1 ns/op	    2688 B/op	       1 allocs/op

/1Write_NoGrow(byte)-8        	29722618	        40.39 ns/op	     112 B/op	       1 allocs/op
/3Write_NoGrow(byte)-8        	 7120958	       170.7 ns/op	     784 B/op	       3 allocs/op
/3Write_Grow(byte)-8          	14372612	        83.13 ns/op	     320 B/op	       1 allocs/op

/1Write_NoGrow(rune)-8        	15920599	        75.18 ns/op	     416 B/op	       1 allocs/op
/3Write_NoGrow(rune)-8        	 3138568	       379.4 ns/op	    3104 B/op	       3 allocs/op
/3Write_Grow(rune)-8          	 5466348	       217.7 ns/op	    1280 B/op	       1 allocs/op

/1Write_NoGrow(int64)-8       	 9155007	       130.3 ns/op	     896 B/op	       1 allocs/op
/3Write_NoGrow(int64)-8       	 1412352	       828.0 ns/op	    6784 B/op	       3 allocs/op
/3Write_Grow(int64)-8         	 2981053	       403.6 ns/op	    2688 B/op	       1 allocs/op

/1Write_NoGrow(uint64)-8      	 8826916	       136.0 ns/op	     896 B/op	       1 allocs/op
/3Write_NoGrow(uint64)-8      	 1393936	       877.6 ns/op	    6784 B/op	       3 allocs/op
/3Write_Grow(uint64)-8        	 2915042	       410.1 ns/op	    2688 B/op	       1 allocs/op

/1Write_NoGrow(string)-8      	 3769698	       290.9 ns/op	    1792 B/op	       1 allocs/op
/3Write_NoGrow(string)-8      	  564712	      2129 ns/op	   14080 B/op	       3 allocs/op
/3Write_Grow(string)-8        	 1587441	       764.8 ns/op	    4864 B/op	       1 allocs/op

/1Write_NoGrow(uint)-8        	 9180537	       144.6 ns/op	     896 B/op	       1 allocs/op
/3Write_NoGrow(uint)-8        	 1407637	       846.4 ns/op	    6784 B/op	       3 allocs/op
/3Write_Grow(uint)-8          	 2984342	       403.5 ns/op	    2688 B/op	       1 allocs/op

/1Write_NoGrow(float64)-8     	 9211239	       134.0 ns/op	     896 B/op	       1 allocs/op
/3Write_NoGrow(float64)-8     	 1393532	       862.2 ns/op	    6784 B/op	       3 allocs/op
/3Write_Grow(float64)-8       	 2951622	       441.0 ns/op	    2688 B/op	       1 allocs/op

/1Write_NoGrow(float32)-8     	15731179	        84.23 ns/op	     416 B/op	       1 allocs/op
/3Write_NoGrow(float32)-8     	 3161264	       406.5 ns/op	    3104 B/op	       3 allocs/op
/3Write_Grow(float32)-8       	 5397398	       219.0 ns/op	    1280 B/op	       1 allocs/op

*/

type (
	IntBuilder   = list.Builder[int, []int]
	FloatBuilder = list.Builder[float64, []float64]
)

const (
	defaultRange = 100
	minRange     = 10
	n            = 100
	listMin      = 15
	listMax      = 85
)

var someInts []int = intList(n, listMin, listMax)

func intList(n, min, max int) []int {
	var rng int
	if max == min {
		rng = defaultRange
	} else {
		if min > max {
			min, max = max, min
		}
		rng = max - min
		if rng < minRange {
			rng = minRange
		}
	}
	retval := make([]int, n)
	for i := 0; i < n; i++ {
		retval[i] = rand.Intn(rng) + min
	}
	return retval
}

var sinkS2 string

func benchmarkIntBuilder(b *testing.B, name string, f func(b *testing.B, numWrite int, grow bool)) {
	b.Run("1Write_NoGrow("+name+")", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 1, false)
	})
	b.Run("3Write_NoGrow", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 3, false)
	})
	b.Run("3Write_Grow", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 3, true)
	})
}

func BenchmarkIntBuildString_Builder(b *testing.B) {
	benchmarkBuilder(b, "IntBuilder", func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf IntBuilder
			if grow {
				buf.Grow(len(someInts) * numWrite)
			}
			for i := 0; i < numWrite; i++ {
				buf.Write(someInts)
			}
			sinkS = buf.String()
		}
	})
}

func toT[T list.Ordered](v int64) T { return *(*T)(unsafe.Pointer(&v)) }

func randList[T list.Ordered, E ~[]T](n int) E {
	retval := make([]T, n)
	for i := 0; i < n; i++ {
		retval[i] = toT[T](rand.Int63())
	}
	return retval
}

func Generic_Builder[T list.Ordered](b *testing.B, numWrite int, grow bool) {
	// t := *new(T)
	some := randList[T, []T](100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf list.Builder[T, []T]
		if grow {
			buf.Grow(len(some) * numWrite)
		}
		for i := 0; i < numWrite; i++ {
			buf.Write(some)
		}
		sinkS = buf.String()
	}
}

var genericFuncs = []struct {
	name string
	fn   func(b *testing.B, numWrite int, grow bool)
}{
	{"int", g[int]},
	{"byte", g[byte]},
	{"rune", g[rune]},
	{"int64", g[int64]},
	{"uint64", g[uint64]},
	{"string", g[string]},
	{"uint", g[uint]},
	{"float64", g[float64]},
	{"float32", g[float32]},
}

func g[T list.Ordered](b *testing.B, numWrite int, grow bool) {
	Generic_Builder[T](b, numWrite, grow)
}

func BenchmarkGeneric_Builder(b *testing.B) {
	for _, ff := range genericFuncs {
		benchmarkBuilder(b, ff.name, func(b *testing.B, numWrite int, grow bool) {
			// Generic_Builder[int](b, numWrite, grow)
			ff.fn(b, numWrite, grow)
		})
	}
}

func BenchmarkIntBuildString_ByteBuffer(b *testing.B) {
	benchmarkBuilder(b, "bytes.Buffer", func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			if grow {
				buf.Grow(len(someBytes) * numWrite)
			}
			for i := 0; i < numWrite; i++ {
				buf.Write(someBytes)
			}
			sinkS = buf.String()
		}
	})
}
