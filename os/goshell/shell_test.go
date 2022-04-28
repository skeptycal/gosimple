package goshell

import (
	"strings"
	"testing"
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

func assertStringPrefix(s string, prefix string) (retval string) {
	if strings.HasPrefix(s, prefix) {
		return s
	} else {
		return "**fail**"
	}
}

func TestShellout(t *testing.T) {
	tests := []struct {
		command    string
		wantStdout string
		wantStderr string
		wantErr    bool
	}{
		// TODO - this is a bunch of strange tests ...
		{
			`echo "hello, world"`,
			func(s string) string { return s }("hello, world\n"),
			"",
			false,
		},
		{
			`git --version`,
			assertStringPrefix(sout("git --version"), "git version"),
			"",
			false,
		},
		{
			`go version`,
			assertStringPrefix(sout("go version"), "go version"),
			func() string { return "" }(),
			false,
		},
		{
			`gh version`,
			assertStringPrefix(sout("gh version"), "gh version"),
			"", // func() string { return "" }(),
			false,
		},
		{
			`go fakeoption`,
			"", // func() string { return "" }(),
			assertStringPrefix(serr("go fakeoption"), "go fakeoption: unknown"),
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
			if gotStdout != tt.wantStdout {
				t.Errorf("Shellout() stdout = %q, want %q", gotStdout, tt.wantStdout)
			}
			if gotStderr != tt.wantStderr {
				t.Errorf("Shellout() stderr = %q, want %q", gotStderr, tt.wantStderr)
			}
		})
	}
}
