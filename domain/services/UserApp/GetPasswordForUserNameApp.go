package services

import (
	"SimonBK_Login_APP/domain/models"
	"errors"

	"gorm.io/gorm"
)

func GetPasswordForUserNameApp(db *gorm.DB, userNameApp string) (*string, error) {
	var user models.AplicationUser
	if err := db.Where("user_name_app = ?", userNameApp).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario aplicacion no existe")
		}
		return nil, err
	}

	if user.Password == nil {
		return nil, errors.New("este usuario de aplicacion no tiene contraseña creada")
	}

	return user.Password, nil
}
