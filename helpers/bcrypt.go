package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (res string) {
	salt := 10
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		log.Println("error generate")
	}

	res = string(hash)
	return
}

func ComparePassword(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		log.Println("Error compare hash and pass")
		return false
	}

	return true
}
