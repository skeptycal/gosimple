package atomic

import "sync/atomic"

type Counter interface {
	Get() int64
	IncDecer
}

type Incer interface {
	Inc()
}

type Decer interface {
	Dec()
}

type IncDecer interface {
	Incer
	Decer
}

type Adder[T comparable] interface {
	Add(v T)
}

type atomicValues interface {
	~int64 | ~uint64
}

type AInt64 interface {
}

type counter struct {
	value int64
}

func (c *counter) Inc()                       { atomic.AddInt64(&c.value, 1) }
func (c *counter) Dec()                       { atomic.AddInt64(&c.value, -1) }
func (c *counter) Add(v int64)                { atomic.AddInt64(&c.value, v) }
func (c *counter) Sub(v int64)                { atomic.AddInt64(&c.value, -v) }
func (c *counter) Get() int64                 { return atomic.LoadInt64(&c.value) }
func (c *counter) Set(v int64)                { atomic.StoreInt64(&c.value, v) }
func (c *counter) Swap(new int64) (old int64) { return atomic.SwapInt64(&c.value, new) }
func (c *counter) CompareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&c.value, old, new)
}
