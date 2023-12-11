package main

import (
	"fmt"
)

func main() {
	var N, S int
	var V float64
	fmt.Scan(&N, &S, &V)
	fmt.Printf("NUMBER = %v\n", N)
	fmt.Printf("SALARY = U$ %.2f\n", float64(S)*V)
}
