package main

import (
	"slices"
	"testing"
)

func TestIntersections(t *testing.T) {
	tests := []struct {
		slice1 []int
		slice2 []int
		flag   bool
		res    []int
	}{
		{
			slice1: []int{1, 2, 3},
			slice2: []int{4, 5, 6},
			flag:   false,
			res:    []int{},
		},
		{
			slice1: []int{1, 2, 3},
			slice2: []int{1, 2, 3},
			flag:   true,
			res:    []int{1, 2, 3},
		},
		{
			slice1: []int{1, 1, 2},
			slice2: []int{4, 1, 1},
			flag:   true,
			res:    []int{1, 1},
		},
		{
			slice1: []int{1, 1, 1},
			slice2: []int{1, 1, 1},
			flag:   true,
			res:    []int{1, 1, 1},
		},
		{
			slice1: []int{},
			slice2: []int{4, 5, 6},
			flag:   false,
			res:    []int{},
		},
	}

	for _, tt := range tests {
		flag, result := Intersections(tt.slice1, tt.slice2)
		if flag != tt.flag || !slices.Equal(result, tt.res) {
			t.Errorf("for the slices test: %v %v expected: %t %v but it was received: %t %v\n", tt.slice1, tt.slice2, tt.flag, tt.res, flag, result)
		}
	}
}
