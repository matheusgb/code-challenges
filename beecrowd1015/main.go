package main

import (
	"fmt"
	"math"
)

func main() {
	var X1, X2, Y1, Y2 float64
	fmt.Scan(&X1, &Y1, &X2, &Y2)
	fmt.Printf("%.4f\n", math.Sqrt(math.Pow(X2-X1, 2)+math.Pow(Y2-Y1, 2)))
}
