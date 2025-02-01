package token

import (
	"fmt"
	"time"
	"auth/main/service/role"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("secret-key")

func CreateToken(username string) (string, error) {
	r, err := role.GetRole(username)
	if err != nil {
		fmt.Println("Error with validationg role!")
		return "", nil
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "app", // TODO: add application
		"aud": r,
		"exp": time.Now().Add(time.Hour).Unix(), // TODO: add parameter
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString(secret)
	if err != nil {
		fmt.Println("Error with signing token!")
		return "", nil
	}

	return token, nil
}