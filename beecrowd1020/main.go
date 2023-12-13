package main

import (
	"fmt"
)

func main() {
	var D, aa, mm, dd int
	fmt.Scan(&D)

	for D > 0 {
		if D >= 365 {
			D = D - 365
			aa++
		} else if D >= 30 {
			D = D - 30
			mm++
		} else {
			D = D - 1
			dd++
		}
	}
	fmt.Printf("%v ano(s)\n%v mes(es)\n%v dia(s)\n", aa, mm, dd)
}
