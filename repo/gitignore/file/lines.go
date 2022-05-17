package file

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/skeptycal/gosimple/repo/gitignore/cli"
)

func (f *GoFile) Lines() []string {
	f.checkDirty()
	if f.lines == nil {

		n := bytes.Count(f.Bytes(), []byte{'\n'})
		retval := make([]string, 0, n+1)

		scanner := bufio.NewScanner(bufio.NewReader(f.buf))
		for scanner.Scan() {
			retval = append(retval, scanner.Text())
			// log.Printf("metric: %s", scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return nil
		}
		f.lines = retval
	}
	return f.lines
}

func (f *GoFile) Tail(n int) string {
	v := f.Lines()
	fmt.Printf("in Tail(%d): %v\n", n, v)
	if v == nil {
		Log.Error(errors.New("error getting file lines").Error())
		return ""
	}

	if n < 1 {
		n = cli.DefaultTailLineLength
	}
	fmt.Printf("in Tail(%d): %v\n", n, v)

	if n > len(v) {
		fmt.Printf("in Tail(%d): returning all\n", n)

		return strings.Join(v, newLine)
	}
	fmt.Printf("  v[len(v)-n:]: %v\n", v[len(v)-n:])
	return strings.Join(v[len(v)-n:], newLine)
}

func (f *GoFile) Head(n int) string {
	v := f.Lines()
	if v == nil {
		Log.Error(errors.New("error getting file lines").Error())
		return ""
	}

	if n < 1 {
		n = cli.DefaultHeadLineLength
	}
	if n > len(v) {
		return strings.Join(v, newLine)
	}
	return strings.Join(v[:n], newLine)
}

func (f *GoFile) Ends(n int) string {
	v := f.Lines()
	if v == nil {
		Log.Error(errors.New("error getting file lines").Error())
		return ""
	}
	if n < 1 {
		n = cli.DefaultTailLineLength
	}
	if len(v) < n*2 {
		return strings.Join(v, newLine)
	}
	return f.Head(n) + "\n ... \n" + f.Tail(n)
}

// AddTrailingSep adds a trailing separator sep to each
// newline in s.
// The actual newline character can be kept or discarded
// based on the keepNewLines bool.
func AddTrailingSep(s, sep string, keepNewLines bool) string {
	if keepNewLines {
		sep = sep + NewLine
	}
	return strings.ReplaceAll(s, "\n", sep)
}

// Lines returns s separated on occurrences of newline.
func Lines(s string) []string {
	return strings.Split(s, NewLine)
}

// Fields returns s separated on occurrences of sep
// after normalizing utf8 whitespace.
func Fields(s string, sep string) []string {
	n := NormalizeWhitespace(s)
	return strings.Split(n, sep)
}

// NormalizeWhitespace splits s on any utf8 whitespace
// and returns each element as a single space
// separated string.
func NormalizeWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// GetFileData returns the contents of the file
// as a string. If any error is encountered,
// the empty string is returned with the error.
func GetFileData(filename string) (string, error) {
	fi, err := os.Stat(*cli.InFile)
	if err != nil {
		return "", err
	}
	inFileName := fi.Name()
	// V("file stat ok: ", inFileName)

	b, err := os.ReadFile(inFileName)
	if err != nil {
		return "", err
	}
	// V("file opened: ", inFileName)

	return B2S(b), nil
}
