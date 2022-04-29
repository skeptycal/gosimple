package testing

import "testing"

func TestArgs2Pairs(t *testing.T) {
	tests := []struct {
		name    string
		in      []string
		want    []string
		wantErr bool
	}{
		{"TES", []string{"", ""}, []string{"", ""}, false},
		{"below min input length", []string{"no"}, []string{"no"}, false},
		// {"TES", []string{"", ""}, []string{"", ""}, false},
		// {"TES", []string{"false", "false"}, []string{""}, true},
		// {"TES", []string{"false"}, []string{""}, false},
		// {"TES", []string{"false"}, []string{""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := StringFields(tt.in...)
			if len(tt.in) < 2 && len(tt.in)%2 == 0 {
				t.Errorf("TestStringFields(%v) input slice length = %v, want %v", tt.name, len(out), len(tt.want))
				return // no further tests if slice lengths are not equal
				// TODO: this should be accounted for in the function
			}
			if len(out) != len(tt.want) {
				t.Errorf("TestStringFields(%v) output slice length = %d, want %d", tt.name, len(out), len(tt.want))
				return // no further tests if slice lengths are not equal
			}
			for i, got := range out {
				if (got != tt.want[i]) != tt.wantErr {
					t.Errorf("TestStringFields(%v) assertion test = %q, want %q", tt.name, got, tt.want[i])
				}
			}
		})
	}
}

func TestStringFields(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{"TES", []string{""}, []string{""}},
		{"TES", []string{"false"}, []string{""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := StringFields(tt.in...)
			if len(out) != len(tt.want) {
				t.Errorf("TestStringFields(%v) slice length = %v, want %v", tt.name, len(out), len(tt.want))
				return // no further tests if slice lengths are not equal
			}
			for i, got := range out {
				if got != tt.want[i] {
					t.Errorf("TestStringFields(%v) assertion test = %v, want %v", tt.name, got, tt.want[i])
				}
			}
		})
	}
}
