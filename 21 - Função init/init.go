package main

import (
	"fmt"
)

var n int

func init() {
	fmt.Println("Executando a função init()")
	n = 10
	fmt.Println("n = ", n)

}

func main() {
	fmt.Println("INIT")

}