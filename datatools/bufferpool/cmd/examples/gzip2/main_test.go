package main

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"fmt"
	"io"
	"math"
	"sync"
	"testing"
)

/* Benchmark Results:

* initial
/BenchmarkWriters-8       	     321	   3703424 ns/op	  817141 B/op	      17 allocs/op
/BenchmarkPoolWriters-8   	     332	   3628577 ns/op	    5613 B/op	       0 allocs/op

*/

type GzipPutWriteCloser[T any] interface {
	io.WriteCloser // *gzip.Writer
	Put(x T)
}

func benchmarkWriter[T any](b *testing.B, w io.Writer, fn func(w io.Writer) GzipPutWriteCloser[T]) {
	d := buf(1024 * 1024)
	rand.Read(d)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		fw := fn(w)
		z, _ := gzip.NewWriterLevel(fw, gzip.BestCompression)
		z.Write(d)
		z.Close()

		// if v, ok := w.(interface{ Put(x any) }); ok {
		fw.Put(z)
	}
}

type gzipPuter[T any] struct {
	*gzip.Writer
}

func (gp *gzipPuter[T]) Put(x T) {
	gp.Put(x)
}

type fakePuter[T any] struct {
	*gzip.Writer
}

func (fp *fakePuter[T]) Put(x T)

func getWriter[T any](w io.Writer) GzipPutWriteCloser[T] {
	z, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
	return &fakePuter[T]{z}
}

func getPoolWriter[T any](w io.Writer) GzipPutWriteCloser[T] {
	z, _ := zippers.Get().(gzipPuter[T])
	if z == nil {
		z, _ = gzipPuter{gzip.NewWriterLevel(w, gzip.BestCompression)}
	} else {
		z.Reset(w)
	}
	return &gzipPuter[T]{z}
}

func BenchmarkWriters(b *testing.B) {
	var writers = append([]bufType{
		{"1 << 8", 1 << 8, bytes.NewBuffer(buf(1 << 8))},
		{"io.Discard", math.MaxInt, io.Discard},
	}, bufs(1, 30, 4)...)

	// normal gzip writers
	for _, bt := range writers {
		name := fmt.Sprintf("%s(%s):", bt.name, "unbuffered")
		b.Run(name, func(b *testing.B) {
			benchmarkWriter(b, bt.w, getWriter)
		})
		name = fmt.Sprintf("%s(%s):", bt.name, "sync.Pool")
		b.Run(name, func(b *testing.B) {
			benchmarkWriter(b, bt.w, getPoolWriter)
		})
	}
}

func GenericPool[T any]() sync.Pool {
	return sync.Pool{
		New: func() any { return *new(T) },
	}
}

var zippers = GenericPool[GzipPutWriteCloser[*gzip.Writer]]()

type bufType struct {
	name string
	size int
	w    io.Writer
}

func buf(size int) []byte {
	return make([]byte, size)
}

func bufs(min, max, step int) []bufType {
	retval := make([]bufType, 0)
	for i := min; i < max; i += step {
		size := 1 << i
		w := bytes.NewBuffer(buf(size))
		retval = append(retval, bufType{fmt.Sprintf("bytes.Buffer(%d)", size), size, w})
	}
	return retval
}
