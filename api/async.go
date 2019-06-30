package api

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	defaultTimeout = 10 // seconds
)

// AsyncCall...
type AsyncCall struct {
	wg      *sync.WaitGroup
	timeout time.Duration
}

type callback func() error

// NewAsyncCall...
func NewAsyncCall() *AsyncCall {
	return &AsyncCall{
		wg:      &sync.WaitGroup{},
		timeout: time.Duration(defaultTimeout) * time.Second,
	}
}

// Run...
func (a *AsyncCall) RunWithTimeout(cb callback) error {
	ch := make(chan error, 1)
	a.wg.Add(1)

	go func(fn callback) {
		fmt.Fprintln(os.Stderr, "Start async")
		ch <- fn()
		fmt.Fprintln(os.Stderr, "End async")
		a.wg.Done()
	}(cb)

	select {
	case res := <-ch:
		return res
	case <-time.After(a.timeout):
		fmt.Fprintln(os.Stderr, "Timeout")
		return errors.New("timeout")
	}
}

// Wait...
func (a *AsyncCall) Wait() {
	a.wg.Wait()
}
