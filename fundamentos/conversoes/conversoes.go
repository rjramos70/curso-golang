package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // tipo float64
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(123)) // 123 será o valor correspondente na tabela ASC, que é {
	fmt.Println("Teste " + string(97))  // 97 será o valor correspondente na tabela ASC, que é a

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123)) // Itoa, converte o valor inteiro 123 para uma string "123"

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string para boolean
	// b, _ := strconv.ParseBool("true")	// OR
	b, _ := strconv.ParseBool("1") // "1" = true; "0" = false
	if b {
		fmt.Println("Verdadeiro")
	}

}
