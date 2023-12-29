package main

import "fmt"

func containsDuplicate(nums []int) bool {
	obj := make(map[int]bool, len(nums))

	for _, n := range nums {
		if obj[n] {
			return true
		} else {
			obj[n] = true
		}
	}
	return false

}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 5}))
}
