package main

// Marshal, UnMarshal

import (
	"bytes"
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
	c := cao{"Raio", "Dálmata", 3}
	fmt.Println(c)
	caoEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	// fmt.Println(caoEmJSON)
	fmt.Println(bytes.NewBuffer(caoEmJSON))

	c2 := map[string]string{
		"nome": "Trovão",
		"raca": "Poodle",
	}
	c2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	// fmt.Println(c2EmJSON)
	fmt.Println(bytes.NewBuffer(c2EmJSON))
}