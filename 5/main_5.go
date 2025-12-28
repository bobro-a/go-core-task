package main

import "fmt"

func Intersections(slice1, slice2 []int) (bool, []int) {
	m := make(map[int]int)
	res := make([]int, 0, min(len(slice1), len(slice2)))
	flag := false
	if len(slice1) == 0 || len(slice2) == 0 {
		return false, []int{}
	}
	for _, val := range slice1 {
		m[val]++
	}
	for _, val := range slice2 {
		if _, ok := m[val]; ok && m[val] > 0 {
			res = append(res, val)
			m[val]--
			flag = true
		}
	}
	return flag, res
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(Intersections(a, b))
}
