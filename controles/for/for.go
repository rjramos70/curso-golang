package main

import (
	"fmt"
	"time"
)

/*
Formar de usar o laço FOR em Go.
Não tem WHILE em Go
*/
func main() {
	// simulando um while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// FOR tradicional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando.....")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("[%d] Par \n", i)
		} else {
			fmt.Printf("[%d] Impar \n", i)
		}
	}

	// laço FOR infinito...
	for {
		fmt.Println("Para sempre...", time.Now())
		time.Sleep(time.Second * 2)
	}
}
