//go:build !windows && !plan9 && !solaris
// +build !windows,!plan9,!solaris

package terminal

// getWinsize contains code from the goterm package
// Reference: https://github.com/buger/goterm (MIT License)
// Reference: https://www.darkcoding.net/software/pretty-command-line-console-output-on-unix-in-python-and-go-lang/

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func getWinsize() (*unix.Winsize, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
