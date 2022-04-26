package main

import (
	crand "crypto/rand"
	"fmt"
	rand "math/rand"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/skeptycal/gosimple/types/constraints"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var c, _ = crand.Prime(crand.Reader, 64)

var counter = 0

type event struct {
	time int
	data struct{}
}

// generate is an infinite loop filling a channel
// with random numbers
func generate[T constraints.Real](ch chan T) {
	var newT T
	for {
		n := rand.Uint64()
		newT = T(atomic.LoadUint64(&n)) // TODO should use pointers
		counter++                       //TODO remove counter after testing
		ch <- newT
	}
}

func fillchanrand[T constraints.Real](commch chan T) {
	for range time.Tick(30 * time.Microsecond) {
		n := rand.Uint64()

		var newT = T(atomic.LoadUint64(&n))

		counter++

		commch <- newT
	}
	fmt.Printf("filled (len: %v)\n", len(commch))
}

func rnd[T constraints.Real](min, max T) T {
	if max == min {
		return max
	}
	if max < min {
		max, min = min, max
	}

	rng := max - min
	fmt.Println("range: ", rng)

	size := unsafe.Sizeof(max)
	fmt.Println("sizeof: ", size)

	return 0

}

func drain(c chan int) {
	go func() {
		for range c {
		}
	}()
}

func main() {
	ch := make(chan int32)
	// fmt.Println("ch: ", ch)
	// fmt.Println("len(ch): ", len(ch))
	// fmt.Println("go fill chan")
	go fillchanrand(ch)
	// fmt.Printf("%v (%v): \n", <-ch, len(ch))
	// fmt.Println()
	go func() {
		for i := 1; ; i++ {
			fmt.Printf("%12v: (%v) %v\n", counter, len(ch), <-ch)
			for j := 1; j < 10; j++ {
				<-ch
			}
		}
	}()
	time.Sleep(5 * time.Second)
}
