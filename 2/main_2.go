package main

import (
	"fmt"
	"math/rand"
)

const (
	countNumber = 10
)

func SliceExample(s []int) []int {
	res := make([]int, 0, len(s))
	for _, val := range s {
		if val%2 == 0 {
			res = append(res, val)
		}
	}
	return res
}

func addElements(s []int, elem int) []int {
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	newSlice = append(newSlice, elem)
	return newSlice
}

func copySlice(s []int) []int {
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	return newSlice
}

func removeElement(s []int, index int) []int {
	newSlice := make([]int, 0, len(s))
	if index >= len(s) || index < 0 {
		return append([]int(nil), s...)
	}
	newSlice = append(newSlice, s[:index]...)
	newSlice = append(newSlice, s[index+1:]...)
	return newSlice
}

func main() {
	originalSlice := make([]int, countNumber)
	for i := range countNumber {
		originalSlice[i] = rand.Intn(10)
	}
	fmt.Println(originalSlice)
	example := SliceExample(originalSlice)
	add := addElements(originalSlice, 100)
	copyS := copySlice(originalSlice)
	removed := removeElement(originalSlice, 9)
	fmt.Println("SliceExample", example)
	fmt.Println("addElements", add)
	fmt.Println("copySlice", copyS)
	fmt.Println("removeElement", removed)
	originalSlice[0] = 5000
	fmt.Println("SliceExample", example)
	fmt.Println("addElements", add)
	fmt.Println("copySlice", copyS)
	fmt.Println("removeElement", removed)
}
