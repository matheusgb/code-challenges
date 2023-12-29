package main

import (
	"fmt"
	"math"
)

func main() {
	var V float64

	fmt.Scan(&V)
	fmt.Printf("VOLUME = %.3f\n", (4.0/3.0)*3.14159*math.Pow(V, 3))
}
