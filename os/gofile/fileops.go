package gofile

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/skeptycal/gosimple/os/basicfile"
)

const (
	smallBufferSize = 64
	maxInt          = int(^uint(0) >> 1)
	minRead         = bytes.MinRead
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return errors.Is(err, os.ErrNotExist)
}

func NotExists(filename string) bool {
	_, err := os.Stat(filename)
	return errors.Is(err, os.ErrNotExist)
}

// Stat returns the os.FileInfo for file if it exists.
//
// It is a convenience wrapper for os.Stat that traps
// and processes errors that may occur using the
// the ErrorLogger package.
//
// If the file does not exist, nil is returned.
// Errors are logged if Err is active.
func Stat(filename string) os.FileInfo {
	fi, err := os.Stat(filename)
	if err != nil {
		Err(basicfile.NewGoFileError("gofile.Stat()", filename, err))
		return nil
	}
	return fi
}

// Mode returns the filemode of file.
// If an error is encountered, it is logged
// with Err (if active) and 0 is returned.
func Mode(filename string) os.FileMode {
	fi, err := os.Stat(filename)
	if err != nil {
		Err(NewGoFileError("gofile.Mode()", filename, err))
		return 0
	}
	return fi.Mode()
}

// StatCheck returns file information (after symlink evaluation)
// using os.Stat(). If the file does not exist, is not a regular file,
// or if the user lacks adequate permissions, an error is returned.
// StatCheck returns file information (after symlink evaluation
// and path cleaning) using os.Stat().
//
// If the file does not exist, is not a regular file,
// or if the user lacks adequate permissions, an error is
// returned.
//
// It is a convenience wrapper for os.Stat that traps
// and processes errors that may occur using the
// the ErrorLogger package.
//
// If the file does not exist, nil is returned.
// Errors are logged if Err is active.
func StatCheck(filename string) (os.FileInfo, error) {

	// EvalSymlinks also calls Abs and Clean as well as
	// checking for existance of the file.
	filename, err := filepath.EvalSymlinks(filename)
	if err != nil {
		return nil, Err(NewGoFileError("gofile.StatCheck()#EvalSymlinks", filename, err))

	}

	fi, err := os.Stat(filename)
	if err != nil {
		return nil, Err(NewGoFileError("gofile.StatCheck()#os.Stat", filename, err))
	}

	//Check 'others' permission
	m := fi.Mode()
	if m&(1<<2) == 0 {
		// err = Wrap(err, "")
		return nil, Err(NewGoFileError("gofile.StatCheck()#insufficient_permissions", filename, ErrPermission))
	}

	if fi.IsDir() {
		return nil, fmt.Errorf("the filename %s refers to a directory", filename)
	}

	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("the filename %s is not a regular file", filename)
	}

	return fi, err
}

// Create creates or truncates the named file and returns an opened file as io.ReadCloser.
//
// If the file already exists, it is truncated. If the file
// does not exist, it is created with mode 0666 (before umask).
// If successful, methods on the returned File can be used
// for I/O; the associated file descriptor has mode O_RDWR. If
// there is an error, it will be of type *PathError.
//
// If the file cannot be created, an error of type *PathError
// is returned.
//
// Errors are logged if gofile.Err is active.
func Create(filename string) io.ReadWriteCloser {

	// OpenFile is the generalized open call; most users will use Open or Create instead. It opens the named file with specified flag (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag is passed, it is created with mode perm (before umask). If successful, methods on the returned File can be used for I/O. If there is an error, it will be of type *PathError.
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		Err(err)
		return nil
	}

	return f
}

// CreateSafe creates the named file and returns an opened file as io.ReadCloser.
//
// If successful, methods on the returned File can be used
// for I/O; the associated file descriptor has mode O_RDWR.
//
// If the file already exists, nil is returned.
// Errors are logged if Err is active.
//
// If the file already exists, of an error occurs, it returns
// nil and an error is sent to Err. If there is an error, it
// will be of type *PathError.
//
func CreateSafe(filename string) io.ReadWriteCloser {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Err(fmt.Errorf("file already exists (%s): %v", filename, err))
		return nil
	}
	return f
}
