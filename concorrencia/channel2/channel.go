package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// Channel é um tipo (tal qual boolean, int, string, entre outros...)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second) // espera 1 segundo
	c <- 2 * base           // multiplica base por 2 e envia para o channel (enquanto esse valor não for lido, o processo não segue para as próximas linhas)

	time.Sleep(time.Second) // espera mais 1 segundo
	c <- 3 * base           // multiplica base por 3 e envia para o channel (enquanto esse valor não for lido, o processo não segue para as próximas linhas)

	time.Sleep(time.Second * 3) // espera mais 3 segundo
	c <- 4 * base               // multiplica base por 4 e envia para o channel (enquanto esse valor não for lido, o processo não finaliza)

}

func main() {
	// criando um channel
	c := make(chan int)

	// chama a função passando um valor base e o canal
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	// lê os dois primeiros valores do canal
	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println("B")
	fmt.Println(a, b) // 4 6

	// recebendo o último valor do canal
	fmt.Println(<-c) // 8

	// Sequência do output:
	// A
	// B
	// 4 6
	// 8

	// fmt.Println(<-c) // como todas as goroutines já foram lidas/consumidas
	// o canal vai entrar em deadlock, porque não tem mais mensagens para serem lidas
	// e foi chamado uma rotina para ler do canal que esta vazio, assim é gerado a mensagem de erro:
	// fatal error: all goroutines are asleep - deadlock!

}
