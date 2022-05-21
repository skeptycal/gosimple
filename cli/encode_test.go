package cli

import (
	"fmt"
	"sync"
	"testing"

	"github.com/skeptycal/gosimple/datatools/random"
)

func TestBasicEncode(t *testing.T) {
	tests := []struct {
		name    string
		input   byte
		want    string
		wantErr bool
	}{
		{"byte", '8', "\033[8m", false},
		{"byte", '2', "\033[2m", false},
		{"byte", '3', "\033[32m", false},
		{"byte", '4', "\033[47m", false},
		{"byte", 'A', "\033[123m", false},
		{"byte", '0', "\033[0m", false},
		{"byte", 'F', "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				got := string(BasicEncode(string(tt.input)))
				if got != tt.want {
					t.Errorf("unexpected ANSI encoding: got %q, want %q", got, tt.want)
				}
			})
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	/* Benchmark results:

	* evaluation:
	- BasicEncode - concat (ansiPrefix + string(in) + ansiSuffix) is SLOW
	- newAnsiColorString - allocates an object and fills value with BasicEncode ... slow again
	- encode is faster, but memory allocation grows quickly
	- simpleEncode is faster and uses less memory

	/encode(1)-8         				19326056	        59.54 ns/op	      16 B/op	       2 allocs/op
	/simpleEncode(1)-8   				37928671	        31.82 ns/op	       8 B/op	       2 allocs/op
	/BasicEncode(1)-8    				13327748	        90.02 ns/op	      16 B/op	       2 allocs/op
	/newAnsiColorString(1)-8         	12719948	        94.35 ns/op	      16 B/op	       2 allocs/op

	/  [...]

	/encode(7)-8                     	  316970	      3769 ns/op	    1024 B/op	     128 allocs/op
	/simpleEncode(7)-8               	  620185	      1926 ns/op	     512 B/op	     128 allocs/op
	/BasicEncode(7)-8                	  190647	      6288 ns/op	     688 B/op	     128 allocs/op
	/newAnsiColorString(7)-8         	  177712	      6728 ns/op	     688 B/op	     128 allocs/op

	* Try BasicEncode with better algorithm
	* Modified newAnsiColorString to use simpleEncode; added unsafeEncode
	- newAnsiColorString is now similar to simpleEncode, as expected
	- unsafeEncode is MUCH faster and has no allocations
	- encode and basicEncode will no longer be tested

	/encode(1)-8         				20037471	     58.89 ns/op	      16 B/op	       2 allocs/op
	/simpleEncode(1)-8   				35872070	     33.14 ns/op	       8 B/op	       2 allocs/op
	/unsafeEncode(1)-8  			   100000000	     10.53 ns/op	       0 B/op	       0 allocs/op
	/basicEncode(1)-8    				13096629	     92.19 ns/op	       8 B/op	       2 allocs/op
	/newAnsiColorString(1)-8         	29760734	     39.76 ns/op	       8 B/op	       2 allocs/op

	/encode(7)-8                     	  320230	      3703 ns/op	    1024 B/op	     128 allocs/op
	/simpleEncode(7)-8               	  459818	      2563 ns/op	     512 B/op	     128 allocs/op
	/unsafeEncode(7)-8               	 1965325	       610 ns/op	       0 B/op	       0 allocs/op
	/basicEncode(7)-8                	  163984	      6216 ns/op	     736 B/op	     128 allocs/op
	/newAnsiColorString(7)-8         	  401575	      2969 ns/op	     512 B/op	     128 allocs/op

	* Modified newAnsiColorString to use unsafeEncode; removed basicEncode and encode from benchmarks
	- the 'unsafe' methods seem to be unbeatable for this particular use case

	/simpleEncode(1)-8         	34375868	        33.17 ns/op	       8 B/op	       2 allocs/op
	/unsafeEncode(1)-8         	125997882	        10.18 ns/op	       0 B/op	       0 allocs/op
	/newAnsiColorString(1)-8   	73412828	        16.32 ns/op	       0 B/op	       0 allocs/op

	/simpleEncode(5)-8         	 2159106	       551.3 ns/op	     128 B/op	      32 allocs/op
	/unsafeEncode(5)-8         	 8011876	       150.5 ns/op	       0 B/op	       0 allocs/op
	/newAnsiColorString(5)-8   	 4619794	       260.0 ns/op	       0 B/op	       0 allocs/op

	/simpleEncode(9)-8         	  118125	     10242 ns/op	    2048 B/op	     512 allocs/op
	/unsafeEncode(9)-8         	  505088	      2382 ns/op	       0 B/op	       0 allocs/op
	/newAnsiColorString(9)-8   	  293617	      4098 ns/op	       0 B/op	       0 allocs/op

	* removed function call from unsafeEncode
	- makes up to ~10% difference with length 2^1 strings, ~2% with length 2^5, ~1.8% with length 2^9
	... around 2% for larger strings

	* ~2%  @ 2^1
	/unsafeEncode(1)-8         	126712488	         9.520 ns/op	       0 B/op	       0 allocs/op
	/unsafeEncode2(1)-8        	124873284	         9.305 ns/op	       0 B/op	       0 allocs/op

	* ~2%  @ 2^5
	/unsafeEncode(5)-8         	 8117953	       147.7 ns/op	       0 B/op	       0 allocs/op
	/unsafeEncode2(5)-8        	 8290264	       144.7 ns/op	       0 B/op	       0 allocs/op

	* ~3% ... @ 2^9
	/unsafeEncode(9)-8         	  504993	      2397 ns/op	       0 B/op	       0 allocs/op
	/unsafeEncode2(9)-8        	  516170	      2325 ns/op	       0 B/op	       0 allocs/op

	* ~2% ... @ 2^13
	/unsafeEncode(13)-8        	   30854	     38061 ns/op	       0 B/op	       0 allocs/op
	/unsafeEncode2(13)-8       	   32211	     37262 ns/op	       0 B/op	       0 allocs/op

	* ~2% ... @ 2^17
	/unsafeEncode(17)-8        	    1981	    605288 ns/op	       0 B/op	       0 allocs/op
	/unsafeEncode2(17)-8       	    2025	    592705 ns/op	       0 B/op	       0 allocs/op


	* Change to string inputs instead of byte
	/encode(1)-8         				32677364	        33.04 ns/op	            8 B/op	       1 allocs/op
	/basicEncode(1)-8    				26807504	        43.04 ns/op	       	    5 B/op	       1 allocs/op
	/simpleEncode(1)-8   				248261799	         4.851 ns/op	        0 B/op	       0 allocs/op
	/newAnsiColorString(1)-8         	100000000	        10.91 ns/op	       		0 B/op	       0 allocs/op
	/unsafeEncode(1)-8               	188197273	         6.384 ns/op	       	0 B/op	       0 allocs/op
	/unsafeEncode2(1)-8              	190628137	         6.297 ns/op	       	0 B/op	       0 allocs/op

	/encode(5)-8                     	19634859	        60.18 ns/op	      	   56 B/op	       2 allocs/op
	/basicEncode(5)-8                	25923874	        45.83 ns/op	      	   48 B/op	       1 allocs/op
	/simpleEncode(5)-8               	245886501	         4.857 ns/op	       	0 B/op	       0 allocs/op
	/newAnsiColorString(5)-8         	100000000	        10.91 ns/op	       		0 B/op	       0 allocs/op
	/unsafeEncode(5)-8               	186311378	         6.466 ns/op	       	0 B/op	       0 allocs/op
	/unsafeEncode2(5)-8              	189064632	         6.379 ns/op	       	0 B/op	       0 allocs/op

	/encode(9)-8                     	10533232	       114.1 ns/op	     	  584 B/op	       2 allocs/op
	/basicEncode(9)-8                	12281014	        98.06 ns/op	     	  576 B/op	       1 allocs/op
	/simpleEncode(9)-8               	248746184	         4.844 ns/op	       	0 B/op	       0 allocs/op
	/newAnsiColorString(9)-8         	100000000	        10.91 ns/op	       		0 B/op	       0 allocs/op
	/unsafeEncode(9)-8               	187336408	         6.413 ns/op	       	0 B/op	       0 allocs/op
	/unsafeEncode2(9)-8              	190801423	         6.340 ns/op	       	0 B/op	       0 allocs/op

	* Added reusable variables and arrays ...
	/unsafeEncode(4)-8         	186454266	         6.339 ns/op	       0 B/op	       0 allocs/op
	/unsafeEncode2(4)-8        	192506998	         6.229 ns/op	       0 B/op	       0 allocs/op
	/blankPtrEncode(4)-8       	181836066	         6.590 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(4)-8          	189714021	         6.325 ns/op	       0 B/op	       0 allocs/op
	/arrayPtrEncode(4)-8       	192136082	         6.253 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(4)-8          	203612048	         6.062 ns/op	       0 B/op	       0 allocs/op

	/unsafeEncode(8)-8         	188934164	         6.348 ns/op	       0 B/op	       0 allocs/op
	/unsafeEncode2(8)-8        	192414396	         6.246 ns/op	       0 B/op	       0 allocs/op
	/blankPtrEncode(8)-8       	181140175	         6.588 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(8)-8          	182102072	         6.335 ns/op	       0 B/op	       0 allocs/op
	/arrayPtrEncode(8)-8       	191928128	         6.274 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(8)-8          	203660974	         5.892 ns/op	       0 B/op	       0 allocs/op

	* size = 2^32 bytes
	/unsafeEncode(32)-8         	189133491	         6.381 ns/op	       0 B/op	       0 allocs/op
	/unsafeEncode2(32)-8        	192350719	         6.322 ns/op	       0 B/op	       0 allocs/op
	/blankPtrEncode(32)-8       	182147818	         6.586 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(32)-8          	189590343	         6.332 ns/op	       0 B/op	       0 allocs/op
	/arrayPtrEncode(32)-8       	192093651	         6.258 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(32)-8          	199930148	         5.907 ns/op	       0 B/op	       0 allocs/op

	* move setup string from bytes outside of loop ;)
	/unsafeEncode2(2)-8         	245204443	         4.654 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(2)-8           	249649645	         4.788 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(2)-8           	286718139	         4.199 ns/op	       0 B/op	       0 allocs/op

	/unsafeEncode2(32)-8        	253468647	         4.730 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(32)-8          	251576634	         4.761 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(32)-8          	287278557	         4.185 ns/op	       0 B/op	       0 allocs/op

	/unsafeEncode2(512)-8       	258524554	         4.637 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(512)-8         	250872980	         4.790 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(512)-8         	287319196	         4.205 ns/op	       0 B/op	       0 allocs/op


	* changed test sample sizing
	/unsafeEncode2(1)-8         	258127544	         4.618 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(1)-8           	252300600	         4.765 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(1)-8           	278681719	         4.179 ns/op	       0 B/op	       0 allocs/op

	/unsafeEncode2(32)-8        	255987802	         4.674 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(32)-8          	253920806	         4.727 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(32)-8          	287370457	         4.176 ns/op	       0 B/op	       0 allocs/op

	/unsafeEncode2(1024)-8      	256837357	         4.763 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(1024)-8        	252236584	         4.758 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(1024)-8        	288356899	         4.151 ns/op	       0 B/op	       0 allocs/op

	/unsafeEncode2(32768)-8     	255659210	         4.688 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(32768)-8       	253322995	         4.739 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(32768)-8       	288352509	         4.170 ns/op	       0 B/op	       0 allocs/op

	* use switch on input length
	/unsafeEncode(32768)-8     	255708174	         4.690 ns/op	       0 B/op	       0 allocs/op
	/blankEncode(32768)-8      	256686138	         4.673 ns/op	       0 B/op	       0 allocs/op
	/arrayEncode(32768)-8      	269549523	         4.455 ns/op	       0 B/op	       0 allocs/op
	/BasicEncode(32768)-8      	329453371	         3.351 ns/op	       0 B/op	       0 allocs/op

	*/

	// args, err := rand.CreateRandomTextSets[string](2, 1, 3)
	// if err != nil {
	// 	b.Fatal(err)
	// }

	args := []string{"8", "32", "123"}

	benchmarks := []struct {
		name string
		fn   func(b string) string
	}{
		// {"encode", fakeEncode},
		// {"basicEncode", basicEncode},
		// {"simpleEncode", blankEncode},
		// {"newAnsiColorString", newAnsiColorString},
		// {"blankPtrEncode", blankPtrEncode},
		// {"arrayPtrEncode", arrayPtrEncode},
		// {"unsafeEncode", unsafeEncode},
		// {"blankEncode", blankEncode},
		{"arrayEncode", arrayEncode},
		{"BasicEncode", BasicEncode},
	}

	for _, arg := range args {
		for _, bb := range benchmarks {
			name := fmt.Sprintf("%v(%d)", bb.name, len(arg))
			// s := UnsafeBytesToString(arg)
			s := arg

			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					globalReturn = bb.fn(s)
				}
			})
		}
	}
}

var globalReturnByte []byte

func getStringArgs(b *testing.B, args []string, n, m int) []string {
	const max = 8

	args1, err := random.CreateTextSets[string](n, m)
	if err != nil {
		b.Fatalf("unable to create random text sets: %v", err)
	}

	return append(args, args1...)
}

type benchmarkRun[In, Out any] struct {
	name string
	// fn   func(b *testing.B)
	fn func(in In) Out
	in In
}

var globalReturnPool = sync.Pool{}

func globalReturnPoolItem[T any]() (item T, deferFunc func(x any)) {
	v := globalReturnPool.Get()
	_ = v
	v = new(T)
	return v.(T), globalReturnPool.Put

}

func globalReturnVar[T any]() T { return *new(T) }

type bm = benchmarkRun[string, []byte]

func multiplexRuns[In, Out any](b *testing.B, count int, done chan bool, ch chan benchmarkRun[In, Out]) chan bool {
	// retval, put := globalReturnPoolItem[Out]()
	// defer put(retval)
	global := globalReturnVar[Out]()
	ret := make(chan bool)

	// wg := &sync.WaitGroup{}
	// wg.Add(routines)
	// starter := make(chan struct{})
	// for i := 0; i < routines; i++ {
	for run := range ch {
		go func(run benchmarkRun[In, Out]) {
			// <-starter
			// run := <-ch
			// b.Run(run.name, run.fn)
			b.Run(run.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					global = run.fn(run.in)
				}
			})
			ret <- true

			// wg.Done()
		}(run)
	}
	for i := 0; i < count; i++ {
		<-ret
	}
	// close(starter)
	// wg.Wait()
	_ = global
	return ret
}

func BenchmarkS2B(b *testing.B) {

	/* Benchmark Results

	* Summary:
	* - for the smallest strings (up to ~32 bytes ... and probably 64) the "unsafe" method is
	*   4 times faster than the "safe" method
	* - even at input sizes of ~1k (seems very common), the "unsafe" methods are 50 times faster.
	* - for 32k strings, the "unsafe" method is 1000 times faster

	* with input strings of sizes 1, 2, and 3 ... unsafe is ~4 times faster
	/safe(1)-8         	52940104	        21.04 ns/op	       8 B/op	       1 allocs/op
	/unsafe(1)-8       	247543324	         4.921 ns/op	       0 B/op	       0 allocs/op
	/safe(2)-8         	56357983	        20.86 ns/op	       8 B/op	       1 allocs/op
	/unsafe(2)-8       	244644535	         4.957 ns/op	       0 B/op	       0 allocs/op
	/safe(3)-8         	56356548	        20.66 ns/op	       8 B/op	       1 allocs/op
	/unsafe(3)-8       	247298966	         4.862 ns/op	       0 B/op	       0 allocs/op


	* as input size increases, the performance drastically descreases and allocations drastically
	* increase for the builtin
	* the "unsafe" pointer swap shows no significant difference until input size rises above
	* approx 32k
	/safe(1)-8         	48336744	        22.05 ns/op	       8 B/op	       1 allocs/op
	/safe(2)-8         	54351210	        27.00 ns/op	       8 B/op	       1 allocs/op
	/safe(3)-8         	35771152	        33.91 ns/op	       8 B/op	       1 allocs/op
	/safe(1)#01-8      	48848498	        21.40 ns/op	       8 B/op	       1 allocs/op
	/safe(32)-8        	48796356	        25.98 ns/op	      32 B/op	       1 allocs/op
	/safe(1024)-8      	 6481688	       190.4 ns/op	    1024 B/op	       1 allocs/op
	/safe(32768)-8     	  320898	      4030 ns/op	   32768 B/op	       1 allocs/op
	/safe(1048576)-8   	   12626	     95902 ns/op	 1048579 B/op	       1 allocs/op
	/unsafe(1)-8       	236882468	         4.866 ns/op	       0 B/op	       0 allocs/op
	/unsafe(2)-8       	236849174	         8.029 ns/op	       0 B/op	       0 allocs/op
	/unsafe(3)-8       	219809071	         4.850 ns/op	       0 B/op	       0 allocs/op
	/unsafe(1)#01-8    	249931875	         4.800 ns/op	       0 B/op	       0 allocs/op
	/unsafe(32)-8      	247573946	         4.849 ns/op	       0 B/op	       0 allocs/op
	/unsafe(1024)-8    	249173250	         4.804 ns/op	       0 B/op	       0 allocs/op
	/unsafe(32768)-8   	246432672	         4.903 ns/op	       0 B/op	       0 allocs/op
	/unsafe(1048576)-8 	238918872	         5.843 ns/op	       0 B/op	       0 allocs/op


	* further testing of "unsafe" function reveals stable performance up to at least
	* string sizes of 1,073,741,824 (~ 1GB strings)
	/unsafe(1)-8         	241697076	         5.047 ns/op	       0 B/op	       0 allocs/op
	/unsafe(2)-8         	239245916	         5.080 ns/op	       0 B/op	       0 allocs/op
	/unsafe(3)-8         	247483934	         4.934 ns/op	       0 B/op	       0 allocs/op
	/unsafe(1)#01-8      	238690602	         4.930 ns/op	       0 B/op	       0 allocs/op
	/unsafe(32)-8        	248208332	         4.838 ns/op	       0 B/op	       0 allocs/op
	/unsafe(1024)-8      	249637830	         4.815 ns/op	       0 B/op	       0 allocs/op
	/unsafe(32768)-8     	247198629	         4.840 ns/op	       0 B/op	       0 allocs/op
	/unsafe(1048576)-8   	249379192	         4.804 ns/op	       0 B/op	       0 allocs/op
	/unsafe(33554432)-8  	246904180	         4.842 ns/op	       0 B/op	       0 allocs/op
	/unsafe(1073741824)-8   250061906	         4.801 ns/op	       0 B/op	       0 allocs/op

	*/

	custom := []string{"8", "32", "123"}
	args := getStringArgs(b, custom, 2, 2)

	benchmarks := []struct {
		name string
		fn   func(in string) []byte
	}{
		{"safe", s2bSafe},
		{"unsafe", s2bUnSafe},
	}

	ch := make(chan bm)
	done := make(chan bool)
	alldone := multiplexRuns(b, len(benchmarks), done, ch)

	for _, bb := range benchmarks {
		for _, arg := range args {
			name := fmt.Sprintf("%v(%d)", bb.name, len(arg))
			ch <- bm{name, bb.fn, arg}
		}
	}
	done <- true
	<-alldone
}
