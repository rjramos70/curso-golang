package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota <= 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}

/*
Refatoração do metodo 'notaParaConceito' para utilizar a
estrutura Switch invês de IF
*/
func notaParaConceito2(nota float64) string {
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota <= 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	// Estrutura 1
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))

	// Estrutura 2
	fmt.Println(notaParaConceito2(9.8))
	fmt.Println(notaParaConceito2(6.9))
	fmt.Println(notaParaConceito2(2.1))

}
