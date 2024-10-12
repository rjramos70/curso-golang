package main

import "fmt"

func main() {
	// Array são estruturas homogêneas (todos os elementos são do mesmo tipo)
	// e estática (fixo), do mesmo tamanho que foi definido na construção.

	var notas [3]float64 // array de 3 posições do tipo float64
	fmt.Println(notas)

	// atribuíndo valores a cada elemento do array na mesma linha.
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	// calculando a média usando o FOR clássico
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média: %.2f\n", media) // imprimir médica mostrando somente 2 casas decimais.

}
