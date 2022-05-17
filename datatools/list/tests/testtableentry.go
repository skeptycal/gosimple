package tests

import (
	"testing"
)

const fmtErrorf = "%s: got %v, want %v (want error: %v): %v"

type (
	TestRunner interface {
		Name() string
		Run(t *testing.T) error
	}

	TableEntryInfo[G any, W comparable] interface {
		In() []G
		Got() W
		Want() W
		WantErr() bool
	}

	TestTableEntry[G any, W comparable] interface {
		TestRunner
		TableEntryInfo[G, W]
	}

	testTableEntry[G any, W comparable] struct {
		name    string
		fn      func(in ...G) W
		in      []G
		want    W
		wantErr bool
	}
)

func (entry *testTableEntry[G, W]) Name() string  { return entry.name }
func (entry *testTableEntry[G, W]) In() []G       { return entry.in }
func (entry *testTableEntry[G, W]) Want() W       { return entry.want }
func (entry *testTableEntry[G, W]) Got() W        { return *new(W) }
func (entry *testTableEntry[G, W]) WantErr() bool { return entry.wantErr }

func (entry *testTableEntry[G, W]) Run(t *testing.T) error {

	return nil
}
