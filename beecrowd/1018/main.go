package main

import (
	"fmt"
)

func main() {
	var A, cem, cin, vin, dez, sin, doi, uno int
	fmt.Scan(&A)
	fmt.Printf("%v\n", A)
	for A > 0 {
		if A >= 100 {
			A = A - 100
			cem++
		} else if A >= 50 {
			A = A - 50
			cin++
		} else if A >= 20 {
			A = A - 20
			vin++
		} else if A >= 10 {
			A = A - 10
			dez++
		} else if A >= 5 {
			A = A - 5
			sin++
		} else if A >= 2 {
			A = A - 2
			doi++
		} else {
			A = A - 1
			uno++
		}
	}
	fmt.Printf("%v nota(s) de R$ 100,00\n", cem)
	fmt.Printf("%v nota(s) de R$ 50,00\n", cin)
	fmt.Printf("%v nota(s) de R$ 20,00\n", vin)
	fmt.Printf("%v nota(s) de R$ 10,00\n", dez)
	fmt.Printf("%v nota(s) de R$ 5,00\n", sin)
	fmt.Printf("%v nota(s) de R$ 2,00\n", doi)
	fmt.Printf("%v nota(s) de R$ 1,00\n", uno)
}
