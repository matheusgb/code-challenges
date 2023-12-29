package main

import (
	"fmt"
)

func main() {
	var N string
	var S, V float64
	fmt.Scan(&N, &S, &V)
	fmt.Printf("TOTAL = R$ %.2f\n", S+((15*V)/100))
}
