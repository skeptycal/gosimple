package file

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/skeptycal/gosimple/os/basicfile"
	"github.com/skeptycal/gosimple/os/gofile"
	"github.com/skeptycal/gosimple/repo/fakecloser"
	. "github.com/skeptycal/gosimple/repo/gitignore/gi_fileparser/cli"
)

const (
	defaultBufferSizeMultiplier = 1.1
	defaultHeadByteLength       = 79
	defaultTailByteLength       = 20
	defaultHeadLineLength       = 5
	defaultTailLineLength       = 5
	newLine                     = "\n"
	defaultStringDelimiter      = ","
)

func init() {
	flag.BoolVar(&Options.FieldsFlag, "fields", false, "print file contents as fields")
	flag.BoolVar(&Options.LinesFlag, "lines", false, "print file contents as lines")

	flag.Parse()
}

type (
	// used as a memfile working GoFile
	GoFile struct {
		filename string      // name provided after Abs() checks and cleanup
		f        *os.File    // file handle
		fi       os.FileInfo // cached file information
		isDirty  bool

		// mu  sync.RWMutex
		buf *bytes.Buffer // memfile
	}
)

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
	f, closer, err := NewFile(f.fi.Name())
	if err != nil {
		return err
	}
	defer closer.Close()

	n, err := io.Copy(f.buf, f.f)
	if err != nil {
		return err
	}
	if n != f.fi.Size() {
		return basicfile.ErrShortWrite
	}

	return nil
}

func (f *GoFile) String() string        { return f.buf.String() }
func (f *GoFile) Bytes() []byte         { return f.buf.Bytes() }
func (f *GoFile) Reader() io.Reader     { return f.f }
func (f *GoFile) Writer() io.Writer     { return f.f }
func (f *GoFile) Closer() io.Closer     { return f.f }
func (f *GoFile) FileInfo() fs.FileInfo { return f.fi }
func (f *GoFile) IsDirty() bool         { return f.isDirty }
func (f *GoFile) Name() string          { return f.FileInfo().Name() }
func (f *GoFile) Mode() fs.FileMode     { return f.FileInfo().Mode() }
func (f *GoFile) Size() int64           { return f.FileInfo().Size() }
func (f *GoFile) IsDir() bool           { return f.FileInfo().IsDir() }
func (f *GoFile) ModTime() time.Time    { return f.FileInfo().ModTime() }

type details struct {
	name  string
	value any
}

// fileDebugDetails returns a list of interesting properties
// that we want to track during development and testing.
func (f *GoFile) fileDebugDetails() []details {
	return []details{
		{"gofile.PWD(): ", gofile.PWD()},
		{"fi name: ", f.FileInfo().Name()},
		{"fi size: ", f.FileInfo().Size()},
		{"fi mode: ", f.FileInfo().Mode()},
		{"fi isdir: ", f.FileInfo().IsDir()},

		{"file type: ", reflect.ValueOf(f.f).Type()},
		{"fi type: ", reflect.ValueOf(f.FileInfo()).Type()},
		{"isDirty: ", f.IsDirty()},
		{"buf head: ", f.HeadBytes(10)},
		{"len(buf): ", f.buf.Len()},
		{"cap(buf): ", f.buf.Cap()},
	}
}

func (f *GoFile) HeadBytes(n int) []byte {
	if n < 1 {
		n = defaultHeadByteLength
	}
	if n > f.buf.Len() {
		return f.buf.Bytes()
	}
	return f.buf.Bytes()[:n]
}

func (f *GoFile) TailBytes(n int) []byte {
	if n < 1 {
		n = defaultTailByteLength
	}
	if n > f.buf.Len() {
		return f.buf.Bytes()
	}
	return f.buf.Bytes()[f.buf.Len()-n:]
}

func (f *GoFile) Lines() []string {
	n := bytes.Count(f.Bytes(), []byte{'\n'})
	retval := make([]string, 0, n+1)

	scanner := bufio.NewScanner(f.Reader())
	for scanner.Scan() {
		retval = append(retval, scanner.Text())
		// log.Printf("metric: %s", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil
	}
	return retval
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

func (f *GoFile) Fields(sep string) []string {
	f.ReadAll()
	if sep == "" {
		sep = defaultStringDelimiter
	}
	list := strings.Fields(f.String())
	fmt.Println("f.String(): ", list)
	s := strings.Join(list, "")
	list = strings.Split(s, sep)
	return list
}

func (f *GoFile) Tail(n int) ([]string, error) {
	v := f.Lines()
	if v == nil {
		return nil, errors.New("error getting file lines")
	}

	if n < 1 {
		n = defaultTailLineLength
	}
	if n > len(v) {
		return v, nil
	}
	return v[len(v)-n:], nil
}

func (f *GoFile) Head(n int) ([]string, error) {
	v := f.Lines()
	if v == nil {
		return nil, errors.New("error getting file lines")
	}

	if n < 1 {
		n = defaultHeadLineLength
	}
	if n > len(v) {
		return v, nil
	}
	return v[:n], nil
}

func (f *GoFile) PrintDebugDetails(detailSet []details) {
	if len(detailSet) == 0 {
		detailSet = f.fileDebugDetails()
	}
	max := 0
	for _, detail := range detailSet {
		if v := len(detail.name); v > max {
			max = v
		}
	}

	max += 2

	format := fmt.Sprintf("%%%ds: %%v\n", max)
	// fmt.Println("format: ", format)

	for _, detail := range detailSet {
		DbEchof(format, detail.name, detail.value)
	}
}
