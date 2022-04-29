package goshell

import (
	"testing"

	tests "github.com/skeptycal/gosimple/tests"
)

func sout(s string) (stout string) {
	stout, _, _ = Shellout(s)
	return
}

func serr(s string) (sterr string) {
	_, sterr, _ = Shellout(s)
	return
}

func cerr(s string) (err error) {
	_, _, err = Shellout(s)
	return err
}

func errString(s string) (e string) {
	return cerr(s).Error()
}

func TestShellout(t *testing.T) {
	tests := []struct {
		command      string
		assertStdout func(...string) bool
		outArg       string
		assertStderr func(...string) bool
		errArg       string
		wantErr      bool
	}{
		// TODO - this is a bunch of strange tests ...
		{
			`echo "hello, world"`,
			tests.AssertStringHasPrefix,
			"hello, world\n",
			tests.AssertTheEmptyString,
			"",
			false,
		},
		{
			`git --version`,
			tests.AssertStringHasPrefix,
			"git version",
			tests.AssertTheEmptyString,
			"",
			false,
		},
		{
			`go version`,
			tests.AssertStringHasPrefix,
			"go version",
			tests.AssertTheEmptyString,
			"",
			false,
		},
		{
			`gh version`,
			tests.AssertStringHasPrefix,
			"gh version",
			tests.AssertTheEmptyString,
			"", // func() string { return "" }(),
			false,
		},
		{
			`go fakeoption`,
			tests.AssertTheEmptyString,
			"",
			tests.AssertStringHasPrefix,
			"go fakeoption: unknown command",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.command, func(t *testing.T) {
			gotStdout, gotStderr, err := Shellout(tt.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shellout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if outGot := tt.assertStdout(tt.outArg, gotStdout); !outGot {
				t.Errorf("Shellout() stdout string assertion = %v, want true (%q)", outGot, gotStdout)
			}
			if outErr := tt.assertStderr(tt.errArg, gotStderr); !outErr {
				t.Errorf("Shellout() stderr string assertion = %v, want true (%q)", outErr, gotStderr)
			}
		})
	}
}
