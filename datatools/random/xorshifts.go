package random

// some random number generators based on xorshift, a subset of LFSR's
// Reference: https://en.wikipedia.org/wiki/Xorshift

func xorshift64() uint64 {
	var x uint64
	x ^= x << 13
	x ^= x >> 7
	x ^= x << 17
	return x
}

func xorshift32() uint32 {
	var x uint32
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	return x
}

const mul = 0x2545F4914F6CDD1D

// A xorshift* generator takes a xorshift generator and applies an invertible multiplication (modulo the word size) to its output as a non-linear transformation, as suggested by Marsaglia.[1] The following 64-bit generator with 64 bits of state has a maximal period of 264−1[8] and fails only the MatrixRank test of BigCrush
//
// 1. https://en.wikipedia.org/wiki/Xorshift#cite_note-marsaglia-1
// 8. https://en.wikipedia.org/wiki/Xorshift#cite_note-vigna-9
func xorshiftStar() int64 {
	var x int64 = 1 /* initial seed must be nonzero, don't use a static variable for the state if multithreaded */
	x ^= x >> 12
	x ^= x << 25
	x ^= x >> 27
	return x * mul // 0x2545F4914F6CDD1DULL;
}

// xoshiro and xoroshiro are other variations of the shift-register generators, using rotations in addition to shifts. According to Vigna, they are faster and produce better quality output than xorshift.[13][14]
//
// 13. https://en.wikipedia.org/wiki/Xorshift#cite_note-xoshiro-web-14
// 14. https://en.wikipedia.org/wiki/Xorshift#cite_note-xoshiro-paper-15
func rol64(x, k int64) int64 {
	return (x << k) | (x >> (64 - k))
}

/*

Reference: https://en.wikipedia.org/wiki/Xorshift

// The state word must be initialized to non-zero
uint32_t xorshift32(struct xorshift32_state *state)
{
	// Algorithm "xor" from p. 4 of Marsaglia, "Xorshift RNGs"
	uint32_t x = state->a;
	x ^= x << 13;
	x ^= x >> 17;
	x ^= x << 5;
	return state->a = x;
}
*/
