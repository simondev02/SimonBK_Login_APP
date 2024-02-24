package services

import (
	"errors"
	"log"
)

func ValidateUserNameApp(userName *string) error {
	if len(*userName) < 5 {
		log.Print("[ValidateUserNameApp] - El nombre de usuario debe tener al menos 5 caracteres")
		return errors.New("el nombre de usuario debe tener al menos 5 caracteres")
	}
	return nil
}
