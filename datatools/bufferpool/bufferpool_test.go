package bufferpool

import (
	"bytes"
	"runtime"
	"runtime/debug"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	. "github.com/skeptycal/gosimple/tests"
)

func TestBufferPool(t *testing.T) {

}

func TestNewBufferPool(t *testing.T) {
	testname := "BufferPool"
	tests := []struct {
		name string
		size int
		want anyPooler[bytes.Buffer]
	}{
		{"0", 0, genPool[bytes.Buffer]{sync.Pool{}}},
		{"100", 100, genPool[bytes.Buffer]{sync.Pool{}}},
		{"1024", 1024, genPool[bytes.Buffer]{sync.Pool{}}},
		// {"0", 0, Pool[bytes.Buffer]{}},
	}
	for _, tt := range tests {
		got := NewBufferPool(tt.size).Get()
		want := tt.want.Get()
		TRunDeep(t, testname, tt.name, got, want)
	}
}

//////////// Smart Benchmarks adapted from the standard library sync.Pool package.

// poolTests is used to compare different implementations of sync.Pool.
var poolTests = []struct {
	name string
	pool anyPooler[any]
}{
	{"sync.Pool", &sync.Pool{}},
	// {"bufferPool", anyPool(&bufferPool{0, sync.Pool{}})},
	// {"genPool", anyPool(NewBufferPool(1))},
	// {"Pool", &Pool[int]{}},
}

func BenchmarkPool(b *testing.B) {
	// for _, bb := range poolTests {
	var p Pool[int]

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
	// }
}

func BenchmarkPoolOverflow(b *testing.B) {
	var p Pool[int]

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for b := 0; b < 100; b++ {
				p.Put(b)
			}
			for b := 0; b < 100; b++ {
				p.Get()
			}
		}
	})
}

// Simulate object starvation in order to force Ps to steal objects
// from other Ps.
func BenchmarkPoolStarvation(b *testing.B) {
	var p Pool[int]
	count := 100
	// Reduce number of putted objects by 33 %. It creates objects starvation
	// that force P-local storage to steal objects from other Ps.
	countStarved := count - int(float32(count)*0.33)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for b := 0; b < countStarved; b++ {
				p.Put(b)
			}
			for b := 0; b < count; b++ {
				p.Get()
			}
		}
	})
}

func BenchmarkPoolSTW(b *testing.B) {
	// Take control of GC.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	var mstats runtime.MemStats
	var pauses []uint64

	var p Pool[any]
	for i := 0; i < b.N; i++ {
		// Put a large number of items into a pool.
		const N = 100000
		var item any = 42
		for i := 0; i < N; i++ {
			p.Put(&item)
		}
		// Do a GC.
		runtime.GC()
		// Record pause time.
		runtime.ReadMemStats(&mstats)
		pauses = append(pauses, mstats.PauseNs[(mstats.NumGC+255)%256])
	}

	// Get pause time stats.
	sort.Slice(pauses, func(i, j int) bool { return pauses[i] < pauses[j] })
	var total uint64
	for _, ns := range pauses {
		total += ns
	}
	// ns/op for this benchmark is average STW time.
	b.ReportMetric(float64(total)/float64(b.N), "ns/op")
	b.ReportMetric(float64(pauses[len(pauses)*95/100]), "p95-ns/STW")
	b.ReportMetric(float64(pauses[len(pauses)*50/100]), "p50-ns/STW")
}

var globalSink any

func BenchmarkPoolExpensiveNew(b *testing.B) {
	// Populate a pool with items that are expensive to construct
	// to stress pool cleanup and subsequent reconstruction.

	// Create a ballast so the GC has a non-zero heap size and
	// runs at reasonable times.
	globalSink = make([]byte, 8<<20)
	defer func() { globalSink = nil }()

	// Create a pool that's "expensive" to fill.
	var p Pool[any]
	var nNew uint64
	p.New = func() any {
		atomic.AddUint64(&nNew, 1)
		time.Sleep(time.Millisecond)
		return 42
	}
	var mstats1, mstats2 runtime.MemStats
	runtime.ReadMemStats(&mstats1)
	b.RunParallel(func(pb *testing.PB) {
		// Simulate 100X the number of goroutines having items
		// checked out from the Pool simultaneously.
		items := make([]any, 100)
		var sink []byte
		for pb.Next() {
			// Stress the pool.
			for i := range items {
				items[i] = p.Get()
				// Simulate doing some work with this
				// item checked out.
				sink = make([]byte, 32<<10)
			}
			for i, v := range items {
				p.Put(&v)
				items[i] = nil
			}
		}
		_ = sink
	})
	runtime.ReadMemStats(&mstats2)

	b.ReportMetric(float64(mstats2.NumGC-mstats1.NumGC)/float64(b.N), "GCs/op")
	b.ReportMetric(float64(nNew)/float64(b.N), "New/op")
}
