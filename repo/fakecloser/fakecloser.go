package fakecloser

import (
	"errors"
)

var (
	errNotImplemented = errors.New("feature not implemented")
)

func New() FakeCloser {
	return &fake{errNotImplemented}
}

func NewFromError(err error) FakeCloser {
	return &fake{err}
}

type (
	// FakeCloser implements io.Closer and the builting error
	// interfaces. Close() will always return some error as
	// defined by Error()
	FakeCloser interface {
		Close() error
		Error() string
	}

	fake struct{ err error }
)

func (f *fake) Close() error { return f.err }

func (f *fake) Error() string {
	return f.err.Error()
}
