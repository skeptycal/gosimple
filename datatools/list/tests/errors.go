package tests

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

const (
	fmtErrorf            = "%s: got %v, want %v: %v"
	fmtErrorfWithWantErr = "%s: got %v, want %v (want error: %v): %v"
	fmtErrorfWithTypes   = "%s: got %v(%T), want %v(%T) (want error: %v): %v"
)

var Wrap = errors.Wrap

func tErrorf[In any, W comparable](t *testing.T, name string, got In, want W, err error) error {
	return fmt.Errorf(fmtErrorf, name, got, want, err)
}

func tErrorfWithWantErr[In any, W comparable](t *testing.T, name string, got In, want W, wantErr bool, err error) error {
	return fmt.Errorf(fmtErrorfWithWantErr, name, got, want, wantErr, err)

}

func tErrorWithTypes[In any, W comparable](t *testing.T, name string, got In, want W, wantErr bool, err error) error {
	return fmt.Errorf(fmtErrorfWithTypes, name, got, got, want, want, wantErr, err)
}
