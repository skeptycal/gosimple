package bufferpool

import (
	"bytes"
	"sync"
)

type (

	// Pooler implements the two standard methods
	// used by sync.Pool implementations.
	Pooler interface {
		Get() any
		Put(x any)
	}

	typedPooler[T any] interface {
		tGet() T
		tPut(x T)
	}

	Resetter interface {
		Reset()
	}

	anyPooler[T any] interface {
		Get() T
		Put(x T)
	}

	genPool[T any] struct {
		sync.Pool
	}
)

func (b genPool[T]) Get() T {
	return b.Pool.Get().(T)
}

func (b genPool[T]) Put(v T) {
	b.Pool.Put(v)
}

// NewPool creates a new sync.Pool implementation
// with the given New function defining the type
// temporary object that will be used.
func NewPool[T any](New func() any) *genPool[T] { return &genPool[T]{sync.Pool{New: New}} }

// Swimmer takes a pointer to a desired pool object to
// be initialized and returns the function that is
// used to Put it back in the pool.
//
// Used with defer statement to Get the pool object
// and defer the matching Put in one line.
func Swimmer[T any](p sync.Pool, object *T) (Put func()) {
	var ok bool
	if object, ok = p.Get().(*T); ok {
		return func() { p.Put(*object) }
	}
	return nil
}

// newBufferPoolFunc creates the New method for a
// generic sync.Pool.
func newPoolFunc[T any]() func() T {
	return func() T { return *new(T) }
}

func (b *Pool[T]) Swimmer(buf T) (Put func()) {
	buf = b.Pool.Get().(T)
	return func() {
		b.Put(buf)
	}
}

func (b Pool[T]) Get() T {
	buf := b.Pool.Get().(T)
	var bb any = buf
	switch v := bb.(type) {
	case Resetter:
		v.Reset()
	}
	// buf.Reset() // TODO test this...
	return buf
}

func (b Pool[T]) Put(v T) {
	b.Pool.Put(v)
}

// newBufferPoolFunc creates the New method for a
// bufferPool sync.Pool.
func newBufferPoolFunc[T *bytes.Buffer](size int) func() interface{} {
	if size < 1 {
		size = defaultBufferSize
	}
	return func() interface{} { return bytes.NewBuffer(make([]byte, 0, size)) }
}
