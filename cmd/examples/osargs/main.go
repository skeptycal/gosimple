package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skeptycal/gosimple/errorlogger"
	osargs "github.com/skeptycal/gosimple/osargs"
)

var OsArgs = osargs.OsArgs

var (
	log = errorlogger.Log
)

func main() {
	var argFlag, compFlag bool
	flag.BoolVar(&argFlag, "args", false, "show details about command line args...")
	flag.BoolVar(&compFlag, "list", false, "list command line args...")
	flag.Parse()

	flag.Usage()
	fmt.Println("")
	fmt.Println("")

	if compFlag {
		compareArgs()
	}
	if argFlag {
		listArgs()
	}
}

func listArgs() {
	fmt.Println("Comparing raw os.Args to the osargs package output:")
	fmt.Println("(Commonly, symlinks are not evaluated by the os version.)")
	fmt.Println("")
	fmt.Printf("raw os.Args[0]:  %s\n", os.Args[0])
	fmt.Printf("osargs.App:  %s\n", osargs.App)
	fmt.Printf("osargs.Args:  %s\n", osargs.Args)
	fmt.Printf("osargs.Here:  %s\n", osargs.Here)
	fmt.Printf("osargs.Me:  %s\n", osargs.Me)
	fmt.Println("")
}

func compareArgs() {
	fmt.Printf("%25.25s %s\n", "raw os.Args[0]:", os.Args[0])

	osargs.Example()
	args := osargs.OsArgs

	fmt.Printf("%25.25s %s\n", "args.App():", args.App())
	fmt.Printf("%25.25s %s\n", "args.ArgString():", args.ArgString())
	fmt.Printf("%25.25s %s\n", "args.Base():", args.Base())
	fmt.Printf("%25.25s %s\n", "args.Dir():", args.Dir())
	fmt.Printf("%25.25s %s\n", "args.Args():", args.Args())
	fmt.Println("")
}
