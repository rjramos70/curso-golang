package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
	f.velocidadeAtual = incrementaVelocidadeViaTurbo(f.velocidadeAtual)
}

func incrementaVelocidadeViaTurbo(velocidadeCorrente int) int {
	velocidadeFinal := float64(velocidadeCorrente) * 1.20
	return int(velocidadeFinal)
}

func main() {
	carro1 := ferrari{"F40", false, 10}
	fmt.Println(carro1)
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F50", false, 20}
	fmt.Println(carro2)
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
