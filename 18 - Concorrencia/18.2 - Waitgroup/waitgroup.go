package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("olá mundo", 5)
		waitGroup.Done()
	}()

	go func() {
		escrever("go go go", 10)
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string, length int) {
	for i := 0; i < int(length); i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
