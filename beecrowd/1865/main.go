package main

import "fmt"

func main() {
	var A, C int
	var B string
	fmt.Scan(&A)
	for i := 1; i <= A; i++ {
		fmt.Scan(&B, &C)
		if B == "Thor" {
			fmt.Println("Y")
		} else {
			fmt.Println("N")
		}
	}
}
