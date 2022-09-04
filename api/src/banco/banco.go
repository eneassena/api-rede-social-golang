package banco

import (
	"database/sql"

	"api/src/config"

	_ "github.com/go-sql-driver/mysql" // Driver Mysql
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
