package main

import (
	"fmt"
	"time"
)

// Função Generator, que com base em um string de entrada, retorna um canal
func falar(pessoa string) <-chan string {
	// cria um canal
	c := make(chan string)

	// cria um função anônima para criar a goroutine
	go func() {
		for i := 0; i < 3; i++ {
			// faz uma pausa de um segundo antes de enviar dado para o canal
			time.Sleep(time.Second)
			// envia a texto para o canal
			c <- fmt.Sprintf("%s - falando: %d", pessoa, i)
		}
	}() // como é uma função anônima, precisamos invacar essa goroutine, para isso os ()
	// retorna o canal
	return c
}

// Função responsável por juntar/multiplexar mais de um canal, retornando um canal multiplexado
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	// cria um canal que será retornado como resultado
	c := make(chan string)

	// cria uma função anônima que pega os canais de entrada, junta/multiplexa e retorna o canal
	go func() {
		// cria um FOR infinito
		for {
			// cria um estrutura select para classificar de qual canal chegar e insere no canal de saída
			select {
			// caso chegue dados na entrada1, enviar para o canal de saída
			case s := <-entrada1:
				// canal 'c' recebe o valor 's'
				c <- s
				// caso chegue dados na entrada2, enviar para o canal de saída
			case s := <-entrada2:
				// canal 'c' recebe o valor 's'
				c <- s
			}
		}
	}() // invoca a função anônima

	// retorna o canal multiplexado
	return c
}

func main() {
	// cria u canal passando duas falas
	c := juntar(falar("João"), falar("Maria"))

	// imprime cada respectivo dado de cada canal
	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)

}
