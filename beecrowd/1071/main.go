package main

import "fmt"

func main() {
	var A, B, C int

	fmt.Scan(&A, &B)
	for i := A - 1; i > B; i-- {
		if i%2 != 0 {
			C += i
		}
	}
	fmt.Println(C)
}
