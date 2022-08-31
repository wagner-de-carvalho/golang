package enderecos

import (
	"strings"
)

// TipoDeEndereco verifica se o endereço é um tipo válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{
		"alameda",
		"avenida",
		"estrada",
		"rodovia",
		"rua",
		"travessa",
		"via",
		"caminho",
		"beco",
	}
	enderecoMinucusculo := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoMinucusculo, " ")[0]
	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return  strings.Title(primeiraPalavraDoEndereco) 
	}
	return "Tipo inválido"
}
