package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// criar um file server passando o diretorio que ele vai ler
	fs := http.FileServer(http.Dir("public"))

	// quando receber uma requisição na raiz '/', passe para o file server 'fs'
	http.Handle("/", fs)

	// mostra uma mensagem so para sabermos que o servidor foi iniciado
	fmt.Println("Executando...")

	// Inicia o servidor não passando o endereço, mas passando só a porta 3000
	log.Fatal(http.ListenAndServe(":3000", nil))
}
