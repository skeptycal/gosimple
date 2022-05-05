// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"bytes"
	"testing"

	"github.com/skeptycal/gosimple/datatools/list"
)

/* Benchmark results

* using Builder = strings.Builder as baseline
/BenchmarkBuildString_Builder/1Write_NoGrow-8         		34722598	        30.94 ns/op	      48 B/op	       1 allocs/op
/BenchmarkBuildString_Builder/3Write_NoGrow-8         		10416933	       116.3 ns/op	     336 B/op	       3 allocs/op
/BenchmarkBuildString_Builder/3Write_Grow-8           		25875894	        46.57 ns/op	     112 B/op	       1 allocs/op
/BenchmarkBuildString_ByteBuffer/1Write_NoGrow-8      		22074208	        59.16 ns/op	     112 B/op	       2 allocs/op
/BenchmarkBuildString_ByteBuffer/3Write_NoGrow-8      	 	 8801134	       133.6 ns/op	     352 B/op	       3 allocs/op
/BenchmarkBuildString_ByteBuffer/3Write_Grow-8        		14100974	        83.93 ns/op	     224 B/op	       2 allocs/op

* Replacing strings.Builder with list.Builder[byte,[]byte]
* (and commenting out / changing any references to WriteString, WriteByte, and WriteRune

/BenchmarkBuildString_Builder/1Write_NoGrow-8       	35467194	        33.57 ns/op	      48 B/op	       1 allocs/op
/BenchmarkBuildString_Builder/3Write_NoGrow-8       	 9813405	       123.4 ns/op	     336 B/op	       3 allocs/op
/BenchmarkBuildString_Builder/3Write_Grow-8         	21060220	        56.45 ns/op	     112 B/op	       1 allocs/op
/BenchmarkBuildString_ByteBuffer/1Write_NoGrow-8    	22193366	        53.82 ns/op	     112 B/op	       2 allocs/op
/BenchmarkBuildString_ByteBuffer/3Write_NoGrow-8    	 8939460	       135.0 ns/op	     352 B/op	       3 allocs/op
/BenchmarkBuildString_ByteBuffer/3Write_Grow-8      	14209744	        84.60 ns/op	     224 B/op	       2 allocs/op

* IntBuilder doesn't fare as well ...
/BenchmarkIntBuildString_Builder/1Write_NoGrow-8         	 8963847	       148.0 ns/op	     896 B/op	       1 allocs/op
/BenchmarkIntBuildString_Builder/3Write_NoGrow-8         	 1000000	      1000 ns/op	    6784 B/op	       3 allocs/op
/BenchmarkIntBuildString_Builder/3Write_Grow-8           	 2817231	       418.1 ns/op	    2688 B/op	       1 allocs/op
/BenchmarkIntBuildString_ByteBuffer/1Write_NoGrow-8      	22152124	        53.93 ns/op	     112 B/op	       2 allocs/op
/BenchmarkIntBuildString_ByteBuffer/3Write_NoGrow-8      	 8936857	       134.0 ns/op	     352 B/op	       3 allocs/op
/BenchmarkIntBuildString_ByteBuffer/3Write_Grow-8        	14221147	        84.01 ns/op	     224 B/op	       2 allocs/op

*/

type (
	ListBuilder = list.Builder[byte, []byte]
)

func check(t *testing.T, b *ListBuilder, want string) {
	t.Helper()
	got := b.String()
	if got != want {
		t.Errorf("String: got %#q; want %#q", got, want)
		return
	}
	if n := b.Len(); n != len(got) {
		t.Errorf("Len: got %d; but len(String()) is %d", n, len(got))
	}
	if n := b.Cap(); n < len(got) {
		t.Errorf("Cap: got %d; but len(String()) is %d", n, len(got))
	}
}

func TestBuilder(t *testing.T) {
	var b ListBuilder
	check(t, &b, "")
	// n, err := b.WriteString("hello")
	// if err != nil || n != 5 {
	// 	t.Errorf("WriteString: got %d,%s; want 5,nil", n, err)
	// }
	// check(t, &b, "hello")
	// if err = b.WriteByte(' '); err != nil {
	// 	t.Errorf("WriteByte: %s", err)
	// }
	// check(t, &b, "hello ")
	// n, err = b.WriteString("world")
	// if err != nil || n != 5 {
	// 	t.Errorf("WriteString: got %d,%s; want 5,nil", n, err)
	// }
	// check(t, &b, "hello world")
}

func TestBuilderString(t *testing.T) {
	// changed WriteString's to Write's and string to []byte
	var b ListBuilder
	b.Write([]byte("alpha"))
	check(t, &b, "alpha")
	s1 := b.String()
	b.Write([]byte("beta"))
	check(t, &b, "alphabeta")
	s2 := b.String()
	b.Write([]byte("gamma"))
	check(t, &b, "alphabetagamma")
	s3 := b.String()

	// Check that subsequent operations didn't change the returned strings.
	if want := "alpha"; s1 != want {
		t.Errorf("first String result is now %q; want %q", s1, want)
	}
	if want := "alphabeta"; s2 != want {
		t.Errorf("second String result is now %q; want %q", s2, want)
	}
	if want := "alphabetagamma"; s3 != want {
		t.Errorf("third String result is now %q; want %q", s3, want)
	}
}

func TestBuilderReset(t *testing.T) {
	var b ListBuilder
	check(t, &b, "")
	b.Write([]byte("aaa"))
	s := b.String()
	check(t, &b, "aaa")
	b.Reset()
	check(t, &b, "")

	// Ensure that writing after Reset doesn't alter
	// previously returned strings.
	b.Write([]byte("bbb"))
	check(t, &b, "bbb")
	if want := "aaa"; s != want {
		t.Errorf("previous String result changed after Reset: got %q; want %q", s, want)
	}
}

func TestBuilderGrow(t *testing.T) {
	for _, growLen := range []int{0, 100, 1000, 10000, 100000} {
		p := bytes.Repeat([]byte{'a'}, growLen)
		allocs := testing.AllocsPerRun(100, func() {
			var b ListBuilder
			b.Grow(growLen) // should be only alloc, when growLen > 0
			if b.Cap() < growLen {
				t.Fatalf("growLen=%d: Cap() is lower than growLen", growLen)
			}
			b.Write(p)
			if b.String() != string(p) {
				t.Fatalf("growLen=%d: bad data written after Grow", growLen)
			}
		})
		wantAllocs := 1
		if growLen == 0 {
			wantAllocs = 0
		}
		if g, w := int(allocs), wantAllocs; g != w {
			t.Errorf("growLen=%d: got %d allocs during Write; want %v", growLen, g, w)
		}
	}
}

func TestBuilderWrite2(t *testing.T) {
	const s0 = "hello 世界"
	for _, tt := range []struct {
		name string
		fn   func(b *ListBuilder) (int, error)
		n    int
		want string
	}{
		{
			"Write",
			func(b *ListBuilder) (int, error) { return b.Write([]byte(s0)) },
			len(s0),
			s0,
		},
		// {
		// 	"WriteRune",
		// 	func(b *Builder) (int, error) { return b.WriteRune('a') },
		// 	1,
		// 	"a",
		// },
		// {
		// 	"WriteRuneWide",
		// 	func(b *Builder) (int, error) { return b.WriteRune('世') },
		// 	3,
		// 	"世",
		// },
		// {
		// 	"WriteString",
		// 	func(b *Builder) (int, error) { return b.WriteString(s0) },
		// 	len(s0),
		// 	s0,
		// },
	} {
		t.Run(tt.name, func(t *testing.T) {
			var b ListBuilder
			n, err := tt.fn(&b)
			if err != nil {
				t.Fatalf("first call: got %s", err)
			}
			if n != tt.n {
				t.Errorf("first call: got n=%d; want %d", n, tt.n)
			}
			check(t, &b, tt.want)

			n, err = tt.fn(&b)
			if err != nil {
				t.Fatalf("second call: got %s", err)
			}
			if n != tt.n {
				t.Errorf("second call: got n=%d; want %d", n, tt.n)
			}
			check(t, &b, tt.want+tt.want)
		})
	}
}

// func TestBuilderWriteByte(t *testing.T) {
// 	var b Builder
// 	if err := b.WriteByte('a'); err != nil {
// 		t.Error(err)
// 	}
// 	if err := b.WriteByte(0); err != nil {
// 		t.Error(err)
// 	}
// 	check(t, &b, "a\x00")
// }

func TestBuilderAllocs(t *testing.T) {
	// Issue 23382; verify that copyCheck doesn't force the
	// Builder to escape and be heap allocated.
	n := testing.AllocsPerRun(10000, func() {
		var b ListBuilder
		b.Grow(5)
		b.Write([]byte("abcde"))
		_ = b.String()
	})
	if n != 1 {
		t.Errorf("Builder allocs = %v; want 1", n)
	}
}

func TestBuilderCopyPanic(t *testing.T) {
	tests := []struct {
		name      string
		fn        func()
		wantPanic bool
	}{
		// {
		// 	name:      "String",
		// 	wantPanic: false,
		// 	fn: func() {
		// 		var a Builder
		// 		a.WriteByte('x')
		// 		b := a
		// 		_ = b.String() // appease vet
		// 	},
		// },
		{
			name:      "Len",
			wantPanic: false,
			fn: func() {
				var a ListBuilder
				a.Write([]byte{'x'})
				b := a
				b.Len()
			},
		},
		{
			name:      "Cap",
			wantPanic: false,
			fn: func() {
				var a ListBuilder
				a.Write([]byte{'x'})
				b := a
				b.Cap()
			},
		},
		{
			name:      "Reset",
			wantPanic: false,
			fn: func() {
				var a ListBuilder
				a.Write([]byte{'x'})
				b := a
				b.Reset()
				a.Write([]byte{'x'})
			},
		},
		{
			name:      "Write",
			wantPanic: true,
			fn: func() {
				var a ListBuilder
				a.Write([]byte("x"))
				b := a
				b.Write([]byte("y"))
			},
		},
		{
			name:      "WriteByte",
			wantPanic: true,
			fn: func() {
				var a ListBuilder
				a.Write([]byte{'x'})
				b := a
				b.Write([]byte{'y'})
			},
		},
		// {
		// 	name:      "WriteString",
		// 	wantPanic: true,
		// 	fn: func() {
		// 		var a Builder
		// 		a.WriteString("x")
		// 		b := a
		// 		b.WriteString("y")
		// 	},
		// },
		// 2````````````````
		{
			name:      "Grow",
			wantPanic: true,
			fn: func() {
				var a ListBuilder
				a.Grow(1)
				b := a
				b.Grow(2)
			},
		},
	}
	for _, tt := range tests {
		didPanic := make(chan bool)
		go func() {
			defer func() { didPanic <- recover() != nil }()
			tt.fn()
		}()
		if got := <-didPanic; got != tt.wantPanic {
			t.Errorf("%s: panicked = %v; want %v", tt.name, got, tt.wantPanic)
		}
	}
}

// func TestBuilderWriteInvalidRune(t *testing.T) {
// 	// Invalid runes, including negative ones, should be written as
// 	// utf8.RuneError.
// 	for _, r := range []rune{-1, utf8.MaxRune + 1} {
// 		var b Builder
// 		b.WriteRune(r)
// 		check(t, &b, "\uFFFD")
// 	}
// }

var someBytes = []byte("some bytes sdljlk jsklj3lkjlk djlkjw")

var sinkS string

func benchmarkBuilder(b *testing.B, name string, f func(b *testing.B, numWrite int, grow bool)) {
	b.Run("1Write_NoGrow("+name+")", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 1, false)
	})
	b.Run("3Write_NoGrow("+name+")", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 3, false)
	})
	b.Run("3Write_Grow("+name+")", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 3, true)
	})
}

func BenchmarkBuildString_Builder(b *testing.B) {
	benchmarkBuilder(b, "string", func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf ListBuilder
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

func BenchmarkBuildString_ByteBuffer(b *testing.B) {
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
