package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("-------Waitgroup-------")
	var waitGroup sync.WaitGroup
	waitGroup.Add(4) //ADICIONA QUANTIDADE DE GO ROUTINES

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Golang!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 3!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
