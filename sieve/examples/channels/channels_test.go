package main

import (
	// _ "net/http/pprof"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime/pprof"
	_ "runtime/pprof"
	"testing"
	"time"
)

/*
Fastest ways to merge channels:
Reference: https://gist.github.com/Xeoncross/3e0328137019b14373ee26701a23ed81

* as posted (when used within a single function and no heap escapes)
/goroutines-8			     218	   5985267 ns/op   	   18020 B/op        405 allocs/op
/atomic-8			         190	   5750518 ns/op   	   17881 B/op        403 allocs/op
/recursion-8			    7416	    440551 ns/op   	   81991 B/op        697 allocs/op

* With global return (intentional heap escape)
/goroutines-8         	    6110	    343351 ns/op	  129185 B/op	     806 allocs/op
/atomic-8             	    5570	    202344 ns/op	  131503 B/op	     802 allocs/op
/recursion-8          	    3716	    290743 ns/op	   86253 B/op	     698 allocs/op

* after making generic (without global return ... but it seems to happen anyway)
/goroutines-8         	    4606	    431865 ns/op	  132857 B/op	     806 allocs/op
/atomic-8             	    5304	    272412 ns/op	  132892 B/op	     802 allocs/op
/recursion-8          	    4040	    433541 ns/op	   81539 B/op	     697 allocs/op

* bring tests into function ...
/goroutines-8         	    6008	    331101 ns/op	  132535 B/op	     805 allocs/op
/atomic-8             	    4696	    281241 ns/op	  131307 B/op	     802 allocs/op

* with pprof running
/goroutines-8         	    4165	    666049 ns/op	  133800 B/op	     806 allocs/op
/atomic-8             	    3942	    288828 ns/op	  135185 B/op	     803 allocs/op

* with http pprof and web server running
/goroutines-8         	    6765	    384056 ns/op	  133527 B/op	     806 allocs/op
/atomic-8             	    3182	    442957 ns/op	  135075 B/op	     802 allocs/op

*/

var (
	globalOutput any
	pwd          string
)

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	go http.ListenAndServe("localhost:8080", nil)

}

func Drain[E any](ch <-chan E) {
	var done bool = false
	for {
		select {
		case i := <-ch:
			_ = i
		default:
			done = true
		}
		if done {
			break
		}
	}
}

func drain2[E any](ch <-chan E) {
	for len(ch) > 0 {
		<-ch
	}
}

func drain3[E any](ch <-chan E) {
L:
	for {
		select {
		case <-ch:
		default:
			break L
		}
	}
}

type namedChan[T any] struct {
	name string
	fun  func(...<-chan T) <-chan T
}

func Merges[T any]() []namedChan[T] {
	return []namedChan[T]{
		{"goroutines", mergeWait[T]},
		{"atomic", mergeAtomic[T]},
		{"recursion", mergeRec[T]},
	}
}

var merges = Merges[int]()

func BenchmarkDrain(b *testing.B) {
	for _, drain := range merges {
		// counter := 0
		b.Run(drain.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				chans := createChannels(100, 100)
				_ = drain.fun(chans...)
				// globalOutput = merge.fun(chans...)
			}
			for i := 0; i < b.N; i++ {
				chans := createChannels(100, 100)
				c := drain.fun(chans...)
				Drain(c)
			}
		})
	}
}

func timestamp() string { return fmt.Sprint(time.Now().Nanosecond()) }

var runtimestamp = timestamp()

func pprofOutPath() string { return filepath.Join(pwd, "pprof", runtimestamp) }

func Heap(testname string) error {
	if true { // pass
		return nil
	}
	filename := pprofOutPath() + "heap.pprof"
	// fmt.Println(filename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	return pprof.WriteHeapProfile(f)
}

func BenchmarkMerge(b *testing.B) {
	// Heap("faketest")
	merges := []namedChan[int]{
		{"goroutines", mergeWait[int]},
		{"atomic", mergeAtomic[int]},
		// {"recursion", mergeRec[int]},
	}
	for _, merge := range merges {
		b.Run(merge.name, func(b *testing.B) {
			defer Heap(merge.name)

			for i := 0; i < b.N; i++ {
				chans := createChannels(100, 100)
				_ = merge.fun(chans...)
				// globalOutput = merge.fun(chans...)
			}
		})
	}
}

// `func makeMergedChannels[T any](cs ...<-chan T) <-chan T {
// 	chans := createChannels(100, 100)
// 	return mergeAtomic(chans...)
// }`
