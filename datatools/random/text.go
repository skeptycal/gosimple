package random

import (
	"crypto/rand"
	"fmt"
)

type text interface {
	~string | ~[]byte
}

// randLength returns a random number between min and max.
// If min == max, min is returned.
// If min > max, they are reversed before processing.
func randLength[T Ints](min, max T) T {

	switch v := max - min; {
	case v == 0:
		return min
	case v < 0:
		return randLength(max, min)
	case v > 0:
		// return 4
		return Int(v) + min
	default:
		panic("number is not zero, less than zero, or greater than zero") // impossible but funny message ...
	}
}

// RandByte returns a random byte.
// It uses the default crypto/rand package source and is
// safe for concurrent use by multiple goroutines.
func RandByte() (byte, error) {
	r := make([]byte, 1)
	nn, err := rand.Read(r)
	if err != nil || nn != 1 {
		return 0, fmt.Errorf("rand: failed to generate random byte: %v", err)
	}
	return r[0], nil
}

// RandText returns a random string or slice of random bytes
// n bytes long.
// It uses the default crypto/rand package source and is
// safe for concurrent use by multiple goroutines.
func RandText[T text](n int) (T, error) {
	r := make([]byte, n)
	nn, err := rand.Read(r)
	if err != nil || nn != n {
		return T(""), fmt.Errorf("rand: failed to generate random text: %v", err)
	}
	return T(r), nil
}

func Rnd() (int, error) {
	b := make([]byte, 8)
	n, err := rand.Reader.Read(b)
	if err != nil || n != 8 {
		return 0, fmt.Errorf("failed to generate random int: %v", err)
	}
	return int(b[0]), nil
}

// blank returns a new zero-value object of the same type as v.
func blank[T any]() T { return *new(T) }

// RandText returns a random string or slice of random bytes
// with a length between min and max.
func CreateRandomLengthText[T text](min, max int) (retval T, err error) {

	if max == min {
		var b byte
		b, err = RandByte()
		if err != nil {
			return // blank[T](), err
		}
		buf := []byte{b}
		return T(buf), nil
	}
	if max < min {
		return CreateRandomLengthText[T](max, min)
	}
	dif := max - min

	size := Int(dif) + min
	// size := rand.Int()
	// size := mathrand.Intn(dif) + min
	return RandText[T](size)
}

func CreateRandomTextSets[T text](n, min, max int) ([]T, error) {
	args := make([]T, n)
	for i := 0; i < n; i++ {
		a, err := CreateRandomLengthText[T](min, max)
		if err != nil {
			return nil, fmt.Errorf("error generating random data: %v", err)
		}
		args[i] = a
	}
	return args, nil
}

// CreateTextSets creates n sets of random strings
// or byte slices that are 1<<(n*mult) in length
func CreateTextSets[T text](n, mult int) ([]T, error) {
	args := make([]T, n)
	for i := 0; i < n; i++ {
		a, err := RandText[T](1 << (i * mult))
		if err != nil {
			return nil, fmt.Errorf("error generating random data: %v", err)
		}
		args[i] = a
	}
	return args, nil
}
