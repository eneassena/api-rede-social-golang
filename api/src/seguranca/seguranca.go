package seguranca

import "golang.org/x/crypto/bcrypt"

// gera a hash a parti da senha passada no paramentro
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// compara a senha com a hash da senha
func VerificarSenha(senhaComHash string, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
