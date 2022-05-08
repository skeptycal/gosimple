package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/skeptycal/gosimple/os/basicfile"
	"github.com/skeptycal/gosimple/os/gofile"
	"github.com/skeptycal/gosimple/repo/fakecloser"
	"github.com/skeptycal/gosimple/types/convert"
)

const (
	defaultInFile               = "./gilist.txt"
	defaultOutFile              = "../gitignore_gen.go"
	defaultBufferSizeMultiplier = 1.1
	defaultHeadByteLength       = 79
	defaultTailByteLength       = 20
	defaultHeadLineLength       = 5
	defaultTailLineLength       = 5
	newLine                     = "\n"
	defaultStringDelimiter      = ","
)

var (
	b2s         = convert.UnsafeBytesToString
	s2b         = convert.UnsafeStringToBytes
	debugFlag   bool
	forceFlag   bool
	verboseFlag bool
	InFile      string
	OutFile     string
)

func init() {
	flag.BoolVar(&debugFlag, "debug", false, "turn on debug mode")
	flag.BoolVar(&forceFlag, "force", false, "force writing to file")
	flag.BoolVar(&verboseFlag, "verbose", false, "turn on verbose mode")
	flag.StringVar(&InFile, "In", defaultInFile, "name of input file")
	flag.StringVar(&OutFile, "Out", defaultOutFile, "name of output file")
	flag.Parse()
}

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
func (f *file) Reader() io.Reader     { return f.f }
func (f *file) Writer() io.Writer     { return f.f }
func (f *file) Closer() io.Closer     { return f.f }
func (f *file) FileInfo() os.FileInfo { return f.fi }
func (f *file) IsDirty() bool         { return f.isDirty }

type details struct {
	name  string
	value any
}

// fileDebugDetails returns a list of interesting properties
// that we want to track during development and testing.
func (f *file) fileDebugDetails() []details {
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

func (f *file) HeadBytes(n int) []byte {
	if n < 1 {
		n = defaultHeadByteLength
	}
	if n > f.buf.Len() {
		return f.buf.Bytes()
	}
	return f.buf.Bytes()[:n]
}

func (f *file) TailBytes(n int) []byte {
	if n < 1 {
		n = defaultTailByteLength
	}
	if n > f.buf.Len() {
		return f.buf.Bytes()
	}
	return f.buf.Bytes()[f.buf.Len()-n:]
}

func (f *file) Lines() []string {
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

func (f *file) ReadAll() error {
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

func (f *file) Fields(sep string) []string {
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

func (f *file) Tail(n int) ([]string, error) {
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

func (f *file) Head(n int) ([]string, error) {
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

func (f *file) printDebugDetails(detailSet []details) {
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
		fmt.Printf(format, detail.name, detail.value)
	}
}

// StatCli returns the os.FileInfo from filename.
// In the Cli version, any error results in log.Fatal().
func StatCli(filename string) os.FileInfo {
	fiIn, err := os.Stat(defaultInFile)
	if err != nil {
		log.Fatal(err)
	}
	return fiIn
}

// getDataCli gets the bytes from filename and returns
// the string version.
// In the Cli version, any error results in log.Fatal().
func getDataCli(filename string) string {
	data, err := os.ReadFile(defaultInFile)
	if err != nil {
		log.Fatal(err)
	}

	return b2s(data)
}

func main() {
	var in *file

	in, closer, err := NewFile(defaultInFile)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	in.ReadAll()
	in.LoadData()

	// data := getDataCli(inFile)

	// _ = data

	if debugFlag {
		in.printDebugDetails(nil)
	}
	// lines := in.Lines()
	// if lines == nil {
	// 	log.Fatal("error getting file lines")
	// }
	// for i, v := range lines {
	// 	fmt.Println(i, ": ", v)
	// }

	fields := in.Fields("")
	if fields == nil {
		log.Fatal("error getting file lines")
	}
	fmt.Println("fields: ", fields)

	// for i, v := range fields {
	// 	fmt.Println(i, ": ", v)
	// }

	out, err := os.OpenFile("../gitignore_gen.go", os.O_RDWR|os.O_CREATE, gofile.NormalMode)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
}
