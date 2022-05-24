package assert

import (
	"runtime"

	"github.com/skeptycal/gosimple/types/constraints"
)

type Ordered constraints.Ordered

type variadic[T any] struct {
	first T
	rest  []T
}

func trace() (file, name string, line int) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[3])
	file, line = f.FileLine(pc[1])
	name = f.Name()
	return
	// fmt.Printf("%s:%d %s\n", file, line, f.Name())
}

func fnName(n int) string {
	if n < 1 {
		n = 1
	}
	callers := make([]uintptr, n+1) // at least 1 entry needed
	// v := runtime.Callers(n, pc)
	f := runtime.CallersFrames(callers)
	counter := 0
	for {
		frame, ok := f.Next()
		{
			if counter == n {
				return frame.Function
			}
			counter++
		}
		if !ok {
			break
		}
	}
	return ""
}

func Var2Slice[T any](in ...T) []T {
	if len(in) == 0 {
		return nil
	}
	if len(in) == 1 {
		return []T{in[0]}
	}
	return in
}
