package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())

	// como usando o método imprimir
	imprimir(coisa)

	// usando o polimorfísmo para mudar o forma/tipo da variável
	// de 'pessoa' para 'produto'
	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// como tanto 'pessoa' quanto 'produto' são imprimivel, podemos
	// passar direto como parâmetro para a função imprimir(x imprimivel)
	p2 := produto{"Camisa de Malha", 13.27}
	imprimir(p2)                       // Camisa de Malha - R$ 13.27
	imprimir(pessoa{"Pedro", "Cunha"}) // Pedro Cunha
}
