package utils

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/frankko93/spaceguru-challenge/src/api/config"
)

func CreateToken(email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	mySecret := config.SecretTokenSpaceGuru
	signedToken, err := token.SignedString([]byte(mySecret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func CheckToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.SecretTokenSpaceGuru), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractEmailFromToken(token *jwt.Token) (string, error) {
	var email string

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return email, errors.New("could not get token data.")
	}

	email, ok = claims["email"].(string)
	if !ok {
		return email, errors.New("could not get token email.")
	}

	return email, nil
}
