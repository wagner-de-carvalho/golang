package main

import (
	"fmt" 
	"errors"
)

func main() {
	var numero int16 = 100
	var numero2 int = 200
	numero3 := -10000000
	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)

	// unsigned int
	var numero4 uint32 = 20000000
	fmt.Println(numero4)

	// alias
	// rune = int32
	var numero5 rune = 123456
	fmt.Println(numero5)

	// byte = uint8
	var numero6 byte = 123
	fmt.Println(numero6)

	// float32, float64
	var numeroReal1 float32 = 132.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1329.454565
	fmt.Println(numeroReal2)

	// strings
	var txt1 string = "Golang go lang"
	fmt.Println(txt1)

	// char
	char := 'B'
	fmt.Println(char)
	
	// valor 0
	var txt2 string
	fmt.Println(txt2)

	var n1 int8
	fmt.Println(n1)

	// boolean
	var booleano1 bool
	fmt.Println(booleano1)

	var booleano2 bool = true
	fmt.Println(booleano2)

	// error
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}