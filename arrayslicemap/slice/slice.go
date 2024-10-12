package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Printf("Array: %v - type: %v\n", a1, reflect.TypeOf(a1))
	fmt.Printf("Slice: %v - type: %v\n", s1, reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5} // Um array

	// Slice não é um array! Slice define um pedaço de um array.
	s2 := a2[1:3] // Slice do Array do índice 1 até o índice anterior a 3
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, pega desde o início do array até o índice anterior a 2, mas apontando para o mesmo array
	fmt.Println(a2, s3)

	// Você pode imaginar um slice como: tem um tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1] // Slice de um slice
	fmt.Println(a2, s2, s4)
}
