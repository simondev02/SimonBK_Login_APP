package services

import (
	views "SimonBK_Login_APP/api/views/Inpunts"
	"SimonBK_Login_APP/domain/models"
	services "SimonBK_Login_APP/domain/services/Login"
	UserApp "SimonBK_Login_APP/domain/services/UserApp"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UpdateAppUser(db *gorm.DB, userApp views.AplicationUserInputs) (string, error) {

	// 1. Validamos jerarquia Solo Usuarios Simon pueden crerar
	err := UserApp.ValidateHierarchy(userApp.FkCompany)
	if err != nil {
		log.Print("[UpdateAppUser]")
		return "", err
	}

	// 2 . Validamos el nombre de usuario
	var aplicationUser models.AplicationUser
	if err := db.First(&aplicationUser, userApp.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("usuario aplicacion no encontrado")
		}
	}

	// 3 . Configurar Usuario que actualizo la contraseña
	aplicationUser.UpdatedBy = userApp.UserID

	// 4 . generamos una contraseña aleatoria
	password := services.GeneratePassword()

	// 5 . Ciframos la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[CreateUserApp]")
		return "", err
	}

	// 6 . Actualizamos la contraseña
	aplicationUser.Password = string(hashedPassword)

	if err := db.Save(&aplicationUser).Error; err != nil {
		return "", err
	}

	return password, nil
}
