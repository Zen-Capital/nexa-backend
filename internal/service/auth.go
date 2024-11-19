package service

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var privateKey *rsa.PrivateKey

func init() {
	keyBytes, err := os.ReadFile("keys/private.key")
	if err != nil {
		panic("Erro ao carregar chave privada: " + err.Error())
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		panic("Erro ao parsear chave privada: " + err.Error())
	}
}
