package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.Carregar()

	r := router.Gerar()

	message := fmt.Sprintf("App Rede Social\nURI: http://127.0.0.1:%v", config.Porta)

	log.Println(message)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
 
