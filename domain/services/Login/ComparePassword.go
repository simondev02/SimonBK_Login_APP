package services

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Print("[ComparePassword] - Contraseña no coincide")
			return false, errors.New("contraseña no coincide")
		}
		return false, err
	}
	return true, nil
}
