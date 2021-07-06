package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {

	/*Manipulo todos os templates .gohtml*/
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))

}

func crudVitao(w http.ResponseWriter, req *http.Request) {

	/***/
	iniciaBanco()
	/***/

	var DadosSlice []Pessoa

	db, err := sql.Open("mysql", "root:4553@/bankVitor")
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados!")
		panic(err)
	}
	defer db.Close()

	if req.Method == http.MethodPost {

		var PessoaInserir Pessoa

		login := req.FormValue("login")
		senha := req.FormValue("pass")

		PessoaInserir.Nome = login
		PessoaInserir.Senha = senha

		retorno := insertPessoa(PessoaInserir)

		if retorno > 0 {

		}

	}
	rows, err := db.Query(`select nome from tbPessoa `)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var DadosAtuais Pessoa
		rows.Scan(&DadosAtuais.Nome)
		DadosSlice = append(DadosSlice, DadosAtuais)
	}

	tpl.ExecuteTemplate(w, "login.gohtml", DadosSlice)

}

func main() {

	http.HandleFunc("/", crudVitao)
	log.Fatal(http.ListenAndServe(":8010", nil))

}
