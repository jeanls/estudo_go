package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64 or only int (based on your pc architecture)
	// uint -> unsigned int

	const numero int64 = 10000000000
	fmt.Println(numero)

	//alias
	//int32 = rune
	var numero2 rune = 123456
	fmt.Println(numero2)

	var numero4 byte = 123
	fmt.Println(numero4)

	//float32, float64, bool

	var erro error = errors.New("deu ruim")
	fmt.Println(erro)
}
