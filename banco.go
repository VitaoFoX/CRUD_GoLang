package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Pessoa é o tipo um objeto
type Pessoa struct {
	CdPessoa int
	Nome     string
	Senha    string
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql) // Isso executa no mysql
	if err != nil {
		//panic(err) // panic para o programa
	}
	return result

}

func iniciaBanco() {

	db, err := sql.Open("mysql", "root:123@/")
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados!QQ")
		panic(err)
	}
	defer db.Close()

	fmt.Println("Banco Atualizado!")

	exec(db, "create database if not exists bankDaniel")
	exec(db, "use bankDaniel")
	exec(db, `
		CREATE TABLE tbPessoa(
			cdPessoa int auto_increment primary key,
			nome varchar(100),
			senha varchar(100)
		);
	`)
}

func insertPessoa(dados Pessoa) int64 {

	db, err := sql.Open("mysql", "root:123@/bankVitor")
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados!")
		panic(err)
	}
	defer db.Close()

	insert, _ := db.Prepare(`
		INSERT INTO tbPessoa (nome, senha)
		VALUES(?,?);
	`)

	resultado, _ := insert.Exec(dados.Nome, dados.Senha)

	linhas, _ := resultado.RowsAffected()

	return linhas

}
