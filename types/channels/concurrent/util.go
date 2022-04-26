package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	// Delay is used to set variable time delays for testing purposes.
	Delay int = 1e3
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type GlobalVar[T any] struct{ v T }

func (v *GlobalVar[T]) String() string { return fmt.Sprintf("%v(%T)", v.v, v.v) }
func (v *GlobalVar[T]) Set(t T)        { v.v = t }
func (v *GlobalVar[T]) Get() T         { return v.v }

// func Delay(delay int, wait chan bool) chan bool {
// 	time.Sleep(time.Duration(rand.Intn(delay)) * time.Millisecond)
// 	wait <- true
// 	return wait
// }

func RandomlyTimedString(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(Delay)) * time.Millisecond)
		}
	}()
	return c
}

func RandomString(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return string(b)
}

/*
package atomic map benchmarks

goos: darwin
goarch: arm64
pkg: sync
BenchmarkLoadMostlyHits/*sync_test.DeepCopyMap-8         	100000000	        11.83 ns/op	       6 B/op	       0 allocs/op
BenchmarkLoadMostlyHits/*sync_test.RWMutexMap-8          	 6637725	       162.9 ns/op	       6 B/op	       0 allocs/op
BenchmarkLoadMostlyHits/*sync.Map-8                      	38954029	        30.55 ns/op	       6 B/op	       0 allocs/op
BenchmarkLoadMostlyMisses/*sync_test.DeepCopyMap-8       	162203781	         6.897 ns/op	       6 B/op	       0 allocs/op
BenchmarkLoadMostlyMisses/*sync_test.RWMutexMap-8        	 6995403	       153.5 ns/op	       6 B/op	       0 allocs/op
BenchmarkLoadMostlyMisses/*sync.Map-8                    	52411282	        23.99 ns/op	       6 B/op	       0 allocs/op
BenchmarkLoadOrStoreBalanced/*sync_test.RWMutexMap-8     	 2406345	       520.5 ns/op	      79 B/op	       1 allocs/op
BenchmarkLoadOrStoreBalanced/*sync.Map-8                 	 3210109	       346.9 ns/op	      64 B/op	       2 allocs/op
BenchmarkLoadOrStoreUnique/*sync_test.RWMutexMap-8       	 1795509	       685.2 ns/op	     196 B/op	       2 allocs/op
BenchmarkLoadOrStoreUnique/*sync.Map-8                   	 1672172	       642.5 ns/op	     118 B/op	       4 allocs/op
BenchmarkLoadOrStoreCollision/*sync_test.DeepCopyMap-8   	215833160	         5.135 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadOrStoreCollision/*sync_test.RWMutexMap-8    	 3634327	       297.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadOrStoreCollision/*sync.Map-8                	56959322	        17.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteBalanced/*sync_test.RWMutexMap-8   	 4427896	       289.6 ns/op	       3 B/op	       0 allocs/op
BenchmarkLoadAndDeleteBalanced/*sync.Map-8               	39282172	        31.48 ns/op	       4 B/op	       0 allocs/op
BenchmarkLoadAndDeleteUnique/*sync_test.RWMutexMap-8     	 4242176	       295.9 ns/op	       7 B/op	       0 allocs/op
BenchmarkLoadAndDeleteUnique/*sync.Map-8                 	37229599	        32.34 ns/op	       8 B/op	       0 allocs/op
BenchmarkLoadAndDeleteCollision/*sync_test.DeepCopyMap-8 	 3728505	       330.0 ns/op	      48 B/op	       1 allocs/op
BenchmarkLoadAndDeleteCollision/*sync_test.RWMutexMap-8  	 4392308	       291.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadAndDeleteCollision/*sync.Map-8              	166873578	         6.967 ns/op	       0 B/op	       0 allocs/op
BenchmarkRange/*sync_test.DeepCopyMap-8                  	  626563	      2688 ns/op	       0 B/op	       0 allocs/op
BenchmarkRange/*sync_test.RWMutexMap-8                   	    7705	    154181 ns/op	   16384 B/op	       1 allocs/op
BenchmarkRange/*sync.Map-8                               	  172954	      6639 ns/op	       0 B/op	       0 allocs/op
BenchmarkAdversarialAlloc/*sync_test.DeepCopyMap-8       	 2033851	       610.8 ns/op	     528 B/op	       1 allocs/op
BenchmarkAdversarialAlloc/*sync_test.RWMutexMap-8        	 6849891	       174.3 ns/op	       8 B/op	       1 allocs/op
BenchmarkAdversarialAlloc/*sync.Map-8                    	 3223450	       362.8 ns/op	      48 B/op	       1 allocs/op
BenchmarkAdversarialDelete/*sync_test.DeepCopyMap-8      	 7244424	       174.8 ns/op	     168 B/op	       1 allocs/op
BenchmarkAdversarialDelete/*sync_test.RWMutexMap-8       	 7019437	       175.1 ns/op	      25 B/op	       1 allocs/op
BenchmarkAdversarialDelete/*sync.Map-8                   	11044850	       108.4 ns/op	      22 B/op	       1 allocs/op
BenchmarkDeleteCollision/*sync_test.DeepCopyMap-8        	 3882878	       331.6 ns/op	      48 B/op	       1 allocs/op
BenchmarkDeleteCollision/*sync_test.RWMutexMap-8         	 4452268	       274.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeleteCollision/*sync.Map-8                     	150914280	         8.417 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 58.6% of statements
ok  	sync	49.982s

*/
