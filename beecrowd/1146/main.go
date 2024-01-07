package main

import "fmt"

func main() {
	var A int
	for {
		fmt.Scan(&A)

		if A == 0 {
			return
		} else {
			for i := 1; i <= A; i++ {
				if A == i {
					fmt.Printf("%v\n", i)
				} else {
					fmt.Printf("%v ", i)
				}
			}
		}
	}
}
