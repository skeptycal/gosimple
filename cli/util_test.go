package cli

import (
	"testing"
)

// check will compare 'got' and 'want' values and report
// true if they are the same or false if they are different.
//
// 'name' is a descriptive name that will display in error messages.
// 'got' and 'want' should be of the same (comparable) type.
// wantErr states whether an error is expected or not.
//
// Use 'nil' for *testing.T to skip reporting (not recommended)
func check[T comparable](name string, got, want T, wantErr bool, t *testing.T) bool {
	if want != got {
		if !wantErr {
			t.Errorf("%s = %v(%T), want %v(%T)", name, got, got, want, want)
			return false
		}
	}
	return true
}

// func TestColumns(t *testing.T) {

// 	t.Run("Columns()", func(t *testing.T) {
// 		// TODO - find out why this function is not returning the
// 		// correct number of columns...
// 		// see issue
// 		got := Columns()
// 		if got < 1 || got > 1000 {
// 			t.Errorf("Columns() - expected int between 1 and 1000, got: %v", got)
// 		}
// 	})
// }

// func TestCheckIfTerminal(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		w     io.Writer
// 		want  bool
// 		wantW string
// 	}{
// 		// TODO: Add test cases.
// 		{"stdout", os.Stdout, true, ""},
// 		{"nil", nil, false, ""},
// 		{"&bytes.Buffer{}", &bytes.Buffer{}, false, ""},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := CheckIfTerminal(tt.w); got != tt.want {
// 				t.Errorf("CheckIfTerminal() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
