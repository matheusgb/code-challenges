package main

import (
	"fmt"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)
	if A+B > C && A+C > B && B+C > A {
		fmt.Printf("Perimetro = %.1f\n", A+B+C)
	} else {
		fmt.Printf("Area = %.1f\n", ((A+B)*C)/2)
	}
}
