package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A int
	fmt.Scan(&A)

	for i := 0; i < A; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		fmt.Println("I am Toorg!")
	}
}
