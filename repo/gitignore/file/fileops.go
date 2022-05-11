package file

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
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
	create bool,
	extraBuffer float64) (out *GoFile, closer io.Closer, err error) {
	const defaultMinFileSize = 1 << 10
	out = &GoFile{}
	out.fi, err = os.Stat(filename)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			// return on error (except for ErrNotExist)
			return nil, nil, err
		}
		if !create {
			// if file does not exist, create it based on bool parameter 'create'
			return nil, nil, err
		}
	}

	out.filename, err = filepath.Abs(filename)
	if err != nil {
		return nil, nil, err
	}

	out.f, err = os.OpenFile(out.filename, os.O_RDONLY|os.O_CREATE, gofile.NormalMode)
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

// NewFile takes a pointer to a file struct and a filename.
// It checks for existance of the file (or creates as necessary),
// creates a new file structure in f, and returns an io.Closer
// for use with defer in the calling code.
func NewFile(filename string) (*GoFile, io.Closer, error) {
	f, closer, err := fileOpenOrCreate(filename, true, defaultBufferSizeMultiplier)
	if err != nil {
		return nil, fakecloser.NewFromError(err), err
	}
	fmt.Println(f)
	return f, closer, nil
}

// LoadData loads all data from the source file into the
// memfile in the *file struct.
func (f *GoFile) LoadData() error {
	fh, closer, err := NewFile(f.fi.Name())
	if err != nil {
		return err
	}
	defer closer.Close()

	n, err := io.Copy(fh.buf, fh.f)
	if err != nil {
		return err
	}
	if n != fh.fi.Size() {
		return basicfile.ErrShortWrite
	}
	f = fh
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
