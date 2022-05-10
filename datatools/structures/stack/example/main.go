package main

import (
	"fmt"
	"os"

	"github.com/pkg/profile"
	"github.com/skeptycal/gosimple/datatools/structures/stack"
)

type intStack stack.Stack[int, []*int]

var new = stack.New[int, []*int]

var done = make(chan bool)

func main() {
	defer profile.Start().Stop()

	s := new(0)

	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	fmt.Println("Len: ", s.Len())
	fmt.Println("Cap: ", s.Cap())
	fmt.Println(s)

	ch := popAll(s)

	// loop:
	select {
	case i := <-ch:
		fmt.Println("pop: ", i)
	case b := <-done:
		if b {
			fmt.Println("b: ", b)
		} else {
			fmt.Println("false: ", b)
		}

		os.Exit(0)
	}

	// select {
	// case i := <-ch:
	// 	fmt.Println("pop: ", i)
	// case b := <-done:
	// 	fmt.Println("done: ", b)
	// 	break loop
	// }

	// time.Sleep(5 * time.Second)

	// for i, err := s.Pop(); err != nil; {
	// 	fmt.Println("pop: ", i)
	// }
}

func popAll(s intStack) (out chan int) {
	go func() {
		for i := 0; i < s.Len(); i++ {
			i, err := s.Pop()
			fmt.Println("in Pop(): ", i)
			fmt.Println("err: ", err)
			fmt.Print(s)
			if err != nil {
				fmt.Println(fmt.Errorf("popAll: %v", err))
			}
			out <- i
			// fmt.Println("pop: ", i)
		}
	}()

	return out
}
