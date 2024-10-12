package main

import "fmt"

func main() {
	// cria uma map chave tipo int com valor tipo string
	// var aprovados map[int]string	// map deve ser inicializado antes de inserirmos qualquer elemento
	aprovados := make(map[int]string)

	// popula o map
	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"

	fmt.Println(aprovados) // map[12345678978:Maria 95135745682:Ana 98765432100:Pedro]

	// laço para percorrer o map com chave e valor
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// imprimindo somente as chaves
	for _, nome := range aprovados {
		fmt.Printf("Nome: %s\n", nome)
	}

	// imprimindo somente os valores
	for cpf, _ := range aprovados {
		fmt.Printf("CPF: %d\n", cpf)
	}

	// com base na chave retorna o valor
	fmt.Println(aprovados[95135745682]) // Ana

	// excluir um elemento, passar o map e a respectiva chave
	delete(aprovados, 95135745682) // remove a Ana que tem o CPF 95135745682
	fmt.Println(aprovados)         // map[12345678978:Maria 98765432100:Pedro]

	// atualizando o valor de um determinado elemento
	aprovados[98765432100] = "Pedro Vasconselos"
	fmt.Println(aprovados) // map[12345678978:Maria 98765432100:Pedro Vasconselos]
}
