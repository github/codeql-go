package main

import (
	"github.com/dgrijalva/jwt-go"
)

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	// GOOD: The secret is not an empty string.
	var hmacSampleSecret2 = []byte("thesecret")
	tokenString, err := token.SignedString(hmacSampleSecret2)
}
