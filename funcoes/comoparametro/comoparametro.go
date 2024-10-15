package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func divisao(a, b int) int {
	return a / b
}

func soma(a, b int) int {
	return a + b
}

/*
função que recebe uma função como parâmetro com dois
parâmetros de entrada e como retorno um int, e mais
dois parâmetros 'p1' e 'p2'
*/
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado1 := exec(multiplicacao, 3, 4)
	fmt.Println(resultado1) // 12

	resultado2 := exec(divisao, 30, 6)
	fmt.Println(resultado2) // 5

	resultado3 := exec(soma, 17, 21)
	fmt.Println(resultado3) // 38
}
