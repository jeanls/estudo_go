package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("olá mundo")
	escrever("go go go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
