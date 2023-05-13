package main

import "fmt"

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que 0")
	}
}
