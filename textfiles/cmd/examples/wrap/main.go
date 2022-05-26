package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/skeptycal/gosimple/textfiles"
)

var (
	sizeFlag  = flag.Int("size", 79, "line length to perform wrapping")
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

func main() {

	if len(fileList) < 1 {
		PrintWrapped(exampleText, size)
		return
	}

	for _, name := range fileList {
		b, err := os.ReadFile(name)
		if err != nil {
			continue
		}
		PrintWrapped(textfiles.B2S(b), size)
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
