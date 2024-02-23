package services

import (
	"SimonBK_Login_APP/domain/models"
	"fmt"

	"gorm.io/gorm"
)

func CreateAplicationUser(db *gorm.DB, UserNameApp *string, HashedPassword string, FkCompany *uint, FkCustomer *uint) (*models.AplicationUser, interface{}) {
	user := &models.AplicationUser{
		UserNameApp: UserNameApp,
		Password:    HashedPassword,
		FkCompany:   FkCompany,
		FkCustomer:  FkCustomer,
	}

	if err := db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("[CreateAplicationUser]-Error en la creacion del Usuario: %w", err)
	}

	return user, "Creación de usuario exitosa"
}
