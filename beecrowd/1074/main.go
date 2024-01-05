package main

import "fmt"

func main() {
	var A, C int
	fmt.Scan(&A)
	for i := 0; i < A; i++ {
		fmt.Scan(&C)
		if C == 0 {
			fmt.Println("NULL")
			continue
		}

		if C%2 != 0 {
			fmt.Printf("ODD ")
		} else {
			fmt.Printf("EVEN ")
		}

		if C > 0 {
			fmt.Printf("POSITIVE\n")
		} else {
			fmt.Printf("NEGATIVE\n")
		}
	}
}
