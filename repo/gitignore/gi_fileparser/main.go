package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/skeptycal/gosimple/os/gofile"
	"github.com/skeptycal/gosimple/repo/fakecloser"
)

const (
	inFile  = "./gilist.txt"
	outFile = "../gitignore_gen.go"
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

		// mu  sync.RWMutex
		buf *bytes.Buffer // memfile
	}
)

func fileOpenOrCreate(
	filename string,
	container *file,
	create bool,
	extraBuffer float64) (io.Closer, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
		if !create {
			return nil, err
		}
		}
		// if file does not exist, create it based on bool parameter 'create'
	}

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

	name, err := filepath.Abs(fi.Name())
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, gofile.NormalMode)
	if err != nil {
		return nil, err
	}

	sizeFloat := float64(fi.Size())
	size := int(sizeFloat * extraBuffer)
	buf := bytes.NewBuffer(make([]byte, 0, size))

	container = &file{
		filename: name,
		f:        f,
		fi:       fi,
		buf:      buf,
	}

	return f, nil
}

// NewFile takes a pointer to a file struct and a filename.
// It checks for existance of the file (or creates as necessary),
// creates a new file structure in f, and returns an io.Closer
// for use with defer in the calling code.
func NewFile(filename string, f *file) io.Closer {
	// var f = &file{}
	closer, err := fileOpenOrCreate(filename, f, true, 1.1)
	if err != nil {
		return fakecloser.New
	}
	return closer
}

func (f *file) LoadData() error {
	r, err := os.Open(f.filename)
	if err != nil {
		return err
	}

	io.Copy(f, r)
	data, err := os.ReadFile(inFile)
	if err != nil {
		return err
	}

	s := b2s(data)
	container = &s

	return nil

}

// var fp FilePrinter = FilePrinter{}

func getFiCli(filename string, container *string) os.FileInfo {
	fiIn, err := os.Stat(inFile)
	if err != nil {
		log.Fatal(err)
	}
	return fiIn
}

func getDataCli(filename string, container *string) error {
	data, err := os.ReadFile(inFile)
	if err != nil {
		return err
	}

	s := b2s(data)
	container = &s

	return nil
}

func main() {

	fp.FileName = ""

	// fmt.Println("$SHELL: ", shell)
	// fmt.Println(" $HOME: ", home)
	// fmt.Println("  $PWD: ", pwd)
	// fmt.Println(" $GOPATH: ", gopath)

	fmt.Println("gofile.PWD(): ", gofile.PWD())

	fiIn, err := os.Stat(inFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input File: %8v %15s %v\n", fiIn.Mode(), fiIn.Name(), fiIn.Size())

	w, err := os.OpenFile("../gitignore_gen.go", os.O_RDWR, gofile.NormalMode)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	b, err := os.ReadFile("../gilist.txt")
	if err != nil {
		log.Fatal(err)
	}

	var data *string

	err = getDataCli(inFile, data)
	if err != nil {
		log.Fatal(err)
	}

	_ = data

	// fmt.Println(data)

}
