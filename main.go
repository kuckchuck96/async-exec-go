// Package async implements a
// generic function to manage async
// calls using goroutines in go applications.
package async

import (
	"errors"
	"sync"
	"time"
)

// Executor pushes the passed function output to the channel passed to this function.
// For output channels are mandatory as this uses goroutine. You can create a channel
// of any type or struct.
// Params:
// in -> input param
// out -> output channel
// wg -> wait group
func Executor[In, Out any](f func(In) Out, in In, out chan Out, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(in In, wg *sync.WaitGroup) {
		out <- f(in)
		wg.Done()
	}(in, wg)
}

// Result returns the actual data by listening to the
// channel for the specific time duration passed.
// Params:
// out -> output channel
// dur -> time duration (seconds, minutes, or hours etc.)
func Result[Out any](out chan Out, dur time.Duration) (Out, error) {
	var result Out
	select {
	case result = <-out:
		return result, nil
	case <-time.After(dur):
		return result, errors.New("channel timeout")
	}
}
