package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Jean",
		"sobrenome": "Leal",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	delete(usuario, "nome")
}
