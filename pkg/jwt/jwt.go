package jwt

import (
	"log"
	"time"

	jwt2 "github.com/dgrijalva/jwt-go"
)

// SecretKey secret key being used to sign tokens
var (
	SecretKey = []byte("rWTzsYnYh7koTlv6Ins6pAStivEuveNz")
)

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(username string) (string, error) {
	token := jwt2.New(jwt2.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt2.MapClaims)
	/* Set token claims */
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt2.Parse(tokenStr, func(token *jwt2.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt2.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
