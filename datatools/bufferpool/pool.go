package bufferpool

import (
	"bytes"
	"sync"
)

var pool = NewPool[int](func() interface{} { return *new(int) })

// NewPool creates a new sync.Pool implementation
// with the given New function defining the type
// temporary object that will be used.
func NewPool[T any](New func() any) Pool[T] {
	if New == nil {
		New = newPoolFunc[any]()
	}

	p := Pool[T]{
		defaultBufferSize: 0,
		Pool:              sync.Pool{New: New},
	}
	// p.setGetFn()
	p.getFn = p.getNoReset
	p.putFn = p.put
	return p
}

// newBufferPoolFunc creates the New() method for a generic sync.Pool.
func newPoolFunc[T any]() func() T { return func() T { return *new(T) } }

type (

	// Pooler is a generic implementation of Pooler
	Pooler[T any] interface {
		Get() T
		Put(x T)
	}

	// resetter is used to implement the Reset
	// method for clearing pool items of previous data.
	resetter[T any] interface {
		Reset()
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
	Pool[T any] struct {
		getFn             func() T
		putFn             func(v T)
		defaultBufferSize int
		sync.Pool
	}

	// BufferPool is a sync.Pool implementation designed to
	// work exclusively with bytes.Buffer objects.
	BufferPool = Pool[*bytes.Buffer]
)

// Get is the generic sync.Pool Get implementation
// but uses a stored function to add additional
// functionality as needed.
func (b Pool[T]) Get() T        { return b.getFn() }
func (b Pool[T]) getNoReset() T { return b.Pool.Get().(T) }

// Put is a generic sync.Pool Put implementation
// but uses a stored function to add additional
// functionality as needed.
func (b Pool[T]) Put(v T) { b.putFn(v) }
func (b Pool[T]) put(v T) { b.Pool.Put(v) }

// Swimmer takes a pointer to a desired pool object to
// be initialized and returns the function that is
// used to Put it back in the pool. This results in a
// ~10 to 30% performance increase over the standard
// Get() ... defer Put() pattern.
//
// If any part of the process fails, the function
// will return nil.
//
// Used with defer statement to Get the pool object
// and defer the matching Put in one line.
func (b Pool[T]) Swimmer(buf T) (Put func()) {
	buf = b.Get()
	return func() {
		b.Put(buf)
	}
}

// Diver wraps an anonymous function in Get()/Put() calls.
func (b Pool[T]) Diver(fn func(buf T)) {
	buf := b.Get()

	// do your stuff in here to avoid "defer penalty"
	// add return value(s) if you need them
	fn(buf)

	// b.Reset()
	b.Put(buf)
}

func diver[T any](b T) {
	// b.Reset()
}

// Swimmer takes a pointer to a desired pool object to
// be initialized and returns the function that is
// used to Put it back in the pool. This results in a
// ~10 to 30% performance increase over the standard
// Get() ... defer Put() pattern.
//
// If any part of the process fails, the function
// will return nil.
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

// getBlank returns a new empty sample pool object (for testing)
func (b Pool[T]) getBlank() T { return *new(T) }
