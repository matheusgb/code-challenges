package main

import "fmt"

func productExceptSelf(nums []int) []int {
	a := make([]int, len(nums))

	p := 1
	for i := 0; i < len(nums); i++ {
		a[i] = p
		p *= nums[i]
	}

	pf := 1
	for i := len(nums) - 1; i >= 0; i-- {
		a[i] *= pf
		pf *= nums[i]
	}

	return a
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
