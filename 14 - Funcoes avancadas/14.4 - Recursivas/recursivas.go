package main

import "fmt"

func fibo(pos uint) uint {
	if pos <= 1 {
		return pos
	}

	return fibo(pos-2) + fibo(pos-1)
}

func main() {
	//0 1 1 2 3 5 8 13 21
	const posicao uint = 5
	fmt.Println(fibo(posicao))
}
