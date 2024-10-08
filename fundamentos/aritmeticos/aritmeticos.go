package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtração => ", a-b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Multiplicação => ", a*b)
	fmt.Println("Módulo => ", a%b)

	// bitwise (operação bit a bit)
	//
	// 0	0	0	0	0	0	0	0	0
	// 256	126	64	32	16	8	4	2	1
	fmt.Println("E => ", a&b)   // 11 & 10 (valor binário de 3 = 00000011 e 2 = 00000010 que é 00000010 --> logo igual a 2 inteiro)
	fmt.Println("Ou => ", a|b)  // 11 | 10 (valor binário de 3 = 00000011 ou 2 = 00000010 que é 00000011 --> logo igual a 3 inteiro)
	fmt.Println("Xor => ", a^b) // 11 ^ 10 (valor binário de 3 = 00000011 ^ 2 = 00000010 que é 00000001 --> logo igual a 1 inteiro)

	// outras operacoes usando math...
	c := 3.0
	d := 2.0

	fmt.Println("Maior entre a e b => ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor entre c e d => ", math.Min(c, d))
	fmt.Println("Exponenciação de c elevado à d => ", math.Pow(c, d))
}
