package main

import (
	"fmt"
)

func main() {
	sh := shellpath.NewPath()
	_ = sh.Clean()

	fmt.Println(sh)
	fmt.Println(sh.Out())

}
