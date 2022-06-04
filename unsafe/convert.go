package unsafe

import (
	"unsafe"
)

type (
	stringHeader struct {
		Data uintptr
		Len  int
	}

	sliceHeader struct {
		Data uintptr
		Len  int
		Cap  int
	}
)

func unsafeBytesToString(bytes []byte) string {
	header := (*sliceHeader)(unsafe.Pointer(&bytes))

	return *(*string)(unsafe.Pointer(&stringHeader{
		Data: header.Data,
		Len:  header.Len,
	}))
}

func unsafeStringToBytes(s string) []byte {
	header := (*stringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: header.Data,
		Len:  header.Len,
		Cap:  header.Len,
	}))
}

///// The following functions are for benchmark comparisons:

func builtinS2B(s string) []byte { return []byte(s) }
func builtinB2S(b []byte) string { return string(b) }
