package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", 5, canal)

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}

	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, length int, canal chan string) {
	for i := 0; i < int(length); i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
