package main

import (
	"fmt"
)

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Tipo genérico")
	generica("Texto")
	generica(123)
	generica(123.456)
	generica(true)

}