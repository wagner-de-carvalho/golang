package main

import(
	"fmt"
)

func diaDaSemana(numero int) string {
	switch numero {
		case 1:
			return "Domingo"
		case 2:
			return "Segunda-feira"
		case 3:
			return "Terça-feira"
		case 4:
			return "Quarta-feira"
		case 5:
			return "Quinta-feira"
		case 6:
			return "Sexta-feira"
		case 7:
			return "Sábado"
		default:
			return "Número inválido"
	}
}

func main() {
	fmt.Println("Estruturas de controle")
	numero := 10

	// IF ELSE
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

	// SWITCH
	fmt.Println(diaDaSemana(6))

}
