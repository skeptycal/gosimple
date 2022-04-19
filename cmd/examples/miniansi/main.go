package main

import (
	"fmt"

	. "github.com/skeptycal/gosimple/miniansi"
)

func ExampleAnsi() {

	fmt.Println(InfoColor, "Info")
	fmt.Println(NewAnsi(34), "Info")
	fmt.Println("")
	fmt.Println(NoticeColor, "Notice")
	fmt.Println(NewAnsi(36), "Notice")
	fmt.Println("")
	fmt.Println(WarningColor, "Warning")
	fmt.Println(NewAnsi(33), "Warning")
	fmt.Println("")
	fmt.Println(ErrorColor, "Error")
	fmt.Println(NewAnsi(31), "Error")
	fmt.Println("")
	fmt.Println(DebugColor, "Debug")
	fmt.Println(NewAnsi(36), "Debug")
	fmt.Println("")

	DbEcho("dbecho ...")
	fmt.Println("")
}

func main() {
	Example()
	ExampleAnsi()
}
