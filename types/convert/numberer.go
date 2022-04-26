package convert

import (
	"fmt"
	"unsafe"

	"github.com/skeptycal/gosimple/types/constraints"
)

type RealSet[T Real] interface {
	List() []T
}

type Realer[T Real] interface {
	// ToString(n any) string
	*T
	fmt.Stringer
}

type real[T constraints.Real] struct {
	n T
}

// CastReal recasts a variable of any type in the
// constraint Number to a specific instantiated
// type of Number.
// This may not have the effect you intend ...
func (r real[N2]) CastReal() N2 {
	return *(*N2)(unsafe.Pointer(&r))
}

func (r real[T]) String() string {
	return ToString(r.n)
}

// func (r real[int]) CastInt() int {
// 	return r.CastReal()
// }

func (r real[uint]) CastUInt() uint {
	return r.CastReal()
}

func (r real[float64]) CastFloat64() float64 {
	return r.CastReal()
}

// func makeList[T constraints.Real](args ...T) []Numberer {
// 	list := make([]Numberer, len(args))
// 	for i, arg := range args {
// 		j := ToNumber(arg)
// 		list[i] = number[T]{j}
// 	}
// 	fmt.Println(list)
// 	return list
// }

func makeSlice[T Real](args ...T) (list []T) {
	return append(list, args...)
}

// func toString[S Stringable](n S) string {
// 	return string(n)
// }
func ToString[S Stringable](s S) string {

	return fmt.Sprintf("%v", s)

	/*

		size := int(unsafe.Sizeof(s[0]))

		alloc := length * size

		buf := make([]byte, alloc)

		for _, v := range s {
			c := []byte{}
			for j := 0; j < size; j++ {
				c = append(c, byte(v))
			}
			buf = append(buf, c...)
		}

		return string(buf)
		// return *(*string)(unsafe.Pointer(&n))
	*/
}

// CastReal recasts a variable of any type in the
// constraint Number to a specific instantiated
// type of Number.
// This may not have the effect you intend ...
func CastReal[N1, N2 Real](n N1) N2 {
	return *(*N2)(unsafe.Pointer(&n))
}

// ToByte recasts a variable of any type in the
// constraint Number to byte.
// This may not have the effect you intend ...
func ToByte[N Real](n N) byte {
	return *(*byte)(unsafe.Pointer(&n))
}

// CastFloat64 recasts a Real number to int.
// This may not have the effect you intend ...
func CastFloat64[T Real](n T) float64 {
	// return *(*float64)(unsafe.Pointer(&n))
	return CastReal[T, float64](n)
}
