package file

import "github.com/skeptycal/gosimple/repo/gitignore/gi_fileparser/cli"

func (f *GoFile) HeadBytes(n int) []byte {
	if n < 1 {
		n = cli.DefaultHeadByteLength
	}
	if n > f.buf.Len() {
		return f.buf.Bytes()
	}
	return f.buf.Bytes()[:n]
}

func (f *GoFile) TailBytes(n int) []byte {
	if n < 1 {
		n = cli.DefaultTailByteLength
	}
	if n > f.buf.Len() {
		return f.buf.Bytes()
	}
	return f.buf.Bytes()[f.buf.Len()-n:]
}
