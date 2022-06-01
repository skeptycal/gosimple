package functions

import (
	"fmt"
	"testing"

	"github.com/skeptycal/gosimple/datatools/list/tests"
)

func TestFactorialr(t *testing.T) {
	testdata := []tests.TestDataDetails[int, int]{
		{"0", 0, 1, false},
		{"1", 1, 1, true},
		{"2", 2, 2, false},
		{"3", 3, 6, false},
		{"4", 4, 24, false},
		{"5", 5, 120, false},
		{"42", 42, 7538058755741581312, false},
		{"-1", -1, 1, false},
		{"-5", -5, 1, true},
	}

	tests.MakeTestRunner("factorialr", factorialr, testdata).Run(t, "factorialr")

	for _, tt := range testdata {
		// tests.NewTestTable("Test_factorialr")
		checkMathFact(t, tt.In, tt.Want)

	}
}

func BenchmarkFactorialr(b *testing.B) {
	testdata := []tests.TestDataDetails[int, int]{
		{"0", 0, 1, false},
		{"1", 1, 1, true},
		{"2", 2, 2, false},
		{"3", 3, 6, false},
		{"4", 4, 24, false},
		{"5", 5, 120, false},
		{"42", 42, 7538058755741581312, false},
		{"-1", -1, 1, false},
		{"-5", -5, 1, true},
	}

	tests.MakeBenchmarkRunner("factorialr", factorialr, testdata).Run(b, "factorialr")

}

func checkMathFact(t *testing.T, in, want int) (ok bool) {
	name := fmt.Sprintf("math.factorial(%v)", in)
	t.Run(name, func(t *testing.T) {
		if got := factorialr(in); got != want {
			t.Errorf("factorial(%v) = %v, want %v", in, got, want)
			ok = false
		} else {
			ok = true
		}
	})
	return
}
