package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	for i < 10 {
		fmt.Println(i)
		time.Sleep(10)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	nomes := [3]string{"jean", "cecilia", "jessica"}

	for k, v := range nomes {
		fmt.Println(k, v)
	}

	for _, v := range nomes {
		fmt.Println(v)
	}

	for k, v := range "jean" {
		fmt.Println(k, v)
	}

	for k, v := range "jean" {
		fmt.Println(k, string(v))
	}

	for {
		fmt.Println("executando infinitamente")
	}
}
