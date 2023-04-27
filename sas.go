package main

import "fmt"

func findMaxMinAvg(numbers []int) (int, int, float64) {
	var max, min, sum int

	// Inicializa o valor máximo e mínimo como o primeiro elemento da lista
	max = numbers[0]
	min = numbers[0]

	// Soma todos os elementos da lista
	for _, number := range numbers {
		sum += number

		// Verifica se o número atual é maior que o máximo
		if number > max {
			max = number
		}

		// Verifica se o número atual é menor que o mínimo
		if number < min {
			min = number
		}
	}

	// Calcula a média e retorna os valores máximo, mínimo e média
	avg := float64(sum) / float64(len(numbers))
	return max, min, avg
}

func main() {
	numbers := []int{10, -2, 5, -8, 3}
	max, min, avg := findMaxMinAvg(numbers)
	fmt.Printf("Valor máximo: %d\n", max)
	fmt.Printf("Valor mínimo: %d\n", min)
	fmt.Printf("Média: %.2f\n", avg)
}
