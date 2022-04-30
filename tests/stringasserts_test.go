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
		{"has prefix", AssertStringHasPrefix, []string{"pre", "prefix"}, false},
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
		want    [][2]string
		wantErr bool
	}{
		{"nil slice", nil, nil, true},
		{"empty slice", []string{}, [][2]string{}, true},
		{"TES", []string{"", ""}, [][2]string{{"", ""}}, true},
		{"below min input length", []string{"no"}, [][2]string{{"no", ""}}, true},
		{"two strings", []string{"one", "two"}, [][2]string{{"one", "two"}}, false},
		// {"four strings", []string{"one", "two", "three", "four"}, [][2]string{{"one", "two"}, {"three", "four"}}, false},
		// {"two strings reversed", []string{"one", "two"}, [][2]string{{"two", "one"}}, false},
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
			out := Args2Pairs(tt.in...)
			for i := 0; i < len(out); i += 2 {
				for i, got := range out {
					if (got[i] != tt.want[i/2][0]) != tt.wantErr {
						t.Errorf("%v(%q) assertion test = %q, want %q", testname, tt.name, out[i], tt.want[i])
					}
					if (got[i+1] != tt.want[i/2][1]) != tt.wantErr {
						t.Errorf("%v(%q) assertion test = %q, want %q", testname, tt.name, out[i+1], tt.want[i])
					}
				}
			}
		})
	}
}

func TestStringFields(t *testing.T) {
	testname := "TestStringFields"

	tests := []struct {
		name    string
		in      []string
		want    []string
		wantErr bool
	}{
		{"TES", []string{""}, []string{""}, true},
		{"equal", []string{"true"}, []string{"true"}, false},
		{"not equal", []string{"false"}, []string{"true"}, true},
		{"length not equal", []string{"1", "2", "3"}, []string{"false"}, true},
		{"two strings", []string{"true", "true"}, []string{"true", "true"}, false},
		{"three strings", []string{"1", "2", "3"}, []string{"1", "2", "3"}, false},
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
				if got != tt.want[i] != tt.wantErr {
					// if !tt.wantErr {
					t.Errorf("%v(%v)[%v] = %q, want %q", testname, tt.name, i, got, tt.want[i])
					// }
				}
			}
		})
	}
}

func TestAssertStringEqualFold(t *testing.T) {
	testname := "AssertStringEqualFold"
	tests := []struct {
		name    string
		in      []string
		want    bool
		wantErr bool
	}{
		{"TES", []string{"", "", ""}, false, false},
		// {"", []string{""}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AssertStringEqualFold(tt.in...)
			if got != tt.want != tt.wantErr {
				t.Errorf("%v(%v) = %v, want %v", testname, tt.name, got, tt.want)
			}
		})
	}
}
