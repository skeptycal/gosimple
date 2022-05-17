package file

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// used as a memfile working GoFile
type GoFile struct {
	filename string      // name provided after Abs() checks and cleanup
	f        *os.File    // file handle
	fi       os.FileInfo // cached file information
	lines    []string    // cached lines
	isDirty  bool        // flag for updated buffer requiring resetting cached items
	create   bool        // flag for creating files that do not exist
	truncate bool        // flag for truncating files instead of appending

	// mu  sync.RWMutex

	buf *bytes.Buffer // memfile
	b   bufio.ReadWriter
}

func (f *GoFile) checkDirty() error {
	if f.IsDirty() {
		err := f.SaveData()
		if err != nil {
			return err
		}
		err = f.refresh()
		if err != nil {
			return err
		}
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
