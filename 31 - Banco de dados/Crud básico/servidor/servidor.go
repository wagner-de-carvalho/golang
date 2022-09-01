package servidor

import (
	"banco-de-dados/banco"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
		w.Write([]byte("Falha ao ler corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário para struct!"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	// PrepareStatement
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	// Inserçaõ
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter id inserido no banco!"))
		return
	}

	mensagem := mensagem{idInserido, "Usuário criado com sucesso!"}
	mensagemJSON, erro := json.Marshal(mensagem)
	if erro != nil {
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
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários!"))
		return
	}
	defer linhas.Close()

	// Montar structs de usuários
	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao ler registro de usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON!"))
		return
	}
}

// Busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}
