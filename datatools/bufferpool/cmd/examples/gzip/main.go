package main

import (
	"compress/gzip"
	"testing"
)

type fakeWriter struct{}

func (f fakeWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func BenchmarkWriters(b *testing.B) {
	w := fakeWriter{}
	for n := 0; n < b.N; n++ {
		gzip.NewWriterLevel(w, gzip.BestCompression)
	}
}

func BenchmarkPoolWriters(b *testing.B) {
	w := fakeWriter{}
	for n := 0; n < b.N; n++ {
		z := zippers.Get().(*gzip.Writer)
		z.Reset(w)
	}
}
