package main

import "fmt"

func majorityElement(nums []int) int {
	m := make(map[int]int)
	g := 0
	in := 0

	for _, num := range nums {
		m[num]++
	}

	for i, mm := range m {
		if mm > g {
			g = mm
			in = i
		}
	}
	return in
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
