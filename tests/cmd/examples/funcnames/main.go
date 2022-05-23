package main

import (
	"fmt"
	"runtime"
)

func main() {
	max := 10

	finalfunc := func(i int) { FrameInfo() }
	t := functree(max, finalfunc, true)
	// printFuncTree(t)
	t[0].fn(6)
}

func originalExample() {
	a()
	// Output:
	// - more:true | runtime.Callers
	// - more:true | runtime_test.ExampleFrames.func1
	// - more:true | runtime_test.ExampleFrames.func2
	// - more:true | runtime_test.ExampleFrames.func3
	// - more:true | runtime_test.ExampleFrames
	ExampleFrames(nil)
}

var (
	b = func() {
		fmt.Println("in b - calling c")
		c()
	}

	a = func() {
		fmt.Println("in a - calling b")
		b()
	}

	c = func() {
		fmt.Println("in c - calling frameinfo")
		FrameInfo()
	}
)

type testfunc struct {
	name string
	msg  string
	fn   func(i int)
}

func functree(n int, final func(i int), verbose bool) []testfunc {
	if n < 1 {
		return nil
	}
	if final == nil {
		final = func(i int) {}
	}
	b := make([]testfunc, n)
	for i := 0; i < n; i++ {
		b[i].name = fmt.Sprintf("func%d()", i)
		b[i].msg = fmt.Sprintf("%v - calling func%d())\n", b[i].name, i+1)
		// each func calls the next in line
		// b[i].fn = makefunc(b[i].name, i)
		b[i].fn = func(i int) {
			if verbose {
				fmt.Print(b[i].msg)
			}
			b[i].fn(i + 1)
		}
	}
	// fmt.Println("len: ", len(b))
	// fmt.Println("cap: ", cap(b))
	// b[len(b)-1].name = fmt.Sprintf("last func%d()", len(b)-1)
	b[len(b)-1].msg = fmt.Sprintf("%v - all done! Calling FrameInfo() \n", b[len(b)-1].name)
	b[len(b)-1].fn = final

	// fmt.Println(b)
	return b
}

func printFuncTree(list []testfunc) {
	fmt.Println("Function Tree List:")
	fmt.Println("------------------------------------------------")
	for i, fn := range list {
		fmt.Printf("%d: %s\n", i, fn.name)
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("running Tree[0] only")
	list[0].fn(0)
}

func FrameInfo() {
	// Ask runtime.Callers for up to 10 PCs, including runtime.Callers itself.
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No PCs available. This can happen if the first argument to
		// runtime.Callers is large.
		//
		// Return now to avoid processing the zero Frame that would
		// otherwise be returned by frames.Next below.
		return
	}

	fmt.Println(pc)

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	// Loop to get frames.
	// A fixed number of PCs can expand to an indefinite number of Frames.
	for {
		frame, more := frames.Next()

		// Process this frame.
		//
		// To keep this example's output stable
		// even if there are changes in the testing package,
		// stop unwinding when we leave package runtime.
		// if !strings.Contains(frame.File, "runtime/") {
		// 	break
		// }
		fmt.Printf("- more:%v | %s\n", more, frame.Function)

		// Check whether there are more frames to process after this one.
		if !more {
			break
		}
	}
}

func ExampleFrames(fn func()) {
	if fn != nil {
		fn()
	}

}
