package random

import (
	"math/big"
	mathrand "math/rand"
	"time"

	"github.com/skeptycal/gosimple/types/constraints"
)

func init() {
	mathrand.Seed(time.Now().UnixNano())
}

// ----------------------------------------------------------------------------

type Ints constraints.Integer

func NewBigInt(x int64) *big.Int { return big.NewInt(x) }

func randInt[T Ints](n int) T {
	return T(mathrand.Intn(n))
}

// Int64 returns a uniform random value in [0, max).
// The value returned is safe for concurrent and cryptographically secure
// applications.
//
// For convenience, if max == 0 or max == nil, 0 is returned.
// If max is < 0, the absolute value of max is used instead.
// (This may lead to unexpected errors.
// It is the caller's responsibility to check the intended functionality
// for max <= 0 and provide a guard if necessary.)
func Int[In, Out Ints](max In) Out {
	if max == 0 {
		return 0
	}
	// if max < 0 {
	// 	return Int(-max)
	// }

	m := big.NewInt(int64(max))
	b, err := bigInt(m)
	if err != nil {
		panic("error generating random number")
	}

	return Out(b.Int64())
}

// bigInt returns a uniform random Int in [0, max). An Int from the
// big package represents a signed multi-precision integer. The Int
// returned is safe for concurrent and cryptographically secure
// applications.
//
// For convenience, if max == 0 or max == nil, a new big.Int with
// the zero value (0) is returned.
// If max is < 0, the absolute value of max is used instead.
// (This may lead to unexpected errors.
// It is the caller's responsibility to check the intended functionality
// for max <= 0 and provide a guard if necessary.)
//
// Operations always take pointer arguments (*Int) rather than Int
// values, and each unique Int value requires its own unique *Int
// pointer. To "copy" an Int value, an existing (or newly allocated)
// Int must be set to a new value using the Int.Set method;
// shallow copies of Ints are not supported and may lead to errors.
func bigInt(max *big.Int) (n *big.Int, err error) {
	if max == nil {
		return new(big.Int), nil
	}

	switch s := max.Sign(); {
	case s == 0:
		return new(big.Int), nil
	case s < 0:
		return bigInt(max.Abs(max))
	default:
	}
	n = new(big.Int)
	n.Sub(max, n.SetUint64(1))

	// bitLen is the maximum bit length needed to encode a value < max.
	bitLen := n.BitLen()
	if bitLen == 0 {
		// the only valid result is 0
		return
	}
	// k is the maximum byte length needed to encode a value < max.
	k := (bitLen + 7) / 8
	// b is the number of bits in the most significant byte of max-1.
	b := uint(bitLen % 8)
	if b == 0 {
		b = 8
	}

	bytes := make([]byte, k)

	for {
		_, err = Read(bytes)
		if err != nil {
			return nil, err
		}

		// Clear bits in the first byte to increase the probability
		// that the candidate is < max.
		bytes[0] &= uint8(int(1<<b) - 1)

		n.SetBytes(bytes)
		if n.Cmp(max) < 0 {
			return
		}
	}
}
