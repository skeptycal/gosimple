// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"unsafe"

	"github.com/skeptycal/gosimple/types/constraints"
)

type (
	Ordered constraints.Ordered

	// orderedList[T Ordered, E ~[]T] struct{}
)

// A Builder is used to efficiently build a list using Write methods.
// It minimizes memory copying. The zero value is ready to use.
// Do not copy a non-zero Builder.
type Builder[T Ordered, E ~[]T] struct {
	addr *Builder[T, E] // of receiver, to detect copies by value
	buf  E
}

// noescape hides a pointer from escape analysis. It is the identity function
// but escape analysis doesn't think the output depends on the input.
// noescape is inlined and currently compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime; see issues 23382 and 7921.
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func (b *Builder[T, E]) copyCheck() {
	if b.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing b to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "b.addr = b".
		b.addr = (*Builder[T, E])(noescape(unsafe.Pointer(b)))
	} else if b.addr != b {
		panic("strings: illegal use of non-zero Builder copied by value")
	}
}

// String returns the accumulated string.
func (b *Builder[T, E]) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}

// Len returns the number of accumulated bytes; b.Len() == len(b.String()).
func (b *Builder[T, E]) Len() int { return len(b.buf) }

// Cap returns the capacity of the builder's underlying byte slice. It is the
// total space allocated for the string being built and includes any bytes
// already written.
func (b *Builder[T, E]) Cap() int { return cap(b.buf) }

// Reset resets the Builder to be empty.
func (b *Builder[T, E]) Reset() {
	b.addr = nil
	b.buf = nil
}

// grow copies the buffer to a new, larger buffer so that there are at least n
// bytes of capacity beyond len(b.buf).
func (b *Builder[T, E]) grow(n int) {
	buf := make(E, len(b.buf), 2*cap(b.buf)+n)
	copy(buf, b.buf)
	b.buf = buf
}

// Grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, Grow panics.
func (b *Builder[T, E]) Grow(n int) {
	b.copyCheck()
	if n < 0 {
		panic("strings.Builder.Grow: negative count")
	}
	if cap(b.buf)-len(b.buf) < n {
		b.grow(n)
	}
}

// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
func (b *Builder[T, E]) Write(p E) (int, error) {
	b.copyCheck()
	b.buf = append(b.buf, p...)
	return len(p), nil
}

// // WriteByte appends the byte c to b's buffer.
// // The returned error is always nil.
// func (b *Builder[T, E]) WriteByte(c byte) error {
// 	b.copyCheck()
// 	b.buf = append(b.buf, T(c))
// 	return nil
// }

// // WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
// // It returns the length of r and a nil error.
// func (b *Builder[T, E]) WriteRune(r rune) (int, error) {
// 	b.copyCheck()
// 	// Compare as uint32 to correctly handle negative runes.
// 	if uint32(r) < utf8.RuneSelf {
// 		b.buf = append(b.buf, T(byte(r)))
// 		return 1, nil
// 	}
// 	l := len(b.buf)
// 	if cap(b.buf)-l < utf8.UTFMax {
// 		b.grow(utf8.UTFMax)
// 	}
// 	n := utf8.EncodeRune(b.buf[l:l+utf8.UTFMax], r)
// 	b.buf = b.buf[:l+n]
// 	return n, nil
// }

// // WriteString appends the contents of s to b's buffer.
// // It returns the length of s and a nil error.
// func (b *Builder[T, E]) WriteString(s string) (int, error) {
// 	b.copyCheck()
// 	if _, ok := s.(fmt.Stringer); ok {

// 	}
// 	b.buf = append(b.buf, T(s).String())
// 	return len(s), nil
// }
