package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	commch := make(chan int, 1000)
	go fillchan(commch)
	for range time.Tick(1000 * time.Millisecond) {
		drainchan(commch)
	}
}

func fillchan(commch chan int) {
	for range time.Tick(300 * time.Millisecond) {
		commch <- rand.Int()
	}
}

func drainchan(commch chan int) {
	for {
		select {
		case e := <-commch:
			fmt.Printf("%v\n", e)
		default:
			return
		}
	}
}
