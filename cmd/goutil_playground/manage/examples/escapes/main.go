package main

import (
	"fmt"
	"os"

	"github.com/skeptycal/gosimple/cli/ansi"
	"github.com/skeptycal/gosimple/cli/miniansi"
)

var (
	debug bool = true
	me         = os.Args[0]
	usage      = `usage: ` + me + ` <package name>`
	red        = ansi.RedText
	reset      = miniansi.ResetColor
)

func dbecho(format string, args ...any) (n int, err error) {
	return fmt.Fprintf(os.Stderr, red+format+reset, args...)
}
func main() {

}

func example() {

	fmt.Printf("This is an %sexample.%s", red, reset)
}
