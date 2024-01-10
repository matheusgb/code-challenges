package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// não aceito pelo avaliador do beecrowd
	var A, B string
	for {
		fmt.Scan(&A, &B)
		if A == "0" && B == "0" {
			break
		}

		B = "0" + B
		r := strings.Replace(B, A, "", -1)
		i, _ := strconv.Atoi(r)
		fmt.Println(i)
	}
}
