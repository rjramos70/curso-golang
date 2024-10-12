package main

import "fmt"

func main() {
	// cria um array
	array1 := [3]int{1, 2, 3}
	// cria um slice sem atribuir valores
	var slice1 []int

	// tenta fazer um append de valores do array ja com as 3 posições preenchidas
	// inserindo os valores do slice
	// array1 = append([3]int(slice1), 4, 5, 6)	// Error: first argument to append must be a slice; have [3]int(slice1) (value of type [3]int)compiler (InvalidAppend)

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, len(array1), cap(array1), slice1, len(slice1), cap(slice1)) // [1 2 3] 3 3 [4 5 6] 3 3

	// cria novo slice usando o make()
	slice2 := make([]int, 2)
	fmt.Println(slice2, len(slice2), cap(slice2))

	// copia tudo do slice2 para o slice1
	copy(slice2, slice1) // copy só funciona passando um slice, e não um array como atributo slice1
	fmt.Println(slice2, len(slice2), cap(slice2))

}
