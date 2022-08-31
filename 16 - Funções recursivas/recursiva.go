package main
import "fmt"

func fibobacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibobacci(posicao - 2) + fibobacci(posicao - 1)
}

func main() {
	fmt.Println("----------Funções recursivas----------")
	posicao := uint(10)
	fmt.Println(fibobacci(posicao))
	
	for i := uint(0); i < posicao; i++ {
		fmt.Print(fibobacci(i), " ")

	}

}