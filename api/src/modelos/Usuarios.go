package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

// Preparar vai chama os metodos de validar e formatar os dados do usuario
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O campo nome é obrigatório")
	}
	if usuario.Nick == "" {
		return errors.New("O campo nick é obrigatório")
	}
	if usuario.Email == "" {
		return errors.New("O campo email é obrigatório")
	}
	if etapa == "cadastro" {
		if usuario.Senha == "" {
			return errors.New("O campo senha é obrigatório")
		}
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {

	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Nick = strings.TrimSpace(usuario.Nick)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}
