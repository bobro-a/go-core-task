package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
)

func PowCube(in <-chan uint8, out chan<- float64) {
	for val := range in {
		newVal := math.Pow(float64(val), 3)
		out <- newVal
	}
	close(out)
}

func main() {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for range 5 {
			val := uint8(rand.UintN(255))
			ch1 <- val
			fmt.Printf("input: %d\n", val)
		}
		close(ch1)
	}()

	go PowCube(ch1, ch2)

	for val := range ch2 {
		fmt.Printf("out: %.f\n", val)
	}
}
