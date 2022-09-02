package servidor

import (
	"banco-de-dados/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint   `json:id`
	Nome  string `json:nome`
	Email string `json:email`
}

type mensagem struct {
	ID       int64  `json:id`
	Mensagem string `json:mensagem`
}

// Insere usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Falha ao ler corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("Erro ao converter usuário para struct!"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	// PrepareStatement
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	// Inserçaõ
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao executar statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Erro ao obter id inserido no banco!"))
		return
	}

	mensagem := mensagem{idInserido, "Usuário criado com sucesso!"}
	mensagemJSON, erro := json.Marshal(mensagem)
	if erro != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("Erro ao converter mensagem para JSON!"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(mensagemJSON))
}

// Busca todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar usuários!"))
		return
	}
	defer linhas.Close()

	// Montar structs de usuários
	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("Erro ao ler registro de usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("Erro ao converter usuários para JSON!"))
		return
	}
}

// Busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("Erro ao converter parâmetro para inteiro!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Erro ao tentar conectar-se ao banco!"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Erro ao tentar buscar usuário no banco!"))
		return
	}
	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("Erro ao escanear usuário!"))
			return
		}
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		mensagem := map[string]string{"mensagem": "Usuário não encontrado!"}
		mensagemJSON, _ := json.Marshal(mensagem)
		w.Write([]byte(mensagemJSON))
		return
	}
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("Erro ao converter usuário para JSON!"))
		return
	}
}

// Atualizar usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		mensagem := map[string]string{"mensagem": "Erro ao converter parâmetro!"}
		mensagemJSON, _ := json.Marshal(mensagem)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(mensagemJSON))
		return
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		mensagem := map[string]string{"mensagem": "Erro ao ler corpo da requisição!"}
		mensagemJSON, _ := json.Marshal(mensagem)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(mensagemJSON))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		mensagem := map[string]string{"mensagem": "Erro ao converter usuário para struct!"}
		mensagemJSON, _ := json.Marshal(mensagem)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(mensagemJSON))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		mensagem := map[string]string{"mensagem": "Erro ao tentar conectar-se ao banco!"}
		mensagemJSON, _ := json.Marshal(mensagem)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(mensagemJSON))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		mensagem := map[string]string{"mensagem": "Erro ao criar statement!"}
		mensagemJSON, _ := json.Marshal(mensagem)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(mensagemJSON))
		return
	}
	defer db.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		mensagem := map[string]string{"mensagem": "Erro ao atualizar usuário!"}
		mensagemJSON, _ := json.Marshal(mensagem)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(mensagemJSON))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
