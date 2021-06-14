package utils

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"time"
)

func GenerateClientToken(email string, client_id uint, role string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = client_id
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	fmt.Println(claims["id"])
	fmt.Println(claims["email"])
	fmt.Println(claims["exp"])

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(GetEnvVariable("JWT_CLIENT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}