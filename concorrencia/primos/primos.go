package main

import (
	"fmt"
	"time"
)

/*
Função que recebe um número e verifica os números primos
existentes de 2 até o número recebido como parâmetro
*/
func isPrimo(num int) bool {
	// verificar todos os números até o número recebido
	for i := 2; i < num; i++ {
		// verifica se cada número não é Primo, não divisível por 'i'
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Função que recebe a quantidade de números primos que eu quero que seja
// retornados nessa função. Os número primos vão ser enviado para o Channel.
func primos(n int, c chan int) {
	fmt.Printf("Qtd. Números Primos: %d\n", n)
	inicio := 2 // apartir de qual número eu vou começar a procurar os números primos.
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 180) // aguarda 100 milesegundos
				break                              // se encontrou um primo ele vai sair do laço
			}
		}
	}
	close(c) // fecha o canal para terminar o laço FOR
}

func main() {
	// cria um canal com um buffer de 30 posições
	c := make(chan int, 30)

	// go primos(cap(c), c)	// passando o tamanho da capacidade do canal
	go primos(60, c) // Os 60 primeiros números primos

	// Iterage sobre o canal
	fmt.Println("Iteragindo com os valores do canal:")
	count := 1
	for primo := range c {
		fmt.Printf("\tPrimo[%d]: %d\n", count, primo)
		count++
	}
	fmt.Println("Fim......")

}
