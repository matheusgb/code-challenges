package main

import (
	"fmt"
)

func main() {
	var raio float64
	fmt.Scan(&raio)
	fmt.Printf("A=%.4f\n", 3.14159*raio*raio)
}
