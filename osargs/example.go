// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package osargs

import (
	"fmt"
)

func Example() {
	here, me := HereMe()

	fmt.Println("Example import acknowledgement from package osargs")
	fmt.Println("The program name is: ", me)
	fmt.Println("The program path is: ", here)
}
