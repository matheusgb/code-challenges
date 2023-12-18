package main

import (
	"fmt"
)

func main() {
	var A float64
	fmt.Scan(&A)

	if A > 0 && A <= 400 {
		P := (15 * A) / 100
		fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %v %%\n", A+P, P, 15)
	} else if A > 400 && A <= 800 {
		P := (12 * A) / 100
		fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %v %%\n", A+P, P, 12)
	} else if A > 800 && A <= 1200 {
		P := (10 * A) / 100
		fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %v %%\n", A+P, P, 10)
	} else if A > 1200 && A <= 2000 {
		P := (7 * A) / 100
		fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %v %%\n", A+P, P, 7)
	} else if A > 2000 {
		P := (4 * A) / 100
		fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %v %%\n", A+P, P, 4)
	}
}
