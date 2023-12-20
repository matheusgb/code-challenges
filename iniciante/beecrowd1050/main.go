package main

import (
	"fmt"
)

func main() {
	var A int
	fmt.Scan(&A)

	ddd := map[int]string{
		61: "Brasilia",
		71: "Salvador",
		11: "Sao Paulo",
		21: "Rio de Janeiro",
		32: "Juiz de Fora",
		19: "Campinas",
		27: "Vitoria",
		31: "Belo Horizonte",
	}

	if ddd[A] == "" {
		fmt.Println("DDD nao cadastrado")
	} else {
		fmt.Println(ddd[A])
	}
}
