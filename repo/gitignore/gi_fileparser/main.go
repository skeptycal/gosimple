package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skeptycal/gosimple/os/gofile"
	. "github.com/skeptycal/gosimple/repo/gitignore/gi_fileparser/cli"
	"github.com/skeptycal/gosimple/repo/gitignore/gi_fileparser/file"
)

func main() {
	var in *file.GoFile

	in, closer, err := file.NewFile(Options.InFile)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	Vprint("file opened: ", in.Name())

	in.ReadAll()
	in.LoadData()
	Vprint("processing file: ", in.Name())

	// data := getDataCli(inFile)
	// _ = data

	in.PrintDebugDetails(nil)

	if Options.LinesFlag {
		lines := in.Lines()
		if lines == nil {
			log.Fatal("error getting file lines")
		}
		for i, v := range lines {
			fmt.Printf("%3d: %v\n", i, v)
		}
	}

	if Options.FieldsFlag {
		fields := in.Fields(",")
		if fields == nil {
			log.Fatal("error getting file lines")
		}
		Vprint("fields: ", fields)
		for i, v := range fields {
			fmt.Printf("%3d: %v\n", i, v)
		}
	}

	out, err := os.OpenFile("../gitignore_gen.go", os.O_RDWR|os.O_CREATE, gofile.NormalMode)
	if err != nil {
		log.Fatal(err)
	}

	// write stuff to output file here ...

	defer out.Close()
}
