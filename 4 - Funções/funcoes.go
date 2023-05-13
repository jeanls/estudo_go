package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	soma := somar(1, 2)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("eu sou uma funcao")

	result1, result2 := calculosMatematicos(10, 5)
	result3, _ := calculosMatematicos(10, 5)

	fmt.Println(result1, result2)
	fmt.Println(result3)
}
