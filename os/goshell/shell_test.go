package goshell

import (
	"testing"

	_ "github.com/skeptycal/gosimple/tests"
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

func TestAssertions(t *testing.T) {
	tests := []struct {
		name      string
		assertion func(...string) bool
		in        []string
		want      bool
	}{
		{"is TES", tests.AssertTheEmptyString, []string{""}, true},
		{"not TES", tests.AssertTheEmptyString, []string{"false"}, false},
		{"has prefix", tests.AssertStringHasPrefix, []string{"pre", "prefix"}, true},
		{"not has prefix", tests.AssertStringHasPrefix, []string{"pre", "false"}, false},
		{"has suffix", tests.AssertStringHasSuffix, []string{"fix", "suffix"}, true},
		{"not has suffix", tests.AssertStringHasSuffix, []string{"fixx", "false"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.assertion(tt.in...)
			if got != tt.want {
				t.Errorf("%v(%v) assertion test = %v, want %v", tt.name, tt.in, got, tt.want)
			}
		})
	}
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
			assertStringHasPrefix,
			"hello, world\n",
			assertTheEmptyString,
			"",
			false,
		},
		{
			`git --version`,
			assertStringHasPrefix,
			"git version",
			assertTheEmptyString,
			"",
			false,
		},
		{
			`go version`,
			assertStringHasPrefix,
			"go version",
			assertTheEmptyString,
			"",
			false,
		},
		{
			`gh version`,
			assertStringHasPrefix,
			"gh version",
			assertTheEmptyString,
			"", // func() string { return "" }(),
			false,
		},
		{
			`go fakeoption`,
			assertTheEmptyString,
			"",
			assertStringHasPrefix,
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
