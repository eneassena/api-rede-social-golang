package repositorios

import (
	"database/sql"
	"fmt"

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

func (repositorio *usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	query := "SELECT id,nome,nick,email,criadoEm FROM usuarios WHERE nome LIKE ? or nick LIKE ?;"

	linhas, erro := repositorio.db.Query(query, nomeOuNick, nomeOuNick)
	if erro != nil {
		return []modelos.Usuario{}, erro
	}

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return []modelos.Usuario{}, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio *usuarios) BuscarPorId(usuarioId uint64) (modelos.Usuario, error) {
	query := "select id,nome,nick,email,criadoEm from usuarios where id=?"

	linhas, erro := repositorio.db.Query(query, usuarioId)
	if erro != nil {
		return modelos.Usuario{}, nil
	}

	if linhas.Next() {
		var usuario modelos.Usuario

		erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		)
		if erro != nil {
			return modelos.Usuario{}, nil
		}
		return usuario, nil
	}
	return modelos.Usuario{}, nil
}

/*
 * BuscarPorEmail busca um usuario pelo email e retorna o id e senha
 */
func (repositorio *usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	query := "select id,senha from devbook.usuarios where email=?"
	linhas, erro := repositorio.db.Query(query, email)
	if erro != nil {
		return modelos.Usuario{}, nil
	}

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro := linhas.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, nil
		}
	}
	return usuario, nil
}

func (repositorio *usuarios) Atualizer(usuarioId uint64, usuario modelos.Usuario) error {
	query := "update usuarios set nome=?,nick=?,email=? where id=?"
	var statement *sql.Stmt
	var erro error
	statement, erro = repositorio.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	linhas, erro := statement.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuarioId,
	)
	if erro != nil {
		return erro
	}
	linhasAfetadas, erro := linhas.RowsAffected()
	if linhasAfetadas != 0 {
		return erro
	}
	return nil
}

func (repositorio *usuarios) Remover(usuarioId uint64) error {
	query := "delete from usuarios where id=?"
	_, erro := repositorio.db.Query(query, usuarioId)
	if erro != nil {
		return erro
	}
	return nil
}
