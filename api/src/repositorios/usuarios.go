package repositorios

import (
	"database/sql"

	"api/src/modelos"
)

type usuarios struct {
	db *sql.DB
}

// A função NovoRepositorioDeUsuario cria um novo repositorio de usuario
func NovoRepositorioDeUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// insere um novo usuario no banco de dados
func (repositorio usuarios) Criar(usuarioInstance modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(INSERT_INTO)
	if erro != nil {
		return 0, erro
	}

	resultado, erro := statement.Exec(
		usuarioInstance.Nome,
		usuarioInstance.Nick,
		usuarioInstance.Email,
		usuarioInstance.Senha,
	)
	if erro != nil {
		return 0, erro
	}

	lastId, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(lastId), nil
}
