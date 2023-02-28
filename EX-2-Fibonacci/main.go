package main

import "fmt"

func main() {
	numeroInserido := NumeroEscolhido()
	fmt.Println(Fibonacci(numeroInserido))
}

func Fibonacci(numeroInserido int) string {

	if numeroInserido == 1 {
		return "O número pertence à sequência de Fibonacci"
	}
	var primeiroNumero = 1
	var ultimoNumero = 1
	for i := 0; i != numeroInserido; i++ {
		soma := primeiroNumero + ultimoNumero
		ultimoNumero = primeiroNumero
		primeiroNumero = soma
		if soma == numeroInserido {

			return "O número pertence à sequência de Fibonacci"

		} else if soma > numeroInserido {
			break
		}
	}
	return "O número não pertence à sequência de Fibonacci"
}

func NumeroEscolhido() int {
	fmt.Println("Insira um número inteiro para verificar se o mesmo pertence à sequência de Fibonacci")
	var numeroInserido int
	fmt.Scan(&numeroInserido)
	return numeroInserido
}
