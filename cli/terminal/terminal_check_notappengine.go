//go:build !appengine && !js && !windows && !nacl && !plan9
// +build !appengine,!js,!windows,!nacl,!plan9

package terminal

// the isTerminal functionality from logrus is used here.
// MIT Licence Copyright (c) 2014 Simon Eskildsen
// https://github.com/sirupsen/logrus

import (
	"fmt"
	"io"
	"os"
)

func checkIfTerminal(w io.Writer) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	switch v := w.(type) {
	case *os.File:
		return isTerminal(int(v.Fd()))
	default:
		return false
	}
}
