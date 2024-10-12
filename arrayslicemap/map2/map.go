package main

import "fmt"

func main() {
	// inicializar um map atribuindo os valores de
	// nome (chave) e salário (valor)
	funcsESalarios := map[string]float64{
		"José João":     11325.45,
		"Gabriel Silva": 15456.78,
		"Pedro Junior":  1200.0,
	}
	fmt.Println(funcsESalarios)

	// adiciona mais um funcionário
	funcsESalarios["Rafael Filho"] = 1350.0

	// atualiza o salário de Pedro Junior em 30%
	funcsESalarios["Pedro Junior"] = funcsESalarios["Pedro Junior"] * 1.3

	fmt.Println(funcsESalarios)

	// remover um elemento que não existe no map
	delete(funcsESalarios, "Inexistente")

	// percorrer o map
	for nome, salario := range funcsESalarios {
		fmt.Printf("%s - %.2f\n", nome, salario)
	}

}
