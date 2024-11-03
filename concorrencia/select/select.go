package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
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
			// faz requisição http para cada url ignorando o erro
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

/*
Esta função recebe uma lista de urls e vai retornar somente a mais rapida
*/
func oMaisRapido(url1, url2, url3 string) string {
	c1 := titulo(url1)
	c2 := titulo(url2)
	c3 := titulo(url3)

	// estrutura de controle especifica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam, demoraram demais!"
		// default:		// default sempre será o mais rápido caso ele não esteja comentado.
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	// passa 3 urls para a função e recebe a mais rápida
	campeao := oMaisRapido(
		"https://www.cod3r.com.br/",
		"https://www.globo.com/",
		"https://www.nba.com/",
	)

	// Imprime o valor que retornou primeiro
	fmt.Println("Campeão : " + campeao)

}
