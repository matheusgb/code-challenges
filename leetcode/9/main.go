package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	i := x
	re := 0
	for i > 0 {
		r := i % 10
		re = (re * 10) + r
		i /= 10
	}

	return re == x
}

func main() {
	fmt.Println(isPalindrome(-101))
}
