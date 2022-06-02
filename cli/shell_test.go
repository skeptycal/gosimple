package cli

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestNewWriteCloserCLI(t *testing.T) {
	tests := []struct {
		name string
		w    any
		want io.WriteCloser
	}{
		// TODO: Add test cases.
		{"os.Stdout", os.Stdout, discardCloser},
		{"bufiowriter os.stdout", bufio.NewWriter(os.Stdout), discardCloser},
		{"discard", io.Discard, discardCloser},
		{"buffered discard", bufio.NewWriter(io.Discard), discardCloser},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewWriteCloserCLI(tt.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriteCloserCLI() = %v, want %v", got, tt.want)
			}
		})
	}
}
