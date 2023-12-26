package main

import (
	"fmt"
)

func main() {
	var C1, C2, Q1, Q2 int
	var V1, V2 float64

	fmt.Scan(&C1, &Q1, &V1)
	fmt.Scan(&C2, &Q2, &V2)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", (float64(Q1)*V1)+(float64(Q2)*V2))
}
