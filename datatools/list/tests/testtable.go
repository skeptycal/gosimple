package tests

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

var errNotImplemented = errors.New("not implemented")

type (
	TestTable[G, W any, E TestTableEntry[G, W], S ~[]E] interface {
		Name() string
		Tests() S
		Run(t *testing.T) error
	}

	testTable[G, W any, E TestTableEntry[G, W], S ~[]E] struct {
		name  string
		tests S
	}
)

func (tbl *testTable[G, W, E, S]) Name() string { return tbl.name }
func (tbl *testTable[G, W, E, S]) Tests() S     { return tbl.tests }
func (tbl *testTable[G, W, E, S]) Run(t *testing.T) error {
	var wrap error
	for i, tt := range tbl.tests {
		t.Run(tbl.Name(), func(t *testing.T) {
			err := tt.Run()
			if err != nil {
				wrap = errWrapper(err, fmt.Sprintf("test %d failed", i))
			}
		})

	}

	return errNotImplemented
}

func errWrapper(err error, msg string) error {
	return errors.Wrap(err, msg)
}
