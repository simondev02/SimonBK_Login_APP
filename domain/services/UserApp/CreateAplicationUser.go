package services

import (
	"SimonBK_Login_APP/domain/models"
	"errors"
	"log"

	"gorm.io/gorm"
)

func CreateAplicationUser(db *gorm.DB, UserNameApp *string, HashedPassword string, FkCompany *uint) (string, string, error) {
	user := &models.AplicationUser{
		UserNameApp: UserNameApp,
		Password:    HashedPassword,
		FkCompany:   FkCompany,
	}

	if err := db.Create(user).Error; err != nil {
		log.Print("[CreateAplicationUser] - Error en la creacion del ususario")
		return "", "", errors.New("error en la creacion del ususario")
	}

	return user.Password, "Creación de usuario exitosa", nil
}
