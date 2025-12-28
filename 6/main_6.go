package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Generator(ctx context.Context, out chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			out <- rand.Int()
		}

	}
}

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Microsecond)
	defer cancel()
	go func() {
		Generator(ctx, ch)
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
