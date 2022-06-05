package testes

import (
	"testing"
)

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
		{"CheckPairs", CheckPairs[string], []string{"fixx", "false"}, true},
		{"not CheckPairs", CheckPairs[string], []string{"fixx", "false", "wrong"}, false},
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

func TestCheckPairs(t *testing.T) {
	testname := "CheckPairs"
	var pairTests = []struct {
		name string
		in   []string
		want bool
	}{
		{"nil slice", nil, false},
		{"empty slice", []string{}, false},
		{"TES", []string{""}, false},
		{"2xTES", []string{"", ""}, true},
		{"<2 args", []string{"no"}, false},
		{"two strings", []string{"one", "two"}, true},
		{"three strings", []string{"one", "two", "three"}, false},
		{"four strings", []string{"one", "two", "three", "four"}, true},
		{"two strings reversed", []string{"one", "two"}, true},
	}
	for _, tt := range pairTests {
		t.Run(testname, func(t *testing.T) {
			if got := CheckPairs(tt.in...); got != tt.want {
				t.Errorf("%v(%v) correct input pairs = %v, want %v", testname, tt.name, got, tt.want)
			}

		})
	}
}

func TestArgs2Pairs(t *testing.T) {
	testname := "Args2Pairs"
	for _, tt := range pairTests {
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
	testname := "StringFields"
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
		{"three strings", []string{"1", "2", "3"}, []string{"1", "2", "3"}, true},
		{"3 strings 4 fields", []string{"1 1", "2", "3"}, []string{"1", "1", "2", "3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := ToFields(tt.in...)
			got := len(out)
			want := len(tt.want)
			if got < want {
				if !tt.wantErr {
					t.Errorf("%v(%v) slice length = %v, want %v", testname, tt.name, got, want)
				}
				return // no further tests
			}
			for i, got := range out {
				if !Contains(got, tt.want) {
					if !tt.wantErr {
						t.Errorf("%v(%v)[%v] = %q, want %q", testname, tt.name, i, got, tt.want[i])
					}
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

var pairTests = []struct {
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

func getT[T any](v T) T {
	return T(v)
}

func TestLen(t *testing.T) {
	testname := "Len"
	i := 42
	arr1 := [2]int{1, 2}
	tests := []struct {
		name  string
		elems any
		want  int
	}{
		{"nil", nil, 0},
		{"empty", []any{}, 0},
		{"1", []any{1}, 1},
		{"2", []any{1, 2}, 2},
		{"bool", true, 1},
		{"array", arr1, 2},
		{"map", map[int]bool{1: true, 2: false}, 2},
		{"slice", []byte{32, 48, 65}, 3},
		{"chan", make(chan int), 0},
		{"string", "string", 6},
		{"ptr int", &i, 2},
		{"ptr array", &arr1, 5},
		{"int", 42, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Len(tt.elems); got != tt.want {
				t.Errorf("%v(%v) (%v) = %v, want %v", testname, tt.name, tt.elems, got, tt.want)
			}
		})
	}
}
