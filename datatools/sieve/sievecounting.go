package sieve

type nothing struct{}

type doneChan chan nothing

type next struct {
	me int
	in chan int
}

func SieveCounting(value int) []int {

	max := MaxSearchValue(value)
	ret := make([]int, 0, max)
	start := 2

	primes := make(chan int)
	nextChan := make(chan next)
	done := make(chan bool)

	nextChan <- next{start, make(chan int)}

Out:
	for {
		// in := countby(start)
		// ch := counter(in, primes, done)
		select {
		case n := <-primes:
			ret = append(ret, n)
			if len(ret) >= max {
				done <- true
				break Out
			}
		case n := <-nextChan:
			Counter(n, primes, nextChan, done)
		}
	}
	return ret

}

// Counter keeps watching for ints on the in channel. If the int matches
// 'me' then it is a prime since it was not filtered out at any other
// point. Otherwise, a new channel is sent to the nextChan channel.
// The done channel is used to end the goroutine.
func Counter(in next, primes chan int, nextChan chan next, done chan bool) {
	go func() {
		for {
			select {
			case <-done:
				return
			case n := <-in.in:
				if n == in.me {
					primes <- n
				} else {
					nextChan <- next{n, CountBy(n)}
				}
			}
		}
	}()
}

// CountBy receives an int, me,  and begins counting by that integer, passing
// all of the ints on except for the ones that are multiples of itself.
func CountBy(me int) chan int {
	c := make(chan int)
	counter := me

	go func() {
		for {
			for i := 0; i < me-1; i++ {
				counter++
				c <- counter
			}
			counter++
			// skip sending this item (multiple of interval) to channel
			// any multiple of interval is not prime ...
		}
	}()

	return c
}
