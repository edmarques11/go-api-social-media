package security_test

import (
	"api/src/security"
	"testing"
)

func TestCreateAndValidateHash(t *testing.T) {
	password := "passwordTest"
	hash, err := security.CreateHash(password)
	if err != nil {
		t.Error("Não foi possível criar o hash")
	}

	if err = security.VerifyPassword(password, string(hash)); err != nil {
		t.Error("Não foi possível validar o hash")
	}
}
