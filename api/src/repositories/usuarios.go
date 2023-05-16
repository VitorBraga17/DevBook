package repositories

import (
	"api/src/models"
	"database/sql"
	"log"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repository Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into usuarios (nome, nick, email, senha ) values (?, ? , ? , ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		log.Fatal(erro)
	}

	return uint64(ultimoIdInserido), nil
}
