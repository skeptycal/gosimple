package reallyunsafe

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	b2sFn func(b []byte) string
	s2bFn func(string)
)
type ByteStringer interface {
	String() string
	Bytes() []byte
}

func NewBS(s string, b []byte) ByteStringer {
	return &byteStringer{b, s, unsafeBytesToString, unsafeStringToBytes}
}

type byteStringer struct {
	b   []byte
	s   string
	b2s func(b []byte) string
	s2b func(s string) []byte
}

func (b *byteStringer) String() string { return b.b2s(b.b) }
func (b *byteStringer) Bytes() []byte  { return b.s2b(b.s) }

var bsList = []struct {
	name string
	bs   ByteStringer
}{
	{"tes", NewBS("", nil)},
	{"tes", NewBS("", nil)},
}

func Test_unsafeStringToBytes(t *testing.T) {
	testFuncs := []struct {
		name string
		fn   func(string) []byte
	}{
		{"unsafe", unsafeStringToBytes},
		{"builtin", builtinS2B},
	}
	tests := []struct {
		name string
		s    string
		want []byte
	}{
		// TODO: Add test cases.
		{"string", "string", []byte("string")},
	}
	for _, tt := range tests {
		for _, ff := range testFuncs {
			name := fmt.Sprintf("%s(%s): ", ff.name, tt.name)
			t.Run(name, func(t *testing.T) {
				if got := ff.fn(tt.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("unsafeStringToBytes() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
