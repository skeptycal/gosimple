package cli

import (
	"io"

	"github.com/pkg/errors"
	"github.com/skeptycal/gosimple/cli/errorlogger"
)

// errorlogger
var (
	Log = errorlogger.New()
	log = Log // sometimes my habit for lowercase is too strong ...
	er  = Log.Err
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrNoForceFlag    = errors.New("no -force flag")
)

func errNotImplemented(msg string) error {
	return er(errors.Wrap(ErrNotImplemented, msg))
}

func textErr(err error, msg string) error {
	return er(errors.Wrap(err, msg))
}

func osErr(err error, msg string) error {
	return er(errors.Wrap(err, msg))
}
