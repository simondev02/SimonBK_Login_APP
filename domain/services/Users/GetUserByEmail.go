package services

import (
	"SimonBK_Login_APP/api/views"
	"SimonBK_Login_APP/domain/models"
	"errors"

	"gorm.io/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) (*views.Users, error) {
	var userModel models.Users
	if err := db.Where("email = ?", email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	// Obtener el nombre y apellido del contacto
	contactName, err := GetContactNameByID(db, userModel.ID)
	if err != nil {
		return nil, err
	}

	// Obtener el ID del rol y la descripción del rol
	userRole, err := GetUserRoleByID(db, userModel.ID)
	if err != nil {
		return nil, err
	}

	// Crear una nueva instancia de views.Users y copiar los datos del modelo
	userView := &views.Users{
		ID:             userModel.ID,
		FkCompany:      userModel.FkCompany,
		FkCustomer:     userModel.FkCustomer,
		Email:          userModel.Email,
		Name:           contactName.Name,
		Lastname:       contactName.LastName,
		FkRole:         userRole.FkRole,
		Role:           userRole.Role,
		Login_attempts: userModel.Login_attempts,
		Last_login:     userModel.Last_login,
	}

	return userView, nil
}
