package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura float32
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}
func main() {
	fmt.Println("Herança")
	pessoa1 := pessoa{"Karlos", "Márcio", 22, 1.82}
	pessoa2 := pessoa{"Muno", "Teles", 21, 1.85}
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	
	estudante1 := estudante{pessoa1, "Programação", "UCL"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.altura)
}