package sort

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	slice := GenerateIntSlice(20)
	sorted := NewQuickSortArray("[]int", slice).Sort()

	if !sort.IntsAreSorted(sorted) {
		t.Errorf("Sort(): slice is not sorted: %v", sorted)
	}
}

func BenchmarkSort(b *testing.B) {
	slice := GenerateIntSlice(20)
	arr := NewQuickSortArray("[]int", slice)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		_ = arr.Shuffle()
		b.StartTimer()
		sorted := arr.Sort()
		_ = sorted
	}
}
func BenchmarkGenerateIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := GenerateIntSlice(20)
		_ = slice
	}
}

func BenchmarkShuffle(b *testing.B) {
	slice := GenerateIntSlice(20)
	arr := NewQuickSortArray("[]int", slice)
	for i := 0; i < b.N; i++ {
		arr.Shuffle()
		// rand.Shuffle(len(slice), func(i, j int) {
		// 	slice[i], slice[j] = slice[j], slice[i]
		// })
	}
}
