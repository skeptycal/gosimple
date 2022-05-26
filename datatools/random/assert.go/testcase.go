package assert

import (
	"fmt"
	"testing"
)

const defaultErrorFormatString = "%v(%v) error: %v(%T) want %v(%T)"

type (
	testRunner Runner

	TestFunc[In any, O Ordered] func(In) O

	AssertFunc[T Ordered] func(got, want T) bool

	testCase[In any, O Ordered, Fn TestFunc[In, O], A AssertFunc[O]] struct {
		name    string
		in      In
		want    O
		fn      Fn
		assert  A
		wantErr bool
	}
)

func NewTestCase[In any, O Ordered, Fn TestFunc[In, O], A AssertFunc[O]](
	name string,
	in In,
	want O,
	fn Fn,
	assert A,
	wantErr bool,
) Runner {
	return &testCase[In, O, Fn, A]{name, in, want, fn, assert, wantErr}
}

func (tc *testCase[In, O, Fn, A]) Run(t *testing.T) (err error) {
	if got := tc.Got(); !tc.assert(got, tc.Want()) != tc.WantErr() {
		err = fmt.Errorf(defaultErrorFormatString, tc.Name(), tc.In(), got, got, tc.Want(), tc.Want())
		t.Errorf(err.Error())
	}
	return
}

func (tc *testCase[In, O, Fn, A]) Name() string  { return tc.name }
func (tc *testCase[In, O, Fn, A]) In() In        { return tc.in }
func (tc *testCase[In, O, Fn, A]) Got() O        { return tc.fn(tc.in) }
func (tc *testCase[In, O, Fn, A]) Want() O       { return tc.want }
func (tc *testCase[In, O, Fn, A]) WantErr() bool { return tc.wantErr }
func (tc *testCase[In, O, Fn, A]) Assert() bool  { return tc.assert(tc.Got(), tc.Want()) }

/// ??
// func (tc *testCase[In, O, Fn, A]) LT(got, want O) bool { return tc.Got() < tc.want }
// func (tc *testCase[In, O, Fn, A]) GT(got, want O) bool { return tc.Got() > tc.want }
// func (tc *testCase[In, O, Fn, A]) NE(got, want O) bool { return tc.Got() != tc.want }
// func (tc *testCase[In, O, Fn, A]) EQ(got, want O) bool { return tc.Got() == tc.want }
