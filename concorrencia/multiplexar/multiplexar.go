package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Titulo obtem o título de uma página html com base na lista de urls passadas como parâmetro.
func titulo(urls ...string) <-chan string {
	// cria o canal
	c := make(chan string)
	// iterage na lista de urls ignorando a índice, pegando somente o valor de
	// cada url recebida
	for _, url := range urls {
		// função anônima que vai gerar a Goroutine
		go func(url string) {
			// faz requisiçõ http para cada url ignorando o erro
			resp, _ := http.Get(url)
			// vamos extrair o body da resposta
			// html, _ := ioutil.ReadAll(resp.Body)	// ate Go 1.16
			html, _ := io.ReadAll(resp.Body) // após Go 1.16

			// vamos transformar o html do bady em string usando um REGEX,
			// e extraindo o conteúdo da tag TITLE da página
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			// extrai o primeiro índice e insere no canal
			c <- r.FindStringSubmatch(string(html))[1]

		}(url) // criar a função anônima, e já chama a mesma passando a url, para que seja criada a cada iteração uma nova Goroutine independente.
	}
	// retorna o canal carregado
	return c
}

// função que recebe um dado de canal de origem e envia o
// dado para um canal de destino.
func encaminhar(origem <-chan string, destino chan string) {
	for {
		// sempre que chegar um dado no canal de origem, enviar para o canal de destino
		destino <- <-origem
	}
}

// função que multiplexa (mistura as mensagens de canais diferentes) num canal só
// input: dois canais distintos de strings
// output: um canal de string
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	// cria um canal do tipo string
	c := make(chan string)
	// cria primeira goroutine que encaminha 'entrada1' para o canal 'c'
	go encaminhar(entrada1, c)
	// cria segunda goroutine que encaminha 'entrada2' para o canal 'c'
	go encaminhar(entrada2, c)
	// retorna o canal multiplexado
	return c
}

func main() {
	// cira um canal usando a função 'juntar' onde passamos as urls
	// para nossa função 'Titulo' que esta dentro do nosso pacote 'html'
	// criado dentro do nosso workspace do GOPATH
	c := juntar(
		titulo("https://www.cod3r.com.br/", "https://www.globo.com/"),
		titulo("https://www.nba.com/", "https://www.youtube.com/"),
	)

	// receber o primeiro valor retornado
	fmt.Printf("Primeiro valor mais rápido: %s\n", <-c)

	// receber o segundo valor retornado
	fmt.Printf("Segundo valor mais rápido: %s\n", <-c)

	// receber o terceiro valor retornado
	fmt.Printf("Terceiro valor mais rápido: %s\n", <-c)

	// receber o quarto valor retornado
	fmt.Printf("Quarto valor mais rápido: %s\n", <-c)

	// Looping para ler todos os dados do canal
	// for msg := range c {
	// 	fmt.Println("Msg: " + msg)
	// }

}
