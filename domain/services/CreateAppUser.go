package services

import (
	services "SimonBK_Login_APP/domain/services/Login"
	UserApp "SimonBK_Login_APP/domain/services/UserApp"
	"SimonBK_Login_APP/infra/db"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserApp(UserNameApp *string, FkCompany *uint, FkCustomer *uint) (string, error) {

	// 1 . Validar UserNameApp
	if UserNameApp == nil || len(*UserNameApp) < 5 {
		return "", fmt.Errorf("el nombre de usuario no puede ser nulo y debe tener al menos 5 caracteres")
	}

	// 2 . generamos una contraseña aleatoria
	password := services.GeneratePassword()

	// 3 . Ciframos la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Manejar el error
		fmt.Println("Error al cifrar la contraseña:", err)
		return "", err
	}

	// 4 . Creacion de usuario aplicacion
	UserApp.CreateAplicationUser(db.DBConn, UserNameApp, string(hashedPassword), FkCompany, FkCustomer)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return password, nil

}
