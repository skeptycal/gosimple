package bufferpool

import (
	"bytes"
	"sync"
)

const defaultBufferSize = 16 * 1024

// NewBufferPool creates a new pool of short-lived bytes.Buffer objects preallocated to the given size.
func NewBufferPool[T bytes.Buffer](size int) anyPooler[T] {
	return NewPool[T](newBufferPoolFunc(size))
}

// A pool is a set of temporary objects that may be individually saved and
// retrieved.
//
// Any item stored in the pool may be removed automatically at any time without
// notification. If the pool holds the only reference when this happens, the
// item might be deallocated.
//
// A pool is safe for use by multiple goroutines simultaneously.
//
// pool's purpose is to cache allocated but unused items for later reuse,
// relieving pressure on the garbage collector. That is, it makes it easy to
// build efficient, thread-safe free lists. However, it is not suitable for all
// free lists.
//
// An appropriate use of a pool is to manage a group of temporary items
// silently shared among and potentially reused by concurrent independent
// clients of a package. pool provides a way to amortize allocation overhead
// across many clients.
//
// An example of good use of a pool is in the fmt package, which maintains a
// dynamically-sized store of temporary output buffers. The store scales under
// load (when many goroutines are actively printing) and shrinks when
// quiescent.
//
// On the other hand, a free list maintained as part of a short-lived object is
// not a suitable use for a pool, since the overhead does not amortize well in
// that scenario. It is more efficient to have such objects implement their own
// free list.
//
// A pool must not be copied after first use.
type (
	Pool[T any] struct {
		blank             T
		defaultBufferSize int
		sync.Pool
	}

	bufferPool = Pool[*bytes.Buffer]
	bpObject   *bytes.Buffer
)

// func (b *bufferPool) Swimmer(buf *bytes.Buffer) (Put func()) {
// 	buf = b.Get()
// 	return func() {
// 		b.Put(buf)
// 	}
// }

// func bpGet[T any](b *bufferPool[T]) {
// 	buf := b.Pool.Get().(T)
// 	buf.Reset() // TODO test this...
// 	return buf
// }

// func (b *bufferPool[T]) Get() *bytes.Buffer {
// 	buf := b.Pool.Get().(*bytes.Buffer)
// 	buf.Reset() // TODO test this...
// 	return buf
// }

// func (b *bufferPool[T]) Put(v any) {
// 	b.Pool.Put(v)
// }

// // newBufferPoolFunc creates the New method for a
// // bufferPool sync.Pool.
// func newBufferPoolFunc(size int) func() interface{} {
// 	if size < 1 {
// 		size = defaultBufferSize
// 	}
// 	return func() interface{} { return bytes.NewBuffer(make([]byte, 0, size)) }
// }
