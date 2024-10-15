package main

import "fmt"

/*
fatorial é uma fórmula matemática representada pelo sinal de exclamação “!”
Exemplos:

	7! = 1 * 2 * 3 * 4 * 5 * 6 * 7 = 5.040
	5! = 1 * 2 * 3 * 4 * 5 = 120
*/
// uint não aceita valor negativo
func fatorial_simples(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial_simples(n-1)
	}
}

func main() {
	// valpor válido
	resultado := fatorial_simples(5) // fatorial é a multiplicação  5 * 4 * 3 * 2 * 1
	fmt.Println(resultado)

	resultado2 := fatorial_simples(0)
	fmt.Println(resultado2)

}
