package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	if A%B == 0 || B%A == 0 {
		fmt.Printf("Sao Multiplos\n")
	} else {
		fmt.Printf("Nao sao Multiplos\n")
	}
}
