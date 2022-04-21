package atomic

import (
	"sync"
	"sync/atomic"
)

func f(start *uint64, finish int, wg *sync.WaitGroup) {
	for i := 0; i < finish; i++ {
		atomic.AddUint64(start, 1)
	}
	wg.Done()
}

// concurrentInc is used to concurrently increment the
// counter from start to finish. It is a model for a
// counter for functions using multiple goroutines.
// Reference: https://golangdocs.com/atomic-operations-in-golang-atomic-package
func concurrentInc(start, finish, workers int) int {
	var wg sync.WaitGroup
	v := uint64(start)

	// this line is a guess ... just messing around
	// MaxWorkers := runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go f(&v, int(finish), &wg)
	}
	wg.Wait()

	return int(v)
}

// concurrentInc is used to concurrently increment the
// counter from start to finish. It is a model for a
// method that does not work because atomic operations
// are not used.
func concurrentIncFail(start, finish, workers int) int {
	var wg sync.WaitGroup
	v := uint64(start)

	// this line is a guess ... just messing around
	// MaxWorkers := runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go fail(&v, finish, &wg)
	}
	// go f(&start, &wg)
	wg.Wait()

	return int(v)
	// fmt.Println(start)
}

func fail(start *uint64, finish int, wg *sync.WaitGroup) {
	for i := 0; i < finish; i++ {
		*start++
	}
	wg.Done()
}

// loopInc is a model implementation of a traditional
// loop counter for comparison in benchmarks.
// The compiler will likely do some magic on this
// and replace it, but ... I'm just messing around.
// The parameter workers is only to make the function
// signature match for table based testing. It serves
// no other purpose.
func loopInc(start, finish, workers int) int {
	_ = workers
	for i := 0; i < finish; i++ {
		start++
	}
	return start
}
