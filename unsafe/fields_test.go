package unsafe_test

import (
	"bytes"
	"testing"
	"unsafe"

	us "github.com/skeptycal/gosimple/unsafe"
)

type readOp int8

type SneakyBuffer struct {
	Buf      []byte // contents are the bytes buf[off : len(buf)]
	Off      int    // read at &buf[off], write at &buf[len(buf)]
	LastRead readOp // last read operation, so that Unread* can work correctly.
}

func TestRecast2(t *testing.T) {

	buf := bytes.Buffer{}
	b := []byte("This is a buffer entry.")
	buf.Write(b)
	buf.Read(make([]byte, 5))
	sneaky := *(*SneakyBuffer)(unsafe.Pointer(&buf))

	tAssertEqual(t, "sneaky.Off", sneaky.Off, 5)
	tAssertEqual(t, "len(sneaky.Buf)", len(sneaky.Buf), 23)

	// 	opRead      readOp = -1 // Any other read operation.
	tAssertEqual(t, "sneaky.LastRead", sneaky.LastRead, -1)

	// contents are the bytes buf[off : len(buf)]
	tAssertEqual(t, "string", string(sneaky.Buf[sneaky.Off:len(sneaky.Buf)]), "is a buffer entry.")

	_ = sneaky

}

func tAssertEqual[T comparable](t *testing.T, name string, got, want T) {
	format := "%s not equal - got: %v, want %v"
	if name == "" {
		name = "value"
	}
	if got != want {
		t.Errorf(format, name, got, want)
	}
}

func TestRecast(t *testing.T) {
	pub := 13
	pri := 42

	// the Fake struct has one public and one private field
	f := us.NewFake(pub, pri)

	// Recast takes the private field and makes it public
	s := us.Recast(*f)

	if s.Public != pub {
		t.Errorf("Public not correct: got %v, want %v", s.Public, pub)
	}

	if s.Private != pri {
		t.Errorf("Private not correct: got %v, want %v", s.Private, pri)
	}
}
