package main

import "fmt"

func main() {
	var A int
	for {
		fmt.Scan(&A)

		if A != 2002 {
			fmt.Println("Senha Invalida")
		} else {
			fmt.Println("Acesso Permitido")
			break
		}

	}
}
