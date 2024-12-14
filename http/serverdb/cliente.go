package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/*
Classe responsável por interagir com o banco de dados e obter a lista dos clientes
*/

// Structs
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// RequestPayload represents the incoming JSON structure
type RequestPayload struct {
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o resquest e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	// obtem a parametro passado no path
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid) // convert string para um int

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorId(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	case r.Method == "POST":
		criarUsuario(w, r)
	case r.Method == "PUT" && id > 0:
		updateUsuario(w, r, id)
	case r.Method == "DELETE" && id > 0:
		removerUsuario(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}

}

func updateUsuario(resp http.ResponseWriter, r *http.Request, id int) {
	// Check if the request method is PUT
	if r.Method != http.MethodPut {
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload from the request body
	var payload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(resp, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// abrir uma conexão com a raiz do banco de dados
	// esta configuração é do 'MySQL Local' que consta no MySQLWorkbench
	db, err := sql.Open("mysql", "root:root@/cursogo")

	// verificar se houve algum erro em conectar ao mysql
	if err != nil {
		log.Fatal(err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, "Erro ao acessar a base de dados... :(")
	}

	// cria uma instacia de uruario
	var u Usuario

	// executa a query para listar todos os usuarios no banco de dados
	result, err := db.Exec("update usuarios set nome = ? where id = ?", payload.Nome, id)

	if err != nil {
		http.Error(resp, fmt.Sprintf("Error updating usuario: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(resp, "Usuario not found", http.StatusNotFound)
		return
	}
	// Prepare the response
	u.ID = id
	u.Nome = payload.Nome
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(u)

}

func removerUsuario(resp http.ResponseWriter, r *http.Request, id int) {
	// abrir uma conexão com a raiz do banco de dados
	// esta configuração é do 'MySQL Local' que consta no MySQLWorkbench
	db, err := sql.Open("mysql", "root:root@/cursogo")

	// verificar se houve algum erro em conectar ao mysql
	if err != nil {
		log.Fatal(err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, "Erro ao acessar a base de dados... :(")
	}

	// atrasa, mas fecha a conexão com o banco de dados, antes de sair dessa função
	defer db.Close()

	// executa a query para listar todos os usuarios no banco de dados
	result, err := db.Exec("delete from usuarios where id = ?", id)
	if err != nil {
		http.Error(resp, fmt.Sprintf("Error deleting usuario: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(resp, "Usuario not found", http.StatusNotFound)
		return
	}

	resp.WriteHeader(http.StatusNoContent)

}

func criarUsuario(resp http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload from the request body
	var requestPayload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(resp, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// abrir uma conexão com a raiz do banco de dados
	// esta configuração é do 'MySQL Local' que consta no MySQLWorkbench
	db, err := sql.Open("mysql", "root:root@/cursogo")

	// verificar se houve algum erro em conectar ao mysql
	if err != nil {
		log.Fatal(err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, "Erro ao acessar a base de dados... :(")
	}

	// cria uma instacia de uruario
	var u Usuario

	// executa a query para listar todos os usuarios no banco de dados
	result, err := db.Exec("insert into usuarios (nome) values (?)", requestPayload.Nome)

	if err != nil {
		http.Error(resp, fmt.Sprintf("Error saving to database: %v", err), http.StatusInternalServerError)
		return
	}

	// Get the ID of the newly created record
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(resp, fmt.Sprintf("Error retrieving last insert ID: %v", err), http.StatusInternalServerError)
		return
	}

	// Prepare the response
	u.ID = int(id)
	u.Nome = requestPayload.Nome
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(u)

}

func usuarioPorId(resp http.ResponseWriter, r *http.Request, id int) {
	// abrir uma conexão com a raiz do banco de dados
	// esta configuração é do 'MySQL Local' que consta no MySQLWorkbench
	db, err := sql.Open("mysql", "root:root@/cursogo")

	// verificar se houve algum erro em conectar ao mysql
	if err != nil {
		log.Fatal(err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, "Erro ao acessar a base de dados... :(")
	}

	// atrasa, mas fecha a conexão com o banco de dados, antes de sair dessa função
	defer db.Close()

	// cria uma instacia de uruario
	var u Usuario

	log.Printf("usuarioPorId - ID : %d", id)
	// executa a query no banco ja scaneando o resultado para a instancia do usuario 'u'
	// QueryRow -> obtem somente uma linha como resposta
	// Query	-> obtem uma lista como resposta
	// db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)
	err = db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)
	switch {
	case err == sql.ErrNoRows:
		// log.Fatalf("no user with id %d", id)
		// log.Fatal(err)
		log.Println("sql.ErrNoRows...")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "Não existe usuário com esse ID :(")
	case err != nil:
		log.Println("Err != nil")
		log.Fatal(err)
	default:
		log.Println("Default...")
		log.Printf("Usuario nome é %s\n", u.Nome)

		// converter a instancia de usuario para json
		json, _ := json.Marshal(u)

		// preencher o header da resposta mostrando que o tipo de dado retornado sera um json
		resp.Header().Set("Content-Type", "applicatio/json")

		// escreve no payload da resposta ja convertendo o json para string
		fmt.Fprint(resp, string(json))
	}

}

func usuarioTodos(resp http.ResponseWriter, req *http.Request) {
	// abrir uma conexão com a raiz do banco de dados
	// esta configuração é do 'MySQL Local' que consta no MySQLWorkbench
	db, err := sql.Open("mysql", "root:root@/cursogo")

	// verificar se houve algum erro em conectar ao mysql
	if err != nil {
		log.Fatal(err)
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, "Erro ao acessar a base de dados... :(")
	}

	// atrasa, mas fecha a conexão com o banco de dados, antes de sair dessa função
	defer db.Close()

	// executa a query para listar todos os usuarios no banco de dados
	rows, _ := db.Query("select id, nome from usuarios")

	// atrasa, mas fecha o resultSet retornado com a lista dos usuarios
	defer rows.Close()

	// cria um slice de usuarios
	var usuarios []Usuario

	// perdorrer o resultSet e preencher o slice
	for rows.Next() {
		// cria um usuario temporario
		var usuario Usuario
		// faz o parce de cada linha para o usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		// faz um append no slice do usuario
		usuarios = append(usuarios, usuario)
	}

	// converte splice de usuarios para json
	json, _ := json.Marshal(usuarios)

	// preencher o header da resposta mostrando que o tipo de dado retornado sera um json
	resp.Header().Set("Content-Type", "applicatio/json")

	// escreve no payload da resposta ja convertendo o json para string
	fmt.Fprint(resp, string(json))
}
