package main

import "fmt"

func main() {
	var A, C int
	fmt.Scan(&A)
	for {
		if A%2 != 0 {
			C++
			fmt.Println(A)
		}
		A++
		if C == 6 {
			break
		}
	}
}
