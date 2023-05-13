package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	// endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var user1 usuario
	user1.idade = 29
	user1.nome = "Jean"

	fmt.Println(user1.idade)

	user2 := usuario{"Jean", 30}
	fmt.Println(user2.nome)

	user3 := usuario{nome: "Jean"}
	fmt.Println(user3)
}
