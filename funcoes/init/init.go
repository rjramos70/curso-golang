package main

import "fmt"

/*
Função init() é executada antes mesmo do que a função main()
*/
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
