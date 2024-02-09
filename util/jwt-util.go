package util

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/ssh"
	"os"
	"time"
)

var PrivateKey *rsa.PrivateKey

func GenerateJwtToken(username string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(PrivateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func init() {
	PrivateKey = loadRsaPrivateKey()
}
func loadRsaPrivateKey() *rsa.PrivateKey {
	pemByte, err := os.ReadFile("certs/private.pem")
	if err != nil {
		panic("Cannot read private key file")
	}
	key, err := ssh.ParseRawPrivateKey(pemByte)
	if err != nil {
		panic("Cannot read private key file")
	}
	return key.(*rsa.PrivateKey)
}
