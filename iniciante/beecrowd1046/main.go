package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	if A > B {
		fmt.Printf("O JOGO DUROU %v HORA(S)\n", 24-A+B)
	} else if B > A {
		fmt.Printf("O JOGO DUROU %v HORA(S)\n", B-A)
	} else {
		fmt.Printf("O JOGO DUROU 24 HORA(S)\n")
	}
}
