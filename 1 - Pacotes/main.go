package main

import(
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main()  {
	fmt.Println("Escrevendo arquivo Principal")
	auxiliar.Escrever()

	erro:= checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}