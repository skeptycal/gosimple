// Package http contains utilities for http requests.
package http

import (
	"bytes"
	"io"
	"net/http"

	"github.com/skeptycal/gosimple/os/gofile"
)

// DownloadURL - download content from a URL to <filename>
func DownloadURL(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f := gofile.Create(filename)
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

type RWC struct {
	*bytes.Buffer
	isClosed bool
	readFn   func(p []byte) (n int, err error) `default:"rwc.Read"`
	writeFn  func(p []byte) (n int, err error) `default:"rwc.Write"`
}

// Close resets the underlying buffer to empty but
// retains the buffer for reuse. In order to use
// the buffer again, call Init()
func (rwc *RWC) Close() error {
	rwc.readFn = rwc.fakeRead
	rwc.writeFn = rwc.fakeWrite
	rwc.Reset()
	return nil
}

// func (rwc *RWC) Init() {
// 	rwc.readFn = rwc.Read
// 	rwc.writeFn = rwc.Write
// }

func (rwc *RWC) Read(p []byte) (n int, err error)      { return rwc.readFn(p) }
func (rwc *RWC) Write(p []byte) (n int, err error)     { return rwc.writeFn(p) }
func (rwc *RWC) fakeRead(p []byte) (n int, err error)  { return len(p), io.ErrClosedPipe }
func (rwc *RWC) fakeWrite(p []byte) (n int, err error) { return len(p), io.ErrClosedPipe }

// GetReader returns an io.Reader with the contents of resp.Body.
func GetReader(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	size := int(resp.ContentLength)
	if size == -1 {
		size = 2048
	}

	buf := &bytes.Buffer{}
	buf.Grow(size)

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
