package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcular a média dentre uma lista de
// valores do tipo float64
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	// recebe a média
	media := total / float64(len(numeros))
	// arredondar a média para 2 casas decimais na base 64
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	// retorna a mádia arredondada
	return mediaArredondada
}
