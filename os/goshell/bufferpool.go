package goshell

import (
	"bytes"
	"sync"
)

const defaultBufferSize = 16 * 1024

var bufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, defaultBufferSize))
	},
}

type Pooler[T any] interface {
	Get() T
	Put(x T)
}

// Swimmer takes a pointer to a desired pool object to
// be initialized and returns the function that is
// used to Put it back in the pool.
func Swimmer[T any](b T) (Put func(any)) {
	// T = *bytes.Buffer
	_ = b
	var ok bool
	if b, ok = bufferPool.Get().(T); ok {
		return bufferPool.Put
	}
	_ = b
	return nil
}
