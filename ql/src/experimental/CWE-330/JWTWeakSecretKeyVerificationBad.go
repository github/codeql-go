package main

import (
	"github.com/dgrijalva/jwt-go"
)

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	// BAD: The secret is an empty string.
	var hmacSampleSecret = []byte("")
	tokenString, err := token.SignedString(hmacSampleSecret)
}
