package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// recebe a string de conexão com o mysql
	StringConexaoBanco = ""
	// recebe a porta default que a api está rodando
	Porta = 0
)

// Carregar vai inicializar as variaveis de ambiente
func Carregar() {
	var err error
	if err = godotenv.Load(".env"); err != nil {
		log.Fatal(err.Error())
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME"))
}
