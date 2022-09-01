package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	// Criar rotas
	http.HandleFunc("/home",
		func(w http.ResponseWriter, r *http.Request) {
			u := usuario{"Pedro Salles", "salles.pe@mail.com"}

			templates.ExecuteTemplate(w, "home.html", u)
		})

	fmt.Println("Servidor executando na porta 5000")
	// Levantar servidor na porta 5000
	log.Fatal(http.ListenAndServe(":5000", nil))

}
