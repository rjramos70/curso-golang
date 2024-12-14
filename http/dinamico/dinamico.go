package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	// cria uma variavel que recebe a hora atual já formatando
	// onde:
	// 02 = dia
	// 01 = mês
	// 2006 = ano com 4 dígito
	// 03 = hora
	// 04 = minuto
	// 05 = segundo
	s := time.Now().Format("02/01/2006 03:04:05")

	// receber um writer como parametro e escreve nele passando a data formatada
	fmt.Fprintf(w, "<h1>Hora certa: %s<h1>", s)

}

func main() {
	// cria a ur de um endpoint que será responsável por chamar essa função
	http.HandleFunc("/horaCerta", horaCerta)

	// mostra uma mensagem so para sabermos que o servidor foi iniciado
	fmt.Println("Executando...")

	// Inicia o servidor não passando o endereço, mas passando só a porta 3000
	log.Fatal(http.ListenAndServe(":3000", nil))

}
