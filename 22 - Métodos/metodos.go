package main

import (
	"fmt"
)

type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Gravando os dados do usuário no banco de dados.\n[%s, %d]\n", u.nome, u.idade)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("MÉTODOS")
	usuario1 := usuario{"Kron", 17}
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
	usuario1.fazerAniversario()
	fmt.Printf("Nome: %s, idade: %d\n", usuario1.nome, usuario1.idade)
	fmt.Println(usuario1.maiorDeIdade())
	
	usuario2 := usuario{"Davi", 33}
	usuario2.salvar()
	usuario2.maiorDeIdade()
	fmt.Println(usuario2.maiorDeIdade())
	usuario2.fazerAniversario()
	fmt.Printf("Nome: %s, idade: %d\n", usuario2.nome, usuario2.idade)
	fmt.Println(usuario2.maiorDeIdade())

}