package main

// Marshal, UnMarshal

import (
	"encoding/json"
	"fmt"
	"log"
)

type cao struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint `json:"idade"`
}

func main()  {
	fmt.Println("JSON")
	caoEmJSON := `{"nome":"Raio","raca":"Dálmata","idade":3}`

	var c cao
	if erro := json.Unmarshal([]byte(caoEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cao2EmJSON := `{"nome":"Bom","raca":"Poodle"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cao2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
	
}