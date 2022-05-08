package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/skeptycal/gosimple/os/basicfile"
	"github.com/skeptycal/gosimple/os/gofile"
	"github.com/skeptycal/gosimple/repo/fakecloser"
	"github.com/skeptycal/gosimple/types/convert"
)

const (
	inFile                      = "./gilist.txt"
	outFile                     = "../gitignore_gen.go"
	defaultBufferSizeMultiplier = 1.1
)

var (
	b2s = convert.UnsafeBytesToString
	s2b = convert.UnsafeStringToBytes
)

type (
	// used as a memfile working file
	file struct {
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
	extraBuffer float64) (out *file, closer io.Closer, err error) {
	const defaultMinFileSize = 1 << 10
	out = &file{}
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
func NewFile(filename string) (*file, io.Closer, error) {
	f, closer, err := fileOpenOrCreate(filename, true, defaultBufferSizeMultiplier)
	if err != nil {
		return nil, fakecloser.NewFromError(err), err
	}
	fmt.Println(f)
	return f, closer, nil
}

// LoadData loads all data from the source file into the
// memfile in the *file struct.
func (f *file) LoadData() error {
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

func (f *file) String() string        { return f.buf.String() }
func (f *file) Bytes() []byte         { return f.buf.Bytes() }
func (f *file) FileInfo() os.FileInfo { return f.fi }

// StatCli returns the os.FileInfo from filename.
// In the Cli version, any error results in log.Fatal().
func StatCli(filename string) os.FileInfo {
	fiIn, err := os.Stat(inFile)
	if err != nil {
		log.Fatal(err)
	}
	return fiIn
}

// getDataCli gets the bytes from filename and returns
// the string version.
// In the Cli version, any error results in log.Fatal().
func getDataCli(filename string) string {
	data, err := os.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
	}

	return b2s(data)
}

func main() {
	var f *file

	f, closer, err := NewFile(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	// defer NewFile(inFile, f)

	fmt.Println("gofile.PWD(): ", gofile.PWD())
	fmt.Println("file name: ", f.FileInfo().Name())
	fmt.Println("file size: ", f.FileInfo().Size())

	// fiIn := f.FileInfo()
	// fmt.Printf("Input File: %8v %15s %v\n", fiIn.Mode(), fiIn.Name(), fiIn.Size())

	w, err := os.OpenFile("../gitignore_gen.go", os.O_RDWR|os.O_CREATE, gofile.NormalMode)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	data := getDataCli(inFile)

	_ = data

	fmt.Println(data)
	fmt.Println(f)

}
