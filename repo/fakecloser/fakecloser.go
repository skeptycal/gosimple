package fakecloser

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

var (
	stdOutCloser       io.WriteCloser = &fake{os.Stdout, nil}
	exampleBytesBuffer io.WriteCloser = &writeCloserReset{&bytes.Buffer{}}
	errNotImplemented                 = errors.New("feature not implemented")
)

// New returns an io.WriteCloser from any
// valid io.Writer. This streamlines code with
//  defer w.Close()
// statements without adding noticeable overhead.
//
// If a
//  Reset()
// method is available, it is called in place
// of close and Close() error returns nil.
// If a
//  Reset() error
// method is available, it is returned directly
// in place of Close() error.
//
// If neither is available, a noop Close() error
// method is used. This method always returns nil.
// If you wish to return a different error, use
//  WithCloser(w any, err error)
// instead to specify an error.
func New(w any) (io.WriteCloser, error) {
	// if w is already an io.WriteCloser
	if v, ok := w.(io.WriteCloser); ok {
		return v, nil
	}

	// if w is a resetter
	if v, err := WithResetCloser(w); err == nil {
		return v, nil
	}

	if v, err := WithCloser(w, nil); err == nil {
		return v, nil
	}

	return nil, interfaceError("New", w)
}

// WithCloser returns an io.WriteCloser from an
// io.Writer using a noop io.Closer implementation.
// Useful for times when os.Stdout, etc may be used as
// a writer but it is not intended to be closed as
// a file or other object might be. The error returned
// by Close() error is always err (which may be nil).
//
// Note that this completely overrides any existing
// Close() method and cannot be used on objects that
// already implement io.Closer. Any io.WriteCloser
// that is passed into this function will be returned
// unchanged.
//
// If you wish to chain Close() methods or increase
// functionality in some way, use
//  ChainCloser()
func WithCloser(w any, err error) (io.WriteCloser, error) {
	if v, ok := w.(io.WriteCloser); ok {
		return v, nil
	}
	if v, ok := w.(io.Writer); ok {
		// if an error is provided ...
		if err != nil {
			return &fake{v, err}, nil
		}
		// if w implements error ...
		if er1, ok := v.(error); ok {
			return &fake{v, er1}, nil
		}
		// writer with no error ...
		return &fake{v, nil}, nil
	}
	return nil, interfaceError("WithCloser", w)
}

// ChainCloser returns an io.Closer that has
// a slice of io.Closer objects that are called in
// order. If any of these objects return errors,
// the errors are wrapped and returned with a single,
// final io.Closer call.
//
// This makes is simple to add additional deferred
// functionality to any io.Closer.
func ChainCloser(w any, funcs ...func() error) (io.Closer, error) {
	return nil, errNotImplemented
}

// WithResetCloser returns an io.WriteCloser from an
// io.Writer using either Reset() or Reset() error as
// a replacement for Close() error.
//
// Useful for situations like bytes.Buffer, etc where
// an io.WriteCloser may be desired but no Close() method
// is available and calling Reset() or Reset() error would
// be a logical replacement.
//
// Input object must be implement io.Writer and one of the
// two types of Reset() methods. If any method is unavailable,
// nil and an error are returned.
func WithResetCloser(w any) (io.WriteCloser, error) {
	if v, ok := w.(resetter); ok {
		return &writeCloserReset{v}, nil
	}

	if v, ok := w.(errResetter); ok {
		return &writeCloserErrReset{v}, nil
	}
	return nil, interfaceError("WithResetCloser", w)
}

type (
	chainCloser struct {
		chain  []func() error
		closer io.Closer
	}
	// fake implements io.WriteCloser, which includes
	//	Close() error
	// Close() will always return err, which may be nil.
	fake struct {
		io.Writer
		err error
	}

	// resetter implements io.Writer with
	//  Reset()
	// In the structs that implement resetter, Reset()
	// is called and nil is returned in place of
	// Close() error.
	resetter interface {
		io.Writer
		Reset()
	}

	// errResetter implements io.Writer with
	//  Reset() error
	// In the structs that implement resetter,
	// Reset() error is called in place of Close() error.
	errResetter interface {
		io.Writer
		Reset() error
	}

	// writeCloserErrReset implements io.WriteCloser
	// using Reset() error in place of Close() error.
	writeCloserErrReset struct{ errResetter }

	// writeCloserErrReset implements io.WriteCloser
	// using Reset(), always returning nil for Close() error.
	writeCloserReset struct{ resetter }
)

func (c *chainCloser) Close() error {
	var er1 error = nil
	chain := append(c.chain, c.closer.(func() error))
	for i, fn := range c.chain {
		err := fn()
		if err != nil {
			er1 = errors.Wrapf(er1, "ChainCloser(%d) error in %q: %v", i, fn, err)
		}
	}
	err := c.closer.Close()
	if err != nil {
		er1 = errors.Wrapf(er1, "ChainCloser(%d) error in %q: %v", -1, "Close()", err)
	}
	return er1
}

func (f *fake) Close() error                { return f.err }
func (w *writeCloserErrReset) Close() error { return w.Reset() }
func (w *writeCloserReset) Close() error {
	w.Reset()
	return nil
}
func interfaceError(functionName string, v any) error {
	return fmt.Errorf("%s: %v(%T) does not implement the interfaces", functionName, v, v)
}
