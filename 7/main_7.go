package main

import (
	"fmt"
	"sync"
)

func MergeChannels(channels ...chan any) chan any {
	out := make(chan any)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	go func() {
		wg.Wait()
		close(out)
	}()

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}
	return out
}

func main() {
	ch1 := make(chan any, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	ch2 := make(chan any, 2)
	ch2 <- "a"
	ch2 <- "b"
	close(ch2)

	ch3 := make(chan any, 10)
	slice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for _, v := range slice {
		ch3 <- v
	}
	close(ch3)

	out := MergeChannels(ch1, ch2, ch3)

	for v := range out {
		fmt.Println(v)
	}
}
