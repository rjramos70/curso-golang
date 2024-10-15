package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// cria um slice (sem tamanho definído)
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}

	imprimirAprovados(aprovados...) // Os ... significar que estamos abrindo o slice, o que torna um array

}
