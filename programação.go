package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "arara"
	nome = strings.ToLower(strings.ReplaceAll(nome, " ", ""))
	for i := 0; i < len(nome)/2; i++ {
		if nome[i] != nome[len(nome)-i-1] {
			fmt.Println("não é um palíndromo")
			return
		}
	}
	fmt.Println("palindromo")
}
