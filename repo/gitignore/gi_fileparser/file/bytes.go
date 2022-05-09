package file

func (f *GoFile) HeadBytes(n int) []byte {
	if n < 1 {
		n = defaultHeadByteLength
	}
	if n > f.buf.Len() {
		return f.buf.Bytes()
	}
	return f.buf.Bytes()[:n]
}

func (f *GoFile) TailBytes(n int) []byte {
	if n < 1 {
		n = defaultTailByteLength
	}
	if n > f.buf.Len() {
		return f.buf.Bytes()
	}
	return f.buf.Bytes()[f.buf.Len()-n:]
}
