// Package async implements a
// generic function to manage async
// calls using goroutines in go applications.
package async

import "sync"

// Executor pushes the passed function output to the channel passed to this function.
// For output channels are mandatory as this uses goroutine. You can create a channel
// of any type or struct.
func Executor[In, Out any](f func(In, *sync.WaitGroup) Out, in In, out chan Out, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		out <- f(in, wg)
	}()
}
