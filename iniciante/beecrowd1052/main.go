package main

import (
	"fmt"
	"time"
)

func main() {
	var A int
	fmt.Scan(&A)

	fmt.Printf("%s\n", time.Month(A))
}
