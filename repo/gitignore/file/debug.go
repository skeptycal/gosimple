package file

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/skeptycal/gosimple/cli"
	"github.com/skeptycal/gosimple/os/gofile"
)

type details struct {
	name  string
	value any
}

// DebugDetails returns a list of interesting properties
// that we want to track during development and testing.
func (f *GoFile) DebugDetails() []details {
	return []details{
		{"gofile.PWD(): ", gofile.PWD()},
		{"fi name: ", f.FileInfo().Name()},
		{"fi size: ", f.FileInfo().Size()},
		{"fi mode: ", f.FileInfo().Mode()},
		{"fi isdir: ", f.FileInfo().IsDir()},

		{"file type: ", reflect.ValueOf(f.f).Type()},
		{"fi type: ", reflect.ValueOf(f.FileInfo()).Type()},
		{"isDirty: ", f.IsDirty()},
		{"Head(0): ", f.Head(0)},
		{"Tail(3): ", f.Tail(3)},
		{"Ends(2): ", f.Ends(2)},
		{"HeadBytes(10): ", f.HeadBytes(10)},
		{"TailBytes(5): ", f.TailBytes(5)},
		{"string(HeadBytes(50)): ", string(f.HeadBytes(50))},
		{"len(buf): ", f.buf.Len()},
		{"cap(buf): ", f.buf.Cap()},
	}
}

func (f *GoFile) PrintDebugDetails(detailSet []details) {
	if len(detailSet) == 0 {
		detailSet = f.DebugDetails()
	}
	max := 0
	for _, detail := range detailSet {
		if v := len(detail.name); v > max {
			max = v
		}
	}

	max += 2

	format := fmt.Sprintf("%%%ds: %%v\n", max)
	// fmt.Println("format: ", format)

	for _, detail := range detailSet {
		cli.DbEchof(format, detail.name, detail.value)
	}
}

var ErrNotImplemented = errors.New("not implemented")
