package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)...
	// 	uint8  = 1 byte
	// 	uint16 = 2 bytes (1 short)
	// 	uint32 = 4 bytes (1 int)
	// 	uint64 = 8 bytes (1 long)
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal (só positivos)...
	// 	int8  = 1 byte
	// 	int16 = 2 bytes (1 short)
	// 	int32 = 4 bytes (1 int)
	// 	int64 = 8 bytes (1 long)
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O valor da letra 'a' na tabela Unicode é", i2)

	// números reais, números com ponto flutuante (float32, float64)
	var x float32 = 49.99
	var x2 = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de x2 é", reflect.TypeOf(x2))
	fmt.Println("O tipo literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println("Negação da variável bo é", !bo)

	// string
	s1 := "Olá meu nome é Fulano"
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string s1 é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Léo`
	fmt.Println("O tamanho da string s2 é", len(s2))

	// char???
	// var x char = 'b'	// O tipo char não existe em Go
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))

}
