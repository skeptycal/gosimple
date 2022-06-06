package reallyunsafe

import (
	"fmt"
	"testing"

	"github.com/skeptycal/gosimple/testes"
)

func TestIntDataSize(t *testing.T) {
	name := "intDataSize"

	for _, tt := range sampleData {
		// test original data ...
		testes.TRunEqual(t, name, tt.Name(), tt.Size(), tt.Want())

		// test slice of testData also ...
		slcLen := 4
		testes.TRunEqual(t, name, "[]"+tt.Name(), tt.SliceSize(slcLen), tt.Want()*slcLen)

		// test pointer also ...
		ptr := tt.Data()
		testes.TRunEqual(t, name, tt.Name(), intDataSize(ptr), tt.Want())
	}

}

var globalSink any

func BenchmarkSliceMaker(b *testing.B) {
	/*
		* Benchmark examples

		* MakeSlice loads all samples into a slice
		* MakeSliceRepeat repeats one sample data point n times

		/MakeSlice(bool)-8         	15658342	        75.10 ns/op	      56 B/op	       2 allocs/op
		/MakeRepeatSlice(bool)-8   	34929667	        42.73 ns/op	      24 B/op	       1 allocs/op

		/MakeSlice(int8)-8         	15870592	        79.65 ns/op	      56 B/op	       2 allocs/op
		/MakeRepeatSlice(int8)-8   	30487998	        35.46 ns/op	      24 B/op	       1 allocs/op

		/MakeSlice(uint8)-8        	13905129	        76.35 ns/op	      56 B/op	       2 allocs/op
		/MakeRepeatSlice(uint8)-8  	34714981	        34.26 ns/op	      24 B/op	       1 allocs/op

	*/

	max := 3            // number of repeats with different slice lengths
	step := 1           // number of stepts to increase length by each repeat
	limitBenchmark := 3 // number of benchmarks to actually run

	slcData := make([][]any, len(sampleData))

	for num := 1; num < max; num <<= step {
		for j, bb := range sampleData[:limitBenchmark] {
			list := make([]any, num)

			for i := 0; i < num; i++ {
				list[i] = bb.Data()
			}
			slcData[j] = list
		}
	}
	b.ResetTimer()

	for num, dd := range slcData {
		for _, bb := range sampleData[:limitBenchmark] {

			name := fmt.Sprintf("%s(%s)", "MakeSlice", bb.Name())
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					globalSink = MakeSlice(dd...)
				}
			})

			name = fmt.Sprintf("%s(%s)", "MakeRepeatSlice", bb.Name())
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					globalSink = MakeRepeatSlice(dd[0], num)
				}
			})
		}
	}
}
