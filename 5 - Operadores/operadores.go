package main

import "fmt"

func main() {
	//Aritméticos
	soma := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	mult := 2 * 2
	resto := 10 % 2

	fmt.Println(soma, sub, div, mult, resto)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//Atribuição
	var var1 string = "string"
	var2 := "string2"

	fmt.Println(var1, var2)

	//Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)

	//Logicos

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//Unarios
	numero := 10
	numero++
	fmt.Println(numero)
	numero--
	numero += 10

}
