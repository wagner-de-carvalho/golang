package main
import (
	"fmt"
)

func funcao() {
	fmt.Println("Executando funcao()")
}

func funcao2() {
	fmt.Println("Executando funcao2()")
}

func alunoEstaAprovado(n1, n2 float32) bool{
	defer fmt.Println("Média calculada. Resultado será retornado!")
	media := (n1 + n2)/2
	return media >= 5
}

func main() {
	fmt.Println("Defer")
	defer funcao()
	funcao2()

	fmt.Println(alunoEstaAprovado(5.0, 4.8))
}
