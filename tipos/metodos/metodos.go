package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomecompleto string) {
	partes := strings.Split(nomecompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto()) // Pedro Silva

	// alterar o nome completo
	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompleto()) // Maria Pereira
}
