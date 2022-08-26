package main
import "fmt"

func soma(numeros...int) int{
	total := 0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func escrever(texto string, numeros...int) {
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções variáticas")
	soma := soma(31, 27, 55, 11, 3, 9)
	fmt.Println(soma)

	escrever("olá mundo ", 11, 27, 29, 19)
}