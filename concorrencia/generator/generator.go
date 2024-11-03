package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Os padrões foram apresentados em:
// Google I/O 2012 - Go Concurrency Patterns
// https://www.youtube.com/watch?v=f6kdp27TYZs
// Github repo: https://github.com/kevinniechen/go-concurrency-patterns
//
// Google I/O 2013 - Advanced Go Concurrency Patterns
// https://www.youtube.com/watch?v=QDDwwePbDtw

// <-chan - canal somente de leitura (não vai passar dados para o canal)

/*
Função que recebe uma lista de URLs, e com base nessa lista faz uma requisição
HTTP para cada url e obtem o respectivo título da página. Retorna um canal de
string de leitura.
*/
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
			html, _ := ioutil.ReadAll(resp.Body)
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

func main() {
	// criando duas chamadas passando duas listas de urls distintas
	// gerando dois canais (channels) distintos
	t1 := titulo("https://www.cod3r.com.br/", "https://www.globo.com/")
	t2 := titulo("https://www.nba.com/", "https://www.youtube.com/")

	// lendo os dados dos canais, mostrando quem foi o primeiro dados de cada canal
	fmt.Println("Primeiros: ", <-t1, "|", <-t2)

	// lendo os dados dos canais, mostrando os segundos colocados
	fmt.Println("Segundos: ", <-t1, "|", <-t2)

}
