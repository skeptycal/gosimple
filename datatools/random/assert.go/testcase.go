package assert

import (
	"fmt"
	"testing"
)

const defaultErrorFormatString = "%v(%v) error: %v(%T) want %v(%T)"

type testCase[In any, O Ordered] struct {
	name    string
	in      In
	want    O
	fn      func(in In) O
	assert  func(got, want O) bool
	wantErr bool
}

func NewTestCase[In any, O Ordered](
	name string,
	in In,
	want O,
	fn func(in In) O,
	assert func(got, want O) bool,
	wantErr bool,
) Runner {
	return &testCase[In, O]{name, in, want, fn, assert, wantErr}
}

type testRunner Runner

func (tc *testCase[In, O]) Run(t *testing.T) (err error) {
	if got := tc.Got(); !tc.assert(got, tc.want) {
		err = fmt.Errorf(defaultErrorFormatString, tc.Name(), tc.In(), got, got, tc.Want(), tc.Want())
		t.Errorf(err.Error())
	}
	return
}

func (tc *testCase[In, O]) Name() string { return tc.name }
func (tc *testCase[In, O]) Got() O       { return tc.fn(tc.in) }
func (tc *testCase[In, O]) In() In       { return tc.in }
func (tc *testCase[In, O]) Want() O      { return tc.want }

/// ??
// func (tc *testCase[In, O]) LT(got, want O) bool { return tc.Got() < tc.want }
// func (tc *testCase[In, O]) GT(got, want O) bool { return tc.Got() > tc.want }
// func (tc *testCase[In, O]) NE(got, want O) bool { return tc.Got() != tc.want }
// func (tc *testCase[In, O]) EQ(got, want O) bool { return tc.Got() == tc.want }
