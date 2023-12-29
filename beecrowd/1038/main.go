package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	C := []float64{4.00, 4.50, 5.00, 2.00, 1.50}
	fmt.Printf("Total: R$ %.2f\n", C[A-1]*float64(B))
}
