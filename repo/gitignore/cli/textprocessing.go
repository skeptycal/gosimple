package cli

// Head returns the first n elements of a sequence.
// If n is longer than s, s is returned unchanged.
// The default value of n is used. If another value
// of n is needed, use HeadN(s S, n int).
func Head[E any](s []E) []E { return HeadN(s, DefaultHeadLineLength) }

// Tail returns the last n elements of a sequence.
// If n is longer than s, s is returned unchanged.
// The default value of n is used. If another value
// of n is needed, use TailN(s S, n int).
func Tail[E any](s []E) []E { return TailN(s, DefaultTailLineLength) }

// Head returns the first n elements of a sequence s.
// If n is longer than s, s is returned unchanged.
func HeadN[E any](s []E, n int) []E {
	if n > len(s) {
		return s
	}
	return s[:n]
}

// Tail returns the last n elements of a sequence s.
// If n is longer than s, s is returned unchanged.
func TailN[E any](s []E, n int) []E {
	if n > len(s) {
		return s
	}
	return s[len(s)-n:]
}
