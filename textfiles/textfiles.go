package textfiles

import (
	"bytes"
	"io"
	"os"
	"sort"
	"strings"
)

func GetTextFile(filename string) (*TextFile, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	t := &TextFile{buf: make([]byte, fi.Size()+8)}

	t.buf, err = io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return t, nil
}

type (
	TextFile struct {
		buf []byte // main data buffer
		// *bytes.Buffer            // used for io
		// ioLock sync.Mutex // protects access to buf while bytes.Buffer is used.

		m     map[string]int // word frequency map
		text  string         // used as a cache for string functions
		lines []string       // used as a cache for list of lines for smaller files
		words []string       // used as a cache for list of words for smaller files

		// fi os.FileInfo
	}
)

func (t *TextFile) String() string { return B2S(t.buf) }
func (t *TextFile) Bytes() []byte  { return t.buf }

func (t *TextFile) Lines() []string {
	if t.lines == nil {
		t.lines = strings.Split(t.String(), newline)
	}
	return t.lines
}

func (t *TextFile) Words() []string {
	if t.words == nil {
		t.words = strings.Fields(t.String())
	}
	return t.words
}

func (t *TextFile) sortLines() []string {
	if !sort.StringsAreSorted(t.Lines()) {
		sort.Strings(t.lines)
	}
	return t.lines
}

// LineSet returns a map of line numbers to lines.
func (t *TextFile) LineMap() map[int]string {
	lines := strings.Split(t.String(), newline) // make sure we have the unsorted original lines
	m := make(map[int]string, len(lines))

	for i, line := range lines {
		m[i] = line
	}

	return m
}

// LineSet returns a set of lines from the given file
// that match the given patterns.
//
// If any of the parameters are the empty string, they
// are ignored.
func (t *TextFile) LineSet(prefix, suffix, contains string) []string {
	return nil
}

func (t *TextFile) sortWords() []string {
	if !sort.StringsAreSorted(t.Words()) {
		sort.Strings(t.words)
	}
	return t.words
}

func (t *TextFile) clearWordSort() []string {
	t.words = strings.Fields(t.String())
	return t.words
}

func (t *TextFile) ToValidUTF8() error {
	t.buf = bytes.ToValidUTF8(t.buf, utf8Replacement)
	return nil
}

func (t *TextFile) Less(i, j int) bool { return t.buf[i] < t.buf[j] }
func (t *TextFile) Swap(i, j int)      { t.buf[i], t.buf[j] = t.buf[j], t.buf[i] }

// Reader returns a new io.Reader that may be used until Close() is called.
// func (t *TextFile) Reader() io.Reader {
// 	t.ioLock.Lock()
// 	return bytes.NewBuffer(t.buf)
// }

// Writer returns a new io.Writer that may be used until Close() is called.
// func (t *TextFile) Writer() io.Writer {
// 	t.ioLock.Lock()
// 	return bytes.NewBuffer(t.buf)
// }

// Close syncs the underlying bytes.Buffer and resets it to zero size.
// func (t *TextFile) Close() error {
// 	t.buf = t.Bytes()
// 	t.Buffer.Reset()
// 	t.ioLock.Unlock()
// 	return nil
// }
