package main

import (
	"fmt"
)

func main() {
	var A, B float64
	fmt.Scan(&A, &B)

	if A == 0 && B == 0 {
		fmt.Printf("Origem\n")
	} else if A == 0 {
		fmt.Printf("Eixo Y\n")
	} else if B == 0 {
		fmt.Printf("Eixo X\n")
	} else if A < 0 && B > 0 {
		fmt.Printf("Q2\n")
	} else if A < 0 && B < 0 {
		fmt.Printf("Q3\n")
	} else if A > 0 && B > 0 {
		fmt.Printf("Q1\n")
	} else {
		fmt.Printf("Q4\n")
	}
}
