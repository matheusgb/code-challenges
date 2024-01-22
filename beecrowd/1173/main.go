package main

import "fmt"

func main() {
	var A int
	fmt.Scan(&A)

	i := 0
	fmt.Printf("N[%v] = %v\n", i, A)
	y := 1
	for i = A; y < 10; i = 2 * i {
		fmt.Printf("N[%v] = %v\n", y, 2*i)
		y++
	}
}
