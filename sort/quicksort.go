package sort

import (
	"fmt"
	"math/rand"

	"github.com/skeptycal/gosimple/types/constraints"
)

type (
	QuickSortArray[E constraints.Ordered, S ~[]E] interface {
		Name() string
		Sort() S
		Reverse() S
		Shuffle() S
		String() string
	}

	quickSortArray[E constraints.Ordered, S ~[]E] struct {
		name    string
		reverse bool
		format  string
		sep     string
		array   S
	}
)

func NewQuickSortArray[E constraints.Ordered, S ~[]E](name string, list S) QuickSortArray[E, S] {
	return &quickSortArray[E, S]{
		name:  name,
		array: list,
	}
}

func (s *quickSortArray[E, S]) Name() string   { return s.name }
func (s *quickSortArray[E, S]) String() string { return fmt.Sprintf(s.format, s.array) }
func (s *quickSortArray[E, S]) Sort() S        { return s.quickSort(0, len(s.array)-1) }
func (s *quickSortArray[E, S]) Reverse() S     { return s.rev() }
func (s *quickSortArray[E, S]) Shuffle() S {
	rand.Shuffle(len(s.array), s.Swap)
	return s.array
}
func (s *quickSortArray[E, S]) Swap(i, j int) { s.array[i], s.array[j] = s.array[j], s.array[i] }

func (s *quickSortArray[E, S]) partition(arr S, low, high int) (S, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func (s *quickSortArray[E, S]) quickSort(low, high int) S {
	if low < high {
		var p int
		s.array, p = s.partition(s.array, low, high)
		s.array = s.quickSort(low, p-1)
		s.array = s.quickSort(p+1, high)
	}
	return s.array
}

func (s *quickSortArray[E, S]) rev() S {
	return nil // TODO: ...
}

// GenerateIntSlice a slice of size filled with random
// integer values between -999 and 999.
func GenerateIntSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

// func GenerateUIntSlice(size, count int) []uint {
// 	slice := make([]uint, size)

// 	for i := 0; i < size; i++ {
// 		for j := 0; j < count; j++ {
// 			slice[i] = uint(xorshift32())
// 		}
// 	}
// 	return slice
// }
