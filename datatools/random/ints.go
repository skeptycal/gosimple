package random

import (
	rand "crypto/rand"
	"math/big"
	mathrand "math/rand"
	"time"

	"github.com/skeptycal/gosimple/types/constraints"
)

func init() {
	mathrand.Seed(time.Now().UnixNano())
}

type Ints constraints.Integer

func randInt[T Ints](n int) T {
	return T(mathrand.Intn(n))
}

func cRandInt[T Ints](n int) T {
	max := new(big.Int)
	max.SetInt64(int64(n))

	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return T(i.Int64())
}
