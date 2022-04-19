package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/skeptycal/gosimple/channels/concurrent"
)

func Boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(concurrent.Delay)) * time.Millisecond)
		}
	}()
	return c
}

func main() {

	fmt.Printf("Random Seed Test (should change from run to run ...): %v, %v\n", rand.Intn(100), rand.Intn(100))

	c := Boring("Boring...")
	// go Boring("time passes ...", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're Boring .. I'm leaving.")
}

/*

// Version 1

func Boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

Version 2

func Boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}


Version 3

Boring --> go Boring()

func main() {
	go Boring("time passes...")
	fmt.Println("I'm listening...")
	time.Sleep(2 * time.Second)
	fmt.Println("You're Boring .. I'm leaving.")
}

Version 4

func Boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go Boring("time passes ...", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're Boring .. I'm leaving.")
}

Version 5


func Boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(delay)) * time.Millisecond)
		}
	}()
	return c
}

func main() {

	c := Boring("Boring...")
	// go Boring("time passes ...", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're Boring .. I'm leaving.")
}


*/
