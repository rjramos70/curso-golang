package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userId int
	itens  []item // slice de itens
}

func (p pedido) valorTotal() (float64, int) {
	total := 0.0
	qtd := 0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
		qtd += item.qtde
	}
	return total, qtd
}

/* SLOWER */
func isEven1(n int) bool {
	return n%2 == 0
}

/* FASTER */
func isEven2(n int) bool {
	fmt.Println(n)
	fmt.Println((n & 1))

	return (n & 1) == 0
}

func main() {
	pedido1 := pedido{
		userId: 1,
		itens: []item{
			item{
				produtoID: 1,
				qtde:      2,
				preco:     12.10,
			},
			item{
				produtoID: 11,
				qtde:      100,
				preco:     3.13,
			},
			item{2, 1, 23.49},
		},
	}
	totalP1, qtdP1 := pedido1.valorTotal()
	fmt.Printf("Valor total do pedido 1 é R$ %.2f [ %v itens ]\n", totalP1, qtdP1)

	// Usando o SimplifyCompositelit do Go
	// https://pkg.go.dev/golang.org/x/tools/gopls/internal/analysis/simplifycompositelit#hdr-Analyzer_simplifycompositelit
	pedido2 := pedido{
		userId: 2,
		itens: []item{
			{
				produtoID: 1,
				qtde:      20,
				preco:     12.10,
			},
			{
				produtoID: 11,
				qtde:      70,
				preco:     3.13,
			},
			{2, 8, 23.49},
		},
	}
	totalP2, qtdP2 := pedido2.valorTotal()
	fmt.Printf("Valor total do pedido 2 é R$ %.2f [ %v itens ]\n", totalP2, qtdP2)

	//
	val := 8
	fmt.Printf("Is Even 1: %v\n", isEven1(val))
	fmt.Printf("Is Even 2: %v\n", isEven2(val))

}
