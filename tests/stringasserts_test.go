package tests

import "testing"

func TestAssertions(t *testing.T) {
	testname := "TestAssertions"
	tests := []struct {
		name      string
		assertion func(...string) bool
		in        []string
		want      bool
	}{
		{"is TES", AssertTheEmptyString, []string{""}, true},
		{"not TES", AssertTheEmptyString, []string{"false"}, false},
		{"has prefix", AssertStringHasPrefix, []string{"pre", "prefix"}, true},
		{"not has prefix", AssertStringHasPrefix, []string{"pre", "false"}, false},
		{"has suffix", AssertStringHasSuffix, []string{"fix", "suffix"}, true},
		{"not has suffix", AssertStringHasSuffix, []string{"fixx", "false"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.assertion(tt.in...)
			if got != tt.want {
				t.Errorf("%v(%v) assertion test = %v, want %v", testname, tt.name, got, tt.want)
			}
		})
	}
}

func TestArgs2Pairs(t *testing.T) {
	testname := "TestArgs2Pairs"
	tests := []struct {
		name    string
		in      []string
		want    []string
		wantErr bool
	}{
		{"nil slice", nil, nil, true},
		{"empty slice", []string{}, []string{}, true},
		{"below min input length", []string{"no"}, []string{"no"}, true},
		{"TES", []string{"", ""}, []string{"", ""}, true},
		{"below min input length", []string{"no"}, []string{"no"}, true},
		// {"TES", []string{"", ""}, []string{"", ""}, false},
		// {"TES", []string{"false", "false"}, []string{""}, true},
		// {"TES", []string{"false"}, []string{""}, false},
		// {"TES", []string{"false"}, []string{""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// tests for panic conditions
			/// the function will panic if either of these are true
			if len(tt.in) < 2 || len(tt.in)%2 != 0 {
				if !tt.wantErr {
					t.Errorf("%v(%v) input slice length = %v, want (a positive even integer)", testname, tt.name, len(tt.in))
				}
				return // no further tests if input slice is invalid
			}
			out := StringFields(tt.in...)
			if len(out) != len(tt.want) {
				if !tt.wantErr {
					t.Errorf("%v(%v) output slice length = %d, want %d", testname, tt.name, len(out), len(tt.want))
				}
				return // no further tests if slice lengths are not equal
			}
			for i, got := range out {
				if (got != tt.want[i]) != tt.wantErr {
					t.Errorf("%v(%q) assertion test = %q, want %q", testname, tt.name, got, tt.want[i])
				}
			}
		})
	}
}

func TestStringFields(t *testing.T) {
	tests := []struct {
		name    string
		in      []string
		want    []string
		wantErr bool
	}{
		{"TES", []string{""}, []string{""}, true},
		{"equal", []string{"true"}, []string{"true"}, false},
		{"not equal", []string{"false"}, []string{"true"}, false},
		{"length not equal", []string{"1", "2", "3"}, []string{"false"}, true},
		{"two strings", []string{"true", "true"}, []string{"true", "true"}, false},
		// {"not equal", []string{"false"}, []string{""}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := StringFields(tt.in...)
			if len(out) != len(tt.want) {
				if !tt.wantErr {
					t.Errorf("TestStringFields(%v) slice length = %v, want %v", tt.name, len(out), len(tt.want))
				}
				return // no further tests if slice lengths are not equal
			}
			for i, got := range out {
				if got != tt.want[i] {
					// if !tt.wantErr {
					t.Errorf("TestStringFields(%v)[%v] = %q, want %q", tt.name, i, got, tt.want[i])
					// }
				}
			}
		})
	}
}
