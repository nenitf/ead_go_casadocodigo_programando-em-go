package main

import (
	"fmt"
)

// segunda função a ser chamada, com retorno
func podeBeber(idade int) string {
	var msg string
	if idade >= 18 {
		msg = "Pode beber\n"
	} else {
		msg = "Não pode beber\n"
	}
	return msg
}

// script inicia aqui
func main() {
	nome := "Felipe"
	idade := 22
	imprimir(nome, idade)
	msg := podeBeber(idade)
	fmt.Printf(msg)
}

// primeira função a ser chamada, void
func imprimir(nome string, idade int) {
	fmt.Printf("Nome: %s. Idade: %d\n",
		nome, idade)
}
