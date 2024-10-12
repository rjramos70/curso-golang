package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	fmt.Println(numeros)
	fmt.Println(len(numeros))

	// pega o índice e o número
	for indice, numero := range numeros {
		fmt.Printf("%d) %d\n", indice, numero)
	}

	// ignora/não define, o índice e só pega o número
	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}

	// se só definir uma variável, esta será o índice
	for n := range numeros {
		fmt.Printf("- %d\n", n)
	}

}
