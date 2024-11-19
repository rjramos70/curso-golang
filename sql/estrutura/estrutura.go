package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// abrir uma conexão com a raiz do banco de dados
	// esta configuração é do 'MySQL Local' que consta no MySQLWorkbench
	db, err := sql.Open("mysql", "root:root@/")

	// verificar se houve algum erro em conectar ao mysql
	if err != nil {
		panic(err)
	}

	// atrase, mas fecha a conexão com o banco de dados, antes de sair dessa função
	defer db.Close()

	// chama a função 'exec' passando a conexão, e o comando de criação do database 'cursogo'
	exec(db, "create database if not exists cursogo")

	// set/use o schema de banco de dados 'cursogo'
	exec(db, "use cursogo")

	// apagar a tabela 'usuarios' caso ela exista (caso queira deletar a tabela a cada execução)
	// exec(db, "drop table if exists usuarios")

	// cria a tabela 'usuarios'
	exec(db, `create table if not exists usuarios(
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)

}
