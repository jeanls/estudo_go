package main

import "fmt"

func main() {
	func(txt string) {
		fmt.Println(txt)
	}("Olá mundo")

	concatFunc := func(txt string) string {
		return fmt.Sprintf("Valor concatenado é.: %s", txt)
	}("Olá mundo")

	fmt.Println(concatFunc)
}
