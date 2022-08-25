package main

import(
	"fmt"
)

func main() {
	fmt.Println("Maps")
	usuario := map[string]string{
		"nome": "Ken",
		"sobrenome": "Masters",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	
	usuario2 := map[int8]map[string]string{
		1: {
			"nome": "Ken",
			"sobrenome": "Masters",
		},
		2: {
			"nome": "Mario",
			"sobrenome": "Bros",
		},
	}
	fmt.Println(usuario2[2])
	delete(usuario2, 2) // REMOVER CHAVE
	fmt.Println(usuario2)

}
