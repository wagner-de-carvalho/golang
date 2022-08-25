package main

import(
	"fmt"
)

func main() {
	fmt.Println("Estruturas de controle")
	numero := 10

	if numero > 10 {
		fmt.Println("Maior que 10")
		}else {
		fmt.Println("Menor ou igual a 10")
	}
	
	if outroNumero := 15; outroNumero > 0 {
		fmt.Println("Maior que zero")
	} else {
		fmt.Println("Menor que zero")
	}

}
