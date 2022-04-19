package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skeptycal/gosimple/errorlogger"
	"github.com/skeptycal/gosimple/osargs"
)

var (
	log     = errorlogger.Log
	verbose bool
)

func main() {
	flag.BoolVar(&verbose, "args", false, "show details about command line args...")
	flag.Parse()
	log.Info("sample log info")
	if verbose {
		args()
	}
}

func args() {
	fmt.Printf("%25.25s %s\n", "raw os.Args[0]:", os.Args[0])

	osargs.Example()
	args := osargs.OsArgs

	fmt.Printf("%25.25s %s\n", "args.App():", args.App())
	fmt.Printf("%25.25s %s\n", "args.ArgString():", args.ArgString())
	fmt.Printf("%25.25s %s\n", "args.Base():", args.Base())
	fmt.Printf("%25.25s %s\n", "args.Dir():", args.Dir())
	fmt.Printf("%25.25s %s\n", "args.Args():", args.Args())
}
