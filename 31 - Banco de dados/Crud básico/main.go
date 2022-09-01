package main

import (
	"banco-de-dados/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("CRUD Básico")

	// cria rotas com mux
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	fmt.Println("Servidor executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
