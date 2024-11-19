package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // import indireto que precisar ter para acessar MySQL
)

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

	// // a partir da transação aberta, cria o statement para a inserção na tabela usuarios
	// stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	// // executar cada update via statement
	// stmt.Exec("Uóxiton Istive", 1)
	// stmt.Exec("Sheristone Uasleska", 2)

	// UPDATE usando transação
	// abrir uma transação
	tx, _ := db.Begin()

	// a partir da transação aberta, cria o statement para a inserção na tabela usuarios
	stmt, _ := tx.Prepare("update usuarios set nome = ? where id = ?")

	// insere um registros usando o statement passando o ID e NOME (mesmo sendo o ID uma chave de autoincremento)
	stmt.Exec("Abreu Botelho III", 1)
	stmt.Exec("Carlos Daniel III", 2)

	// UPDATE - update elemento com base no ID inexistente
	// Cria query
	updateQuery := "update usuarios set nome = ? where id = ?"

	// cria nome e ID que não existe na base de dados
	nomeUpdate := "Pedro de Lara"
	idUpdate := 2000

	// executar o update do registro com base no ID que não existe na base
	resultUpdate, errUpdate := tx.Exec(updateQuery, nomeUpdate, idUpdate) // chave inexistente

	// verifica se existe erro ao fazr update
	if errUpdate != nil {
		// Rollback the transaction on error
		tx.Rollback()
		log.Fatalf("Error executing update query: %v", errUpdate)
	}

	// Check the number of rows affected
	rowsAffectedUpdate, errUpd := resultUpdate.RowsAffected()
	if errUpd != nil {
		// Rollback the transaction on error
		tx.Rollback()
		log.Fatalf("Error checking rows affected update: %v", errUpd)
	}

	// Handle the case where no rows were affected
	if rowsAffectedUpdate == 0 {
		// Rollback the transaction
		tx.Rollback()
		fmt.Printf("Error: Update usuario with ID %d does not exist in the 'usuarios' table.\n", idUpdate)
		return
	}

	// DELETE - remove elemento com base no ID
	// Cria query
	deleteQuery := "delete from usuarios where id = ?"

	// cria ID inexistente na base de dados
	idDelete := 2000

	// executar a remoção do registro com base no ID que não existe na base
	result, errDel := tx.Exec(deleteQuery, idDelete) // chave inexistente

	fmt.Printf("ErrorDel: %v\n", errDel)

	// verifica se existe erro, e havendo erro, fazer o roolback da transação
	if errDel != nil {
		// Rollback the transaction on error
		tx.Rollback()
		log.Fatalf("Error executing delete query: %v", errDel)
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Rollback the transaction on error
		tx.Rollback()
		log.Fatalf("Error checking rows affected: %v", err)
	}

	// Handle the case where no rows were affected
	if rowsAffected == 0 {
		// Rollback the transaction
		tx.Rollback()
		fmt.Printf("Error: Record with ID %d does not exist in the 'usuarios' table.\n", idDelete)
		return
	}

	// Commit the transaction if successful
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Error committing transaction: %v", err)
	}

	// fmt.Printf("Record with ID %d deleted successfully.\n", idDelete)

	// caso não tenha erro, fazer o commit da transação com sucesso
	// tx.Commit()

}
