package atomic

import (
	"sync/atomic"
	"unsafe"
)

// NewValue returns a new atomic.Value that provides an atomic load and store of a consistently typed value.
//
// NewValue performs the initial Store of a new, zero-value
// variable. The example provides the type only and any value
// is ignored.
//
// A Value must not be copied after first use.
func NewValue[T any](example T) atomicvalue[T] {
	v := atomicvalue[T]{atomic.Value{}}
	v.Reset()
	return v
}

// A Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type atomicvalue[T any] struct{ atomic.Value }

// ifaceWords is interface{} internal representation.
// Reference: standard library (1.18) sync/atomic
type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

// New clears the value and stores a new, zero-value object.
func (at atomicvalue[T]) Reset() {
	at.Store(zeroValue[T]())
}

func (at atomicvalue[T]) Load() T {
	if at.Value.Load() == nil {
		at.Reset()
	}
	return at.Value.Load().(T)
}

// Store sets the value of the Value to x.
func (at atomicvalue[T]) Store(val T) {
	at.Value.Store(val)
}

// zeroValue returns a new, zero value instance of type T.
func zeroValue[T any]() T { return *new(T) }

// A Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value[T any] interface {

	// Load returns the value set by the most recent Store.
	// It returns nil if there has been no call to Store for this Value.
	Load() T

	// Store sets the value of the Value to x.
	// All calls to Store for a given Value must use values of the same concrete type.
	// Store of an inconsistent type panics, as does Store(nil).
	Store(val T)

	// Swap stores new into Value and returns the previous value. It returns nil if
	// the Value is empty.
	//
	// All calls to Swap for a given Value must use values of the same concrete
	// type. Swap of an inconsistent type panics, as does Swap(nil).
	Swap(new any) (old T)

	// CompareAndSwap executes the compare-and-swap operation for the Value.
	//
	// All calls to CompareAndSwap for a given Value must use values of the same
	// concrete type. CompareAndSwap of an inconsistent type panics, as does
	// CompareAndSwap(old, nil).
	CompareAndSwap(old, new T) bool
}

// Disable/enable preemption, implemented in runtime.
func runtime_procPin()
func runtime_procUnpin()
