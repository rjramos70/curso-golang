package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // import indireto que precisar ter para acessar MySQL
)

// cria um estrutura para mapear os usuários que vamos fazer select no banco de dados
type usuario struct {
	id   int
	nome string
}

func main() {
	// abrir uma conexão com a raiz do banco de dados
	// esta configuração é do 'MySQL Local' que consta no MySQLWorkbench
	db, err := sql.Open("mysql", "root:root@/cursogo")

	// verificar se houve algum erro em conectar ao mysql
	if err != nil {
		log.Fatal(err)
	}

	// garante o fechamento da conexão com o banco de dados, antes de sair dessa função
	defer db.Close()

	// cria e executa query select
	rows, _ := db.Query("select id, nome from usuarios where id > ?", 0)

	// fecha as rows
	defer rows.Close()

	// faz um FOR para iteragir com a lista de usuarios (valor determinado),
	// para quantidade de registros indeterminados usamos o WHILE.
	for rows.Next() {
		// cria uma variavel 'user' do tipo 'usuario'
		var user usuario
		// chamar o método scan para pegar um elemento e transformar ele e uma estrutura(struct)
		// from 'rows' to 'user'
		rows.Scan(&user.id, &user.nome)
		// imprime cada usuário 'u'
		fmt.Printf("Usuário : %v\n", user)
	}

}
