package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua das Flores", "Rua"},
		{"Alameda dos Anjos", "Alameda"},
		{"Estrada da Posse", "Estrada"},
		{"Via do centro", "Via"},
		{"Travessa Ribeira", "Travessa"},
		{"Beco do Ribeira", "Beco"},
		{"Rodovia do contorno", "Rodovia"},
		{"Caminho das Pedras", "Caminho"},
		{"Praça da Paz", "Tipo inválido"},
		{"RUA ABC", "Rua"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido = %s é diferente do esperado = %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}
