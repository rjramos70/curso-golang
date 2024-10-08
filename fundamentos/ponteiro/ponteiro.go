package main

import "fmt"

/*
Ponteiro não nada a mais que uma referência para um
espaço/endereço na memória. Sempre que você mudar o
espaço na memória que esteja atrelado a um ponteiro,
pelo ponteiro se consegue chegar a esse novo espaço
em outro eendereço de memória.
*/
func main() {
	i := 1

	var p *int = nil // cria um ponteiro do tipo int que recebe o valor nulo (ni).

	p = &i // pegue o endereço de memória da variável 'i' e atribua ao ponteiro.
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não tem aritmética de ponteiro
	// p++

	fmt.Println(i, p, *p, &i)
}
