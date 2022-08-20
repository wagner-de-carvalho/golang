package main
import ("fmt")
func main() {
	soma := somar(10, 4)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f("SAPO")

	resultSoma, resultSub := calculosMatematicos(20, 8)
	fmt.Println(resultSoma)
	fmt.Println(resultSub)
	
	resultadoSoma, _ := calculosMatematicos(10, 8)
	fmt.Println(resultadoSoma)
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func somar(n1 int8, n2 int8) int8 { 
	return n1 + n2
}