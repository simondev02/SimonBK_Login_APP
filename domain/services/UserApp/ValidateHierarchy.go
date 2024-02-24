package services

import (
	"errors"
	"log"
)

func ValidateHierarchy(FkCompany *uint) error {
	if FkCompany != nil {
		log.Print("[ValidateHierarchy] - Este usuario no puede crear ususarios de aplicacion")
		return errors.New("este usuario no puede crear ususarios de aplicacion")
	}
	return nil
}
