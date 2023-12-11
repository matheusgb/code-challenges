package main

import (
	"fmt"
	"sort"
)

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	arr := []int{A, B, C}
	sort.Ints(arr)
	fmt.Printf("%v eh o maior\n", arr[2])
}
