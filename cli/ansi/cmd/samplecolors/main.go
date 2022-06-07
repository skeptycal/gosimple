package main

import (
	"fmt"

	"github.com/skeptycal/gosimple/cli/ansi"
)

func main() {
	a := ansi.BasicEncode("2")

	fmt.Println(a, "Hello, World!", ansi.Reset)
}
