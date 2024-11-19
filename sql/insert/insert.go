package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// abrir uma conexão com a raiz do banco de dados
	// esta configuração é do 'MySQL Local' que consta no MySQLWorkbench
	db, err := sql.Open("mysql", "root:root@/cursogo")

	// verificar se houve algum erro em conectar ao mysql
	if err != nil {
		panic(err)
	}

	// atrase, mas fecha a conexão com o banco de dados, antes de sair dessa função
	defer db.Close()

	// cria o statement para a inserção na tabela usuarios
	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")

	// insere um registros usando o statement (todo insert retorna o 'id')
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	// obter o último 'id' inserido no banco de dados
	id, _ := res.LastInsertId()
	fmt.Printf("Último ID: %v\n", id)

	// Quantidade de linhas que foram afetadas
	linhas, _ := res.RowsAffected()
	fmt.Printf("Linhas afetadas: %v\n", linhas)
}
