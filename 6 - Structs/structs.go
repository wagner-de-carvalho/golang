package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	// Create a new endereço
	fmt.Println("Structs")
	endereco1 := endereco{"Rua dos Cravos", 150}
	var u usuario
	u.nome = "Go Lang"
	u.idade = 21
	u.endereco = endereco1
	fmt.Println(u)
	
	// inferência de tipo
	u2 := usuario{"Go Lang 2", 22, endereco1}
	fmt.Println(u2)
	
	u3 := usuario{nome: "Go Lang 3"}
	fmt.Println(u3)

	u4 := usuario{idade: 33}
	fmt.Println(u4)

	u5 := usuario{idade: 35, endereco: endereco1}
	fmt.Println(u5)

	// 
}