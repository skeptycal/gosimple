package file

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/skeptycal/gosimple/repo/gitignore/gi_fileparser/cli"
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
