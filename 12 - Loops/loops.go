package main
import(
	"fmt"
	// "time"
)

func main() {
	fmt.Println("Loops")
	// i := 0
	// for i < 10 {
	// 	fmt.Println("Incrementando i = ", i)
	// 	i++
	// 	time.Sleep(time.Second)
	// }


	// for i := 0; i < 10; i++ {
	// 	fmt.Print("Incrementando i = ", i)
	// }

	// ITERAÇÃO
	// nomes := []string{"Go Lang", "Loops", "Slice", "Int"}

	// for index, nome := range nomes {
	// 	fmt.Println(index, nome)
	// }

	// for index, nome := range "PALAVRAS" {
	// 	fmt.Print(index, nome, " ")
	// 	fmt.Println(index, string(nome))
	// }

	usuario := map[string]string{
		"nome": "Palavras",
		"sobrenome": "Refrão",
	}

	for chave, valor := range usuario {
		fmt.Print(chave, ": ", valor, " ")
	}
}