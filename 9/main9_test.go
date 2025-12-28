package main

import (
	"sync"
	"testing"
)

func TestPowCube(t *testing.T) {
	tests := []struct {
		input  uint8
		output float64
	}{
		{input: 2, output: 8},
		{input: 0, output: 0},
		{input: 1, output: 1},
		{input: 5, output: 125},
	}

	in := make(chan uint8)
	out := make(chan float64)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for _, tt := range tests {
			in <- tt.input
		}
		close(in)
	}()

	go PowCube(in, out)

	for _, tt := range tests {
		val := <-out
		if tt.output != val {
			t.Errorf("error for the argument: %d expected: %.f, received: %.f\n", tt.input, tt.output, val)
		}
	}
}
