package list

import (
	"reflect"
	"testing"
)

func Test_prependOneLine(t *testing.T) {
	tests := []struct {
		name    string
		element E
		list    S
		want    S
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prependOneLine(tt.args.element, tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prependOneLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
