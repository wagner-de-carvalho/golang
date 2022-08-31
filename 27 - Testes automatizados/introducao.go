package main

import(
	"fmt"
	"introducao-testes/enderecos"
)

func main()  {
	fmt.Println("Introdução testes automatizados")
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}