package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

// teste unitário da função Media
func TestMedia(t *testing.T) {
	// informa que esse teste pode ser executado de forma paralela
	t.Parallel()

	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
