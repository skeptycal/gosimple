package main

import (
	"fmt"
	"math"
	"net/http"
	"sync"

	"github.com/skeptycal/gosimple/cli/errorlogger"
)

/* http status code categories

1xx: Informational - Request received, continuing process
2xx: Success - The action was successfully received, understood, and accepted
3xx: Redirection - Further action must be taken in order to complete the request
4xx: Client Error - The request contains bad syntax or cannot be fulfilled
5xx: Server Error - The server failed to fulfill an apparently valid request
*/

type Category int

const (
	Invalid       Category = iota // 0: Invalid Status Code
	Informational                 // 1xx: Informational - Request received, continuing process
	Success                       // 2xx: Success - The action was successfully received, understood, and accepted
	Redirection                   // 3xx: Redirection - Further action must be taken in order to complete the request
	ClientError                   // 4xx: Client Error - The request contains bad syntax or cannot be fulfilled
	ServerError                   // 5xx: Server Error - The server failed to fulfill an apparently valid request
)

func CatagoryFromStatus(status int) Category {
	n := Category(math.Floor(float64(status) / 100))
	if n < 1 || n > 5 {
		return 0 // errors.New("invalid status category")
	}
	return n
}

func (c Category) String() string {
	switch c {
	// case 0:
	// 	return "0: Invalid Status Code"
	case 1:
		return "1xx: Informational - Request received, continuing process"
	case 2:
		return "2xx: Success - The action was successfully received, understood, and accepted"
	case 3:
		return "3xx: Redirection - Further action must be taken in order to complete the request"
	case 4:
		return "4xx: Client Error - The request contains bad syntax or cannot be fulfilled"
	case 5:
		return "5xx: Server Error - The server failed to fulfill an apparently valid request"
	default:
		return "0: Invalid Status Code"
	}
}

var log = errorlogger.New()

func main() {
	fmt.Println()
	fmt.Println("This works because each instance of v is sent to the channel. The variable is not 'captured by func literal'")
	chNoCapture()
	fmt.Println()
	fmt.Println("This does not work because the same variable v is sent with each channel send. The variable is 'captured by func literal'")
	chCapture()
	fmt.Println()
	fmt.Println("These are the numbers [0..7] in the order that they are received and processed on the channel.")
	wgCountTo8()
	checkURLs()
	for k, v := range urlcache {
		fmt.Printf("%30s%10s... %-20s\n", k, " ", v)
	}
}

func isDone(done chan struct{}) {
	done <- blank
}

var urlcache map[string]string = make(map[string]string)

func checkURLs() {
	// g := new(errgroup.Group)
	done := make(chan struct{})

	var urls = []string{"http://www.golang.org/", "http://www.google.com/", "http://www.somestupidname.com/", "http://www.skeptycal.com/"}
	for _, url := range urls {
		go func(in string) error {
			defer isDone(done)

			// fmt.Printf("%s\n", in)
			resp, err := http.Get(in)
			if err != nil {
				log.Errorf("url error: %v", err)
				return err
			}
			defer resp.Body.Close()
			cat := CatagoryFromStatus(resp.StatusCode)
			if cat > 2 {
				err := http.ErrAbortHandler
				log.Errorf("url error: %v", err)
				log.Errorf("http error category: %v", cat)

			}
			urlcache[in] = resp.Status

			defer resp.Body.Close()

			// stuff ...

			return nil
		}(url)
	}
	for range urls {
		<-done
	}
	fmt.Println("Successfully checked all URLs.")
}

func wgCountTo8() {
	var wg sync.WaitGroup

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(x int) {
			fmt.Println("Thread: ", x)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// This works because each instance of v is sent to
// the channel.
// The variable is not "captured by func literal"
func chNoCapture() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for range values {
		<-done
	}
}

var blank = struct{}{}

// This does not work because the same variable v
// is sent with each channel send.
// The variable is "captured by func literal"
//
// Reference: https://go.dev/doc/faq#closures_and_goroutines
func chCapture() {
	done := make(chan struct{})
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- blank
		}()
	}

	// wait for all goroutines to complete before exiting
	for range values {
		<-done
	}
}
