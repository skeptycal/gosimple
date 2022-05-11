package file

import (
	"bytes"
	"io"
	"os"
)

type RWCPlus interface {
	io.ReadWriteCloser
	io.Seeker
	io.StringWriter
	io.ByteWriter
}

// bufioReader is the interface for bufio.Reader
type bufioReader interface {
	Size() int
	Reset(r io.Reader)
	Peek(n int) ([]byte, error)
	Discard(n int) (discarded int, err error)
	Read(p []byte) (n int, err error)
	ReadByte() (byte, error)
	UnreadByte() error
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
	Buffered() int
	ReadSlice(delim byte) (line []byte, err error)
	ReadLine() (line []byte, isPrefix bool, err error)
	ReadBytes(delim byte) ([]byte, error)
	ReadString(delim byte) (string, error)
	WriteTo(w io.Writer) (n int64, err error)
}

// bufioWriter is the interface for bufio.Writer
type bufioWriter interface {
	Size() int
	Reset(w io.Writer)
	Flush() error
	Available() int
	AvailableBuffer() []byte
	Buffered() int
	Write(p []byte) (nn int, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (size int, err error)
	WriteString(s string) (int, error)
	ReadFrom(r io.Reader) (n int64, err error)
}

type MemFile struct {
	name    string
	isDirty bool
	buf     io.ReadWriteCloser
}

type GoFile2 struct {
	filename string      // name provided after Abs() checks and cleanup
	f        *os.File    // file handle
	fi       os.FileInfo // cached file information
	lines    []string    // cached lines
	isDirty  bool        // flag for updated buffer requiring resetting cached items

	// mu  sync.RWMutex
	buf *bytes.Buffer // memfile
}
