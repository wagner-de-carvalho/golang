package main

import (
	"fmt"
	"race-patterns/utils"
)

func main() {
	fmt.Println("Padrão de concorrência worker pool")
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados) //processo 1
	go worker(tarefas, resultados) //processo 2
	go worker(tarefas, resultados) //processo 3
	go worker(tarefas, resultados) //processo 4

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}
}

//Especifica canal que só recebe ou só envia
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- utils.Fibonacci(numero)
	}
}
