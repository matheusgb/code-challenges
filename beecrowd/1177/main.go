package main

import "fmt"

func main() {
	var T, ii int
	fmt.Scan(&T)

	fmt.Println("N[0] = 0")
	for i := 1; i < 1000; i++ {
		if ii < T-1 {
			ii++
		} else {
			ii = 0
		}
		fmt.Printf("N[%v] = %v\n", i, ii)
	}
}
