package main 

import "fmt"
func main()  {
	// declaração de variável explícita
	var str1 = "Golang"
	fmt.Println(str1)
	// declaração de variável implícita
	str2 := "Golang 2"
	fmt.Println(str2)

	// múltiplas declarações
	var (
		str3 = "Golang 3"
		str4 = "Golang 4"
	)
	fmt.Println(str3)
	fmt.Println(str4)

	str5, str6 := "Golang 5", "Golang 6"
	fmt.Println(str5)
	fmt.Println(str6)

	const constante1 = "Const1"
	fmt.Println(constante1)

	// Inverter valores de variáveis
	str1, str2 = str2, str1
	fmt.Println(str1, str2)
}