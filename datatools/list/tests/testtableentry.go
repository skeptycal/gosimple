package tests

import (
	"fmt"
	"testing"
)

type (
	TestTableEntry[In any, Out comparable] interface {
		In() []In
		Want() Out
		Got() Out
		Name() string
		WantErr() bool
		Run() error
	}

	testTableEntry[In any, Out comparable] struct {
		name    string
		fn      func(in ...In) Out
		in      []In
		want    Out
		wantErr bool
	}
)

func (entry *testTableEntry[In, Out]) Name() string  { return entry.name }
func (entry *testTableEntry[In, Out]) WantErr() bool { return entry.wantErr }
func (entry *testTableEntry[In, Out]) In() []In      { return entry.in }
func (entry *testTableEntry[In, Out]) Want() Out     { return entry.want }
func (entry *testTableEntry[In, Out]) Got() Out {

	return *new(Out)
}

const fmtErrorf = "%s(%s): got %v, want %v (want error: %v"

func (entry *testTableEntry[In, Out]) Run(t *testing.T) error {
	if entry.Got() != entry.Want() != entry.WantErr() {
		err := fmt.Errorf(fmtErrorf, entry.Name(), entry.In(), entry.Got(), entry.Want(), entry.WantErr())
		t.Error(err.Error())
		return err
	}
	return nil
}
