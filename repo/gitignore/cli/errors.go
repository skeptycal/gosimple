package cli

import "github.com/pkg/errors"

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
