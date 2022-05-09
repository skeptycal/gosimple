package file

import (
	"io"
	"io/fs"
	"os"
	"time"
)

// Functions for returning cached struct values

func (f *GoFile) String() string { return f.buf.String() }
func (f *GoFile) Bytes() []byte  { return f.buf.Bytes() }

// func (f *GoFile) Reader() io.Reader     { return f.f }
func (f *GoFile) Writer() io.Writer     { return f.f }
func (f *GoFile) Closer() io.Closer     { return f.f }
func (f *GoFile) FileInfo() fs.FileInfo { return f.fi }
func (f *GoFile) IsDirty() bool         { return f.isDirty }
func (f *GoFile) Name() string          { return f.FileInfo().Name() }
func (f *GoFile) Mode() fs.FileMode     { return f.FileInfo().Mode() }
func (f *GoFile) Size() int64           { return f.FileInfo().Size() }
func (f *GoFile) IsDir() bool           { return f.FileInfo().IsDir() }
func (f *GoFile) ModTime() time.Time    { return f.FileInfo().ModTime() }

func (f *GoFile) Reader() (io.ReadCloser, error) {
	fh, err := os.Open(f.Name())
	if err != nil {
		return nil, err
	}
	return fh, nil
}
