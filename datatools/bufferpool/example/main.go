package main

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/skeptycal/gosimple/datatools/bufferpool"
)

var pool = bufferpool.NewPool[int]()
var bpool = bufferpool.NewPool[*bytes.Buffer]()

var sum1, sum2 int64

var globalSink any

// func swimmer() {
// 	var i int
// 	defer pool.Swimmer(i)

// 	i += rand.Intn(3) - 1
// 	sum1 += i

// }

// func getput() {
// 	i := pool.Get()
// 	defer pool.Put(i)

// 	i += rand.Intn(3) - 1
// 	sum2 += i

// }

func swimmer(n int) {
	var w bytes.Buffer
	defer bpool.Swimmer(&w)

	// for i := 0; i < n/10; i++ {
	num := byte(rand.Intn(255))
	// fmt.Printf("%02x\n", num)
	w.WriteByte(num)
	// }

	// out = append(out, w.Bytes()[0])
	// globalSink = w.Bytes()
	sum1 += int64(len(w.Bytes()))
}

var sinkWriter = bytes.Buffer{}

func getput(n int) {
	var w = bpool.Get()
	defer bpool.Put(w)

	// buf := make([]byte, 0, n)

	// for i := 0; i < n/10; i++ {
	num := byte(rand.Intn(255))
	// fmt.Printf("%02x\n", num)
	w.WriteByte(num)
	// }

	// rand.Read(buf)

	// nn, err := w.Write(buf)
	// if err != nil {
	// 	fmt.Printf("w.Write(buf) (n=%v): %v\n", nn, err)
	// }
	// fmt.Println(len(w.Bytes()))
	// fmt.Println(w.String())
	// out = append(out, w.Bytes()[0])
	// globalSink = w.Bytes()
	sum2 += int64(len(w.Bytes()))
}

const size = 1 << 8

var out = make([]byte, 0, size)

func main() {
	for j := 0; j < 10000; j++ {
		go swimmer(size)
		// go getput(size)
	}

	fmt.Println()
	fmt.Printf("sum1: %d\n", sum1)
	fmt.Printf("sum2: %d\n", sum2)
	fmt.Printf("out: %v\n", out)
}
