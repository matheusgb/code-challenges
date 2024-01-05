package main

import "fmt"

func main() {
	var A, m, c int
	for i := 0; i < 3; i++ {
		fmt.Scan(&A)
		if A > m {
			m = A
			c = i + 1
		}
	}
	fmt.Println(m)
	fmt.Println(c)
}
