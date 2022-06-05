package reallyunsafe_test

import (
	"bytes"
	"testing"
	"unsafe"

	"github.com/skeptycal/gosimple/reallyunsafe"
	"github.com/skeptycal/gosimple/testes"
)

var tAssertEqual = testes.AssertEqual

type readOp int8

type SneakyBuffer struct {
	Buf      []byte // contents are the bytes buf[off : len(buf)]
	Off      int    // read at &buf[off], write at &buf[len(buf)]
	LastRead readOp // last read operation, so that Unread* can work correctly.
}

func TestRecastSneaky(t *testing.T) {

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

func TestRecastPubPri(t *testing.T) {
	pub := 13
	pri := 42

	// the Fake struct has one public and one private field
	f := reallyunsafe.NewFake(pub, pri)

	// Recast takes the private field and makes it public
	s := reallyunsafe.Recast(*f)

	if s.Public != pub {
		t.Errorf("Public not correct: got %v, want %v", s.Public, pub)
	}

	if s.Private != pri {
		t.Errorf("Private not correct: got %v, want %v", s.Private, pri)
	}
}
