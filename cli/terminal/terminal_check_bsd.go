//go:build (darwin || dragonfly || freebsd || netbsd || openbsd) && !js
// +build darwin dragonfly freebsd netbsd openbsd
// +build !js

package terminal

// the isTerminal functionality from logrus is used here.
// MIT Licence Copyright (c) 2014 Simon Eskildsen
// https://github.com/sirupsen/logrus

import (
	"fmt"
	"golang.org/x/sys/unix"
)

const ioctlReadTermios = unix.TIOCGETA

// isTerminal returns true if the given file descriptor is a terminal.
func isTerminal(fd int) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}
