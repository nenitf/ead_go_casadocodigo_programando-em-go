package main

/*
go run p11-if/conversor.go 32 27.4 -3 0 celsius

resultado:
32.00 celsius = 89.60 fahrenheit
27.40 celsius = 81.32 fahrenheit
-3.00 celsius = 26.60 fahrenheit
0.00 celsius = 32.00 fahrenheit
*/

// declaração de pacotes externos
import (
	"fmt"     //formatação de strings para o print
	"os"      //sistema operacional, inputs
	"strconv" //conversor de strings para decimais
)

// inicio
func main() {
	// são necesssário no minimo 3 parametros: conversor.go, valor e unidade
	if len(os.Args) < 3 {
		fmt.Println("uso correto: go run path/conversor.go <valores> <unidade>\nunidade = celsius ou quilometros")
		os.Exit(1)
	}

	// descobre pelo ultimo valor no array a unidade -> conversor.go[0] <valor>[1] <valor>[2]... <unidade>[total-1]
	unidadeOrigem := os.Args[len(os.Args)-1]

	// fatia array os valores entre a segunda posição até a penultima
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	// percore valores origem, com retorno duplo da func range. A var i recebe a chave e o v o valor da posição
	for i, v := range valoresOrigem {

		// retorno duplo, sendo o resultado para valorOrigem e erro para err
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {

			// %s string e %d inteiro
			fmt.Printf(
				"O valor %s na posição %d não está correto!\n",
				v, i)
			os.Exit(1)
		}

		// unidade destino para o printf, deve estar no escopo globa da main
		var unidadeDestino string

		// valor que recebe calculo, deve estar no escopo globa da main
		var valorDestino float64

		// testa tipo e calcula a situação
		if unidadeOrigem == "celsius" {
			unidadeDestino = "fahrenheit"
			valorDestino = valorOrigem*1.8 + 32
		} else if unidadeOrigem == "quilometros" {
			unidadeDestino = "milhas"
			valorDestino = valorOrigem / 1.60934
		} else {
			fmt.Printf("%s não é uma unidade conhecida!", unidadeDestino)
			os.Exit(1)
		}

		// %.2f float com duas casas decimais
		fmt.Printf("%.2f %s = %.2f %s\n",
			valorOrigem, unidadeOrigem,
			valorDestino, unidadeDestino)
	}
}
