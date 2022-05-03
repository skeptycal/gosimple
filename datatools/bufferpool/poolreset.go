package bufferpool

// The following methods are used to manage resetting of
// pool items so that they do not contain old data.
// They are a work in progress and not needed for the
// general generic Pool implementation.

// resetPool is the standard sync.Pool implementation
// but has additional functionality to reset objects
// retrieved from the pool:
//
// Items with their own Reset() method will use it.
// Items without a Reset() method will be reset by
// the Pool's Reset() method, if one is available.
// Any other items are the caller's responsibility.
//
// All reset behavior can be disabled by using
// Enable() and Disable(). It is enabled by default.
type resetPool[T any] struct {
	getFn         func() T
	resetDisabled bool

	Pool[T]
}

// setGetFn is used to choose the correct Get() method
// and should be part of the constructor implementation.
func (b resetPool[T]) setGetFn() {
	a := b.Pool.Get()
	defer b.Pool.Put(a)

	// if items have their own Reset() method...
	if _, ok := any(a).(resetter[T]); ok {
		b.getFn = b.getItemReset
		_ = b.getFn
		return
	}

	// if the pool has a Reset() method...
	if _, ok := any(b).(resetter[T]); ok {
		b.getFn = b.getPoolReset
		_ = b.getFn
		return
	}

	// no reset ... caller is responsible
	b.getFn = b.getNoReset
	_ = b.getFn
}

/// Methods to reset the pool items.

// getNoReset is a Get() method performs no reset. It is the
// caller's responsibility to reset the pool items.
func (b resetPool[T]) getNoReset() T { return b.Pool.Get() }

// Reset is used to reset a pool item after Get if
// the items have no Reset() method of their own.
func (b resetPool[T]) Reset(a T) { _ = a; a = *new(T); _ = a }

// getPoolReset is a Get() method that uses the pool's Reset() method
func (b resetPool[T]) getPoolReset() T {
	a := b.getNoReset()
	b.Reset(a)
	return a
}

// getItemReset is a Get() method that uses the item's Reset() method
func (b resetPool[T]) getItemReset() T {
	a := b.getNoReset()
	if v, ok := any(a).(resetter[T]); ok {
		v.Reset()
	}
	return a
}

/// Methods to enable or disable the automatic reset of the pool items.

func (b resetPool[T]) Enable()  { b.setResetDisabled(false) }
func (b resetPool[T]) Disable() { b.setResetDisabled(true) }
func (b resetPool[T]) setResetDisabled(disable bool) {
	if disable {
		b.resetDisabled = true
		b.getFn = b.getNoReset
	} else {
		b.resetDisabled = false
		b.setGetFn()
	}
	_ = b.getFn
}
