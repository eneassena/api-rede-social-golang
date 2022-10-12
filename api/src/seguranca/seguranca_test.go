package seguranca_test

import (
	"api/src/seguranca"
	"log"
	"testing"
)

func TestHash(t *testing.T) {

	senhaUser := "senha123"
	// senhaComHash, _ := seguranca.Hash(senhaUser)
	t.Run("verifica senha com hash", func(t *testing.T) {
		erro := seguranca.VerificarSenha(string("$2a$10$0SigwL.ZH92Wnl4hT5uh4efvYkrvqu1VRjnJIA11SZ/dPH4qkPEXC"), "usuario100")
		if erro != nil {
			t.Errorf("seguranca.VerificarSenha() = %v", erro.Error())
		}
		log.Println("senha ", senhaUser, " esta correta")
	})

}
