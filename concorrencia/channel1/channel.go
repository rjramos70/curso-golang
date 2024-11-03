package main

import "fmt"

func main() {
	// criando um channel
	ch := make(chan int, 1)

	// enviando dados para o channel (processo de escrita)
	ch <- 1

	// recebendo dados de um channel (processo de leitura)
	<-ch

	// enviando novo dado para o channel (processo de escrita)
	ch <- 2

	// recebendo e imprimindo dados do channel (leitura)
	fmt.Println(<-ch) // Output: 2
}
