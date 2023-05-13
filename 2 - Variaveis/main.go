package main

import "fmt"

func main() {
	var variavel1 string = "var 1"
	fmt.Println(variavel1)

	variavel2 := "var 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "212"
		variavel4 string = "32313"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "1212", "21212"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "const 1"

	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
}
