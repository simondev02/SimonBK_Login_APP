package services

import (
	"SimonBK_Login_APP/api/views"
	Token "SimonBK_Login_APP/domain/services/AccesToken"
	Login "SimonBK_Login_APP/domain/services/Login"
	UserApp "SimonBK_Login_APP/domain/services/UserApp"
	Users "SimonBK_Login_APP/domain/services/Users"

	"gorm.io/gorm"
)

func HandleUser(db *gorm.DB, userNameApp, username, password string) (*views.Users, bool, string, error) {

	// 1. Validamos que el usuario de aplicación exista y retorna password cifrado.
	storedPassword, err := UserApp.GetPasswordForUserNameApp(db, userNameApp)
	if err != nil {
		return nil, false, "", err
	}

	// 2. Comparamos la contraseña ingresada con la contraseña cifrada almacenada.
	match, err := Login.ComparePasswords(*storedPassword, password)
	if err != nil {
		return nil, false, "", err
	}

	// 3. Obtenemos los datos del usuario por su email.
	user, err := Users.GetUserByEmail(db, username)
	if err != nil {
		return nil, false, "", err
	}

	// 4. Generamos el token de acceso.
	token, err := Token.GenerateAccessToken(*user)
	if err != nil {
		return nil, false, "", err
	}

	return user, match, token, nil
}
