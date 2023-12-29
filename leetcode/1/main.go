package main

import (
	"fmt"
)

// bruteforce
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// solução melhor usando hashmap https://www.code-recipe.com/post/two-sum
// func twoSum(nums []int, target int) []int {
// 	indexMap := make(map[int]int)
// 	for currIndex, currNum := range nums {
// 		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
// 			return []int{requiredIdx, currIndex}
// 		}
// 		indexMap[currNum] = currIndex
// 	}
// 	return []int{}
// }

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
