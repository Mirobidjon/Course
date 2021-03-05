package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	gp             = "jksJHKJA565AJUmksk"
	tokenStringKey = "as54sadjUHHKAsa"
)

type tokenClaims struct {
	jwt.StandardClaims
	ID   int
	Role string
}

func GetToken(id int, role string) (string, error) {
	tk := tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
		role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	return token.SignedString([]byte(tokenStringKey))
}

func (r *AuthMasterService) ParseToken(tk string) (string, int, error) {
	token, err := jwt.ParseWithClaims(tk, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(tokenStringKey), nil
	})

	if err != nil {
		return "", 0, err
	}

	cl, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", 0, errors.New("token claims are not type of *tokenClaims")
	}

	return cl.Role, cl.ID, nil
}

func generatePassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(gp)))
}
