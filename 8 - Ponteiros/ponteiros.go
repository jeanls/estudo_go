package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++

	fmt.Println(var1, var2)

	//ponteiro
	var ponteiro *int
	var var3 int = 100
	fmt.Println(ponteiro, var3)

	ponteiro = &var3

	fmt.Println(ponteiro, var3)
	fmt.Println(*ponteiro, var3)
}
