package main

import "fmt"

func soma(numeros ...int) int {
	soma := 0
	for _, v := range numeros {
		soma += v
	}
	return soma
}

func escrever(texto string, numeros ...int) {
	for _, v := range numeros {
		fmt.Println(texto, v)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5))
	escrever("olá mundo", 3, 5, 10)
}
