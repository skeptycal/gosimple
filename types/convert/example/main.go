package main

import (
	"fmt"

	"github.com/skeptycal/gosimple/types/convert"
)

func main() {
	// var f = 2.5

	for i := 1; i < 32; i++ {
		// printSample(i, 1<<i)
		fmt.Println("-------------------------------------")
	}

}

func fmtSample(args ...any) {
	list := make([]string, len(args))
	for i, arg := range args {
		list[i] = convert.ToString(arg)
	}
	fmt.Println(list)
}

// func printSample[T constraints.Real](i int, n T) {
// 	fmt.Println("i float, int(), ToInt(), ToInt2()")
// 	// fmt.Printf(" %v, %v, %v, %v\n", i, n, int(n), convert.ToInt(n), convert.ToInt(n))

// 	makeSlice(i, int(n), convert.ToInt(n), convert.ToInt2(n), convert.CastFloat64(n), convert.ToByte(n))
// 	fmtSample(i, n-1)
// 	fmtSample(i, n+1)
// }
