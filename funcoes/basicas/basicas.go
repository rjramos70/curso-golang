package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

// retorna mais de um valor
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()

	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	// função 'f5' retorna dois parametros
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)

	// podemos omitir um dos retornos da função 'f5'
	r5_1, _ := f5()
	fmt.Println("F5_1:", r5_1)

	_, r5_2 := f5()
	fmt.Println("F5_2:", r5_2)

}
