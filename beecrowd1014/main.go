package main

import (
	"fmt"
)

func main() {
	var A int
	var B float64
	fmt.Scan(&A, &B)
	fmt.Printf("%.3f km/l\n", float64(A)/B)
}
