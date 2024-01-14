package main

import "fmt"

func main() {
	var A int
	fmt.Scan(&A)
	for i := 1; i <= A; i++ {
		if i == A {
			fmt.Println("Ho!")
		} else {
			fmt.Printf("Ho ")
		}
	}
}
