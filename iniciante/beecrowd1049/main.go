package main

import (
	"fmt"
)

func main() {
	var A, B, C string
	fmt.Scan(&A, &B, &C)

	animais := map[string]map[string]map[string]string{
		"vertebrado": {
			"ave": {
				"carnivoro": "aguia",
				"onivoro":   "pomba",
			},
			"mamifero": {
				"herbivoro": "vaca",
				"onivoro":   "homem",
			},
		},
		"invertebrado": {
			"inseto": {
				"hematofago": "pulga",
				"herbivoro":  "lagarta",
			},
			"anelideo": {
				"hematofago": "sanguessuga",
				"onivoro":    "minhoca",
			},
		},
	}

	fmt.Println(animais[A][B][C])
}
