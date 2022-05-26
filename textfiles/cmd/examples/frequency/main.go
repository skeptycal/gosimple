package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/skeptycal/gosimple/textfiles"
)

var (
	sizeFlag  = flag.Int("size", 0, "min word size to count")
	forceFlag = flag.Bool("force", false, "force writing to file (default is os.Stdout)")

	fileList = []string{}
	size     int
)

func init() {
	flag.Parse()
	size = *sizeFlag
	if len(flag.Args()) > 0 {
		fileList = flag.Args()
	}
}

func PWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return pwd
}

func main() {
	// fmt.Println("PWD: ", PWD())

	filename := "the-king-james-bible.txt"
	filename = "textfiles/cmd/examples/data/romeo_and_juliet.txt"

	fmt.Println("\nFrequency1:")
	textfiles.Frequency(filename, 0)

	fmt.Println("\nFrequency2:")
	textfiles.Frequency2(filename)

	fmt.Println("\nFrequency3:")
	textfiles.Frequency3(filename)

	/*
		exampleFile := "../data/romeo_and_juliet.txt"
		exampleFile = "../data/short.txt"

		if len(fileList) < 1 {
			fileList = []string{exampleFile}
		}

		for _, name := range fileList {
			t, err := textfiles.GetTextFile(name)
			if err != nil {
				continue
			}
			// fmt.Println(t.String())
			// fmt.Println(t.LineMap())
			t.CleanAlphaNumeric()
			PrintFrequency(t, size)
			unique, total := t.WordCount()
			fmt.Printf("unique words: %v\ntotal words: %v\n", unique, total)
			t.Top(5)
		}
	*/
}

func PrintFrequency(t *textfiles.TextFile, min int) {

	m := t.Frequency()

	for k, v := range m {
		if v < min {
			continue
		}
		fmt.Printf("%v: %v\n", k, v)
	}
}

func PrintWrapped(text string, size int) {
	PrintScale(size/10 + 1)
	fmt.Println(strings.Repeat("-", size))
	fmt.Println(textfiles.ReWrap(text, size))
	fmt.Println()
}

func PrintScale(size int) {
	numbers := "1234567890"
	repeatNumbers := strings.Repeat(numbers, 10)
	for i := 0; i < size; i++ {
		fmt.Print(strings.Repeat(" ", 9), repeatNumbers[i]-48)
	}
	fmt.Println()
	fmt.Println(strings.Repeat(numbers, size))

}

const (
	exampleText = "This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. This is an example text. "
)
