package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	commch := make(chan int, 100)
	go fillchan(commch)
	drainchan(commch)
}

func fillchan(commch chan int) {
	for {
		select {
		case <-time.Tick(30 * time.Millisecond):
			commch <- rand.Int()
		}
	}
}

func drainchan(commch chan int) {
	for {
		chanlen := len(commch) // get number of entries in channel
		time.Sleep(1 * time.Second)
		for i := 0; i <= chanlen; i++ { //flush them based on chanlen
			fmt.Printf("chan len: %v num: %v\n", chanlen, <-commch)
		}
	}
}
