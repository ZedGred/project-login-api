package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(userpassword string, providepassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userpassword), []byte(providepassword))
	return err
}
