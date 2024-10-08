package main

import (
	"fmt"
	"time"
)

func main() {
	// comparações básicas
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	//
	d1 := time.Unix(0, 0)
	fmt.Println("d1 = ", d1)
	d2 := time.Unix(0, 0)
	fmt.Println("d2 = ", d2)

	// compara datas pelo valores e não pela posição de memória
	fmt.Println("Datas: ", d1 == d2)
	fmt.Println("Datas: ", d1.Equal(d2))

	// compara tipos struct pelo valor e não pela posição de memória
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoas: ", p1 == p2)
}
