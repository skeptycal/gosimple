package shpath

import (
	"fmt"
	"strings"
	"testing"

	"github.com/skeptycal/gosimple/datatools/list/tests"
)

var fakeString = ""

var (
	C8  = strings.Repeat("Repeat", 1<<8)
	C16 = strings.Repeat("Repeat", 1<<16)
	C24 = strings.Repeat("Repeat", 1<<24)
	// C48 = strings.Repeat("Repeat", 1<<48) // overflows string cap
)

var stringModifyFuncs = []struct {
	name string
	fn   func(string, string) string
}{
	{"DropDupes", DropDupeSeps},
}

/* Benchmark results
/(cache_loading_..._ignore_this_one)-8         	 5062321	       310.8 ns/op	     224 B/op	       2 allocs/op
/using_bytes_package-8                         	 3359154	       374.0 ns/op	     224 B/op	       2 allocs/op
/for_loop-8                                    	 3393852	       328.8 ns/op	     336 B/op	       3 allocs/op
/bytes.Replace-8                               	 3850560	       315.6 ns/op	     448 B/op	       4 allocs/op
/strings.Replace-8                             	 4760284	       252.3 ns/op	     224 B/op	       2 allocs/op
/aliased-8                                     	 4771462	       253.4 ns/op	     224 B/op	       2 allocs/op
/wrapped-8                                     	 4706868	       255.0 ns/op	     224 B/op	       2 allocs/op
/use_strings.Builder-8                         	 2104315	       589.7 ns/op	     224 B/op	       2 allocs/op

* after streamlining a bit ...
/(cache_loading_..._ignore_this_one)-8         	 4232391	       273.3 ns/op	     224 B/op	       2 allocs/op
/using_bytes_package_(unsafe_conversion)-8     	 3621552	       353.8 ns/op	     224 B/op	       2 allocs/op
/using_bytes_package-8                         	 2904105	       414.3 ns/op	     448 B/op	       4 allocs/op
/strings.Replace-8                             	 4628241	       258.3 ns/op	     224 B/op	       2 allocs/op
/for_loop-8                                    	 2780509	       392.4 ns/op	     352 B/op	       3 allocs/op
/use_strings.Builder_with_loop-8               	 1927165	       610.9 ns/op	     224 B/op	       2 allocs/op
/wrapped_essential_Go_book_(wrapped)-8         	 3853898	       324.1 ns/op	     448 B/op	       4 allocs/op
/aliased-8                                     	 4565563	       259.4 ns/op	     224 B/op	       2 allocs/op

/(cache_loading_..._ignore_this_one)#01-8      	11576916	       100.8 ns/op	      16 B/op	       1 allocs/op
/using_bytes_package_(unsafe_conversion)#01-8  	 9021549	       128.2 ns/op	      16 B/op	       1 allocs/op
/using_bytes_package#01-8                      	 8336128	       143.4 ns/op	      32 B/op	       2 allocs/op
/strings.Replace#01-8                          	12015027	        98.21 ns/op	      16 B/op	       1 allocs/op
/... error stopped further benchmarks ...

/(cache_loading_..._ignore_this_one)-8         	 4410556	       302.3 ns/op	     224 B/op	       2 allocs/op
/using_bytes_package_(unsafe_conversion)-8     	 3232776	       371.9 ns/op	     224 B/op	       2 allocs/op
/using_bytes_package-8                         	 2617440	       437.8 ns/op	     448 B/op	       4 allocs/op
/strings.Replace-8                             	 4622386	       264.0 ns/op	     224 B/op	       2 allocs/op
/for_loop-8                                    	 3410743	       345.5 ns/op	     336 B/op	       3 allocs/op
/use_strings.Builder_with_loop-8               	 2069804	       590.7 ns/op	     224 B/op	       2 allocs/op
/wrapped_essential_Go_book_(wrapped)-8         	 3831040	       315.9 ns/op	     448 B/op	       4 allocs/op
/aliased-8                                     	 4572475	       256.6 ns/op	     224 B/op	       2 allocs/op

/(cache_loading_..._ignore_this_one)#01-8      	11324400	       105.4 ns/op	      16 B/op	       1 allocs/op
/using_bytes_package_(unsafe_conversion)#01-8  	 9021105	       145.1 ns/op	      16 B/op	       1 allocs/op
/using_bytes_package#01-8                      	 7883144	       145.2 ns/op	      32 B/op	       2 allocs/op
/strings.Replace#01-8                          	11732355	       102.5 ns/op	      16 B/op	       1 allocs/op
/for_loop#01-8                                 	17830664	        67.62 ns/op	      32 B/op	       2 allocs/op
/use_strings.Builder_with_loop#01-8            	 5625138	       214.1 ns/op	      16 B/op	       1 allocs/op
/wrapped_essential_Go_book_(wrapped)#01-8      	 7574922	       159.5 ns/op	      48 B/op	       3 allocs/op
/aliased#01-8                                  	11430462	       104.0 ns/op	      16 B/op	       1 allocs/op

/(cache_loading_..._ignore_this_one)#02-8      	 7569651	       167.3 ns/op	      32 B/op	       2 allocs/op
/using_bytes_package_(unsafe_conversion)#02-8  	 5101627	       220.3 ns/op	      32 B/op	       2 allocs/op
/using_bytes_package#02-8                      	 4746549	       245.4 ns/op	      48 B/op	       3 allocs/op
/strings.Replace#02-8                          	 6563895	       158.7 ns/op	      32 B/op	       2 allocs/op
/for_loop#02-8                                 	16517322	        74.73 ns/op	      32 B/op	       2 allocs/op
/use_strings.Builder_with_loop#02-8            	 2990830	       389.4 ns/op	      48 B/op	       3 allocs/op
/wrapped_essential_Go_book_(wrapped)#02-8      	 6420805	       181.9 ns/op	      48 B/op	       3 allocs/op
/aliased#02-8                                  	 7304424	       165.2 ns/op	      32 B/op	       2 allocs/op

/(cache_loading_..._ignore_this_one)#03-8      	 2660059	       424.8 ns/op	       0 B/op	       0 allocs/op
/using_bytes_package_(unsafe_conversion)#03-8  	 2075952	       616.7 ns/op	       0 B/op	       0 allocs/op
/using_bytes_package#03-8                      	  698600	      1795 ns/op	   12288 B/op	       2 allocs/op
/strings.Replace#03-8                          	 2745063	       415.8 ns/op	       0 B/op	       0 allocs/op
/for_loop#03-8                                 	   79858	     14410 ns/op	   18432 B/op	       3 allocs/op
/use_strings.Builder_with_loop#03-8            	 1000000	      1245 ns/op	       0 B/op	       0 allocs/op
/wrapped_essential_Go_book_(wrapped)#03-8      	  381115	      3024 ns/op	   24576 B/op	       4 allocs/op
/aliased#03-8                                  	 2702842	       412.9 ns/op	       0 B/op	       0 allocs/op
*/

func BenchmarkNormalize(b *testing.B) {
	for _, arg := range normalizeInput {
		for _, bb := range stringFuncs {
			name := fmt.Sprintf("%s(%q):", bb.NameFunc, arg.Name)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					fakeString = bb.Fn(arg.In)
				}
			})
		}
	}
}

func BenchmarkNormalizeNew(b *testing.B) {
	tests.MakeBenchmarkRunner("NormalizeNewlines", true, true, stringFuncs, normalizeInput)
}

// []string
var normalizeInput = []tests.BenchmarkInput[string, string]{
	{"mix", "asdlfkn2;leja-9cv8yh	-2piouej4b-	2u9hnasdj;lasdjflkasnvj8q92nn2den\rasdfklw\r\nl;jkqw;cijhpoiqjwd\n\njl-9c8vn-	wd"},
	{"only \\r", "\r asdfsa \r"},
	{"all", "\r\nasdfa\t\n\n\r\r"},
	{"none", strings.Repeat("Repeat", 1024)},
}

// []struct {
// 	name string
// 	fn   func(string) string
// }
var stringFuncs = []tests.BenchmarkFunc[string, string]{
	{"(cache loading ... ignore this one)", NormalizeNL},
	{"normalize (unsafe conversion)", normalize},
	{"normalizeBytesTester", normalizeBytesTester},
	{"normalizeNewlinesString", normalizeNewlinesString},
	{"normalizeNLForLoop", normalizeNLForLoop},
	{"normalizeLoop", normalizeLoop},
	// {"use strings.Builder", normalizeStringsBuilder},
	{"normalizeNewlinesBytesWrapper", normalizeNewlinesBytesWrapper},
	// {"aliased", NormalizeNL},
}

func TestNormalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1234567890", args{"1234567890"}, "1234567890"},
		{"\n1234567890", args{"\n1234567890"}, "\n1234567890"},
		{"\r1234567890", args{"\r1234567890"}, "\n1234567890"},
		{"\r\n1234567890", args{"\r\n1234567890"}, "\n1234567890"},
		{"12345\n67890", args{"12345\n67890"}, "12345\n67890"},
		{"12345\r67890", args{"12345\r67890"}, "12345\n67890"},
		{"12345\r\n67890", args{"12345\r\n67890"}, "12345\n67890"},
		{"fake", args{"fake"}, "fake"},
		{"fa\rke", args{"fa\rke"}, "fa\nke"},
		{"fa\r\nke", args{"fa\r\nke"}, "fa\nke"},
		{"Repeat x 1<<8", args{C8}, C8},
		{"Repeat x 1<<8", args{C16}, C16},
		{"Repeat x 1<<8", args{C24}, C24},
		// {"Repeat x 1<<8", args{C48}, C48},
	}
	for _, ff := range stringFuncs {
		for _, tt := range tests {
			name := ff.NameFunc + "(" + tt.name + ")"
			t.Run(name, func(t *testing.T) {
				if got := ff.Fn(tt.args.s); got != tt.want {
					t.Errorf("%q = %q, want %q", name, got, tt.want)
				}
			})
		}
	}
}

func TestDropDupes(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1234567890", args{"1234567890", "5"}, "1234567890"},
		{"fake", args{"fake", "e"}, "fake"},
		{"newlines", args{"new\n\nline\n", "\n"}, "new\nline\n"},
		{"eee's", args{"slender feet", "e"}, "slender fet"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DropDupeSeps(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("DropDupes() = %v, want %v", got, tt.want)
			}
		})
	}
}
