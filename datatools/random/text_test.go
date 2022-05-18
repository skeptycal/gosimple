package random

import (
	"math/rand"
	"testing"
	"time"

	"gonum.org/v1/gonum/stat"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Test_randLength(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"0,0", 0, 0},
		{"5,5", 5, 5},
		{"3,5", 3, 5},
		{"42,42", 42, 42},
		{"44,444444444", 44, 444444444},
		{"444444444,44", 444444444, 44},
		{"0,-1", 0, -1},
		{"-3,-1", -3, -1},
		{"-3,-14", -3, -14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randLength(tt.min, tt.max)
			max := tt.max
			min := tt.min
			if max < min {
				max, min = min, max
			}
			if got < min || got > max {
				t.Logf("got: %d, min: %d, max: %d", got, tt.min, tt.max)
				t.Errorf("randLength() = %v, not between min(%v) and max(%v)", got, min, max)
			}
		})
	}
}

func TestRandByte(t *testing.T) {
	const n = 100000
	list := make([]float64, n)

	for i := 0; i < n; i++ {
		t.Run("RandByte", func(t *testing.T) {
			got, err := RandByte()
			if err != nil {
				t.Errorf("RandByte() error = %v", err)
				return
			}
			list[i] = float64(got)
		})
	}

	mean, sd := stat.MeanStdDev(list, nil)
	ent := stat.Entropy(list)

	if mean < 127 || mean > 128 {
		t.Errorf("series (n=%v) is not within spec:  mean: %4.2f (sd: %3.2f) entropy: %f", n, mean, sd, ent)
	}
}
