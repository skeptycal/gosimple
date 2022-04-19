package io

import (
	"compress/gzip"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

const (
	longstring = "lasjdkf;ijpil2nep	2kj nv-0u9ashc90uvdhwqn3ekjfna;skjdf;ljansdfnpuvdjpq"
	filler1    = `Now is the time for all good men to come to the aid of their country.`
	filler2    = `The quick brown fox jumps over the lazy dog`
)

var (
	DefaultHeadLength int       = 5
	sampleReader      io.Reader = strings.NewReader(longstring)
	ReplacementChar   rune      = unicode.ReplacementChar
	u8                          = strings.ToValidUTF8(longstring, ".")
)

type inc struct {
	done  bool
	max   int
	count int
	incFn func(n int)
}

func (i *inc) Inc(n int) {
	i.incFn(n)
}

func (i *inc) inc(n int) {
	i.count += n
	if i.count+n > i.max {
		i.count = i.max
		i.incFn = i.noInc
	}
}

func (i *inc) noInc(n int) {
	// noop
}

// ReadCloser returns an io.ReadCloser from a file.
// The caller is responsible for closing the file
// when finished with it.
func ReadCloser(filename string) (io.ReadCloser, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, errors.Wrapf(err, "open file error: %v", err)
	}
	return f, nil
}

// CloseIf checks a reader to see if it also
// has a close method. If so, it calls r.Close().
// Useful for places where any io.Reader can be used,
// but a file (or other ReadCloser that requires
// closing) may be substituted.
func CloseIf(r io.Reader) error {
	// TODO: if a file (or other ReadCloser) was used,
	// the Reader may also have a Close method ...
	// probably should call that ...
	if v, ok := r.(io.Closer); ok {
		return v.Close()
	}
	return nil
}

// CloseWriterIf checks a writer to see if it also
// has a close method. If so, it calls r.Close().
// Useful for places where any io.Writer can be used,
// but a file (or other WriteCloser that requires
// closing) may be substituted.
func CloseWriterIf(w io.Writer) error {
	// TODO: if a file (or other WriteCloser) was used,
	// the Writer may also have a Close method ...
	// probably should call that ...
	if v, ok := w.(io.Closer); ok {
		return v.Close()
	}
	return nil
}

func PrintGzip(r io.Reader) (n int64, err error) {
	return FprintGzip(os.Stdout, r)
}

func ReadGzip(dst io.Writer, src io.Reader) (n int64, err error) {
	defer w.Close()
	return io.Copy(gzip.NewWriter(dst), src)

}

func FprintGzip(w io.Writer, r io.Reader) (n int64, err error) {
	return io.Copy(gzip.NewWriter(w), r)
}
