package main

import "fmt"

func main() {
	var A, B, C, D, E, F, CC, DD float64
	fmt.Scan(&A, &B, &C, &D, &E, &F)

	arr := []float64{A, B, C, D, E, F}
	for _, s := range arr {
		if s > 0 {
			DD += s
			CC++
		}
	}
	fmt.Printf("%.f valores positivos\n", CC)
	fmt.Printf("%.1f\n", DD/CC)
}
