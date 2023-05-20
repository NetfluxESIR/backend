package server

import (
	"github.com/NetfluxESIR/backend/internal/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func generateToken(userId string) (models.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	stringToken, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return models.Token{}, err
	}
	return models.Token{
		Token:          stringToken,
		AccountId:      userId,
		ExpirationDate: time.Now().Add(24 * time.Hour),
		Expire:         false,
	}, nil
}
