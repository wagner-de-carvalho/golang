package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Padrão de concorrência Generator")
	canal := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Print(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s\n", texto)
			time.Sleep(time.Millisecond * 500) 
		}
	}()

	return canal
}