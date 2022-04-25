package rand

import (
	"crypto/rand"
	"fmt"
	mathrand "math/rand"
)

type text interface {
	~string | ~[]byte
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

// RandText returns a random string or slice of random bytes
// n bytes long.
func CreateRandomLengthText[T text](min, max int) (T, error) {
	dif := max - min
	if dif < 1 {
		dif = 1
	}

	size := mathrand.Intn(dif) + min
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
