package services

import (
	"SimonBK_Login_APP/domain/models"
	"errors"
	"log"

	"gorm.io/gorm"
)

func GetUserApp(db *gorm.DB, userNameApp *string) (string, uint, *uint, error) {

	var user models.AplicationUser

	if err := db.Where("user_name_app = ?", userNameApp).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Print("[GetUserApp] - Usuario aplicacion no existe")
			return "", 0, nil, errors.New("usuario aplicacion no existe")
		}
		return "", 0, nil, err
	}

	if user.Password == "" {
		log.Print("[GetUserApp] - Este usuario de aplicacion no tiene contraseña creada")
		return "", 0, nil, errors.New("este usuario de aplicacion no tiene contraseña creada")
	}

	return user.Password, user.ID, user.FkCompany, nil
}
