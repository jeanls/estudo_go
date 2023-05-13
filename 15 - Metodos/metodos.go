package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando usuario %s\n", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	user := usuario{"Jean", 1}
	user.salvar()
	user.fazerAniversario()

	fmt.Println(user)
}
