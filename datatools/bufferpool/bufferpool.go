package bufferpool

/// TODO check this out ... inconsistent results
// const defaultBufferSize = 16 * 1024

// // NewBufferPool creates a new pool of short-lived bytes.Buffer objects preallocated to the given size.
// func NewBufferPool[T bytes.Buffer](size int) Pooler[T] {
// 	return NewPool[T](newBufferPoolFunc(size))
// }

// // newBufferPoolFunc creates the New method for a
// // bufferPool sync.Pool.
// func newBufferPoolFunc[T *bytes.Buffer](size int) func() interface{} {
// 	if size < 1 {
// 		size = defaultBufferSize
// 	}
// 	return func() interface{} { return bytes.NewBuffer(make([]byte, 0, size)) }
// }
