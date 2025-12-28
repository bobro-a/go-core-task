package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type MyWaitGroup struct {
	semaphore chan struct{}
	wait      chan struct{}
	size      uint64
	count     int64
	closed    uint32
}

func NewMyWaitGroup(size uint64) *MyWaitGroup {
	return &MyWaitGroup{semaphore: make(chan struct{}, size),
		wait: make(chan struct{}),
		size: size}
}

func (wg *MyWaitGroup) Add(i int64) {
	if atomic.CompareAndSwapUint32(&wg.closed, 1, 0) {
		wg.wait = make(chan struct{})
	}

	if i <= 0 {
		return
	}

	for range i {
		select {
		case wg.semaphore <- struct{}{}:
		default:
			panic("too many channels to wait for")
		}
	}

	atomic.AddInt64(&wg.count, i)
}

func (wg *MyWaitGroup) Done() {
	select {
	case <-wg.semaphore:
	default:
		panic("there are already no active goroutines")
	}

	if atomic.AddInt64(&wg.count, -1) == 0 {
		if atomic.CompareAndSwapUint32(&wg.closed, 0, 1) {
			close(wg.wait)
		}
	}
}

func (wg *MyWaitGroup) Wait() {
	<-wg.wait
}

func main() {

	wg := NewMyWaitGroup(10)
	wg.Add(10)

	for val := range 10 {
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("%d\n", val)
		}()
	}

	wg.Wait()
}
