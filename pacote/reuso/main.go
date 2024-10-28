package main

import (
	"fmt"

	"github.com/rjramos70/area"
)

func main() {
	// vamos utilizar as funções que fora declaradas dentro
	// da nossa workspace da pasta "/Users/rjramos70/go/src/github.com/rjramos70/area"
	fmt.Println(area.Circ(6.0))      // 113.0976
	fmt.Println(area.Rect(5.0, 2.0)) // 10
	// fmt.Println(area._TrianguloEq(5.0, 2.0)) // Essa função não é visível (é privada) por iniciar com _
}
