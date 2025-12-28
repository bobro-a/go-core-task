package main

import (
	"slices"
	"testing"
)

func TestUniqueElem(t *testing.T) {
	tests := []struct {
		slice1 []string
		slice2 []string
		result []string
	}{
		{
			slice1: []string{"1", "2"},
			slice2: []string{},
			result: []string{"1", "2"},
		},
		{
			slice1: []string{},
			slice2: []string{"1", "2", "3"},
			result: []string{},
		},
		{
			slice1: []string{},
			slice2: []string{},
			result: []string{},
		},
		{
			slice1: []string{"1"},
			slice2: []string{"1"},
			result: []string{},
		},
	}

	for _, tt := range tests {
		res := UniqueElem(tt.slice1, tt.slice2)
		if !slices.Equal(res, tt.result) {
			t.Errorf("for the slices test: %v %v expected: %v but it was received: %v\n", tt.slice1, tt.slice2, tt.result, res)
		}
	}
}
