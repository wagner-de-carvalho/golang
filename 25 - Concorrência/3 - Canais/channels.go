package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-------Channels-------")
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	// for {
	// 	mensagem, aberto := <- canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal{
		fmt.Println(mensagem)	
	}
	fmt.Println("Fim da execução")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
