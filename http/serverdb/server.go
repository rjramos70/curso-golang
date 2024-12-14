package main

import (
	"log"
	"net/http"
)

/*
Classe main do projeto para podermos executar o nosso web server cliente.go
*/

func main() {
	// faz o direcionamento do request para a função UsuarioHandler
	http.HandleFunc("/usuarios/", UsuarioHandler)

	// cria um log para mostrar que nosso servidor esta rodando
	log.Println("Executando server...")

	// define a porta que o servidor estará rodando
	log.Fatal(http.ListenAndServe(":3000", nil))

}
