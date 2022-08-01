package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetHash(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, 14)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
