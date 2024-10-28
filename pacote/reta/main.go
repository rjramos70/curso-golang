package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	// por estar dentro do contexto do mesmo pacote se
	// pode chamar um método privado.
	fmt.Println(catetos(p1, p2)) // 0 2

	// executar o método público Distancia
	fmt.Println(Distancia(p1, p2)) // 2
}
