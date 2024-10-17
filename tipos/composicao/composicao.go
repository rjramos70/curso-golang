package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos...
	ligarLimpadores()
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func (b bmw7) ligarLimpadores() {
	fmt.Println("Limpadores...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	// executar os métodos do modelo esportivoLuxuoso
	b.ligarTurbo()
	b.fazerBaliza()
	b.ligarLimpadores()
}
