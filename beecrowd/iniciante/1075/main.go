package main

import "fmt"

func main() {
	var A int
	fmt.Scan(&A)
	for i := 1; i <= 10000; i++ {
		if i%A == 2 {
			fmt.Println(i)
		}
	}
}
