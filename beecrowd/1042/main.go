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
	fmt.Printf("%v\n%v\n%v\n\n%v\n%v\n%v\n", arr[0], arr[1], arr[2], A, B, C)
}
