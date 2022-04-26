package gofile

import (
	"math"

	"github.com/skeptycal/gosimple/types/constraints"
)

const (
	defaultBufSize = 4096
	chunk          = 512.0
)

// chunkMultiple returns a multiple of chunk size closest to but greater than size.
func chunkMultiple[T constraints.Real](size T) T {
	return T(math.Ceil(float64(size)/chunk) * chunk)
}

// InitialCapacity returns the multiple of 'chunk' one more than needed to
// accomodate the given capacity.
func InitialCapacity[T constraints.IntegerNo8Bit](capacity T) T {
	if capacity <= defaultBufSize {
		return defaultBufSize
	}
	return chunkMultiple(capacity)
}
