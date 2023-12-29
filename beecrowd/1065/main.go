package main

import (
	"fmt"
)

func main() {
	var A, B, C, D, E, CC int
	fmt.Scan(&A, &B, &C, &D, &E)

	arr := []int{A, B, C, D, E}

	for i := 0; len(arr) > i; i++ {
		if arr[i]%2 == 0 {
			CC++
		}
	}
	fmt.Printf("%v valores pares\n", CC)
}
