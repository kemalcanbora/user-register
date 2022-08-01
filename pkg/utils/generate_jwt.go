package utils

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"rollic/internal/models"
	"time"
)

func GenerateJWT(user models.User) (string, error) {
	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("Something Went Wrong: %s\n", err.Error())
		return "", err
	}
	return tokenString, nil
}
