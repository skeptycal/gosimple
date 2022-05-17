package file

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/skeptycal/gosimple/os/basicfile"
	"github.com/skeptycal/gosimple/os/gofile"
	"github.com/skeptycal/gosimple/repo/fakecloser"
)

func (f *GoFile) refresh() error {
	return nil
}

func (f *GoFile) ReadAll() error {
	fh, err := os.Open(f.filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	n, err := io.Copy(f.buf, fh)
	if err != nil || n != f.fi.Size() {
		return bufio.ErrBadReadCount
	}
	return nil
}

func fileOpenOrCreate(
	filename string,
	create,
	truncate bool,
	extraBuffer float64) (out *GoFile, closer io.WriteCloser, err error) {
	const defaultMinFileSize = 1 << 10
	// out = &GoFile{}

	out.fi, err = os.Stat(filename)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, nil, err
		}
		if !create {
			return nil, nil, err
		}
	}

	out.filename, err = filepath.Abs(filename)
	if err != nil {
		return nil, nil, err
	}

	fileFlag := os.O_WRONLY

	if create {
		fileFlag |= os.O_CREATE
	}

	if truncate {
		fileFlag |= os.O_TRUNC
	} else {
		fileFlag |= os.O_APPEND
	}

	out.f, err = os.OpenFile(out.filename, fileFlag, gofile.NormalMode)
	if err != nil {
		return nil, nil, err
	}
	// out.isDirty = false // this is the 'useful' default

	// extraBuffer should be positive
	if extraBuffer < 0 {
		extraBuffer *= -1
	}

	// make sure extraBuffer is between 0 and 10
	for extraBuffer > 10 {
		extraBuffer /= 10
	}

	// extraBuffer should be above 1
	// (it is not a percentage increase, but a size multiplier)
	// e.g. 10% larger should be 1.1 as a multiplier
	// this allows inputs of 0.1 and 1.1 to mean the same thing...
	if extraBuffer < 1 {
		extraBuffer += 1
	}

	sizeFloat := float64(out.fi.Size())
	size := int(sizeFloat * extraBuffer)
	if size < defaultMinFileSize {
		size = defaultMinFileSize
	}
	out.buf = bytes.NewBuffer(make([]byte, 0, size))
	out.b = bufio.ReadWriter{
		Reader: bufio.NewReaderSize(out.buf, size),
		Writer: bufio.NewWriterSize(out.buf, size),
	}

	return out, out.f, nil
}

// NewGoFile returns a file object and an io.WriteCloser from a filename.
//
// It checks for existance of the file and creates it as necessary.
// The returned file object is opened for writing only, and
// truncates any existing file. For append mode operations,
// use NewGoFileAppend().
//
// Callers should use the io.WriteCloser
// for standard operations. The returned file object has
// buffering and additional capabilities for advanced operations.
func NewGoFile(filename string) (*GoFile, io.WriteCloser, error) { // (*GoFile, io.Closer, error) {
	f, wc, err := fileOpenOrCreate(filename, true, true, defaultBufferSizeMultiplier)
	if err != nil {
		return nil, nil, err
	}

	return f, wc, nil
}

// NewGoFileAppend returns a file object and an io.WriteCloser
// from a filename.
//
// It checks for existance of the file and creates it as necessary.
// The returned file object is opened for writing only, and
// is in append mode. For truncate mode operations,
// use NewGoFile().
//
// Callers should use the io.WriteCloser
// for standard operations. The returned file object has
// buffering and additional capabilities for advanced operations.
func NewGoFileAppend(filename string) (*GoFile, io.WriteCloser, error) { // (*GoFile, io.Closer, error) {
	f, wc, err := fileOpenOrCreate(filename, true, false, defaultBufferSizeMultiplier)
	if err != nil {
		return nil, nil, err
	}

	return f, wc, nil
}

// GetWriteCloser returns an io.WriteCloser from
// the given filename if the file was opened successfully.
func GetWriteCloser(filename string) (io.WriteCloser, error) {
	_, wc, err := NewGoFile(filename)
	if err != nil {
		return nil, err
	}

	return fakecloser.New(wc)
}

// LoadData loads all data from the source file into the
// memfile in the *file struct.
func (f *GoFile) LoadData() error {
	gf, fh, err := NewGoFile(f.fi.Name())
	if err != nil {
		return err
	}
	defer fh.Close()

	n, err := io.Copy(gf.buf, gf.f)
	if err != nil {
		return err
	}
	if n != gf.fi.Size() {
		return basicfile.ErrShortWrite
	}
	f = gf
	_ = f
	return nil
}

func (f *GoFile) Handle() (*os.File, io.Closer, error) {
	var err error
	f.f, err = os.OpenFile(f.filename, os.O_RDONLY|os.O_CREATE, gofile.NormalMode)

	return f.f, f.f, err
}

// SaveData saves all changes in the memfile buffer to the
//' source file.
func (f *GoFile) SaveData() error {
	// f, closer, err := NewFile(f.fi.Name())
	fh, closer, err := f.Handle()
	if err != nil {
		return err
	}
	defer closer.Close()

	n, err := io.Copy(fh, f.buf)
	if err != nil {
		return err
	}
	if n != f.fi.Size() {
		return io.ErrShortWrite
	}
	return nil
}
