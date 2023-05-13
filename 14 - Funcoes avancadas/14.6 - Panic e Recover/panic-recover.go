package main

import "fmt"

func recuperarExec() {
	if r := recover(); r != nil {
		fmt.Println("execução recuperada", r)
	}

	// fmt.Println("Tentativa de recuperar a execução")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExec()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic(media) // chama o que tiver o defer antes
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	// fmt.Println("Pós exec.")
}
