package assert

import (
	"testing"

	"github.com/pkg/errors"
)

// NewTestSet returns a Runner that runs the given slice
// of tests as a group. Any errors returned are wrapped
// and returned as a single error by Run()
func NewTestSet(name string, tc []Runner) Runner {
	return &testSet{name, tc}
}

// Runner runs a test, set, or other functionality, which may
// or may not be generic, and returns any errors encountered.
type Runner interface {
	Name() string
	Run(t *testing.T) error
}

type testSet struct {
	name string
	list []Runner
}

func (ts *testSet) Name() string { return ts.name }

func (ts *testSet) Run(t *testing.T) (wrap error) {
	for _, tt := range ts.list {
		err := tt.Run(t)
		if err != nil {
			wrap = errors.Wrap(err, tt.Name())
		}
	}
	return
}
