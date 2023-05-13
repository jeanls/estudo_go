package main

import "fmt"

func main() {
	var array1 [5]int
	fmt.Println(array1)

	array2 := [3]string{"", "", ""}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4}
	fmt.Println(array3)

	//slice

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4)

	fmt.Println(slice)

	fmt.Println(slice[1:2])

	// arrays internos

	slice2 := make([]float32, 10, 15)

	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
