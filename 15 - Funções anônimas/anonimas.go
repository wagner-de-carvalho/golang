package main
import "fmt"

func main() {
	fmt.Println("----------Funções anônimas----------")
	
	func(texto string) {
		fmt.Println(texto)
	}("Esta é uma função anônima")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido: %s", texto)
	}("Esta é uma outra função anônima")

	fmt.Println(retorno)

}