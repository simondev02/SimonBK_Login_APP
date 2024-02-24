package services

import (
	services "SimonBK_Login_APP/domain/services/Login"
	UserApp "SimonBK_Login_APP/domain/services/UserApp"
	"SimonBK_Login_APP/infra/db"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserApp(UserNameApp *string, FkCompany *uint, ParamCompany *uint) (string, error) {

	// 1. Validamos jerarquia Solo Ususarios Simon pueden crerar
	err := UserApp.ValidateHierarchy(FkCompany)
	if err != nil {
		log.Print("[CreateUser]")
		return "", err
	}

	// 2 . Validamos el nombre de usuario
	err = UserApp.ValidateUserNameApp(UserNameApp)
	if err != nil {
		log.Print("[CreateUserApp]")
		return "", err
	}

	// 3 . generamos una contraseña aleatoria
	password := services.GeneratePassword()

	// 4 . Ciframos la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[CreateUserApp]")
		return "", err
	}

	// 5 . Creacion de usuario aplicacion
	UserApp.CreateAplicationUser(db.DBConn, UserNameApp, string(hashedPassword), ParamCompany)
	if err != nil {
		log.Print("[CreateUserApp]")
		return "", err
	}
	return password, nil

}
