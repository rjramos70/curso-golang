package main

import "fmt"

func main() {
	// cria um slice com valores defauts com tamanho definido de 10 elementos
	s := make([]int, 10)
	s[9] = 12      // seta o útimo elemento para 12
	fmt.Println(s) // [0 0 0 0 0 0 0 0 0 12]

	// reatribuindo um novo slice de inteiros 10 elementos com um array interno de 20 elementos
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))              // cap(s) -> pega a capacidade máxima do slice
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // anexa mais 10 elementos ao slice
	fmt.Println(s)                              // [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0]

	s = append(s, 11, 12, 13, 14, 15, 16)
	fmt.Println(s, len(s), cap(s)) // [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0 11 12 13 14 15 16] 26 40

}
