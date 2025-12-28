package main

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

const countChan = 20

func TestMergeChannels(t *testing.T) {
	slice := make([]chan any, countChan)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := range slice {
		slice[i] = make(chan any, 100)
		go func(ch chan any) {
			for {
				select {
				case <-ctx.Done():
					close(ch)
					return
				default:
					ch <- rand.Int()
				}
			}
		}(slice[i])
	}

	count := 0
	out := MergeChannels(slice...)
	for {
		select {
		case v, ok := <-out:
			if !ok {
				if count == 0 {
					t.Fatal("no values received")
				}
				return
			}
			count++
			t.Logf("got: %v", v)
		}
	}
}
