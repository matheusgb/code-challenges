package main

import (
	"fmt"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A)

	if A > 4500 {
		C = A - 4500
		B = (28 * C) / 100
	}

	if A > 3000 && A <= 4500 {
		C = A - 3000
		B += (18 * C) / 100
	} else if A > 3000 {
		B += (18 * 1500) / 100
	}

	if A > 2000 && A <= 3000 {
		C = A - 2000
		B += (8 * C) / 100
	} else if A > 2000 {
		B += (8 * 1000) / 100
	} else {
		fmt.Println("Isento")
	}

	if B > 1 {
		fmt.Printf("R$ %.2f\n", B)
	}

}
