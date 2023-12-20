package main

import (
	"fmt"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	if A < B {
		B, A = A, B
	}
	if A < C {
		A, C = C, A
	}
	if B < C {
		C, B = B, C
	}

	if A >= B+C {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if A*A == B*B+C*C {
			fmt.Println("TRIANGULO RETANGULO")
		}
		if A*A > B*B+C*C {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		if A*A < B*B+C*C {
			fmt.Println("TRIANGULO ACUTANGULO")
		}
		if A == B && B == C {
			fmt.Println("TRIANGULO EQUILATERO")
		} else if A == B || B == C || A == C {
			fmt.Println("TRIANGULO ISOSCELES")
		}
	}
}
