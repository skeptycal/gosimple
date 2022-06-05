// Package terminal provides information about the state of the terminal.
package terminal

import (
	"io"
)

type winSize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

type winsize = winSize

type WinSize interface {
	Row() int
	Col() int
	Xpixel() int
	Ypixel() int
}

func GetWinSize() WinSize {
	_, err := getWinSize()
	if err != nil {
		return nil
	}
	return &blank{}
}

// getWinSize converts the machine specific
// struct to the standard *winsize type.
func getWinSize() (*winsize, error) {
	var w any
	w, err := getWinsize()
	if err != nil {
		return nil, err
	}
	return w.(*winsize), nil
}

// w is a helper to return either the *winsize
// or nil if there is an error.
func w() *winSize {
	w, err := getWinSize()
	if err != nil {
		return nil
	}
	return w
}

// blank implements the WinSize interface using
// direct calls to the w() function instead of
// struct values.
type blank struct{}

func (*blank) Row() int    { return int(w().Row) }
func (*blank) Col() int    { return int(w().Col) }
func (*blank) Xpixel() int { return int(w().Xpixel) }
func (*blank) Ypixel() int { return int(w().Ypixel) }

func CheckIfTerminal(w io.Writer) bool {
	return checkIfTerminal(w)
}

func IsTerminal(fd int) bool {
	return isTerminal(fd)
}
