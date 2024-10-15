package main

import "fmt"

/*
fatorial é uma fórmula matemática representada pelo sinal de exclamação “!”
Exemplos:

	7! = 1 * 2 * 3 * 4 * 5 * 6 * 7 = 5.040
	5! = 1 * 2 * 3 * 4 * 5 = 120
*/
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	// valpor válido
	resultado, err1 := fatorial(5) // fatorial é a multiplicação  5 * 4 * 3 * 2 * 1
	fmt.Println(resultado, err1)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	resultado2, _ := fatorial(0)
	fmt.Println(resultado2)

	// Uma solução melhor seria...

}
