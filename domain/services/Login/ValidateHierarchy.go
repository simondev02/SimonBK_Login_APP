package services

import (
	"errors"
	"fmt"
	"log"
)

func ValidateHierarchyCompany(UserFkCompany *uint, FkCompany *uint) error {
	fmt.Println("UserFkCompany", *UserFkCompany, "FkCompany", *FkCompany)
	if *UserFkCompany != *FkCompany {
		log.Print("[ValidateHierarchyCompany] - No tiene permisos para acceder a esta empresa")
		return errors.New("no tiene permisos para acceder a esta empresa")
	}
	return nil
}
