package main

import (
	"context"
	"testing"
	"time"
)

func TestGenerator_Simple(t *testing.T) {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		Generator(ctx, ch)
		close(ch)
	}()

	count := 0
	for range ch {
		count++
		if count > 10 {
			break
		}
	}
	if count == 0 {
		t.Error("no values generated")
	}
}
