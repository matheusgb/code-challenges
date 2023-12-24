package main

import "fmt"

func main() {
	var A int
	fmt.Scan(&A)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = %v\n", i, A, i*A)
	}
}
