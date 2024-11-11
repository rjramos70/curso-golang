package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	// informa que esse teste pode ser executado de forma paralela
	t.Parallel()

	// testes iram rodar a depender da arquitetura do processador da maquina
	// Se for 'amd64' não fazer nada
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}

	//
	t.Fail()
}
