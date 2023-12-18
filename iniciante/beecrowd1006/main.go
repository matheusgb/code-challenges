package main

import (
	"fmt"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)
	fmt.Printf("MEDIA = %.1f\n", ((A*2)+(B*3)+(C*5))/10)
}
