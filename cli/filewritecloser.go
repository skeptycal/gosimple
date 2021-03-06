package cli

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/skeptycal/gosimple/reallyunsafe"
	"github.com/skeptycal/gosimple/repo/fakecloser"
)

const (
	NormalMode os.FileMode = 0644
	DirMode    os.FileMode = 0755
)

// TODO ignoring the error below may be foolish ...

var (
	stdOutCloser, _  = fakecloser.New(os.Stdout)
	discardCloser, _ = fakecloser.New(discard)
	discard          = io.Discard
	s2b              = reallyunsafe.S2B
	b2s              = reallyunsafe.B2S
)

func NewWriteCloserCLI(w any) io.WriteCloser {
	fk, err := fakecloser.New(w)
	if err != nil {
		err = log.Err(err)
		log.Error(err)
		return nil
	}
	return fk
}

// WriteFile writes the string to filename.
// As a precaution, the writer uses os.Stdout unless
// the -force (ForceFlag) CLI option is enabled.
func WriteFile(filename, s string) (n int, err error) {
	if !*ForceFlag {
		return os.Stdout.Write(s2b(s))
	}

	w, err := FileWriteCloser(filename, true)
	if err != nil {
		return 0, err
	}
	defer w.Close()

	return w.Write(s2b(s))
}

// WriteString writes the string to w. It implements the
// io.StringWriter interface.
//
// As a precaution, the writer uses os.Stdout unless
// the -force (ForceFlag) CLI option is enabled.
func WriteString(w io.Writer, s string) (n int, err error) {
	if !*ForceFlag {
		return os.Stdout.Write(s2b(s))
	}
	return w.Write(s2b(s))
}

// FileWriteCloser returns an io.FileWriteCloser from the given
// filename. The file is truncated upon opening if truncate
// is true. Otherwise, the file is opened in append mode.
// If the file does not exist, a new file is created.
//
// As a precaution, the writer uses os.Stdout unless
// the -force (ForceFlag) CLI option is enabled.
func FileWriteCloser(filename string, truncate bool) (io.WriteCloser, error) {
	if !*ForceFlag {
		return stdOutCloser, nil
	}
	fi, err := os.Stat(filename)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, osErr(err, "WriteCloser() failed to stat file")
	}

	var trunc int

	if truncate {
		trunc = os.O_TRUNC
	} else {
		trunc = os.O_APPEND
	}

	fileflag := os.O_WRONLY | os.O_CREATE | trunc
	f, err := os.OpenFile(fi.Name(), fileflag, NormalMode)
	if err != nil {
		return nil, osErr(err, "WriteCloser() failed to open file")
	}
	return f, err
}

// Getenv returns the value of the string while
// replaces ${var} or $var in the string according
// to the values of the current environment variables.
// References to undefined variables are replaced by
// defaultValue.
//  d := Getenv("${HOME}/.config")
//  fmt.Println(d)
//  // /Users/skeptycal/.config
func Getenv(envVarName string, defaultValue string) (retval string) {
	retval = os.ExpandEnv(envVarName)
	if retval == "" {
		retval, _ = os.LookupEnv(envVarName)
	}
	if retval == "" {
		osErr(os.ErrInvalid, fmt.Errorf("Getenv(%s) error: %s (using default value: %s", envVarName, retval, defaultValue).Error())
		return defaultValue
	}
	return
}
