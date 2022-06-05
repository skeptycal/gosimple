package io

import (
	"bufio"
	"io"
	"log"
	"strings"
	"text/scanner"

	"github.com/pkg/errors"
)

func scanFile(r io.Reader, n int) {
	defer CloseReaderIf(r)

	// rd := bufio.NewReader(r)

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		_ = sc.Text() // GET the line string
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}

func readFileLines(r io.Reader) (lines []string, err error) {
	defer CloseReaderIf(r)

	rd := bufio.NewReader(r)

	for {
		line, err1 := rd.ReadString('\n')
		lines = append(lines, line)
		if err1 != nil {
			if err1 != io.EOF {
				return nil, errors.Wrapf(err1, "read file line error: %v", err1)
			}
			break
		}
	}
	return lines, nil
}

type head struct {
	r io.Reader // source Reader
	n int       // number of lines
}

const sampleText = `sample text
This stuff is boring.
It reads lines.

It reads more lines.
It puts the lotion on its skin...

last line
`

func sample() string {
	sampleReader := strings.NewReader(sampleText)
	sampleHead := &head{sampleReader, DefaultHeadLength}
	b := make([]byte, 0, DefaultHeadLength*80)
	_, err := sampleHead.Read(b)
	if err != nil {
		return ""
	}
	return string(b)
}

// Head returns an io.Reader that always
// returns the first n lines of a source reader
func Head(r io.Reader, n int) ([]byte, error) {
	return io.ReadAll(&head{r, n})
}

func (h *head) Read(p []byte) (n int, err error) {
	pr, pw := io.Pipe()
	var s scanner.Scanner
	s.Init(h.r)
	s.Whitespace ^= 1 << '\n' // don't skip new lines
	// s.Whitespace ^= 1<<'\t' | 1<<'\n' // don't skip tabs and new lines

	for tok := s.Scan(); tok != scanner.EOF && s.Line < h.n; tok = s.Scan() {
		// if tok == '\n' {
		// 	count++
		// }
		n++
		pw.Write([]byte(string(tok)))
		pr.Read(p)
	}
	return n, pr.Close()
}

func ByteHead() {
	// lr, err := io.LimitReader(h.r, int64(n)).Read(buf)
	// if err != nil {
	// 	return nil, err
	// }
}
