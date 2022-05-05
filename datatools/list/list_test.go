package list

import (
	"reflect"
	"testing"
)

/* Benchmark Results for slice remove functions

/RemoveOrdered-8         			253738635	         4.817 ns/op	       0 B/op	       0 allocs/op
/RemoveOrderedWithCheck-8         	255259587	         4.696 ns/op	       0 B/op	       0 allocs/op
/RemoveUnOrdered-8                	361103577	         3.325 ns/op	       0 B/op	       0 allocs/op
/removeUnOrderedWithClear-8       	353929098	         3.454 ns/op	       0 B/op	       0 allocs/op

* interestingly, the extra step of the 'clear' added in step 4 seems to increases performance ???

/removeUnOrderedWithClear-8         	341171955	         3.431 ns/op	       0 B/op	       0 allocs/op
/RemoveUnOrdered-8                  	361536258	         3.323 ns/op	       0 B/op	       0 allocs/op
/RemoveOrdered-8                    	255241557	         4.704 ns/op	       0 B/op	       0 allocs/op
/RemoveOrderedWithCheck-8           	255097747	         4.730 ns/op	       0 B/op	       0 allocs/op

* after changing the order, it seems it was a cache or warmup artifact, as was the slower RemoveOrdered in the first set

/RemoveUnOrdered-8                  	360841346	         3.332 ns/op	       0 B/op	       0 allocs/op
/removeUnOrderedInts-8              	360677396	         3.320 ns/op	       0 B/op	       0 allocs/op
/RemoveOrdered-8                    	255700069	         4.691 ns/op	       0 B/op	       0 allocs/op
/removeOrderedInts-8                	255230564	         4.690 ns/op	       0 B/op	       0 allocs/op
/removeOrderedInterface-8           	 35236254	         30.45 ns/op	      48 B/op	       1 allocs/op

* generic function is comparable to typed function; interfaces are slower (in my horrible implementation)

/RemoveUnOrdered-8                  	360841346	         3.332 ns/op	       0 B/op	       0 allocs/op
/RemoveOrdered-8                    	254830700	         4.697 ns/op	       0 B/op	       0 allocs/op

* maintaining order is ~30% slower with the functions tested so far

* The check for sorting takes a long time ... (using standard library sort package / sort Slice)
* maybe a generic sort check and/or sort is better?
/RemoveOrdered-8                    	254830700	         4.697 ns/op	       0 B/op	       0 allocs/op
/RemoveOrderedWithCheck-8           	 23204566	         51.08 ns/op	      24 B/op	       1 allocs/op

*/

var removeBenchmarkFuncs = []struct {
	name string
	fn   func(slice []int, pos int) []int
}{
	{"removeUnOrderedWithClear", removeUnOrderedWithClear[int, []int]},
	{"RemoveUnOrdered", removeUnOrderedGeneric[int, []int]},
	{"removeUnOrderedInts", removeUnOrderedInts},
	{"RemoveOrdered", removeOrderedGeneric[int, []int]},
	{"RemoveOrderedWithCheck", removeOrderedWithCheckGeneric[int, []int]},
	{"removeOrderedInts", removeOrderedInts},
	// {"removeOrderedInterface", removeInterfaceWrapper},
}

func BenchmarkRemove(b *testing.B) {
	tt := intBenchmarkExamples[0]
	for _, ff := range removeBenchmarkFuncs {

		in := tt.in
		b.Run(ff.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ff.fn(in, tt.pos)
			}
		})
	}
}

func TestRemoveOrdered(t *testing.T) {
	var intTests = []struct {
		name string
		in   []int
		pos  int
		want []int
	}{
		{"12345-3", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3, 5}},
		{"12345-1", []int{1, 2, 3, 4, 5}, 1, []int{1, 3, 4, 5}},
		{"12345-2", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
		{"12345-4", []int{1, 2, 3, 4, 5}, 4, []int{1, 2, 3, 4}},
		{"12345-3", []int{4, 3, 5, 1, 2}, 4, []int{4, 3, 5, 1}},
	}
	for _, tt := range intTests {
		in := tt.in
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOrderedGeneric(in, tt.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveOrdered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveUnOrdered(t *testing.T) {
	var intTests = []struct {
		name string
		in   []int
		pos  int
		want []int
	}{
		{"12345-3", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3, 5}},
		{"12345-1", []int{1, 2, 3, 4, 5}, 1, []int{1, 5, 3, 4}},
		{"12345-2", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 5, 4}},
		{"12345-4", []int{1, 2, 3, 4, 5}, 4, []int{1, 2, 3, 4}},
		{"12345-3", []int{4, 3, 5, 1, 2}, 0, []int{2, 3, 5, 1}},
	}
	for _, tt := range intTests {
		in := tt.in
		t.Run(tt.name, func(t *testing.T) {
			got := removeUnOrderedGeneric(in, tt.pos)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveOrdered(%v) = %v, want %v", tt.pos, got, tt.want)
			}
			got2 := removeUnOrderedGeneric(got, 0)
			want := append([]int{got[len(got)-1]}, got[1:len(got)-1]...)
			if !reflect.DeepEqual(got2, want) {
				t.Errorf("RemoveOrdered[0] = %v, want %v", got2, want)
			}
		})
	}
}

var intBenchmarkExamples = []struct {
	name string
	in   []int
	pos  int
	want []int
}{
	{"12345-3", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3, 5}},
	{"12345-1", []int{1, 2, 3, 4, 5}, 1, []int{1, 3, 4, 5}},
	{"12345-2", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
	{"12345-4", []int{1, 2, 3, 4, 5}, 4, []int{1, 2, 3, 4}},
	{"12345-3", []int{4, 3, 5, 1, 2}, 4, []int{4, 3, 5, 1}},
}
