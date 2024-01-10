package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// não aceito pelo avaliador do beecrowd
	var A int
	fmt.Scan(&A)

	for i := 0; i < A; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Bytes()

		for i := 0; i < len(line); i++ {
			if line[i] >= 'A' && line[i] <= 'Z' || line[i] >= 'a' && line[i] <= 'z' {
				line[i] = line[i] + 3
			}
		}

		for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
			line[i], line[j] = line[j], line[i]
		}

		for i, ll := range line[len(line)/2:] {
			line[len(line)/2:][i] = ll - 1
		}

		fmt.Println(string(line))
	}
}
