package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// abrir uma conexão com a raiz do banco de dados
	// esta configuração é do 'MySQL Local' que consta no MySQLWorkbench
	db, err := sql.Open("mysql", "root:root@/cursogo")

	// verificar se houve algum erro em conectar ao mysql
	if err != nil {
		log.Fatal(err)
	}

	// atrase, mas fecha a conexão com o banco de dados, antes de sair dessa função
	defer db.Close()

	// abrir uma transação
	tx, _ := db.Begin()

	// a partir da transação aberta, cria o statement para a inserção na tabela usuarios
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	// insere um registros usando o statement passando o ID e NOME (mesmo sendo o ID uma chave de autoincremento)
	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	// agora vamos gerar um erro por tentar fazer um insert de um ID duplicado
	pessoa, err := stmt.Exec(1, "Thiago") // chave duplicada

	// verifica se existe erro, e havendo erro, fazer o roolback da transação
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	fmt.Printf("Pessoa: %v\n", pessoa)

	// caso não tenha erro, fazer o commit da transação com sucesso
	tx.Commit()

}
