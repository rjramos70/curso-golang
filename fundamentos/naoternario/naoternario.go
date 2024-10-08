package main

import "fmt"

/*
Em Go NÃO tem operador Ternário
*/
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	// Se tivesse operador ternário seria como abaixo
	// return nota >= 6 ? "Aprovado" : "Reprovado"
}

func main() {
	resultado := obterResultado(6.2)
	fmt.Println("Resultado:", resultado)
}
