package cli

import (
	"errors"
	"fmt"
	"strings"
)

func Br() { fmt.Println() }
func Hr() { fmt.Println(headerString()) }
func Fr() { fmt.Println(footerString()) }

func DbEcho(args ...any) (n int, err error) {
	if !DEBUG {
		return 0, nil
	}
	if len(args) < 1 {
		return 0, errors.New("no arguments provided")
	}
	if len(args) == 1 {
		return fmt.Println(args[0])
	}

	// if first arg is format string - best guess =)
	if v, ok := args[0].(string); ok {
		if strings.Contains(v, "%") {
			return fmt.Printf(v+"\n", args[1:]...)
		}
	}

	return fmt.Println(args...)
}

func Box(args ...any) (n int, err error) {
	Hr()
	n, err = fmt.Print(args...)
	Fr()
	return
}
