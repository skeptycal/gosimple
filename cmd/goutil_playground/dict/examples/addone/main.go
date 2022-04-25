package main

import (
	"os"
	"strconv"

	"github.com/skeptycal/goutil/errorlogger"
	"github.com/skeptycal/goutil_playground/dict/examples/addone/examples"
)

func main() {
	var log = errorlogger.New()
	var err error
	inInt := 0
	inUint := 0
	inFloat := 0

	if len(os.Args) < 2 {
		inInt, err = strconv.Atoi(os.Args[1])
		if err != nil {

		}

	}

	examples.ExampleAddOne(42)

}
