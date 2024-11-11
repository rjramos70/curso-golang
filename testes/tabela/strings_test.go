package strings

import (
	"strings"
	"testing"
)

// criar uma constante
const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	// informa que esse teste pode ser executado de forma paralela
	t.Parallel()

	// criar o DataSet dos testes
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é show", "Cod3r", 0}, // buscar Cod3r na posição zero (0)
		{"", "", 0},                  // buscar vazio na posição zero (0)
		{"Opa", "opa", -1},           // não achar 'opa' no texto e retornar -1
		{"leonardo", "o", 2},         // buscar a letra 'o' na palavra no índice 2
	}

	// iterage nos cenários de teste da struct
	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		// verifica se índice retornado na variavel 'atual' é diferente do esperado
		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
