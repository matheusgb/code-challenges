package main

import "fmt"

func main() {
	var A, B int
	for {
		fmt.Scan(&A, &B)

		if A > B {
			fmt.Println("Decrescente")
		} else if B > A {
			fmt.Println("Crescente")
		} else {
			break
		}

	}
}
