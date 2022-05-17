package rand

import (
	"fmt"
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
	const n = 10000
	var sum int

	m := make(map[byte]int)

	for i := 0; i < n; i++ {
		name := fmt.Sprintf("RandByte(%d)", i)

		t.Run(name, func(t *testing.T) {
			got, err := RandByte()
			if err != nil {
				t.Errorf("RandByte() error = %v", err)
				return
			}
			m[got]++
			if got > 255 {
				t.Errorf("RandByte() = %v", got)
			}
		})
	}
	for k, v := range m {
		sum += int(k) * v
	}
	avg := sum / n

	m, s := stat.MeanStdDev(data, nil)

	if avg != n/2 {
		t.Errorf("average of series (n=%v) is not within spec:  %v", n, avg)
	}
}
