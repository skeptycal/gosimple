package main

import (
	"fmt"
	"time"

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

loop:
	for {
		select {
		case i := <-ch:
			fmt.Println("pop: ", i)
		case b := <-done:
			fmt.Println("done: ", b)
			break loop
		}
	}

	time.Sleep(5 * time.Second)

	// for i, err := s.Pop(); err != nil; {
	// 	fmt.Println("pop: ", i)
	// }
}

func popAll(s intStack) (out chan int) {
	go func() {
		for {
			i, err := s.Pop()
			if err != nil {
				done <- true
				return
			}
			out <- i
			// fmt.Println("pop: ", i)
		}
	}()

	return out
}
