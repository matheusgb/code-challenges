package main

import (
	"fmt"
	"math"
)

func main() {
	var NO float64
	fmt.Scan(&NO)
	N := int(NO)
	C := int(math.Round((NO - float64(N)) * 100))
	fmt.Printf("NOTAS:\n%v nota(s) de R$ 100.00\n%v nota(s) de R$ 50.00\n%v nota(s) de R$ 20.00\n%v nota(s) de R$ 10.00\n%v nota(s) de R$ 5.00\n%v nota(s) de R$ 2.00\n", N/100, (N%100)/50, ((N%100)%50)/20, ((N%50)%20)/10, ((N%20)%10)/5, ((N%10)%5)/2)
	fmt.Printf("MOEDAS:\n%v moeda(s) de R$ 1.00\n%v moeda(s) de R$ 0.50\n%v moeda(s) de R$ 0.25\n%v moeda(s) de R$ 0.10\n%v moeda(s) de R$ 0.05\n%v moeda(s) de R$ 0.01\n", ((N%5)%2)/1, C/50, (C%50)/25, ((C%50)%25)/10, ((C%25)%10)/5, ((C%10)%5)/1)
}
