package lib

import (
	"runtime"
	"strconv"
)

// NMap works on every integer in the rage from 0 to ... n-1 in parallel
//
//	two slices. The first is a slice of interface{}
//	the second contains errors
//	n is a number, typically the length of a list
func NMap(n int, parmax int, fn func(m int) (r interface{}, err error)) (resultlist []interface{}, errorlist []error) {
	if n == 0 {
		return
	}
	resultlist = make([]interface{}, n)
	errorlist = make([]error, n)
	if n == 1 {
		resultlist[0], errorlist[0] = fn(0)
		return
	}
	if parmax < 1 {
		parmax = -1
	}
	pmax := ""
	if pmax == "" {
		pmax = "1"
	}
	max := runtime.GOMAXPROCS(-1)
	if max > 1 {
		max = max - 1
	}
	if x, _ := strconv.Atoi(pmax); x != -1 && max > x {
		max = x
	}

	if parmax != -1 && max > parmax {
		max = parmax
	}
	if n < max {
		max = n
	}

	// Channel for blocking until work done

	done := make(chan bool, max)

	// Setup producer

	jobs := make(chan int)
	go func(n int, jobs chan<- int) {
		for i := 0; i < n; i++ {
			jobs <- i
		}
		close(jobs)
	}(n, jobs)

	// Setup max consumers
	for i := 0; i < max; i++ {
		go func(jobs <-chan int, done chan<- bool) {
			for job := range jobs {
				resultlist[job], errorlist[job] = fn(job)
			}
			done <- true
		}(jobs, done)
	}
	for i := 0; i < max; i++ {
		<-done
	}
	close(done)
	return
}
