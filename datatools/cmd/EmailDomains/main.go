package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/skeptycal/gosimple/datatools/format"
)

func main() {
	list := strings.Join(os.Args[1:], " ")
	out := format.GetDomainNames(list)
	fmt.Println(out)
}
