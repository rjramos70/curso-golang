package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 1; i <= qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i)
	}
}

func main() {
	// Processamento sequenciado
	// 	fale("Maria", "Pq vc não fala comigo?", 3)
	// 	fale("João", "Só posso falar depois de vc!", 1)

	// Cria uma Goroutine, executa cada linha de forma independente.
	// Goroutine só é executada enquanto a thread/função principal estiver rodando.
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	//
	go fale("Maria", "Enetndi!!!", 10)
	fale("João", "Parabéns!", 5)

}
