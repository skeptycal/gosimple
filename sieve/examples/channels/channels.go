package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type none struct{}

type noneChannel chan none

var noChan = make(noneChannel)

// Fill Channel with int values
func fillChan[T any](number int, blank T) <-chan T {
	c := make(chan T)
	go func() {
		for i := 0; i < number; i++ {
			c <- blank
		}
		close(c)
	}()
	return c
}

// Create multiple channels and fill them
func createChannels[T any](number int, fill T) (chans []<-chan T) {
	chans = make([]<-chan T, number)
	for i := 0; i < number; i++ {
		chans[i] = fillChan(number, fill)
	}
	return
}

func mergeTwo[T any](a, b <-chan T) (c chan T) {
	c = make(chan T)
	go func() {
	loop:
		for {
			select {
			case c <- <-a:
				//
			case c <- <-b:
				//
			default:
				break loop
			}
		}
		close(c)
	}()
	return c
}

func mergeRec[T any](chans ...<-chan T) <-chan T {
	switch len(chans) {
	case 0:
		c := make(chan T)
		close(c)
		return c
	case 1:
		return chans[0]
	default:
		m := len(chans) / 2
		return mergeTwo(
			mergeRec(chans[:m]...),
			mergeRec(chans[m:]...))
	}
}

func mergeWait[T any](cs ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan T) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func mergeAtomic[T any](cs ...<-chan T) <-chan T {
	out := make(chan T)
	var i int32
	atomic.StoreInt32(&i, int32(len(cs)))
	for _, c := range cs {
		go func(c <-chan T) {
			for v := range c {
				out <- v
			}
			if atomic.AddInt32(&i, -1) == 0 {
				close(out)
			}
		}(c)
	}
	return out
}

func Merge[T any](cs ...<-chan T) <-chan T {
	return mergeAtomic(cs...)
}

func main() {
	a := fillChan(2, 2)
	b := fillChan(2, 2)
	c := fillChan(2, 2)
	d := mergeWait(a, b, c)
	for v := range d {
		fmt.Println(v)
	}
}
