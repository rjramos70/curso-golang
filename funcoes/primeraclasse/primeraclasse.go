package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3)) // 5

	// cria uma função atribuíndo a uma variável
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(10, 3)) // 7
}
