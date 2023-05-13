package main

import "fmt"

func init() {
	fmt.Println("Executando a função init") // executa antes do main
}

func main() {
	fmt.Println("Executando a função main")
}
