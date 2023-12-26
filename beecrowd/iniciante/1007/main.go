package main

import (
	"fmt"
)

func main() {
	var A, B, C, D float64
	fmt.Scan(&A, &B, &C, &D)
	fmt.Printf("DIFERENCA = %v\n", (A*B)-(C*D))
}
