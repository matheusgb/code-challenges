package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Printf("%.3f\n", float64(A*B)/12)
}
