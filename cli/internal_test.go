package cli

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"
	"unsafe"
)

var n int
var arr []byte
var nums []uint32

/* Benchmark results ... this is just ridiculous.

/BenchmarkCast-8                  1000000000             1.041 ns/op	   0 B/op	       0 allocs/op
/BenchmarkReaderNoAlloc-8       	   25947	     46176 ns/op	       0 B/op	       0 allocs/op
/BenchmarkReaderManual-8        	   24493	     49017 ns/op	   32768 B/op	       1 allocs/op
/BenchmarkReaderAll-8           	   12321	     97293 ns/op	   65608 B/op	       4 allocs/op
/BenchmarkReadBytesNoAlloc-8    	   81763	     14693 ns/op	       0 B/op	       0 allocs/op
/BenchmarkReadBytesOneAlloc-8   	   65926	     18215 ns/op	       0 B/op	       0 allocs/op
/BenchmarkReadBytes-8           	   59835	     20072 ns/op	   32768 B/op	       1 allocs/op

*/

// Reference: https://kokes.github.io/blog/2019/03/19/deserialising-ints-from-bytes.html
func init() {
	n = 1024 * 8
	arr = make([]byte, n*4)
	nums = make([]uint32, n)
	for j := 0; j < n; j++ {
		rnd := rand.Int31()
		nums[j] = uint32(rnd)
		binary.LittleEndian.PutUint32(arr[4*j:4*(j+1)], uint32(rnd))
	}
}

func BenchmarkCast(b *testing.B) {
	for it := 0; it < b.N; it++ {
		narr := *(*[]uint32)(unsafe.Pointer(&arr))
		narr = narr[:n]
		_ = narr
	}
}

func BenchmarkReaderNoAlloc(b *testing.B) {
	for it := 0; it < b.N; it++ {
		rd := bytes.NewReader(arr)
		buf := make([]byte, 4)
		var sum uint64
		for j := 0; j < n; j++ {
			n, err := rd.Read(buf)
			if n != 4 {
				log.Fatalf("not enough bytes read, %d", n)
			}
			if err != nil {
				log.Fatal(err)
			}
			sum += uint64(binary.LittleEndian.Uint32(buf))
		}
	}
}
func BenchmarkReaderManual(b *testing.B) {
	for it := 0; it < b.N; it++ {
		bnums := make([]uint32, n)
		rd := bytes.NewReader(arr)
		buf := make([]byte, 4)
		for j := 0; j < n; j++ {
			n, err := rd.Read(buf)
			if n != 4 {
				log.Fatalf("not enough bytes read, got %d, expected 4", n)
			}
			if err != nil {
				log.Fatal(err)
			}
			bnums[j] = binary.LittleEndian.Uint32(buf)
		}
	}
}
func BenchmarkReaderAll(b *testing.B) {
	for it := 0; it < b.N; it++ {
		bnums := make([]uint32, n)
		err := binary.Read(bytes.NewReader(arr), binary.LittleEndian, &bnums)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func BenchmarkReadBytesNoAlloc(b *testing.B) {
	for it := 0; it < b.N; it++ {
		var sum uint64
		for j := 0; j < n; j++ {
			sum += uint64(binary.LittleEndian.Uint32(arr[4*j : 4*(j+1)]))
		}
	}
}
func BenchmarkReadBytesOneAlloc(b *testing.B) {
	bnums := make([]uint32, n)
	for it := 0; it < b.N; it++ {
		for j := 0; j < n; j++ {
			bnums[j] = binary.LittleEndian.Uint32(arr[4*j : 4*(j+1)])
		}
	}
}
func BenchmarkReadBytes(b *testing.B) {
	for it := 0; it < b.N; it++ {
		bnums := make([]uint32, n)
		for j := 0; j < n; j++ {
			bnums[j] = binary.LittleEndian.Uint32(arr[4*j : 4*(j+1)])
		}
	}
}
