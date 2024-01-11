package main

import "fmt"

func main() {
	var T float64

	for i := 0; i < 100; i++ {
		fmt.Scan(&T)
		if T <= 10 {
			fmt.Printf("A[%v] = %.1f\n", i, T)
		}
	}
}
