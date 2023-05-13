package main

import "fmt"

func closure() func() {
	texto := "Dentro da função"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	// texto := "Oláaaa"
	funcaoNova := closure()
	funcaoNova()
}
