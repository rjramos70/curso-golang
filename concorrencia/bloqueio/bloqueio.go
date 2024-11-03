package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante, só irá continuar quando alguem consumir o dado do canal
	fmt.Println("Só depois que for lido...")
}

func main() {
	// cria um canl sem buffer
	c := make(chan int)

	// cria a goroutine
	go rotina(c)

	// receber/ler o dado, que também é uma opração bloqueante
	fmt.Println(<-c)

	// função só vai ser chamada depois que o dado for lido
	fmt.Println("Dado já foi lido...")

	// Tentar ler o dado do canal que já esta vazio nesse ponto
	fmt.Println(<-c) // gera o deadlock (fatal error: all goroutines are asleep - deadlock!)

	// imprimi ao fim do processo
	fmt.Println("Fim")

}
