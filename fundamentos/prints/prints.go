package main

import "fmt"

func main() {

	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	fmt.Println(" Nova")
	fmt.Println("Linha.")

	x := 3.141516
	// fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x) // transforma de float64 para string
	fmt.Println("O valor de xs é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x)
	// fmt.Printf("O valor de x é %.4f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
	// Onde:
	// %d 	--> inteiro
	// %f 	--> float
	// %.1f --> float mostrand somente uma casa decimal
	// %t 	--> boolean
	// %s 	--> string
	// %v 	--> imprime quase todos os tipos

}
