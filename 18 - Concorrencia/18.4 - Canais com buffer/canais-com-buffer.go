package main

import "fmt"

func main() {
	canal := make(chan string, 2) //não bloqueia a execução, somente quando atigir sua cap max de 2

	canal <- "Olá"
	canal <- "Mundo"

	msg1 := <-canal
	msg2 := <-canal

	fmt.Println(msg1)
	fmt.Println(msg2)

}
