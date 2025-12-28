package main

import "fmt"

func UniqueElem(slice1, slice2 []string) []string {
	m := make(map[string]struct{})
	res := make([]string, 0, len(slice1))
	for _, val := range slice1 {
		m[val] = struct{}{}
	}
	for _, val := range slice2 {
		delete(m, val)
	}
	for _, val := range slice1 {
		if _, ok := m[val]; ok {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(UniqueElem(slice1, slice2))
}
