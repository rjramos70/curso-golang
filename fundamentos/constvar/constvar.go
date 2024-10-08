package main

import (
	"fmt"
	m "math" // pode-se usar um alias para o pacote importado, neste caso 'm' para math
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2) // raio ao quadrado

	fmt.Println("A área da circunferência é", area)

	// Outro modo de declarar constantes e variaveis
	const (
		A = 1
		B = 2
	)

	var (
		C = 3
		D = 4
	)

	fmt.Println(A, B, C, D)

	// declarando mais de uma variavel e atribuindo valores
	var e, f bool = true, false

	fmt.Println(e, f)

	// declarando e atribuindo variaveis atribuindo valores usando sintaxe reduzida
	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
