package main

import (
	"fmt"
	"log"

	shellpath "github.com/skeptycal/gosimple/os/shellpath"
)

func main() {
	sh, err := shellpath.NewPath()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sh)

	n := sh.Clean()

	fmt.Println()
	fmt.Printf("%v\n", n)
	fmt.Println()

	fmt.Println(sh)
	fmt.Println(sh.Out())
	fmt.Println()
	fmt.Printf("%v\n", n)
}
