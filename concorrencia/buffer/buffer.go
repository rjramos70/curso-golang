package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou!") // como o buffer é de 3, a linha 4 não será executada,
	// e nem esse linha do print, até que seja consumido/lido pelo menos uma linha do buffer.
	ch <- 5
	ch <- 6

}

func main() {
	// cria o cana com um buffer de 3
	ch := make(chan int, 3)

	// criar a goroutine
	go rotina(ch)

	// faz a função aguardar 1 segundo
	time.Sleep(time.Second)

	// ler o primeiro dado do canal
	fmt.Println(<-ch)

}
