package compare

import "github.com/skeptycal/gosimple/cli/errorlogger"

var log = errorlogger.New()

// InterfaceEqual protects against panics from doing equality tests on
// two interfaces with non-comparable underlying types.
// adapted from:
//
// /usr/local/go/src/os/exec/exec.go (go 1.15.6)
func InterfaceEqual(a, b interface{}) bool {
	defer func() {
		err := recover()
		log.Errorf("panic recovered: %v", err)
	}()
	return a == b
}
