package main

import (
	"fmt"
)

// há solução melhor com o bucketsort https://leetcode.com/problems/top-k-frequent-elements/solutions/81602/java-o-n-solution-bucket-sort/
// https://leetcode.com/problems/top-k-frequent-elements/solutions/1930575/hashmap/
func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	mr := map[int][]int{}
	for n, count := range m {
		mr[count] = append(mr[count], n)
	}

	o := []int{}
	for i := len(nums); len(o) != k; i-- {
		for _, n := range mr[i] {
			if len(o) != k {
				o = append(o, n)
			}
		}
	}

	return o
}

func main() {
	fmt.Println(topKFrequent([]int{3, 2, 1, 2}, 1))
}
