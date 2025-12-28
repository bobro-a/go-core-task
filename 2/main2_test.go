package main

import (
	"slices"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "all zeros",
			input:  []int{0, 0, 0, 0, 0, 0},
			output: []int{0, 0, 0, 0, 0, 0},
		},
		{
			name:   "all ones",
			input:  []int{1, 1, 1, 1, 1},
			output: []int{},
		},
		{
			name:   "nil slice",
			input:  nil,
			output: []int{},
		},
		{
			name:   "one elem",
			input:  []int{9, 9, 9, 9, 2},
			output: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SliceExample(tt.input)
			if !slices.Equal(tt.output, res) {
				t.Errorf("SliceExample(%v) = %v, want %v", tt.input, res, tt.output)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
		index  int
	}{
		{
			name:   "delete first",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{2, 3, 4, 5},
			index:  0,
		},
		{
			name:   "delete last",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 2, 3, 4},
			index:  4,
		},
		{
			name:   "delete middle",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 2, 4, 5},
			index:  2,
		},
		{
			name:   "delete non-existent",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 2, 3, 4, 5},
			index:  100,
		},
		{
			name:   "delete non-existent",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 2, 3, 4, 5},
			index:  -100,
		},
		{
			name:   "empty slice",
			input:  []int{},
			output: []int{},
			index:  1,
		},
	}
	for _, tt := range tests {
		res := removeElement(tt.input, tt.index)
		if !slices.Equal(res, tt.output) {
			t.Errorf("removeElement(%v) with index(%d) = %v, want %v", tt.input, tt.index, res, tt.output)
		}
	}
	tests[0].output[1] = 9999
	if tests[0].input[1] == 9999 {
		t.Errorf("method change input slice\n")
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5, 6}
	slice = addElements(slice, 6)
	if !slices.Equal(slice, expected) {
		t.Error("add method did not work")
	}

	newSlice := slice[:4]
	newSlice = addElements(newSlice, 10)
	if slice[4] == 10 {
		t.Error("add method change original slice")
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := copySlice(slice)
	newSlice[0] = 6
	if slice[0] == 6 {
		t.Errorf("copy continues to link to the previous slice\n")
	}
}
