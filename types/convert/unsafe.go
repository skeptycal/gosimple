package convert

import (
	"encoding/binary"
	"unsafe"
)

// func BytesToInt64(b *[]byte) int64 {
// 	return atomic.LoadInt64(b)
// }

type StringHeader struct {
	Data uintptr
	Len  int
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// func unsafeSlice(buf []byte, len int) {
// 	s := unsafe.Slice((*byte)(buf), desiredSliceLen)
// 	unsafe.Slice()
// }

// Reference: https://go101.org/article/unsafe.html
func unsafeByteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func unsafeBytesToString2(bytes []byte) string {
	// sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))

	return *(*string)(unsafe.Pointer(&bytes))
}

func UnsafeBytesToString(bytes []byte) string {
	sliceHeader := (*SliceHeader)(unsafe.Pointer(&bytes))

	return *(*string)(unsafe.Pointer(&StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}))
}

func UnsafeStringToBytes(s string) []byte {
	stringHeader := (*StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}))
}

// intDataSize returns the size of the data required to represent the data when encoded.
// It returns zero if the type cannot be implemented by the fast path in Read or Write.
func intDataSize(data any) int {
	switch data := data.(type) {
	case bool, int8, uint8, *bool, *int8, *uint8:
		return 1
	case []bool:
		return len(data)
	case []int8:
		return len(data)
	case []uint8:
		return len(data)
	case int16, uint16, *int16, *uint16:
		return 2
	case []int16:
		return 2 * len(data)
	case []uint16:
		return 2 * len(data)
	case int32, uint32, *int32, *uint32:
		return 4
	case []int32:
		return 4 * len(data)
	case []uint32:
		return 4 * len(data)
	case int64, uint64, *int64, *uint64:
		return 8
	case []int64:
		return 8 * len(data)
	case []uint64:
		return 8 * len(data)
	case float32, *float32:
		return 4
	case float64, *float64:
		return 8
	case []float32:
		return 4 * len(data)
	case []float64:
		return 8 * len(data)
	}
	return 0
}

// Size returns how many bytes Write would generate to encode the value v, which
// must be a fixed-size value or a slice of fixed-size values, or a pointer to such data.
// If v is neither of these, Size returns -1.
var Size = binary.Size
