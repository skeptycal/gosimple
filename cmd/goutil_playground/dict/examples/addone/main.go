package main

import (
	"os"
	"strconv"

	"github.com/skeptycal/gosimple/cli/errorlogger"
	"github.com/skeptycal/gosimple/cmd/goutil_playground/dict/examples/addone/examples"
)

var log = errorlogger.New()

func main() {
	var err error
	inInt := 0
	// inUint := 0
	// inFloat := 0

	if len(os.Args) < 2 {
		inInt, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		inInt = 42
	}

	examples.ExampleAddOne(inInt)

}
