package testes

import (
	"testing"

	"github.com/skeptycal/gosimple/types/constraints"
)

type Ordered constraints.Ordered

func TRun[T comparable](t *testing.T, name, arg string, got, want T) {
	if !AssertEqual(got, want) && !AssertDeepEqual(got, want) {
		t.Errorf("%v(%v) = %v, want %v\n", name, arg, got, want)
	}
}

func TRunEqual[T comparable](t *testing.T, name, arg string, got, want T) {
	if !AssertEqual(got, want) {
		t.Errorf("%v(%v) = %v, want %v", name, arg, got, want)
	}
}

func TRunDeep[T any](t *testing.T, name, arg string, got, want T) {
	if !AssertDeepEqual(got, want) {
		t.Errorf("%v(%v) = %v, want %v", name, arg, got, want)
	}
}

// t.Run(tt.name, func(t *testing.T) {
// 	if got := NewBufferPool(tt.size); !reflect.DeepEqual(got, tt.want) {
// 		t.Errorf("%v() = %v, want %v", testname, got, tt.want)
// 	}
// })
