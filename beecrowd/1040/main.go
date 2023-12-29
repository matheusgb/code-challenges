package main

import (
	"fmt"
)

func main() {
	var A, B, C, D, E float64
	fmt.Scan(&A, &B, &C, &D)

	M1 := (A*2 + B*3 + C*4 + D*1) / 10
	fmt.Printf("Media: %.1f\n", M1)

	if M1 >= 7.0 {
		fmt.Printf("Aluno aprovado.\n")
	} else if M1 >= 5.0 && M1 <= 6.9 {
		fmt.Printf("Aluno em exame.\n")
		fmt.Scan(&E)

		E1 := (M1 + E) / 2
		fmt.Printf("Nota do exame: %.1f\n", E)
		if E1 >= 5.0 {
			fmt.Printf("Aluno aprovado.\n")
		} else {
			fmt.Printf("Aluno reprovado.\n")
		}
		fmt.Printf("Media final: %.1f\n", E1)
	} else {
		fmt.Printf("Aluno reprovado.\n")
	}
}
