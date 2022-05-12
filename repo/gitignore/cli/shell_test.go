package cli

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestNewWriteCloser(t *testing.T) {
	tests := []struct {
		name string
		w    any
		want io.WriteCloser
	}{
		// TODO: Add test cases.
		{"writer", bufio.NewWriter(os.Stdout), io.WriteCloser(bufio.NewWriter(os.Stdout))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWriteCloser(tt.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriteCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
