package miniansi

import (
	"io"
)

/* Benchmarks for Fprint implementations

/basic_32(Fprintf)-8         	 5878192	       204.6 ns/op	      48 B/op	       3 allocs/op
/basic_32(Fprintf2)-8        	 7530549	       161.5 ns/op	      48 B/op	       2 allocs/op
/basic_32(Fprintf3)-8        	24017612	        50.11 ns/op	       0 B/op	       0 allocs/op
/basic_33(Fprintf)-8         	 5860226	       203.7 ns/op	      48 B/op	       3 allocs/op
/basic_33(Fprintf2)-8        	 7517073	       158.7 ns/op	      48 B/op	       2 allocs/op
/basic_33(Fprintf3)-8        	24072799	        49.91 ns/op	       0 B/op	       0 allocs/op

/basic_32(Fprintf1)-8         	 5778409	       210.0 ns/op	      48 B/op	       3 allocs/op
/basic_32(Fprintf2)-8         	 7332867	       168.6 ns/op	      48 B/op	       2 allocs/op
/basic_32(Fprintf3)-8         	22178139	        56.09 ns/op	       0 B/op	       0 allocs/op
/basic_33(Fprintf1)-8         	 5061067	       273.1 ns/op	      48 B/op	       3 allocs/op
/basic_33(Fprintf2)-8         	 7154754	       164.3 ns/op	      48 B/op	       2 allocs/op
/basic_33(Fprintf3)-8         	22243744	        54.97 ns/op	       0 B/op	       0 allocs/op
*/

type testStruct struct {
	name   string
	w      io.Writer
	format string
	args   []interface{}
	a      Stringer
}

type FprintFunc = func(w io.Writer, format string, args ...interface{}) (n int, err error)

// type funcList[T AnsiConstraint] struct {
// 	name string
// 	fn FprintFunc
// }

func genTests() []testStruct {
	tests := []testStruct{
		{"basic 32", io.Discard, "%v", []any{"This basic stuff"}, NewAnsi(32)},
		{"basic 33", io.Discard, "%v", []any{"This other stuff"}, NewAnsi(33)},
	}
	return tests
}

// func BenchmarkFprint(b *testing.B) {
// 	for _, bb := range genTests(32) {
// 		var testFuncs = []struct {
// 			name string
// 			fn   FprintFunc
// 		}{
// 			{"Fprintf1", bb.a.Fprintf1},
// 			{"Fprintf2", bb.a.Fprintf2},
// 			{"Fprintf3", bb.a.Fprintf3},
// 		}
// 		for _, ff := range testFuncs {
// 			b.Run(bb.name+"("+ff.name+")", func(b *testing.B) {
// 				for i := 0; i < b.N; i++ {
// 					ff.fn(bb.w, bb.format, bb.args...)
// 				}
// 			})
// 		}
// 	}
// }
