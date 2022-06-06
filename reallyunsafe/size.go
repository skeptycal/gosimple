package reallyunsafe

import (
	"encoding/binary"
)

// Size returns how many bytes Write would generate to encode the value v, which
// must be a fixed-size value or a slice of fixed-size values, or a pointer to such data.
// If v is neither of these, Size returns -1.
var Size = binary.Size

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

func MakeSlice[T any](x ...T) []T {
	slc := make([]T, len(x))
	copy(slc, x)
	return slc
}

func MakeRepeatSlice[T any](x T, n int) []T {
	slc := make([]T, n)
	for i := 0; i < n; i++ {
		slc[i] = x
	}
	return slc
}

// sampleDataType is a generic set of data of various types used for
// type specific functionality. The interface Sizer allows a list to
// be created with any number of types, something that is difficult
// in traditional table based tests.
type sampleDataType[In any, Out comparable] struct {
	argname string
	data    In
	want    Out
}

type sizer[Out comparable] interface {
	Name() string
	Data() any
	Size() int
	SliceSize(n int) int
	Want() Out
}

func (tdt *sampleDataType[In, Out]) Name() string { return tdt.argname }
func (tdt *sampleDataType[In, Out]) Data() any    { return tdt.data }
func (tdt *sampleDataType[In, Out]) Want() Out    { return tdt.want }
func (tdt *sampleDataType[In, Out]) Size() int    { return intDataSize(tdt.data) }
func (tdt *sampleDataType[In, Out]) SliceSize(n int) int {

	slc := MakeRepeatSlice(tdt.data, n)
	retval := intDataSize(slc)
	slc = nil
	return retval
}

var sampleData = []sizer[int]{
	&sampleDataType[bool, int]{"bool", true, 1},
	&sampleDataType[int8, int]{"int8", int8(42), 1},
	&sampleDataType[uint8, int]{"uint8", uint8(42), 1},
	&sampleDataType[int16, int]{"int16", int16(42), 2},
	&sampleDataType[uint16, int]{"uint16", uint16(42), 2},
	&sampleDataType[int32, int]{"int32", int32(42), 4},
	&sampleDataType[uint32, int]{"uint32", uint32(42), 4},
	&sampleDataType[int64, int]{"int64", int64(42), 8},
	&sampleDataType[uint64, int]{"uint64", uint64(42), 8},
	&sampleDataType[float32, int]{"float32", float32(42), 4},
	&sampleDataType[float64, int]{"float64", float64(42), 8},
}

var sampleDataRaw = []struct {
	argname string
	data    any
	want    int
}{
	{"bool", true, 1},
	{"int8", int8(42), 1},
	{"uint8", uint8(42), 1},
	{"int16", int16(42), 2},
	{"uint16", uint16(42), 2},
	{"int32", int32(42), 4},
	{"uint32", uint32(42), 4},
	{"int64", int64(42), 8},
	{"uint64", uint64(42), 8},
	{"float32", float32(42), 4},
	{"float64", float64(42), 8},
}
