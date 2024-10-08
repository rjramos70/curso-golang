package main

import "fmt"

func main() {
	// atribuição simples
	var b byte = 3
	fmt.Println(b)

	i := 3 // define e inicializa a variável
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	// define e atribui mais de uma variável na mesma linha
	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

}
