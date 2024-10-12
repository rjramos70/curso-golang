package main

import "fmt"

func main() {
	// cria um map de map
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}
	fmt.Println(funcsPorLetra)

	// removendo todos os funcionário que iniciam com a letra 'P'
	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	// iteraje sobre o map imprimindo a letra e o funcionario
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)                 // imprimi a letra
		for nome, salario := range funcs { // para cada letra, imprimir nome e salario do funcionário
			fmt.Printf("\t%s\t- %.2f\n", nome, salario)
		}
	}
}
