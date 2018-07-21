package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// entrada recebe parametros  => go run quicksort.go <valor> <valor> ...
	entrada := os.Args[1:]

	// cria slice vazio para os numeros
	numeros := make([]int, len(entrada))

	// itera por todos argumentos de entrada
	// range retorna posição (i) e valor (n)
	for i, n := range entrada {
		// converte string para inteiro
		numero, err := strconv.Atoi(n)
		// testa erro
		if err != nil {
			fmt.Printf("%s não é um numero válido!\n", n)
			os.Exit(1)
		}
		// slice vazio recebe numero
		numeros[i] = numero
	}
	// printa valor da função quicksort
	fmt.Println(quicksort())
}

// func nome(parametro tipo) retorno
func quicksort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	// cria slice vazio para os numeros
	n := make([int], len(numeros))

	// copia os valores dos numeros para o slice vazio
	copy(n, numeros)

	// excolhe valor na metade do slice
	indicePivo := len(n)/2
	pivo := n[indicePivo]

	//
	novoSlice := append(slice, value)
	return 1
}
