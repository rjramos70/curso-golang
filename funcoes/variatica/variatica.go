package main

import "fmt"

/*
Função variática é o conceito que a função pode receber um ou N
parâmetros na mesma assinatura do método.
*/
func media(numeros ...float64) float64 {
	total := 0.0
	// verifica se o array numeros está vazio e retorna 0.0
	if len(numeros) == 0 {
		return total
	}
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média 1: %.2f\n", media(7.7, 8.1, 5.9))

	fmt.Printf("Média 2: %.2f\n", media(7.7, 8.1, 5.9, 9.9))

	fmt.Printf("Média 3: %.2f\n", media(7.7))

	fmt.Printf("Média 4: %.2f\n", media())
}
