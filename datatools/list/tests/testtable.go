package tests

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/skeptycal/gosimple/errorhandling"
)

var errNotImplemented = errorhandling.ErrNotImplemented

type (
	TestTable[G any, W comparable, S ~[]TestTableEntry[G, W]] interface {
		TestRunner
		Tests() S
	}

	testTable[G any, W comparable, S ~[]TestTableEntry[G, W]] struct {
		name  string
		tests S
	}
)

func (tbl *testTable[G, W, S]) Name() string { return tbl.name }
func (tbl *testTable[G, W, S]) Tests() S     { return tbl.tests }
func (tbl *testTable[G, W, S]) Run(t *testing.T) error {
	var wrap error = nil
	for i, tt := range tbl.tests {
		name := tbl.Name() + "(" + tt.Name() + ")"

		t.Run(name, func(t *testing.T) {
			if tt.Got() != tt.Want() != tt.WantErr() {
				err := tErrorf(t, name, tt)
				if err != nil {
					wrap = errWrapper(err, fmt.Sprintf("test %d failed", i))
				}
			}
		})

	}

	return wrap
}

func errWrapper(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func tErrorf[G any, W comparable](t *testing.T, name string, tt TestTableEntry[G, W]) error {
	return fmt.Errorf(fmtErrorf, name, tt.Got(), tt.Want(), tt.WantErr())
}
