package tests

import (
	"fmt"
	"testing"

	"github.com/skeptycal/gosimple/errorhandling"
)

var errNotImplemented = errorhandling.ErrNotImplemented

// MakeTestRunner returns an interface that can be Run()
// to perform all tests on a single function.
//
// The function takes a name, func, and slice of data
// adds them to a test table.
// The entire set is returned as a TestRunner ready for
// immediate use in testing.
//
// The function name and fn are used for the entire set.
// Each of the items in testdata (In, Want, and WantErr)
// are used as the input, output, and bool to indicate
// if an error is desired in separate tests.
//
// The anonymous test table is not accessible for any other
// functionality.
//
func MakeTestRunner[In any, W comparable](
	name string,
	fn func(In) W,
	testdata []TestDataDetails[In, W],
) TestRunner {
	return NewTestTable[In, W](name).AddSet(name, fn, testdata)
}

func (tbl *TestTable[In, W]) AddSet(name string, fn func(In) W, entries []TestDataDetails[In, W]) TestRunner {
	if len(entries) == 0 {
		return nil
	}
	for _, entry := range entries {
		tt := TestDataType[In, W]{name, fn, entry}
		tbl.Add(tt)
	}
	return tbl
}

func (tbl *TestTable[In, W]) Add(entry TestDataType[In, W]) {
	tbl.Tests = append(tbl.Tests, entry)
}

type (
	TestTable[In any, W comparable] struct {
		Name  string
		Tests []TestDataType[In, W]
	}
)

func (tbl *TestTable[In, W]) Run(t *testing.T, name string) (wrap error) {
	for i, tt := range tbl.Tests {
		if err := tt.Run(t); err != nil {
			if !tt.WantErr {
				wrap = Wrap(err, fmt.Sprintf("test %d failed", i))
			}
		}
	}
	return wrap
}
