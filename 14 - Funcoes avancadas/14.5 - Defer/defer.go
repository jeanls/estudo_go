package main

import "fmt"

func funcao1() {
	fmt.Println("Exec func 1")
}

func funcao2() {
	fmt.Println("Exec func 2")
}

func media(n1, n2 float32) bool {
	defer fmt.Println("Média calculada com sucesso.")
	fmt.Println("iniciando calculo da média.")

	media := n1 + n2/2

	if media > 6 {
		return true
	}
	return false
}

func main() {
	// defer funcao1() // adia a execução da função até o último momento possível
	// funcao2()

	media(5, 6)
}
