package goshell

import (
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

func TestShellout(t *testing.T) {
	testname := "Shellout"
	tests := []struct {
		in         string
		wantSout   string
		outLen     int // 0 = ignore
		wantSerr   string
		errLen     int // 0 = ignore
		wantSerror error
		wantErr    bool
	}{
		// TODO - this is a bunch of strange tests ...
		{`echo "hello, world"`, "hello, world\n", 0, "", 0, nil, false},
		{`git --version`, "git version", 11, "", 11, nil, false},
		{`go version`, "go version", 10, "", 0, nil, false},
		{`gh version`, "gh version", 10, "", 0, nil, false},
		{`go fakeoption`, "", 0, "go fakeoption: unknown command", 0, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			gotStdout, gotStderr, gotErr := Shellout(tt.in)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("%v() error = %v, wantErr %v", testname, gotErr, tt.wantErr)
				return
			}
			if tt.outLen > 0 {
				gotStdout = gotStdout[:tt.outLen]
			}
			if tt.errLen > 0 {
				gotStderr = gotStderr[:tt.errLen]
			}
			if gotStdout != tt.wantSout {
				t.Errorf("%v() stderr string = %v want %v", testname, gotStderr, tt.wantSout)
			}
			if gotStderr != tt.wantSerr {
				t.Errorf("%v() stderr string = %v want %v", testname, gotStderr, tt.wantSout)
			}
			if gotErr != tt.wantSerror != tt.wantErr {
				t.Errorf("%v() error = %v, want %v", testname, gotErr, gotStdout)
			}
		})
	}
}
