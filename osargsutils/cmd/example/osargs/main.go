package main

import (
	"fmt"
	"os"

	"github.com/skeptycal/gorepos/osargs"
)

var OsArgs = osargs.OsArgs

func main() {
	fmt.Println("Comparing raw os.Args to the osargs package output:")
	fmt.Println("(Commonly, symlinks are not evaluated by the os version.)")
	fmt.Println("")
	fmt.Printf("raw os.Args[0]:  %s\n", os.Args[0])
	fmt.Printf("osargs.App:  %s\n", osargs.App)
	fmt.Printf("osargs.Args:  %s\n", osargs.Args)
	fmt.Printf("osargs.Here:  %s\n", osargs.Here)
	fmt.Printf("osargs.Me:  %s\n", osargs.Me)
}
