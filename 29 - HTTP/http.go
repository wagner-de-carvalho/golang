package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bem-vindos usuários!"))
}

func main() {
	
	// Criar rotas
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	// Levantar servidor na porta 5000
	log.Fatal(http.ListenAndServe(":5000", nil))


}