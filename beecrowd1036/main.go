package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	delta := math.Pow(B, 2) - (4 * A * C)
	R1 := (-B + math.Sqrt(delta)) / (2 * A)
	R2 := (-B - math.Sqrt(delta)) / (2 * A)

	if math.IsNaN(R1) || math.IsInf(R1, 0) || math.IsNaN(R2) || math.IsInf(R2, 0) {
		fmt.Printf("Impossivel calcular\n")
	} else {
		fmt.Printf("R1 = %.5f\n", R1)
		fmt.Printf("R2 = %.5f\n", R2)
	}
}
