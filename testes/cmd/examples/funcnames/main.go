package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

const defaultMaxframes = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	treeExample(20)
	// originalExample() m2222
}

func treeExample(n int) {
	if n < 0 {
		n = defaultMaxframes
	}
	finalfunc := func(i int) {
		fmt.Println(FrameInfo(n + 10))
		fmt.Println(GetFuncName(4))
	}
	t := functree(n, finalfunc, true)
	t[0].fn(0)
}

func originalExample() {
	var (
		max = 10
		c   = func() {
			fmt.Println("in c - calling frameinfo")
			fmt.Println(FrameInfo(max))
		}

		b = func() {
			fmt.Println("in b - calling c")
			c()
		}

		a = func() {
			fmt.Println("in a - calling b")
			b()
		}
	)
	_, _ = b, c
	a()
	// Output:
	// - more:true | runtime.Callers
	// - more:true | runtime_test.ExampleFrames.func1
	// - more:true | runtime_test.ExampleFrames.func2
	// - more:true | runtime_test.ExampleFrames.func3
	// - more:true | runtime_test.ExampleFrames
}

type testfunc struct {
	name string
	msg  string
	fn   func(i int)
}

func makeTestFunc(name, msg string, fn func(i int)) testfunc {
	return testfunc{
		name: "func()",
		msg:  "single, lonely function ... ",
		fn:   fn,
	}
}

func functree(n int, final func(i int), verbose bool) []testfunc {
	if n < 1 {
		return nil
	}
	if n == 1 {
		return []testfunc{{
			name: "func()",
			msg:  "single, lonely function ... ",
			fn:   final,
		}}
	}
	if final == nil {
		final = func(i int) { os.Exit(n) }
	}
	b := make([]testfunc, n)
	for j := 0; j < n; j++ {
		r := rand.Intn(n)

		t := testfunc{}

		t.name = fmt.Sprintf("func%d()", j)
		t.msg = fmt.Sprintf("%v - calling func%d() (random: %v)\n", t.name, j+1, r)

		t.fn = func(i int) {
			if verbose {
				fmt.Print(t.msg)
			}
			t.fn(i + 1)
		}
		b[j] = t
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

// FrameInfo returns a string summary of the package
// path-qualified functions name of this call frame.
func FrameInfo(max int) string {
	frames := getFrames(0, max)
	if frames == nil {
		return ""
	}

	sb := &strings.Builder{}

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
		fmt.Fprintf(sb, "- more:%v | %s\n", more, frame.Function)

		// Check whether there are more frames to process after this one.
		if !more {
			break
		}
	}
	return sb.String()
}

// GetFuncName returns the package path-qualified function
// name of a call frame, skipping 'skip' initial frames.
//
// If non-empty, this string uniquely identifies a single
// function in the program. This may be the empty string
// if not known.
func GetFuncName(skip int) string {
	return getFrame(skip).Function
}

func getFrames(skip, max int) *runtime.Frames {

	if max < 0 || max > defaultMaxframes {
		max = defaultMaxframes
	}

	if skip < 0 {
		skip = 0
	}

	if skip >= max {
		max = skip + 1
	}

	// Ask runtime.Callers for up to max PCs, including runtime.Callers itself.
	callers := make([]uintptr, max+1)
	n := runtime.Callers(skip, callers)
	if n == 0 {
		// No PCs available. This can happen if the first argument to
		// runtime.Callers is large.
		//
		// Return now to avoid processing the zero Frame that would
		// otherwise be returned by frames.Next below.
		return nil
	}

	// pass only valid pcs to runtime.CallersFrames
	return runtime.CallersFrames(callers[:n])
}

func getFrame(skip int) runtime.Frame {
	frames := getFrames(skip, skip+1)
	frame, _ := frames.Next()
	return frame
}
