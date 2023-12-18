package main

import (
	"fmt"
)

func main() {
	var A, B, C, D, hora, minuto int

	fmt.Scan(&A, &B, &C, &D)
	if A > C {
		hora = 24 - A + C
	} else if C > A {
		hora = C - A
	} else {
		if D > B {
			hora = 0
		} else {
			hora = 24
		}
	}

	if B > D {
		minuto = (D - B) + 60
		hora--
	} else if D > B {
		minuto = D - B
	} else {
		minuto = 0
	}

	fmt.Printf("O JOGO DUROU %v HORA(S) E %v MINUTO(S)\n", hora, minuto)
}
